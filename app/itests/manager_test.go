package itests

import (
	"app/repo"
	"app/services"
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/sirupsen/logrus"
	"testing"
	"time"
)

func TestGetByNameSurnameManager(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	pool, err := pgx.Connect(ctx, "user=admin password=pahst13_82 host=172.24.185.225 port=5432 dbname=test-db")
	defer pool.Close(ctx)
	if err != nil {
		fmt.Printf("conncet error: %v", err.Error())
		t.Fail()
	}

	ar := repo.NewPostgresManagerRepo(pool)
	as := services.CreateNewManagerService(ar, time.Second*5)
	lg := logrus.Logger{}

	clnts, err := as.GetByNameSurname("Иван", "Иванов", &lg)
	if err != nil {
		t.Errorf("get client by name and surname error: %v", err.Error())
	}
	if len(clnts) != 1 {
		t.Errorf("get client by name and surname error: not enough")
	}
}

func TestGetByDepartment(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	pool, err := pgx.Connect(ctx, "user=admin password=pahst13_82 host=172.24.185.225 port=5432 dbname=test-db")
	defer pool.Close(ctx)
	if err != nil {
		fmt.Printf("conncet error: %v", err.Error())
		t.Fail()
	}

	ar := repo.NewPostgresManagerRepo(pool)
	as := services.CreateNewManagerService(ar, time.Second*5)
	lg := logrus.Logger{}

	clnts, err := as.GetByDepartment("Москва Северная", &lg)
	if err != nil {
		t.Errorf("get client by name and surname error: %v", err.Error())
	}
	if len(clnts) != 1 {
		t.Errorf("get client by name and surname error: not enough")
	}
}

func TestDeleteManager(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	pool, err := pgx.Connect(ctx, "user=admin password=pahst13_82 host=172.24.185.225 port=5432 dbname=test-db")
	defer pool.Close(ctx)
	if err != nil {
		fmt.Printf("conncet error: %v", err.Error())
		t.Fail()
	}

	ar := repo.NewPostgresManagerRepo(pool)
	as := services.CreateNewManagerService(ar, time.Second*5)
	lg := logrus.Logger{}

	err = as.Delete(6, &lg)
	if err != nil {
		t.Errorf("delete manager test failed: %v", err.Error())
	}
}
