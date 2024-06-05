package repo

import (
	"app/domain"
	"context"
	"database/sql"
	"github.com/jackc/pgx/v5"
	"github.com/sirupsen/logrus"
	"golang.org/x/xerrors"
)

type postgresClientRepo struct {
	db *pgx.Conn
}

func NewPostgresClientRepo(conn *pgx.Conn) domain.IClientRepo {
	return &postgresClientRepo{db: conn}
}

func (p *postgresClientRepo) Add(c context.Context, clnt *domain.Client, lg *logrus.Logger) (domain.Client, error) {
	query := `insert into clients (user_id, name, surname, mail, phone) values ($1, $2, $3, $4, $5)`

	lg.Info("client repo add")

	_, err := p.db.Exec(c, query, clnt.AccID, clnt.Name, clnt.Surname, clnt.Mail, clnt.Phone)
	if err != nil {
		lg.Errorf("bad client insert")
		return domain.Client{}, xerrors.Errorf("client repo: add error: %v", err.Error())
	}

	return *clnt, nil
}

func (p postgresClientRepo) Delete(c context.Context, id int, lg *logrus.Logger) error {
	tx, err := p.db.BeginTx(c, pgx.TxOptions{})
	if err != nil {
		lg.Warnf("client repo: delete err: begin transaction")
		return xerrors.Errorf("client repo: delete error: %v", err.Error())
	}

	defer func() {
		if err != nil {
			rollbackErr := tx.Rollback(c)
			if rollbackErr != nil {
				err = xerrors.Errorf("client repo: delete error: %v", err.Error())
			}
		}
	}()
	query := `select * from clients where id = $1`
	row := tx.QueryRow(c, query, id)
	var clnt domain.Client
	err = row.Scan(&clnt.ID, &clnt.AccID, &clnt.Name, &clnt.Surname,
		&clnt.Mail, &clnt.Phone)
	if err != nil {
		lg.Error("bad client delete")
		return xerrors.Errorf("client repo: delete error: %v", err.Error())
	}

	query = `delete from clients where id = $1`

	lg.Info("client repo delete")

	_, err = tx.Exec(c, query, id)
	if err != nil {
		lg.Error("bad client delete")
		return xerrors.Errorf("client repo: delete error: %v", err.Error())
	}

	query = `delete from accounts where id=$1`

	lg.Info("cascade delete account from client repo")

	_, err = tx.Exec(c, query, clnt.AccID)
	if err != nil {
		lg.Error("bad cascade delete account from client repo")
		return xerrors.Errorf("client repo: delete error: %v", err.Error())
	}

	if err = tx.Commit(c); err != nil {
		lg.Error("bad commit transaction")
		return xerrors.Errorf("manager repo: delete error: %v", err.Error())
	}

	return nil
}

func (p *postgresClientRepo) GetByNameSurname(c context.Context, name string, surname string, lg *logrus.Logger) ([]domain.Client, error) {
	query := `select * from clients where name=$1 and surname=$2`
	lg.Info("client repo: get by name and surname")

	rows, err := p.db.Query(c, query, name, surname)
	defer rows.Close()

	if err != nil {
		lg.Errorf("bad select clients")
		return nil, xerrors.Errorf("client repo: get by name and surname error: %v", err.Error())
	}

	clnts := []domain.Client{}
	for rows.Next() {
		clnt := domain.Client{}
		err := rows.Scan(&clnt.ID, &clnt.AccID, &clnt.Name, &clnt.Surname, &clnt.Mail, &clnt.Phone)
		if err != nil {
			return nil, xerrors.Errorf("sale repo: get by name and surname error: %v", err.Error())
		}
		clnts = append(clnts, clnt)
	}

	return clnts, nil
}

func (p *postgresClientRepo) GetByPhone(c context.Context, phone string, lg *logrus.Logger) ([]domain.Client, error) {
	query := `select * from clients where phone=$1`
	lg.Info("client repo: get by phone")

	rows, err := p.db.Query(c, query, phone)
	defer rows.Close()

	if err != nil {
		lg.Errorf("bad select clients")
		return nil, xerrors.Errorf("client repo: get by phone error: %v", err.Error())
	}

	clnts := []domain.Client{}
	for rows.Next() {
		clnt := domain.Client{}
		err := rows.Scan(&clnt.ID, &clnt.AccID, &clnt.Name, &clnt.Surname, &clnt.Mail, &clnt.Phone)
		if err != nil {
			return nil, xerrors.Errorf("sale repo: get by phone error: %v", err.Error())
		}
		clnts = append(clnts, clnt)
	}

	return clnts, nil
}

func (p *postgresClientRepo) Update(c context.Context, id int, newState *domain.Client, lg *logrus.Logger) error {
	query := `update clients set id = $1,
                   name = $2,
                   surname = $3,
                   mail = $4,
                   phone = $5
                   where id = $6`

	lg.Info("client repo update")

	_, err := p.db.Exec(c, query, newState.ID, newState.Name, newState.Surname,
		newState.Mail, newState.Phone)
	if err != nil {
		lg.Error("bad update client")
		return xerrors.Errorf("client repo: update error: %v", err.Error())
	}

	return nil
}

func (p *postgresClientRepo) GetActiveRequestsByID(c context.Context, id int, lg *logrus.Logger) ([]domain.Request, error) {
	query := `select * from requests where clnt_id = $1 and 
                             (status='подтверждена' or 
                              status='принята' or status='обрабатывается')`
	var tourID sql.NullInt64
	lg.Info("client repo get active requests")

	rows, err := p.db.Query(c, query, id)
	defer rows.Close()

	if err != nil {
		lg.Errorf("bad select active requests")
		return nil, xerrors.Errorf("client repo: get active requests error: %v", err.Error())
	}

	reqs := []domain.Request{}
	for rows.Next() {
		req := domain.Request{}
		err := rows.Scan(&req.ID, &tourID, &req.ClntID, &req.MngrID,
			&req.Status, &req.CreateTime, &req.ModifyTime, &req.Data)
		if err != nil {
			return nil, xerrors.Errorf("client repo: get active requests error: %v", err.Error())
		}
		if tourID.Valid {
			req.TourID = int(tourID.Int64)
		} else {
			req.TourID = domain.DefaultEmptyValue
		}
		reqs = append(reqs, req)
	}

	return reqs, nil
}

func (p *postgresClientRepo) GetDoneRequestsByID(c context.Context, id int, lg *logrus.Logger) ([]domain.Request, error) {
	query := `select * from requests where clnt_id = $1 and 
                             status='оплачена'`
	var tourID sql.NullInt64
	lg.Info("client repo get done requests")

	rows, err := p.db.Query(c, query, id)
	defer rows.Close()

	if err != nil {
		lg.Errorf("bad select done requests")
		return nil, xerrors.Errorf("client repo: get done requests error: %v", err.Error())
	}

	reqs := []domain.Request{}
	for rows.Next() {
		req := domain.Request{}
		err := rows.Scan(&req.ID, &tourID, &req.ClntID, &req.MngrID,
			&req.Status, &req.CreateTime, &req.ModifyTime, &req.Data)
		if err != nil {
			return nil, xerrors.Errorf("client repo: get done requests error: %v", err.Error())
		}
		if tourID.Valid {
			req.TourID = int(tourID.Int64)
		} else {
			req.TourID = domain.DefaultEmptyValue
		}
		reqs = append(reqs, req)
	}

	return reqs, nil
}

func (p *postgresClientRepo) GetById(c context.Context, id int, lg *logrus.Logger) (domain.Client, error) {
	query := `select * from clients where id=$1`
	lg.Info("client repo: get by id")

	var clnt domain.Client
	err := p.db.QueryRow(c, query, id).Scan(&clnt.ID, &clnt.AccID, &clnt.Name, &clnt.Surname,
		&clnt.Mail, &clnt.Phone)
	if err != nil {
		lg.Errorf("bad select by id")
		return domain.Client{}, xerrors.Errorf("client repo: getbyid error: %v", err.Error())
	}

	return clnt, nil
}
