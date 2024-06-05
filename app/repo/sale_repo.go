package repo

import (
	"app/domain"
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/sirupsen/logrus"
	"golang.org/x/xerrors"
	"strings"
	"time"
)

type postgresSaleRepo struct {
	db *pgx.Conn
}

func NewPostgresSaleRepo(conn *pgx.Conn) domain.ISaleRepo {
	return &postgresSaleRepo{db: conn}
}

func (p *postgresSaleRepo) Add(c context.Context, sale *domain.Sale, lg *logrus.Logger) error {
	query := `insert into sales (name, expiredtime, percent) values ($1, $2, $3)`

	lg.Info("sale repo add")

	_, err := p.db.Exec(c, query, sale.Name, sale.ExpiredTime, sale.Percent)
	if err != nil {
		lg.Errorf("bad add sale")
		return xerrors.Errorf("sale repo: add error: %v", err.Error())
	}

	return nil
}

func (p *postgresSaleRepo) GetById(c context.Context, id int, lg *logrus.Logger) (domain.Sale, error) {
	query := `select * from sales where id = $1`

	lg.Info("sale repo get by id")

	var sale domain.Sale
	err := p.db.QueryRow(c, query, id).Scan(&sale.ID, &sale.Name, &sale.ExpiredTime, &sale.Percent)
	if err != nil {
		lg.Errorf("bad select by id")
		return domain.Sale{}, xerrors.Errorf("sale repo: getbyid error: %v", err.Error())
	}

	return sale, nil
}

func (p *postgresSaleRepo) Update(c context.Context, id int, newState *domain.Sale, lg *logrus.Logger) error {
	query := `update sales set id = $1,
                 				name = $2,
                 				expiredtime = $3,
                 				percent = $4
                 				where id = $5`

	lg.Info("sale repo update")
	_, err := p.db.Exec(c, query, newState.ID, newState.Name,
		newState.ExpiredTime, newState.Percent, id)
	if err != nil {
		lg.Errorf("bad update sale")
		return xerrors.Errorf("sale repo: update error: %v", err.Error())
	}

	return nil
}

func (p *postgresSaleRepo) Delete(c context.Context, id int, lg *logrus.Logger) error {
	query := `delete from sales where id = $1`

	lg.Info("sale repo delete")
	_, err := p.db.Exec(c, query, id)
	if err != nil {
		lg.Errorf("bad delete sale")
		return xerrors.Errorf("sale repo: delete error: %v", err.Error())
	}

	return nil
}

func (p *postgresSaleRepo) GetByCriteria(c context.Context, offset int, limit int, saleCriteria *domain.Sale, lg *logrus.Logger) ([]domain.Sale, error) {
	query := `select * from sales`
	values := make([]interface{}, 0)
	values = append(values, limit, offset)
	where, critValues := addSaleParams(saleCriteria)

	lg.Info("sale repo: get by criteria")

	if len(where) > 0 {
		query += ` where ` + strings.Join(where, ` and `)
	}

	values = append(values, critValues...)

	query += ` limit $1 offset $2`

	rows, err := p.db.Query(c, query, values...)
	defer rows.Close()

	if err != nil {
		lg.Errorf("bad select by criteria")
		return nil, xerrors.Errorf("sale repo: getbycriteria error: %v", err.Error())
	}

	sales := []domain.Sale{}
	for rows.Next() {
		sale := domain.Sale{}
		err := rows.Scan(&sale.ID, &sale.Name, &sale.ExpiredTime, &sale.Percent)
		if err != nil {
			return nil, xerrors.Errorf("sale repo: getbycriteria error: %w", err.Error())
		}
		sales = append(sales, sale)
	}

	return sales, nil
}

func addSaleParams(criteria *domain.Sale) ([]string, []interface{}) {
	where := make([]string, 0)
	values := make([]interface{}, 0)
	cnt := 3
	if criteria.Name != "" {
		where = append(where, fmt.Sprintf(`name=$%d`, cnt))
		cnt += 1
		values = append(values, criteria.Name)
	}
	if criteria.ExpiredTime != time.Date(1970, 1,
		1, 0, 0, 0, 0, time.UTC) {
		where = append(where, fmt.Sprintf(`expiredtime=$%d`, cnt))
		cnt += 1
		values = append(values, criteria.ExpiredTime)
	}
	if criteria.Percent != domain.DefaultEmptyValue {
		where = append(where, fmt.Sprintf(`percent=$%d`, cnt))
		values = append(values, criteria.Percent)
	}

	return where, values
}

func (p *postgresSaleRepo) GetLimit(c context.Context, offset int, limit int, lg *logrus.Logger) ([]domain.Sale, error) {
	query := `select * from sales limit $1 offset $2`

	lg.Info("sale repo getlimit")

	rows, err := p.db.Query(c, query, limit, offset)
	defer rows.Close()

	if err != nil {
		lg.Errorf("bad select sales")
		return nil, xerrors.Errorf("sale repo: getlimit error: %v", err.Error())
	}

	sales := []domain.Sale{}
	for rows.Next() {
		sale := domain.Sale{}
		err := rows.Scan(&sale.ID, &sale.Name, &sale.ExpiredTime, &sale.Percent)
		if err != nil {
			return nil, xerrors.Errorf("sale repo: getbycriteria error: %w", err.Error())
		}
		sales = append(sales, sale)
	}

	return sales, nil
}
