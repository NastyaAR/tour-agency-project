package comm_parse

import (
	"app/domain"
	"encoding/json"
	"github.com/sirupsen/logrus"
	"golang.org/x/xerrors"
	"time"
)

const ParseSaleString = `{
"id": %d,
"name": "%s",
"expired_time": "%s",
"percent": %d
}`

type SaleJson struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	ExpiredTime string `json:"expired_time"`
	Percent     int    `json:"percent"`
}

func FromStringToSaleJson(str string, lg *logrus.Logger) (SaleJson, error) {
	var sale SaleJson
	err := json.Unmarshal([]byte(str), &sale)
	if err != nil {
		lg.Warnf("bad parsing from terminal to sale json")
		return SaleJson{}, xerrors.Errorf("sale: from string to json error: %v", err.Error())
	}
	return sale, nil
}

func (s *SaleJson) ToDomainSale(lg *logrus.Logger) (domain.Sale, error) {
	var sale domain.Sale
	layout := "2006-01-02 15:04"
	timeLocal, err := time.Parse(layout, s.ExpiredTime)
	if err != nil {
		lg.Warnf("bad salejson to domain sale")
		return domain.Sale{}, xerrors.Errorf("tour: todomainsale error: %v", err.Error())
	}

	sale.ID = s.ID
	sale.Name = s.Name
	sale.ExpiredTime = timeLocal
	sale.Percent = s.Percent

	return sale, nil
}
