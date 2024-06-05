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

func TestCreateTour(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	pool, err := pgx.Connect(ctx, "user=admin password=pahst13_82 host=172.24.185.225 port=5432 dbname=test-db")
	defer pool.Close(ctx)
	if err != nil {
		fmt.Printf("conncet error: %v", err.Error())
		t.Fail()
	}

	tr := repo.NewPostgresTourRepo(pool)
	ts := services.NewTourService(tr, 5*time.Second)
	lg := logrus.Logger{}

	tour := domain.Tour{
		ID:             0,
		ChillPlace:     "Владивосток",
		FromPlace:      "Москва",
		Date:           time.Date(2025, 10, 11, 12, 30, 0, 0, time.UTC),
		Duration:       10,
		Cost:           100000,
		TouristsNumber: 2,
		ChillType:      "городской",
	}
	err = ts.Create(&tour, &lg)
	if err != nil {
		t.Errorf("create tour test failed: %v", err.Error())
	}
}

func TestCreateBadTour(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	pool, err := pgx.Connect(ctx, "user=admin password=pahst13_82 host=172.24.185.225 port=5432 dbname=test-db")
	defer pool.Close(ctx)
	if err != nil {
		fmt.Printf("conncet error: %v", err.Error())
		t.Fail()
	}

	tr := repo.NewPostgresTourRepo(pool)
	ts := services.NewTourService(tr, 5*time.Second)
	lg := logrus.Logger{}

	tour := domain.Tour{
		ID:             0,
		ChillPlace:     "Владивосток",
		FromPlace:      "Москва",
		Date:           time.Now(),
		Duration:       -10,
		Cost:           100000,
		TouristsNumber: 2,
		ChillType:      "городской",
	}
	err = ts.Create(&tour, &lg)
	if err == nil {
		t.Errorf("create bad tour test failed")
	}
}

func TestGetByIdTour(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	pool, err := pgx.Connect(ctx, "user=admin password=pahst13_82 host=172.24.185.225 port=5432 dbname=test-db")
	defer pool.Close(ctx)
	if err != nil {
		fmt.Printf("conncet error: %v", err.Error())
		t.Fail()
	}

	tr := repo.NewPostgresTourRepo(pool)
	ts := services.NewTourService(tr, 5*time.Second)
	lg := logrus.Logger{}

	tourExpect := domain.Tour{
		ID:             11,
		ChillPlace:     "Владивосток",
		FromPlace:      "Москва",
		Date:           time.Date(2025, 10, 11, 12, 30, 0, 0, time.UTC),
		Duration:       10,
		Cost:           100000,
		TouristsNumber: 2,
		ChillType:      "городской",
	}
	tour, err := ts.GetById(11, &lg)
	if err != nil {
		t.Errorf("get by id test failed: %v", err.Error())
	}
	if tour != tourExpect {
		t.Errorf("get by id test failed: data doesnt same")
	}
}

func TestDeleteTour(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	pool, err := pgx.Connect(ctx, "user=admin password=pahst13_82 host=172.24.185.225 port=5432 dbname=test-db")
	defer pool.Close(ctx)
	if err != nil {
		fmt.Printf("conncet error: %v", err.Error())
		t.Fail()
	}

	tr := repo.NewPostgresTourRepo(pool)
	ts := services.NewTourService(tr, 5*time.Second)
	lg := logrus.Logger{
		Out:          os.Stdout,
		Hooks:        nil,
		Formatter:    nil,
		ReportCaller: false,
		Level:        0,
		ExitFunc:     nil,
		BufferPool:   nil,
	}
	err = ts.Delete(11, &lg)
	if err != nil {
		t.Errorf("delete tour error")
	}
}

func TestGetTours(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	pool, err := pgx.Connect(ctx, "user=admin password=pahst13_82 host=172.24.185.225 port=5432 dbname=test-db")
	defer pool.Close(ctx)
	if err != nil {
		fmt.Printf("conncet error: %v", err.Error())
		t.Fail()
	}

	tr := repo.NewPostgresTourRepo(pool)
	ts := services.NewTourService(tr, 5*time.Second)
	lg := logrus.Logger{
		Out:          os.Stdout,
		Hooks:        nil,
		Formatter:    nil,
		ReportCaller: false,
		Level:        0,
		ExitFunc:     nil,
		BufferPool:   nil,
	}

	tours, err := ts.GetTours(0, 5, &lg)
	if err != nil {
		t.Errorf("gettours error: %v", err.Error())
	}
	if len(tours) != 5 {
		t.Errorf("gettours error: not enough tours")
	}
	fmt.Println(tours)
}
