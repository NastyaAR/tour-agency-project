package comm_parse

import (
	"app/domain"
	"encoding/json"
	"github.com/sirupsen/logrus"
	"golang.org/x/xerrors"
)

const ParseAccountString = `{
"id": %d,
"login": "%s",
"password": "%s"
}`

const ParseClientString = `{
"id": %d,
"name": "%s",
"surname": "%s",
"mail": "%s",
"phone": "%s"
}`

const ParseManagerString = `{
"id": %d,
"name": "%s",
"surname": "%s",
"department": "%s"
}`

type AccountJson struct {
	ID       int    `json:"id"`
	Login    string `json:"login"`
	Password string `json:"password"`
}

type NewAccountDTOJson struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Surname    string `json:"surname"`
	Mail       string `json:"mail"`
	Phone      string `json:"phone"`
	Department string `json:"department"`
	Role       string `json:"role"`
}

func FromStringToAccountJson(str string, lg *logrus.Logger) (AccountJson, error) {
	var acc AccountJson
	err := json.Unmarshal([]byte(str), &acc)
	if err != nil {
		lg.Warnf("bad parsing from terminal to acc json")
		return AccountJson{}, xerrors.Errorf("account: from string to json error: %v", err.Error())
	}
	return acc, nil
}

func FromStringToUserJson(str string, lg *logrus.Logger) (NewAccountDTOJson, error) {
	var acc NewAccountDTOJson
	err := json.Unmarshal([]byte(str), &acc)
	if err != nil {
		lg.Warnf("bad parsing from terminal to acc json")
		return NewAccountDTOJson{}, xerrors.Errorf("user: from string to json error: %v", err.Error())
	}
	return acc, nil
}

func (s *AccountJson) ToDomainAccount(lg *logrus.Logger) domain.Account {
	var acc domain.Account
	acc.ID = s.ID
	acc.Login = s.Login
	acc.Password = s.Password
	return acc
}

//func (ns *NewAccountDTOJson) ToDomainClient(lg *logrus.Logger) domain.Client {
//	var clnt domain.Client
//	clnt.ID = ns.ID
//	clnt.Name = ns.Name
//	clnt.Surname = ns.Surname
//	clnt.Mail = ns.Mail
//	clnt.Phone = ns.Phone
//	return clnt
//}
//
//func (ns *NewAccountDTOJson) ToDomainManager(lg *logrus.Logger) domain.Manager {
//	var mngr domain.Manager
//	mngr.ID = ns.ID
//	mngr.Name = ns.Name
//	mngr.Surname = ns.Surname
//	mngr.Department = ns.Department
//	return mngr
//}

func (ns *NewAccountDTOJson) ToDtoClientAccount(lg *logrus.Logger) domain.NewAccountDTO {
	var dto domain.NewAccountDTO
	dto.ID = ns.ID
	dto.Name = ns.Name
	dto.Surname = ns.Surname
	dto.Mail = ns.Mail
	dto.Phone = ns.Phone
	dto.Role = "client"
	return dto
}

func (ns *NewAccountDTOJson) ToDtoManagerAccount(lg *logrus.Logger) domain.NewAccountDTO {
	var dto domain.NewAccountDTO
	dto.ID = ns.ID
	dto.Name = ns.Name
	dto.Surname = ns.Surname
	dto.Department = ns.Department
	dto.Role = "manager"
	return dto
}

func ToDomainClient(newAcc *domain.NewAccountDTO) domain.Client {
	var clnt domain.Client
	clnt.ID = newAcc.ID
	clnt.Name = newAcc.Name
	clnt.Surname = newAcc.Surname
	clnt.Mail = newAcc.Mail
	clnt.Phone = newAcc.Phone
	return clnt
}

func ToDomainManager(newAcc *domain.NewAccountDTO) domain.Manager {
	var mngr domain.Manager
	mngr.ID = newAcc.ID
	mngr.Name = newAcc.Name
	mngr.Surname = newAcc.Surname
	mngr.Department = newAcc.Department
	return mngr
}
