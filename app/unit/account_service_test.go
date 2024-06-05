package unit

import (
	"app/domain"
	"app/services"
	mock_domain "app/unit/mocks"
	"go.uber.org/mock/gomock"
	"reflect"
	"testing"
	"time"

	"github.com/sirupsen/logrus"
)

func Test_accountService_Register(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockAccountRepo := mock_domain.NewMockIAccountRepo(ctrl)
	acc := domain.Account{
		ID:       0,
		Login:    "ivanov_login_work",
		Password: "mypass",
	}
	cAccDTO := domain.NewAccountDTO{
		ID:         0,
		Name:       "Иван",
		Surname:    "Иванов",
		Mail:       "ivanov@mail.ru",
		Phone:      "81234567890",
		Department: "",
		Role:       "client",
	}
	mAccDTO := domain.NewAccountDTO{
		ID:         0,
		Name:       "Иван",
		Surname:    "Иванов",
		Mail:       "ivanov@mail.ru",
		Phone:      "81234567890",
		Department: "",
		Role:       "client",
	}

	mockAccountRepo.EXPECT().GetByLogin(gomock.Any(), acc.Login, &logrus.Logger{}).Return(domain.Account{}, nil).Times(2)
	mockAccountRepo.EXPECT().Add(gomock.Any(), &acc, &cAccDTO, &logrus.Logger{}).Return(nil)
	mockAccountRepo.EXPECT().Add(gomock.Any(), &acc, &mAccDTO, &logrus.Logger{}).Return(nil)

	type fields struct {
		accountRepo mock_domain.MockIAccountRepo
		timeout     time.Duration
	}
	fields_ := fields{
		accountRepo: *mockAccountRepo,
		timeout:     time.Second,
	}
	type args struct {
		acc       *domain.Account
		newAccDTO *domain.NewAccountDTO
		lg        *logrus.Logger
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{"normal client", fields_, args{&acc, &cAccDTO, &logrus.Logger{}}, false},
		{"normal manager", fields_, args{&acc, &mAccDTO, &logrus.Logger{}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			as := services.CreateNewAccountService(mockAccountRepo, time.Second)
			if err := as.Register(tt.args.acc, tt.args.newAccDTO, tt.args.lg); (err != nil) != tt.wantErr {
				t.Errorf("accountService.Register() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_accountService_GetByLogin(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockAccountRepo := mock_domain.NewMockIAccountRepo(ctrl)
	acc := domain.Account{
		ID:       0,
		Login:    "ivanov_login_work",
		Password: "mypass",
	}
	mockAccountRepo.EXPECT().GetByLogin(gomock.Any(), acc.Login, &logrus.Logger{}).Return(acc, nil)
	type fields struct {
		accountRepo mock_domain.MockIAccountRepo
		timeout     time.Duration
	}
	fields_ := fields{
		accountRepo: *mockAccountRepo,
		timeout:     time.Second,
	}
	type args struct {
		login string
		lg    *logrus.Logger
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    domain.Account
		wantErr bool
	}{
		{"normal", fields_, args{acc.Login, &logrus.Logger{}}, acc, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			as := services.CreateNewAccountService(mockAccountRepo, time.Second)
			got, err := as.GetByLogin(tt.args.login, tt.args.lg)
			if (err != nil) != tt.wantErr {
				t.Errorf("accountService.GetByLogin() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("accountService.GetByLogin() = %v, want %v", got, tt.want)
			}
		})
	}
}
