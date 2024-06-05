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

func Test_clientService_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockClientRepo := mock_domain.NewMockIClientRepo(ctrl)

	clnt := domain.Client{
		ID:      0,
		Name:    "Иван",
		Surname: "Иванов",
		Mail:    "ivanov@mail.ru",
		Phone:   "89237317273",
	}

	mockClientRepo.EXPECT().Add(gomock.Any(), &clnt, &logrus.Logger{}).Return(clnt, nil)

	type fields struct {
		clientRepo mock_domain.MockIClientRepo
		timeout    time.Duration
	}
	fields_ := fields{
		clientRepo: *mockClientRepo,
		timeout:    time.Second,
	}
	type args struct {
		clnt *domain.Client
		lg   *logrus.Logger
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    domain.Client
		wantErr bool
	}{
		{"normal", fields_, args{&clnt, &logrus.Logger{}}, clnt, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cs := services.CreateNewClientService(mockClientRepo, time.Second)
			got, err := cs.Create(tt.args.clnt, tt.args.lg)
			if (err != nil) != tt.wantErr {
				t.Errorf("clientService.Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("clientService.Create() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_clientService_Delete(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockClientRepo := mock_domain.NewMockIClientRepo(ctrl)
	mockClientRepo.EXPECT().Delete(gomock.Any(), 1, &logrus.Logger{}).Return(nil)
	type fields struct {
		clientRepo mock_domain.MockIClientRepo
		timeout    time.Duration
	}
	fields_ := fields{
		clientRepo: *mockClientRepo,
		timeout:    time.Second,
	}
	type args struct {
		id int
		lg *logrus.Logger
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{"normal", fields_, args{1, &logrus.Logger{}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cs := services.CreateNewClientService(mockClientRepo, time.Second)
			if err := cs.Delete(tt.args.id, tt.args.lg); (err != nil) != tt.wantErr {
				t.Errorf("clientService.Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_clientService_GetByNameSurname(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockClientRepo := mock_domain.NewMockIClientRepo(ctrl)
	clnt := domain.Client{
		ID:      0,
		Name:    "Иван",
		Surname: "Иванов",
		Mail:    "ivanov@mail.ru",
		Phone:   "89237317273",
	}
	clnts := make([]domain.Client, 0)
	clnts = append(clnts, clnt)

	mockClientRepo.EXPECT().GetByNameSurname(gomock.Any(), "Иван", "Иванов", &logrus.Logger{}).Return(clnts, nil)
	type fields struct {
		clientRepo mock_domain.MockIClientRepo
		timeout    time.Duration
	}
	fields_ := fields{
		clientRepo: *mockClientRepo,
		timeout:    time.Second,
	}
	type args struct {
		name    string
		surname string
		lg      *logrus.Logger
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []domain.Client
		wantErr bool
	}{
		{"normal", fields_, args{"Иван", "Иванов", &logrus.Logger{}}, clnts, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cs := services.CreateNewClientService(mockClientRepo, time.Second)
			got, err := cs.GetByNameSurname(tt.args.name, tt.args.surname, tt.args.lg)
			if (err != nil) != tt.wantErr {
				t.Errorf("clientService.GetByNameSurname() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("clientService.GetByNameSurname() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_clientService_GetByPhone(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockClientRepo := mock_domain.NewMockIClientRepo(ctrl)
	clnt := domain.Client{
		ID:      0,
		Name:    "Иван",
		Surname: "Иванов",
		Mail:    "ivanov@mail.ru",
		Phone:   "89237317273",
	}
	clnts := make([]domain.Client, 0)
	clnts = append(clnts, clnt)

	mockClientRepo.EXPECT().GetByPhone(gomock.Any(), "89237317273", &logrus.Logger{}).Return(clnts, nil)
	type fields struct {
		clientRepo mock_domain.MockIClientRepo
		timeout    time.Duration
	}
	fields_ := fields{
		clientRepo: *mockClientRepo,
		timeout:    time.Second,
	}
	type args struct {
		phone string
		lg    *logrus.Logger
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []domain.Client
		wantErr bool
	}{
		{"normal", fields_, args{"89237317273", &logrus.Logger{}}, clnts, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cs := services.CreateNewClientService(mockClientRepo, time.Second)
			got, err := cs.GetByPhone(tt.args.phone, tt.args.lg)
			if (err != nil) != tt.wantErr {
				t.Errorf("clientService.GetByPhone() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("clientService.GetByPhone() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_clientService_Update(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockClientRepo := mock_domain.NewMockIClientRepo(ctrl)
	clnt := domain.Client{
		ID:      0,
		Name:    "Иван",
		Surname: "Иванов",
		Mail:    "ivanov@mail.ru",
		Phone:   "89237317273",
	}

	mockClientRepo.EXPECT().Update(gomock.Any(), 0, &clnt, &logrus.Logger{}).Return(nil)
	type fields struct {
		clientRepo mock_domain.MockIClientRepo
		timeout    time.Duration
	}
	fields_ := fields{
		clientRepo: *mockClientRepo,
		timeout:    time.Second,
	}
	type args struct {
		id       int
		newState *domain.Client
		lg       *logrus.Logger
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{"normal", fields_, args{0, &clnt, &logrus.Logger{}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cs := services.CreateNewClientService(mockClientRepo, time.Second)
			if err := cs.Update(tt.args.id, tt.args.newState, tt.args.lg); (err != nil) != tt.wantErr {
				t.Errorf("clientService.Update() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_clientService_GetById(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockClientRepo := mock_domain.NewMockIClientRepo(ctrl)
	clnt := domain.Client{
		ID:      0,
		Name:    "Иван",
		Surname: "Иванов",
		Mail:    "ivanov@mail.ru",
		Phone:   "89237317273",
	}

	mockClientRepo.EXPECT().GetById(gomock.Any(), 0, &logrus.Logger{}).Return(clnt, nil)
	type fields struct {
		clientRepo mock_domain.MockIClientRepo
		timeout    time.Duration
	}
	fields_ := fields{
		clientRepo: *mockClientRepo,
		timeout:    time.Second,
	}
	type args struct {
		id int
		lg *logrus.Logger
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    domain.Client
		wantErr bool
	}{
		{"normal", fields_, args{0, &logrus.Logger{}}, clnt, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cs := services.CreateNewClientService(mockClientRepo, time.Second)
			got, err := cs.GetById(tt.args.id, tt.args.lg)
			if (err != nil) != tt.wantErr {
				t.Errorf("clientService.GetById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("clientService.GetById() = %v, want %v", got, tt.want)
			}
		})
	}
}
