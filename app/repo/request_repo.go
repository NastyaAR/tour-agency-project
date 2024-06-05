package repo

import (
	"app/domain"
	"context"
	"database/sql"
	"github.com/jackc/pgx/v5"
	"github.com/sirupsen/logrus"
	"golang.org/x/xerrors"
	"math/rand"
	"time"
)

type postgresRequestRepo struct {
	db *pgx.Conn
}

func NewPostgresRequestRepo(conn *pgx.Conn) domain.IRequestRepo {
	return &postgresRequestRepo{db: conn}
}

func (p *postgresRequestRepo) GetByStatus(c context.Context, status string, offset int, limit int, lg *logrus.Logger) ([]domain.Request, error) {
	var tourID sql.NullInt64
	query := `select * from requests where status = $1 limit $2 offset $3`

	lg.Info("request repo get by status")

	rows, err := p.db.Query(c, query, status, limit, offset)
	defer rows.Close()

	if err != nil {
		lg.Errorf("bad select requests")
		return nil, xerrors.Errorf("request repo: getbystatus error: %v", err.Error())
	}

	requests := []domain.Request{}
	for rows.Next() {
		req := domain.Request{}
		err = rows.Scan(&req.ID, &tourID, &req.ClntID, &req.MngrID,
			&req.Status, &req.CreateTime, &req.ModifyTime, &req.Data)
		if tourID.Valid {
			req.TourID = int(tourID.Int64)
		} else {
			req.TourID = domain.DefaultEmptyValue
		}
		if err != nil {
			return nil, xerrors.Errorf("request repo: getbystatus error: %w", err.Error())
		}
		requests = append(requests, req)
	}

	return requests, nil
}

func (p *postgresRequestRepo) Add(c context.Context, req *domain.Request, lg *logrus.Logger) error {
	rand.Seed(time.Now().UnixNano())

	query := `select id from managers`

	rows, err := p.db.Query(c, query)
	defer rows.Close()

	if err != nil {
		lg.Errorf("bad select managers id")
		return xerrors.Errorf("request repo: add request error: %v", err.Error())
	}

	ids := make([]int, 0)
	for rows.Next() {
		mngr := domain.Manager{}
		err := rows.Scan(&mngr.ID)
		if err != nil {
			return xerrors.Errorf("request repo: add error: %v", err.Error())
		}
		ids = append(ids, mngr.ID)
	}

	randomIndex := rand.Intn(len(ids))
	lg.Warnf("request repo add")

	if req.TourID == domain.DefaultEmptyValue {
		query = `insert into requests (clnt_id, mngr_id, status,
                      create_time, modify_time, data) values ($1,
                    	$2, $3, $4, $5, $6)`
		_, err = p.db.Exec(c, query, req.ClntID, ids[randomIndex],
			domain.Accepted, time.Now(), time.Now(), req.Data)
	} else {
		query = `insert into requests (tour_id, clnt_id, mngr_id, status,
                      create_time, modify_time, data) values ($1, $2,
                    	$3, $4, $5, $6, $7)`
		_, err = p.db.Exec(c, query, req.TourID, req.ClntID, ids[randomIndex],
			domain.Accepted, time.Now(), time.Now(), req.Data)
	}
	if err != nil {
		lg.Errorf("bad insert req repo")
		return xerrors.Errorf("request repo: add error: %v", err.Error())
	}

	return nil
}

func (p *postgresRequestRepo) Update(c context.Context, id int, reqCriteria *domain.Request, lg *logrus.Logger) error {
	query := `update requests set
                    tour_id = $1,
                    clnt_id = $2,
                    mngr_id = $3,
                    status = $4,
                    create_time = $5,
                    modify_time = $6,
                    data = $7
							where id = $8`

	lg.Info("request repo update")
	_, err := p.db.Exec(c, query, reqCriteria.TourID, reqCriteria.ClntID,
		reqCriteria.MngrID, reqCriteria.Status, reqCriteria.CreateTime, time.Now(),
		reqCriteria.Data, id)
	if err != nil {
		lg.Errorf("update request error")
		return xerrors.Errorf("request repo: update error: %v", err.Error())
	}
	return nil
}

func (p *postgresRequestRepo) GetById(c context.Context, id int, lg *logrus.Logger) (domain.Request, error) {
	var tourID sql.NullInt64
	query := `select * from requests where id = $1`

	lg.Info("request repo getbyid")

	var req domain.Request
	err := p.db.QueryRow(c, query, id).Scan(&req.ID, &tourID, &req.ClntID, &req.MngrID,
		&req.Status, &req.CreateTime, &req.ModifyTime, &req.Data)
	if tourID.Valid {
		req.TourID = int(tourID.Int64)
	} else {
		req.TourID = domain.DefaultEmptyValue
	}
	if err != nil {
		lg.Errorf("bad select by id")
		return domain.Request{}, xerrors.Errorf("request repo: getbyid error: %v", err.Error())
	}

	return req, nil
}

func (p *postgresRequestRepo) Reject(c context.Context, id int, lg *logrus.Logger) error {
	query := `update requests set status ='отклонена',
                    				modify_time = $1
                					where id = $2`

	lg.Info("request repo reject")

	_, err := p.db.Exec(c, query, time.Now(), id)
	if err != nil {
		lg.Errorf("bad update request status")
		return xerrors.Errorf("request repo: reject error: %v", err.Error())
	}
	return nil
}

func (p *postgresRequestRepo) Approve(c context.Context, id int, lg *logrus.Logger) error {
	query := `update requests set status = 'подтверждена',
                					modify_time = $1
                					where id = $2`

	lg.Info("request repo approve")

	_, err := p.db.Exec(c, query, time.Now(), id)
	if err != nil {
		lg.Errorf("bad update request status")
		return xerrors.Errorf("request repo: approve error: %v", err.Error())
	}
	return nil
}

func (p *postgresRequestRepo) CountFinalCost(c context.Context, id int, lg *logrus.Logger) (int, error) {
	query := `select count(*)
		from tours_sales join sales
		on tours_sales.sale_id = sales.id
		where tours_sales.tour_id = $1`
	var cnt int
	err := p.db.QueryRow(c, query, id).Scan(&cnt)
	if err != nil {
		lg.Errorf("bad count cost")
		return -1, xerrors.Errorf("request repo: count final cost error: %v", err.Error())
	}
	var cost int
	if cnt != 0 {
		query = `select tours.cost * (100 - (select sum(percent)
               				from tours_sales join sales
               				on tours_sales.sale_id = sales.id
               				where tours_sales.tour_id = $1)) / 100
			from tours
			where id = $2`
		err = p.db.QueryRow(c, query, id, id).Scan(&cost)
	} else {
		query = `select cost from tours where id = $1`
		err = p.db.QueryRow(c, query, id).Scan(&cost)
	}

	lg.Info("request repo count final cost")

	if err != nil {
		lg.Errorf("bad count cost: %v", err.Error())
		return -1, xerrors.Errorf("request repo: count final cost error: %v", err.Error())
	}

	return cost, nil
}

func (p *postgresRequestRepo) GetLimit(c context.Context, offset int, limit int, lg *logrus.Logger) ([]domain.Request, error) {
	query := `select * from requests limit $1 offset $2`
	var tourID sql.NullInt64

	lg.Info("request repo getlimit")

	rows, err := p.db.Query(c, query, limit, offset)
	defer rows.Close()

	if err != nil {
		lg.Errorf("bad select requests")
		return nil, xerrors.Errorf("request repo: getbystatus error: %v", err.Error())
	}

	requests := []domain.Request{}
	for rows.Next() {
		req := domain.Request{}
		err := rows.Scan(&req.ID, &tourID, &req.ClntID, &req.MngrID,
			&req.Status, &req.CreateTime, &req.ModifyTime, &req.Data)
		if err != nil {
			return nil, xerrors.Errorf("request repo: getbystatus error: %w", err.Error())
		}
		if tourID.Valid {
			req.TourID = int(tourID.Int64)
		} else {
			req.TourID = domain.DefaultEmptyValue
		}
		requests = append(requests, req)
	}

	return requests, nil
}

func (p *postgresRequestRepo) AtomicPay(c context.Context, finalCost int, req *domain.Request, lg *logrus.Logger) error {
	tx, err := p.db.BeginTx(c, pgx.TxOptions{})
	if err != nil {
		lg.Warn("bad begin transaction")
		return xerrors.Errorf("request repo: atomicpay error: %v", err.Error())
	}

	defer func() {
		if err != nil {
			rollbackErr := tx.Rollback(c)
			if rollbackErr != nil {
				err = xerrors.Errorf("request repo: atomicpay error: %v", err.Error())
			}
		}
	}()

	query := `update requests set status = 'оплачена',
                    				modify_time = $1
                					where id = $2`

	lg.Info("change status of request on payed")
	_, err = tx.Exec(c, query, time.Now(), req.ID)
	if err != nil {
		lg.Error("bad change status of request on payed")
		return xerrors.Errorf("request repo: atomicpay error: %v", err.Error())
	}

	query = `insert into request_outbox (req_id, sum, state)
				values ($1, $2, 'non-send')`

	lg.Info("inserting into request_outbox")
	_, err = tx.Exec(c, query, req.ID, finalCost)
	if err != nil {
		lg.Error("bad insert into request-outbox")
		return xerrors.Errorf("request repo: atomicpay error: %v", err.Error())
	}

	if err = tx.Commit(c); err != nil {
		lg.Error("bad commit transaction")
		return xerrors.Errorf("request repo: atomicpay error: %v", err.Error())
	}

	return nil
}

func (p *postgresRequestRepo) UpdateOutbox(c context.Context, id int, lg *logrus.Logger) error {
	query := `update request_outbox set state = 'send' where req_id = $1`

	lg.Info("update outbox")

	_, err := p.db.Exec(c, query, id)
	if err != nil {
		lg.Error("bad update outbox")
		return xerrors.Errorf("pay repo: updateoutbox error: %v", err.Error())
	}

	return nil
}

func (p *postgresRequestRepo) GetNonSendEvents(c context.Context, limit int, lg *logrus.Logger) ([]domain.PayEvent, error) {
	query := `select * from request_outbox where state='non-send' limit $1`

	lg.Info("pay repo get nonsend events")

	rows, err := p.db.Query(c, query, limit)
	defer rows.Close()

	if err != nil {
		lg.Errorf("bad select request events")
		return nil, xerrors.Errorf("request repo: get nonsend events error: %v", err.Error())
	}

	events := []domain.PayEvent{}
	for rows.Next() {
		event := domain.PayEvent{}
		err := rows.Scan(&event.Id, &event.ReqID,
			&event.Sum, &event.State)
		if err != nil {
			return nil, xerrors.Errorf("request repo: get nonsend events error: %w", err.Error())
		}
		events = append(events, event)
	}

	return events, nil
}

func (p *postgresRequestRepo) AddTour(c context.Context, id int, tour_id int, lg *logrus.Logger) error {
	query := `update requests set tour_id = $1, 
                    			status = 'обрабатывается',
                    			modify_time = $3
                where id = $2`
	lg.Info("add tour into request")

	_, err := p.db.Exec(c, query, tour_id, id, time.Now())
	if err != nil {
		lg.Errorf("bad add tour: %v", err.Error())
		return xerrors.Errorf("request repo: addtour error: %v", err.Error())
	}

	return nil
}

func (p *postgresRequestRepo) GetRequestsForClient(c context.Context, clnt_id int, offset int, limit int, lg *logrus.Logger) ([]domain.Request, error) {
	query := `select * from requests 
			where clnt_id = $1
			limit $2 offset $3`
	lg.Info("get requests for client")

	var tourID sql.NullInt64

	rows, err := p.db.Query(c, query, clnt_id, limit, offset)
	defer rows.Close()

	if err != nil {
		lg.Errorf("bad select requests")
		return nil, xerrors.Errorf("request repo: getbystatus error: %v", err.Error())
	}

	requests := []domain.Request{}
	for rows.Next() {
		req := domain.Request{}
		err := rows.Scan(&req.ID, &tourID, &req.ClntID, &req.MngrID,
			&req.Status, &req.CreateTime, &req.ModifyTime, &req.Data)
		if err != nil {
			return nil, xerrors.Errorf("request repo: get requests for client error: %w", err.Error())
		}
		if tourID.Valid {
			req.TourID = int(tourID.Int64)
		} else {
			req.TourID = domain.DefaultEmptyValue
		}
		requests = append(requests, req)
	}

	return requests, nil
}
