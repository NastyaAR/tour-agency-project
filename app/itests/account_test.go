package itests

import (
	"app/domain"
	"app/repo"
	"app/services"
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/sirupsen/logrus"
	"testing"
	"time"
)

func TestCreateAccountManager(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	pool, err := pgx.Connect(ctx, "user=admin password=pahst13_82 host=172.24.185.225 port=5432 dbname=test-db")
	defer pool.Close(ctx)
	if err != nil {
		fmt.Printf("conncet error: %v", err.Error())
		t.Fail()
	}

	ar := repo.NewPostgresAccountRepo(pool)
	as := services.CreateNewAccountService(ar, time.Second*5)
	lg := logrus.Logger{}

	acc := domain.Account{
		ID:       0,
		Login:    "ivanov_login_work",
		Password: "mypass",
	}
	accDTO := domain.NewAccountDTO{
		ID:         0,
		Name:       "Иван",
		Surname:    "Иванов",
		Mail:       "",
		Phone:      "",
		Department: "Москва Северная",
		Role:       "manager",
	}

	err = as.Register(&acc, &accDTO, &lg)
	if err != nil {
		t.Errorf("create client account test failed: %v", err.Error())
	}
}

func TestCreateAccountClient(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	pool, err := pgx.Connect(ctx, "user=admin password=pahst13_82 host=172.24.185.225 port=5432 dbname=test-db")
	defer pool.Close(ctx)
	if err != nil {
		fmt.Printf("conncet error: %v", err.Error())
		t.Fail()
	}

	ar := repo.NewPostgresAccountRepo(pool)
	as := services.CreateNewAccountService(ar, time.Second*5)
	lg := logrus.Logger{}

	acc := domain.Account{
		ID:       0,
		Login:    "ivanov_login_client",
		Password: "mypass",
	}
	accDTO := domain.NewAccountDTO{
		ID:         0,
		Name:       "Иван",
		Surname:    "Иванов",
		Mail:       "ivanov@mail.ru",
		Phone:      "81234567890",
		Department: "",
		Role:       "client",
	}

	err = as.Register(&acc, &accDTO, &lg)
	if err != nil {
		t.Errorf("create client account test failed: %v", err.Error())
	}
}

func TestGetByLogin(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	pool, err := pgx.Connect(ctx, "user=admin password=pahst13_82 host=172.24.185.225 port=5432 dbname=test-db")
	defer pool.Close(ctx)
	if err != nil {
		fmt.Printf("conncet error: %v", err.Error())
		t.Fail()
	}

	ar := repo.NewPostgresAccountRepo(pool)
	as := services.CreateNewAccountService(ar, time.Second*5)
	lg := logrus.Logger{}

	acc, err := as.GetByLogin("ivanov_login_work", &lg)
	fmt.Println(acc)
	if err != nil {
		t.Errorf("get by login test failed: %v", err.Error())
	}
}
