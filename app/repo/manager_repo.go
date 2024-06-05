package repo

import (
	"app/domain"
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/sirupsen/logrus"
	"golang.org/x/xerrors"
	"time"
)

type postgresManagerRepo struct {
	db *pgx.Conn
}

func NewPostgresManagerRepo(conn *pgx.Conn) domain.IManagerRepo {
	return &postgresManagerRepo{db: conn}
}

func (p *postgresManagerRepo) Add(c context.Context, mngr *domain.Manager, lg *logrus.Logger) (domain.Manager, error) {
	query := `insert into managers (acc_id, name, surname, department)
				values ($1, $2, $3, $4)`

	lg.Info("manager repo add")

	_, err := p.db.Exec(c, query, mngr.AccID, mngr.Name, mngr.Surname, mngr.Department)
	if err != nil {
		lg.Error("bad insert manager")
		return domain.Manager{}, xerrors.Errorf("manager repo: add error: %v", err.Error())
	}

	return *mngr, nil
}

func (p *postgresManagerRepo) Delete(c context.Context, id int, lg *logrus.Logger) error {
	tx, err := p.db.BeginTx(c, pgx.TxOptions{})
	if err != nil {
		lg.Warnf("manager repo: delete err: begin transaction")
		return xerrors.Errorf("manager repo: delete error: %v", err.Error())
	}

	defer func() {
		if err != nil {
			rollbackErr := tx.Rollback(c)
			if rollbackErr != nil {
				err = xerrors.Errorf("client repo: delete error: %v", err.Error())
			}
		}
	}()

	query := `select * from managers where id = $1`
	row := tx.QueryRow(c, query, id)
	var mngr domain.Manager
	err = row.Scan(&mngr.ID, &mngr.AccID, &mngr.Name, &mngr.Surname, &mngr.Department)
	if err != nil {
		lg.Error("bad delete manager")
		return xerrors.Errorf("manager repo: delete error: %v", err.Error())
	}

	query = `delete from managers where id = $1`

	lg.Info("manager repo delete")

	_, err = tx.Exec(c, query, id)
	if err != nil {
		lg.Error("bad delete manager")
		return xerrors.Errorf("manager repo: delete error: %v", err.Error())
	}

	query = `delete from accounts where id=$1`

	lg.Info("cascade delete account from manager repo")

	_, err = tx.Exec(c, query, mngr.AccID)
	if err != nil {
		lg.Error("bad cascade delete account from manager repo")
		return xerrors.Errorf("manager repo: delete error: %v", err.Error())
	}

	if err = tx.Commit(c); err != nil {
		lg.Error("bad commit transaction")
		return xerrors.Errorf("manager repo: delete error: %v", err.Error())
	}

	return nil
}

func (p *postgresManagerRepo) GetByNameSurname(c context.Context, name string, surname string, lg *logrus.Logger) ([]domain.Manager, error) {
	query := `select * from managers where name=$1 and surname=$2`
	lg.Info("manager repo: get by name and surname")

	rows, err := p.db.Query(c, query, name, surname)
	defer rows.Close()

	if err != nil {
		lg.Errorf("bad select managers")
		return nil, xerrors.Errorf("manager repo: get by name and surname error: %v", err.Error())
	}

	mngrs := []domain.Manager{}
	for rows.Next() {
		mngr := domain.Manager{}
		err := rows.Scan(&mngr.ID, &mngr.AccID, &mngr.Name, &mngr.Surname, &mngr.Department)
		if err != nil {
			return nil, xerrors.Errorf("manager repo: get by name and surname error: %w", err.Error())
		}
		mngrs = append(mngrs, mngr)
	}

	return mngrs, nil
}

func (p *postgresManagerRepo) GetByDepartment(c context.Context, department string, lg *logrus.Logger) ([]domain.Manager, error) {
	query := `select * from managers where department=$1`
	lg.Info("manager repo: get by department")

	rows, err := p.db.Query(c, query, department)
	defer rows.Close()

	if err != nil {
		lg.Errorf("bad select managers")
		return nil, xerrors.Errorf("manager repo: get by department error: %v", err.Error())
	}

	mngrs := []domain.Manager{}
	for rows.Next() {
		mngr := domain.Manager{}
		err := rows.Scan(&mngr.ID, &mngr.AccID, &mngr.Name, &mngr.Surname, &mngr.Department)
		if err != nil {
			return nil, xerrors.Errorf("manager repo: get by department error: %w", err.Error())
		}
		mngrs = append(mngrs, mngr)
	}

	return mngrs, nil
}

func (p *postgresManagerRepo) GetById(c context.Context, id int, lg *logrus.Logger) (domain.Manager, error) {
	query := `select * from managers where id=$1`
	lg.Info("manager repo: get by id")

	var mngr domain.Manager
	err := p.db.QueryRow(c, query, id).Scan(&mngr.ID, &mngr.AccID, &mngr.Name, &mngr.Surname, &mngr.Department)
	if err != nil {
		lg.Errorf("bad select by id")
		return domain.Manager{}, xerrors.Errorf("manager repo: getbyid error: %v", err.Error())
	}

	return mngr, nil
}

func (p *postgresManagerRepo) Update(c context.Context, id int, newState *domain.Manager, lg *logrus.Logger) error {
	query := `update managers set id = $1,
                    acc_id = $2,
                    name = $3,
                    surname = $4,
                    department = $5
                    where id = $6`

	lg.Info("manager repo add")

	_, err := p.db.Exec(c, query, newState.ID, newState.AccID, newState.Name, newState.Surname,
		newState.Department, id)
	if err != nil {
		lg.Error("bad update manager")
		return xerrors.Errorf("manager repo: update error: %v", err.Error())
	}

	return nil
}

func (p *postgresManagerRepo) GetNumberServedRequests(c context.Context, id int, lg *logrus.Logger) (int, error) {
	query := `select count(*) from requests where mngr_id = $1 and status = 'оплачена'`

	lg.Info("manager repo get number served reqs")
	var req int
	err := p.db.QueryRow(c, query, id).Scan(&req)
	if err != nil {
		lg.Errorf("bad select by id")
		return -1, xerrors.Errorf("manager repo: get number served reqs error: %v", err.Error())
	}

	return req, nil
}

func (p *postgresManagerRepo) GetAllRequests(c context.Context, id int, lg *logrus.Logger) (int, error) {
	query := `select count(*) from requests where mngr_id = $1`

	lg.Info("manager repo get all reqs")
	var req int
	err := p.db.QueryRow(c, query, id).Scan(&req)
	if err != nil {
		lg.Errorf("bad select by id")
		return -1, xerrors.Errorf("manager repo: get all reqs error: %v", err.Error())
	}

	return req, nil
}

func (p *postgresManagerRepo) GetSumOnPeriod(c context.Context, id int, from time.Time, to time.Time, lg *logrus.Logger) (int, error) {
	query := `select sum(tours.cost)
			from requests join tours on requests.tour_id = tours.id
			join managers on requests.mngr_id = managers.id
			where status='оплачена' and
            modify_time between $1 and $2 and managers.id = $3`

	lg.Info("manager repo get sum on period")

	var sum int
	err := p.db.QueryRow(c, query, from, to, id).Scan(&sum)
	if err != nil {
		lg.Errorf("bad select sum")
		return -1, xerrors.Errorf("manager repo: get sum on period error: %v", err.Error())
	}

	return sum, nil
}

func (p *postgresManagerRepo) GetLimit(c context.Context, offset int, limit int, lg *logrus.Logger) ([]domain.Manager, error) {
	query := `select * from managers limit $1 offset $2`

	lg.Info("manager repo getlimit")

	rows, err := p.db.Query(c, query, limit, offset)
	defer rows.Close()

	if err != nil {
		lg.Errorf("bad select managers")
		return nil, xerrors.Errorf("manager repo: getlimit error: %v", err.Error())
	}

	mngrs := []domain.Manager{}
	for rows.Next() {
		mngr := domain.Manager{}
		err := rows.Scan(&mngr.ID, &mngr.AccID, &mngr.Name, &mngr.Surname, &mngr.Department)
		if err != nil {
			return nil, xerrors.Errorf("manager repo: getlimit error: %w", err.Error())
		}
		mngrs = append(mngrs, mngr)
	}

	return mngrs, nil
}
