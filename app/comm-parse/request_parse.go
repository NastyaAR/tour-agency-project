package comm_parse

import (
	"app/domain"
	"encoding/json"
	"github.com/sirupsen/logrus"
	"golang.org/x/xerrors"
	"time"
)

const ParseRequestString = `
"id": %d,
"tourID": %d,
"clntID": %d,
"mngrID": %d,
"status": %s,
"createTime": %s,
"modifyTime": %s,
"data": %d
`

type RequestJson struct {
	ID         int    `json:"id"`
	TourID     int    `json:"tourID"`
	ClntID     int    `json:"clntID"`
	MngrID     int    `json:"mngrID"`
	Status     string `json:"status"`
	CreateTime string `json:"createTime"`
	ModifyTime string `json:"modifyTime"`
	Data       string `json:"data"`
}

func FromStringToRequestJson(str string, lg *logrus.Logger) (RequestJson, error) {
	var req RequestJson
	err := json.Unmarshal([]byte(str), &req)
	if err != nil {
		lg.Warnf("bad parsing from terminal to request json")
		return RequestJson{}, xerrors.Errorf("request: from string to json error: %v", err.Error())
	}
	return req, nil
}

func (r *RequestJson) ToDomainRequest(lg *logrus.Logger) (domain.Request, error) {
	var req domain.Request
	layout := "2006-01-02 15:04"
	timeCreate, err := time.Parse(layout, r.CreateTime)
	if err != nil {
		lg.Warnf("bad requestjson to domain request")
		return domain.Request{}, xerrors.Errorf("request: todomainrequest error: %v", err.Error())
	}
	timeModify, err := time.Parse(layout, r.ModifyTime)
	if err != nil {
		lg.Warnf("bad requestjson to domain request")
		return domain.Request{}, xerrors.Errorf("request: todomainrequest error: %v", err.Error())
	}
	req.ID = r.ID
	req.TourID = r.TourID
	req.ClntID = r.ClntID
	req.MngrID = r.MngrID
	req.Status = r.Status
	req.CreateTime = timeCreate
	req.ModifyTime = timeModify
	req.Data = r.Data

	return req, nil
}
