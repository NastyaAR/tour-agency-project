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

func TestGetByNameSurnameClient(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	pool, err := pgx.Connect(ctx, "user=admin password=pahst13_82 host=172.24.185.225 port=5432 dbname=test-db")
	defer pool.Close(ctx)
	if err != nil {
		fmt.Printf("conncet error: %v", err.Error())
		t.Fail()
	}

	cr := repo.NewPostgresClientRepo(pool)
	cs := services.CreateNewClientService(cr, 5*time.Second)
	lg := logrus.Logger{}

	clnts, err := cs.GetByNameSurname("Иван", "Иванов", &lg)
	if err != nil {
		t.Errorf("get client by name and surname error: %v", err.Error())
	}
	if len(clnts) != 1 {
		t.Errorf("get client by name and surname error: not enough")
	}
}

func TestGetByPhone(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	pool, err := pgx.Connect(ctx, "user=admin password=pahst13_82 host=172.24.185.225 port=5432 dbname=test-db")
	defer pool.Close(ctx)
	if err != nil {
		fmt.Printf("conncet error: %v", err.Error())
		t.Fail()
	}

	cr := repo.NewPostgresClientRepo(pool)
	cs := services.CreateNewClientService(cr, 5*time.Second)
	lg := logrus.Logger{}

	clnts, err := cs.GetByPhone("81234567890", &lg)
	if err != nil {
		t.Errorf("get client by name and surname error: %v", err.Error())
	}
	if len(clnts) != 1 {
		t.Errorf("get client by name and surname error: not enough")
	}
}

func TestDeleteClient(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	pool, err := pgx.Connect(ctx, "user=admin password=pahst13_82 host=172.24.185.225 port=5432 dbname=test-db")
	defer pool.Close(ctx)
	if err != nil {
		fmt.Printf("conncet error: %v", err.Error())
		t.Fail()
	}

	cr := repo.NewPostgresClientRepo(pool)
	cs := services.CreateNewClientService(cr, 5*time.Second)
	lg := logrus.Logger{}

	err = cs.Delete(6, &lg)
	if err != nil {
		t.Errorf("get client by name and surname error: %v", err.Error())
	}
}
