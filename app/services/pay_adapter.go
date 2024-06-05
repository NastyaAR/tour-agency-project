package services

import (
	"app/domain"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"golang.org/x/xerrors"
	"net/http"
)

type payAdapter struct{}

func NewPayAdapter() domain.IPayAdapter {
	return &payAdapter{}
}

func (p *payAdapter) SendPaymentRequest(pay *domain.PayEvent, lg *logrus.Logger) error {
	lg.Info("send payment request")
	requestBody := map[string]interface{}{
		"sum": pay.Sum,
	}
	requestBodyBytes, err := json.Marshal(requestBody)
	if err != nil {
		lg.Error("payment error: marshal")
		return xerrors.Errorf("send payment request error: %v", err.Error())
	}

	httpClient := &http.Client{}
	request, err := http.NewRequest("POST", "http://127.0.0.1:8080/payment", bytes.NewBuffer(requestBodyBytes))
	if err != nil {
		lg.Error("payment error: new http request: %v", err.Error())
		return xerrors.Errorf("send payment error: %v", err.Error())
	}

	request.Header.Set("Content-Type", "application/json")

	response, err := httpClient.Do(request)
	if err != nil {
		lg.Errorf("doing http request error: %v", err.Error())
		return xerrors.Errorf("send payment request: %v", err.Error())
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("получен неправильный код ответа: %d", response.StatusCode)
	}

	return nil
}
