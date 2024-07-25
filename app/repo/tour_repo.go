package repo

import (
	"app/domain"
	"context"
	"database/sql"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/sirupsen/logrus"
	"golang.org/x/xerrors"
	"strings"
	"time"
)

type postgresTourRepo struct {
	db *pgx.Conn
}

func NewPostgresTourRepo(conn *pgx.Conn) domain.ITourRepo {
	return &postgresTourRepo{db: conn}
}

func (p *postgresTourRepo) GetByCriteria(c context.Context, offset int, limit int, criteria *domain.Tour, lg *logrus.Logger) ([]domain.Tour, error) {
	query := `select * from tours`
	values := make([]interface{}, 0)
	values = append(values, limit, offset)
	where, critValues := addParams(criteria)

	lg.Info("postgres repo: get by criteria")

	if len(where) > 0 {
		query += ` where ` + strings.Join(where, ` and `)
	}

	values = append(values, critValues...)

	query += ` limit $1 offset $2`

	rows, err := p.db.Query(c, query, values...)
	defer rows.Close()

	if err != nil {
		lg.Errorf("bad select by criteria")
		return nil, xerrors.Errorf("tour repo: getbycriteria error: %v", err.Error())
	}

	tours := []domain.Tour{}
	for rows.Next() {
		tour := domain.Tour{}
		err := rows.Scan(&tour.ID, &tour.ChillPlace, &tour.FromPlace,
			&tour.Date, &tour.Duration, &tour.Cost,
			&tour.TouristsNumber, &tour.ChillType)
		if err != nil {
			return nil, xerrors.Errorf("tour repo: getbycriteria error: %v", err.Error())
		}
		tours = append(tours, tour)
	}
	return tours, nil
}

func addParams(criteria *domain.Tour) ([]string, []interface{}) {
	where := make([]string, 0)
	values := make([]interface{}, 0)
	cnt := 3
	if criteria.ChillPlace != "" {
		where = append(where, fmt.Sprintf(`chillplace=$%d`, cnt))
		cnt += 1
		values = append(values, criteria.ChillPlace)
	}
	if criteria.FromPlace != "" {
		where = append(where, fmt.Sprintf(`fromplace=$%d`, cnt))
		cnt += 1
		values = append(values, criteria.FromPlace)
	}
	if criteria.Date != time.Date(1970, 1,
		1, 0, 0, 0, 0, time.UTC) {
		where = append(where, fmt.Sprintf(`date=$%d`, cnt))
		cnt += 1
		values = append(values, criteria.Date)
	}
	if criteria.ChillType != "" {
		where = append(where, fmt.Sprintf(`chilltype=$%d`, cnt))
		cnt += 1
		values = append(values, criteria.ChillType)
	}
	if criteria.Duration != domain.DefaultEmptyValue {
		where = append(where, fmt.Sprintf(`duration=$%d`, cnt))
		cnt += 1
		values = append(values, criteria.Duration)
	}
	if criteria.Cost != domain.DefaultEmptyValue {
		where = append(where, fmt.Sprintf(`cost<$%d`, cnt))
		cnt += 1
		values = append(values, criteria.Cost)
	}
	if criteria.TouristsNumber != domain.DefaultEmptyValue {
		where = append(where, fmt.Sprintf(`tourists_number=$%d`, cnt))
		cnt += 1
		values = append(values, criteria.TouristsNumber)
	}

	return where, values
}

func (p *postgresTourRepo) GetById(c context.Context, id int, lg *logrus.Logger) (domain.Tour, error) {
	query := `select * from tours where id=$1`
	lg.Info("tour repo: get by id")

	var tour domain.Tour
	err := p.db.QueryRow(c, query, id).Scan(&tour.ID, &tour.ChillPlace, &tour.FromPlace,
		&tour.Date, &tour.Duration, &tour.Cost,
		&tour.TouristsNumber, &tour.ChillType)
	if err != nil {
		lg.Errorf("bad select by id")
		return domain.Tour{}, xerrors.Errorf("tour repo: getbyid error: %v", err.Error())
	}

	return tour, nil
}

func (p *postgresTourRepo) Add(c context.Context, tour *domain.Tour, lg *logrus.Logger) error {
	query := `insert into tours (chillplace, fromplace, date, duration,
                   cost, tourists_number, chilltype) values ($1, $2, $3, $4, $5, $6, $7)`

	lg.Info("tour repo add")

	_, err := p.db.Exec(c, query, tour.ChillPlace, tour.FromPlace,
		tour.Date, tour.Duration, tour.Cost,
		tour.TouristsNumber, tour.ChillType)
	if err != nil {
		lg.Errorf("add tour error")
		return xerrors.Errorf("tour repo: add error: %v", err.Error())
	}

	return nil
}

func (p *postgresTourRepo) Delete(c context.Context, id int, lg *logrus.Logger) error {
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

	query := `update requests set tour_id = null where tour_id = $1`
	_, err = tx.Exec(c, query, id)
	if err != nil {
		lg.Errorf("update req connected with tour error")
		return xerrors.Errorf("tour repo: delete error: %v", err.Error())
	}

	query = `delete from tours_sales where tour_id = $1`
	_, err = tx.Exec(c, query, id)
	if err != nil {
		lg.Errorf("delete from connected table tours_sales")
		return xerrors.Errorf("tour repo: delete error: %v", err.Error())
	}

	query = `delete from tours where id = $1`

	lg.Info("tour repo delete")

	_, err = tx.Exec(c, query, id)
	if err != nil {
		lg.Errorf("delete tour error")
		return xerrors.Errorf("tour repo: delete error: %v", err.Error())
	}

	if err = tx.Commit(c); err != nil {
		lg.Error("bad commit transaction")
		return xerrors.Errorf("tour repo: delete error: %v", err.Error())
	}

	return nil
}

func (p *postgresTourRepo) Update(c context.Context, id int, newState *domain.Tour, lg *logrus.Logger) error {
	query := `update tours set id = $1,
							   	chillplace = $2,
								fromplace = $3,
								date = $4,
								duration = $5,
								cost = $6,
								tourists_number = $7
								chilltype = $8 where id = $9`

	lg.Info("tour repo update")

	_, err := p.db.Exec(c, query, newState.ID, newState.ChillType,
		newState.FromPlace, newState.Date,
		newState.Duration, newState.Cost,
		newState.TouristsNumber, newState.ChillType, id)

	if err != nil {
		lg.Errorf("update tour error")
		return xerrors.Errorf("tour repo: update error: %v", err.Error())
	}

	return nil
}

func (p *postgresTourRepo) UpdateSale(c context.Context, id int, newSale *domain.Sale, lg *logrus.Logger) error {
	query := `insert into tours_sales (tour_id, sale_id) values ($1, $2)`

	lg.Info("tour repo updatesale")

	_, err := p.db.Exec(c, query, id, newSale.ID)
	if err != nil {
		lg.Errorf("update sale error")
		return xerrors.Errorf("tour repo: updatesale error: %v", err.Error())
	}

	return nil
}

func (p *postgresTourRepo) GetLimit(c context.Context, offset int, limit int, lg *logrus.Logger) ([]domain.Tour, error) {
	query := `select * from tours limit $1 offset $2`

	lg.Info("tour repo getlimit")

	rows, err := p.db.Query(c, query, limit, offset)
	defer rows.Close()

	if err != nil {
		lg.Errorf("bad select tours")
		return nil, xerrors.Errorf("tour repo: getlimit error: %v", err.Error())
	}

	tours := []domain.Tour{}
	for rows.Next() {
		tour := domain.Tour{}
		err := rows.Scan(&tour.ID, &tour.ChillPlace, &tour.FromPlace,
			&tour.Date, &tour.Duration, &tour.Cost,
			&tour.TouristsNumber, &tour.ChillType)
		if err != nil {
			return nil, xerrors.Errorf("tour repo: getlimit error: %w", err.Error())
		}
		tours = append(tours, tour)
	}

	return tours, nil
}

func (p *postgresTourRepo) GetHotTours(c context.Context, offset int, limit int, lg *logrus.Logger) ([]domain.HotTourDto, error) {
	query := `select tours.*, sales.name, sales.percent
				from tours join tours_sales on tours.id = tours_sales.tour_id
				join sales on sales.id = tours_sales.sale_id limit $1 offset $2`

	lg.Info("tour repo gethot")

	rows, err := p.db.Query(c, query, limit, offset)
	defer rows.Close()

	if err != nil {
		lg.Errorf("bad select tours")
		return nil, xerrors.Errorf("tour repo: gethot error: %v", err.Error())
	}

	tours := []domain.HotTourDto{}
	for rows.Next() {
		tour := domain.HotTourDto{}
		err := rows.Scan(&tour.ID, &tour.ChillPlace, &tour.FromPlace,
			&tour.Date, &tour.Duration, &tour.Cost,
			&tour.TouristsNumber, &tour.ChillType, &tour.SaleName, &tour.SalePercemt)
		if err != nil {
			return nil, xerrors.Errorf("tour repo: gethot error: %w", err.Error())
		}
		tours = append(tours, tour)
	}

	return tours, nil
}

func (p *postgresTourRepo) GetNumberOfTours(c context.Context, lg *logrus.Logger) (int, error) {
	query := `select count(*) from tours`
	lg.Info("tour repo get number of tours")

	row := p.db.QueryRow(c, query)

	var number sql.NullInt32
	err := row.Scan(&number)
	if err != nil {
		lg.Errorf("bad scan result from getnumberoftours: %v", err.Error())
		return domain.DefaultEmptyValue, xerrors.Errorf("tour repo: getnumberof tours error: %v", err.Error())
	}

	if !number.Valid {
		lg.Errorf("not valid number (null)")
		return domain.DefaultEmptyValue, xerrors.Errorf("tour repo: getnumbertours tours error: not valid number")
	}

	intNumber := int(number.Int32)
	return intNumber, nil
}
