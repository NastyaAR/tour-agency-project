package itests

import (
	"app/domain"
	"app/repo"
	"app/services"
	mock_domain "app/unit/mocks"
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/sirupsen/logrus"
	"go.uber.org/mock/gomock"
	"testing"
	"time"
)

func TestCreateRequest(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	pool, err := pgx.Connect(ctx, "user=admin password=pahst13_82 host=172.24.185.225 port=5432 dbname=test-db")
	defer pool.Close(context.Background())
	if err != nil {
		fmt.Printf("conncet error: %v", err.Error())
		t.Fail()
	}

	r := repo.NewPostgresRequestRepo(pool)
	p := services.NewPayAdapter()
	done := make(chan bool, 1)
	done <- true
	lg := logrus.Logger{}
	rs := services.CreateNewRequestService(r, p, 5*time.Second, done, 10, time.Second, &lg)

	req := domain.Request{
		ID:         0,
		TourID:     1,
		ClntID:     1,
		MngrID:     1,
		Status:     "принята",
		CreateTime: time.Date(2025, 10, 11, 10, 0, 0, 0, time.UTC),
		ModifyTime: time.Time{},
		Data:       "{}",
	}

	err = rs.Create(&req, &lg)
	if err != nil {
		t.Errorf("create request test error: %v", err.Error())
	}
}

func TestGetByStatus(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	pool, err := pgx.Connect(ctx, "user=admin password=pahst13_82 host=172.24.185.225 port=5432 dbname=test-db")
	defer pool.Close(context.Background())
	if err != nil {
		fmt.Printf("conncet error: %v", err.Error())
		t.Fail()
	}

	r := repo.NewPostgresRequestRepo(pool)
	p := services.NewPayAdapter()
	done := make(chan bool, 1)
	done <- true
	lg := logrus.Logger{}
	rs := services.CreateNewRequestService(r, p, 5*time.Second, done, 10, time.Second, &lg)

	reqs, err := rs.GetByStatus("принята", 0, 10, &lg)
	if len(reqs) != 3 {
		t.Errorf("get by id test failed: len doesnt same")
	}
}

func TestGetById(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	pool, err := pgx.Connect(ctx, "user=admin password=pahst13_82 host=172.24.185.225 port=5432 dbname=test-db")
	defer pool.Close(context.Background())
	if err != nil {
		fmt.Printf("conncet error: %v", err.Error())
		t.Fail()
	}

	expectReq := domain.Request{
		ID:         11,
		TourID:     1,
		ClntID:     1,
		MngrID:     1,
		Status:     "принята",
		CreateTime: time.Date(2025, 10, 11, 10, 0, 0, 0, time.UTC),
		ModifyTime: time.Time{},
		Data:       "{}",
	}

	r := repo.NewPostgresRequestRepo(pool)
	p := services.NewPayAdapter()
	done := make(chan bool, 1)
	done <- true
	lg := logrus.Logger{}
	rs := services.CreateNewRequestService(r, p, 5*time.Second, done, 10, time.Second, &lg)

	req, err := rs.GetById(11, &lg)
	if err != nil {
		t.Errorf("get by id request test failed: %v", err.Error())
	}
	if expectReq != req {
		t.Errorf("get by id request test failed: reqs doesnt same")
	}
}

func TestReject(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	pool, err := pgx.Connect(ctx, "user=admin password=pahst13_82 host=172.24.185.225 port=5432 dbname=test-db")
	defer pool.Close(context.Background())
	if err != nil {
		fmt.Printf("conncet error: %v", err.Error())
		t.Fail()
	}

	expectReq := domain.Request{
		ID:         11,
		TourID:     1,
		ClntID:     1,
		MngrID:     1,
		Status:     "отклонена",
		CreateTime: time.Date(2025, 10, 11, 10, 0, 0, 0, time.UTC),
		ModifyTime: time.Time{},
		Data:       "{}",
	}

	r := repo.NewPostgresRequestRepo(pool)
	p := services.NewPayAdapter()
	done := make(chan bool, 1)
	done <- true
	lg := logrus.Logger{}
	rs := services.CreateNewRequestService(r, p, 5*time.Second, done, 10, time.Second, &lg)

	err = rs.Reject(11, &lg)
	if err != nil {
		t.Errorf("reject test failed: %v", err.Error())
	}
	req, err := rs.GetById(11, &lg)
	if req != expectReq {
		t.Errorf("reject test failed: status doesnt update")
	}
}

func TestApprove(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	pool, err := pgx.Connect(ctx, "user=admin password=pahst13_82 host=172.24.185.225 port=5432 dbname=test-db")
	defer pool.Close(context.Background())
	if err != nil {
		fmt.Printf("conncet error: %v", err.Error())
		t.Fail()
	}

	expectReq := domain.Request{
		ID:         11,
		TourID:     1,
		ClntID:     1,
		MngrID:     1,
		Status:     "подтверждена",
		CreateTime: time.Date(2025, 10, 11, 10, 0, 0, 0, time.UTC),
		ModifyTime: time.Time{},
		Data:       "{}",
	}

	r := repo.NewPostgresRequestRepo(pool)
	p := services.NewPayAdapter()
	done := make(chan bool, 1)
	done <- true
	lg := logrus.Logger{}
	rs := services.CreateNewRequestService(r, p, 5*time.Second, done, 10, time.Second, &lg)

	err = rs.Approve(11, &lg)
	if err != nil {
		t.Errorf("approve test failed: %v", err.Error())
	}
	req, err := rs.GetById(11, &lg)
	if req != expectReq {
		t.Errorf("approve test failed: status doesnt update")
	}
}

func TestPay(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	pool, err := pgx.Connect(ctx, "user=admin password=pahst13_82 host=172.24.185.225 port=5432 dbname=test-db")
	defer pool.Close(context.Background())
	if err != nil {
		fmt.Printf("conncet error: %v", err.Error())
		t.Fail()
	}

	r := repo.NewPostgresRequestRepo(pool)
	p := services.NewPayAdapter()
	done := make(chan bool, 1)
	done <- true
	lg := logrus.Logger{}
	rs := services.CreateNewRequestService(r, p, 5*time.Second, done, 10, time.Second, &lg)

	rs.Approve(1, &lg)
	err = rs.Pay(1, &lg)
	if err != nil {
		t.Errorf("pay test failed: %v", err.Error())
	}

	events, err := r.GetNonSendEvents(ctx, 10, &lg)
	if len(events) != 1 {
		t.Errorf("pay test failed: number in req_outbox bad")
	}
}

func TestPaying(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockPayAdapter := mock_domain.NewMockIPayAdapter(ctrl)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	pool, err := pgx.Connect(ctx, "user=admin password=pahst13_82 host=172.24.185.225 port=5432 dbname=test-db")
	defer pool.Close(context.Background())
	if err != nil {
		fmt.Printf("conncet error: %v", err.Error())
		t.Fail()
	}

	r := repo.NewPostgresRequestRepo(pool)
	done := make(chan bool, 1)
	lg := logrus.Logger{}
	mockPayAdapter.EXPECT().SendPaymentRequest(gomock.Any(), &lg).Return(nil)
	_ = services.CreateNewRequestService(r, mockPayAdapter, 5*time.Second, done, 10, time.Second, &lg)

	time.Sleep(time.Second * 5)

	events, err := r.GetNonSendEvents(ctx, 10, &lg)
	if len(events) != 0 {
		t.Errorf("pay test failed: number in req_outbox bad")
	}
}
