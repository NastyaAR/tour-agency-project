package itests

import (
	"app/domain"
	"app/repo"
	"app/services"
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/sirupsen/logrus"
	"os"
	"testing"
	"time"
)

func TestCreateSale(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	pool, err := pgx.Connect(ctx, "user=admin password=pahst13_82 host=172.24.185.225 port=5432 dbname=test-db")
	defer pool.Close(ctx)
	if err != nil {
		fmt.Printf("conncet error: %v", err.Error())
		t.Fail()
	}

	sr := repo.NewPostgresSaleRepo(pool)
	s := services.NewSaleService(sr, 5*time.Second)
	lg := logrus.Logger{}

	sale := domain.Sale{
		ID:          0,
		Name:        "Осенняя скидка",
		ExpiredTime: time.Date(2025, 9, 22, 0, 0, 0, 0, time.UTC),
		Percent:     40,
	}
	err = s.Create(&sale, &lg)
	if err != nil {
		t.Errorf("create sale test failed: %v", err.Error())
	}
}

func TestGetByIdSale(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	pool, err := pgx.Connect(ctx, "user=admin password=pahst13_82 host=172.24.185.225 port=5432 dbname=test-db")
	defer pool.Close(ctx)
	if err != nil {
		fmt.Printf("conncet error: %v", err.Error())
		t.Fail()
	}

	expectSale := domain.Sale{
		ID:          11,
		Name:        "Осенняя скидка",
		ExpiredTime: time.Date(2025, 9, 22, 0, 0, 0, 0, time.UTC),
		Percent:     40,
	}

	sr := repo.NewPostgresSaleRepo(pool)
	s := services.NewSaleService(sr, 5*time.Second)
	lg := logrus.Logger{}

	sale, err := s.GetById(11, &lg)
	if err != nil {
		t.Errorf("test get by id sale error: %v", err.Error())
	}
	if sale != expectSale {
		t.Errorf("get by id test failed: doesnt same")
	}
}

func TestUpdateSale(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	pool, err := pgx.Connect(ctx, "user=admin password=pahst13_82 host=172.24.185.225 port=5432 dbname=test-db")
	defer pool.Close(ctx)
	if err != nil {
		fmt.Printf("conncet error: %v", err.Error())
		t.Fail()
	}

	sr := repo.NewPostgresSaleRepo(pool)
	s := services.NewSaleService(sr, 5*time.Second)
	lg := logrus.Logger{}

	newSale := domain.Sale{
		ID:          11,
		Name:        "Весенняя скидка",
		ExpiredTime: time.Date(2025, 01, 01, 0, 0, 0, 0, time.UTC),
		Percent:     30,
	}

	err = s.Update(11, &newSale, &lg)
	if err != nil {
		t.Errorf("update sale test failed: %v", err.Error())
	}

	sale, err := s.GetById(11, &lg)
	if sale.Percent != 30 {
		t.Errorf("update sale test failed: value doesnt change")
	}
	if sale != newSale {
		t.Errorf("update sale test failed: doesnt update")
	}
}

func TestDelete(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	pool, err := pgx.Connect(ctx, "user=admin password=pahst13_82 host=172.24.185.225 port=5432 dbname=test-db")
	defer pool.Close(ctx)
	if err != nil {
		fmt.Printf("conncet error: %v", err.Error())
		t.Fail()
	}

	sr := repo.NewPostgresSaleRepo(pool)
	s := services.NewSaleService(sr, 5*time.Second)
	lg := logrus.Logger{
		Out:          os.Stdout,
		Hooks:        nil,
		Formatter:    nil,
		ReportCaller: false,
		Level:        0,
		ExitFunc:     nil,
		BufferPool:   nil,
	}

	err = s.Delete(11, &lg)
	if err != nil {
		t.Errorf("delete sale test failed: %v", err.Error())
	}
}
