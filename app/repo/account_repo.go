package repo

import (
	"app/domain"
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/sirupsen/logrus"
	"golang.org/x/xerrors"
)

type postgresAccountRepo struct {
	db *pgx.Conn
}

func NewPostgresAccountRepo(conn *pgx.Conn) domain.IAccountRepo {
	return &postgresAccountRepo{db: conn}
}

func (p *postgresAccountRepo) Add(c context.Context, acc *domain.Account, newAccDTO *domain.NewAccountDTO, lg *logrus.Logger) error {
	query := `insert into accounts (login, password) values ($1, $2)`

	lg.Warn("account repo add")
	_, err := p.db.Exec(c, query, acc.Login, acc.Password)
	if err != nil {
		lg.Errorf("bad insert account")
		return xerrors.Errorf("account repo: add error: %v", err.Error())
	}
	accTmp, err := p.GetByLogin(c, acc.Login, lg)

	newAccDTO.ID = accTmp.ID
	switch newAccDTO.Role {
	case "client":
		_, err = p.AddClient(c, newAccDTO, lg)
	case "manager":
		_, err = p.AddManager(c, newAccDTO, lg)
	default:
		lg.Warnf("unexpected role %s", newAccDTO.Role)
		return xerrors.New("account repo: Add error: unexpected role")
	}

	if err != nil {
		lg.Errorf("bad add user")
		return xerrors.Errorf("account repo: add error: %v", err.Error())
	}

	return nil
}

func (p *postgresAccountRepo) AddClient(c context.Context, dto *domain.NewAccountDTO, lg *logrus.Logger) (int, error) {
	var clnt domain.Client

	query := `insert into clients (acc_id, name, surname, mail, phone) values ($1, $2, $3, $4, $5)`
	_, err := p.db.Exec(c, query, dto.ID, dto.Name, dto.Surname, dto.Mail, dto.Phone)
	if err != nil {
		lg.Warnf("account repo: addClient error: %v", err.Error())
		return -1, xerrors.Errorf("account repo: addClient error: %v", err.Error())
	}

	return clnt.ID, nil
}

func (p *postgresAccountRepo) AddManager(c context.Context, dto *domain.NewAccountDTO, lg *logrus.Logger) (int, error) {
	var mngr domain.Manager

	query := `insert into managers (acc_id, name, surname, department) values ($1, $2, $3, $4)`
	_, err := p.db.Exec(c, query, dto.ID, dto.Name, dto.Surname, dto.Department)
	if err != nil {
		lg.Warnf("account repo: addManager error: %v", err.Error())
		return -1, xerrors.Errorf("account repo: addManager error: %v", err.Error())
	}

	return mngr.ID, nil
}

func (p *postgresAccountRepo) Delete(c context.Context, id int, lg *logrus.Logger) error {
	query := `delete from account where id = $1`

	lg.Warn("account repo delete")

	_, err := p.db.Exec(c, query, id)
	if err != nil {
		lg.Errorf("bad delete account")
		return xerrors.Errorf("account repo: delete error: %v", err.Error())
	}

	return nil
}

func (p *postgresAccountRepo) GetByLogin(c context.Context, login string, lg *logrus.Logger) (domain.Account, error) {
	query := `select * from accounts where login=$1`
	lg.Info("account repo: get by id")

	var acc domain.Account
	err := p.db.QueryRow(c, query, login).Scan(&acc.ID, &acc.Login, &acc.Password)
	if err == pgx.ErrNoRows {
		return domain.Account{}, nil
	}
	if err != nil {
		lg.Errorf("bad select by login")
		return domain.Account{}, xerrors.Errorf("account repo: getbylogin error: %v", err.Error())
	}

	return acc, nil
}

func (p *postgresAccountRepo) GetClientById(c context.Context, userID int, lg *logrus.Logger) (domain.Client, error) {
	query := `select * from clients where acc_id=$1`
	lg.Info("account repo: get client by acc id")

	var acc domain.Client
	err := p.db.QueryRow(c, query, userID).Scan(&acc.ID, &acc.AccID, &acc.Name, &acc.Surname,
		&acc.Mail, &acc.Phone)
	if err != nil {
		lg.Errorf("bad client select by acc id")
		return domain.Client{}, xerrors.Errorf("account repo: get client error: %v", err.Error())
	}

	return acc, nil
}

func (p *postgresAccountRepo) GetManagerById(c context.Context, userID int, lg *logrus.Logger) (domain.Manager, error) {
	query := `select * from managers where acc_id=$1`
	lg.Info("account repo: get manager by user id")

	var acc domain.Manager
	err := p.db.QueryRow(c, query, userID).Scan(&acc.ID, &acc.AccID, &acc.Name, &acc.Surname,
		&acc.Department)
	if err != nil {
		lg.Errorf("bad manager select by user id")
		return domain.Manager{}, xerrors.Errorf("account repo: get client error: %v", err.Error())
	}

	return acc, nil
}
