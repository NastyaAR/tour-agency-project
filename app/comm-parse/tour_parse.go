package comm_parse

import (
	"app/domain"
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"golang.org/x/xerrors"
	"time"
)

const ParseTourString = `{
"id": %d,
"chillPlace": "%s",
"fromPlace": "%s",
"date": "%s",
"duration": %d,
"cost": %d,
"touristsNumber": %d,
"chillType": "%s"
}`

type TourJson struct {
	ID             int    `json:"id"`
	ChillPlace     string `json:"chillPlace"`
	FromPlace      string `json:"fromPlace"`
	Date           string `json:"date"`
	Duration       int    `json:"duration"`
	Cost           int    `json:"cost"`
	TouristsNumber int    `json:"touristsNumber"`
	ChillType      string `json:"chillType"`
}

func FromStringToTourJson(str string, lg *logrus.Logger) (TourJson, error) {
	var tour TourJson
	err := json.Unmarshal([]byte(str), &tour)
	if err != nil {
		lg.Warnf("bad parsing from terminal to tour json")
		return TourJson{}, xerrors.Errorf("tour: from string to json error: %v", err.Error())
	}
	return tour, nil
}

func (t *TourJson) ToDomainTour(lg *logrus.Logger) (domain.Tour, error) {
	var tour domain.Tour
	layout := "2006-01-02 15:04"
	timeLocal, err := time.Parse(layout, t.Date)
	if err != nil {
		lg.Warnf("bad tourjson to domain tour")
		return domain.Tour{}, xerrors.Errorf("tour: todomaintour error: %v", err.Error())
	}
	tour.ID = t.ID
	tour.ChillPlace = t.ChillPlace
	tour.FromPlace = t.FromPlace
	tour.Date = timeLocal
	tour.Duration = t.Duration
	tour.Cost = t.Cost
	tour.TouristsNumber = t.TouristsNumber
	tour.ChillType = t.ChillType

	return tour, nil
}

func FromDomainTourToJsonString(tour domain.Tour) string {
	tourJson := fmt.Sprintf(ParseTourString, 0, tour.ChillPlace, tour.FromPlace, tour.Date,
		tour.Duration, tour.Cost, tour.TouristsNumber, tour.ChillType)
	return tourJson
}
