package unit

import (
	"app/domain"
	"app/services"
	mock_domain "app/unit/mocks"
	"go.uber.org/mock/gomock"
	"golang.org/x/xerrors"
	"reflect"
	"testing"
	"time"

	"github.com/sirupsen/logrus"
)

func Test_managerService_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockManagerRepo := mock_domain.NewMockIManagerRepo(ctrl)

	mngr := domain.Manager{
		ID:         0,
		Name:       "Иван",
		Surname:    "Иванов",
		Department: "Москва Северная",
	}

	mockManagerRepo.EXPECT().Add(gomock.Any(), &mngr, &logrus.Logger{}).Return(mngr, nil)

	type fields struct {
		managerRepo mock_domain.MockIManagerRepo
		timeout     time.Duration
	}
	fields_ := fields{
		managerRepo: *mockManagerRepo,
		timeout:     time.Second,
	}
	type args struct {
		mngr *domain.Manager
		lg   *logrus.Logger
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    domain.Manager
		wantErr bool
	}{
		{"normal", fields_, args{&mngr, &logrus.Logger{}}, mngr, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ms := services.CreateNewManagerService(mockManagerRepo, time.Second)
			got, err := ms.Create(tt.args.mngr, tt.args.lg)
			if (err != nil) != tt.wantErr {
				t.Errorf("managerService.Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("managerService.Create() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_managerService_Delete(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockManagerRepo := mock_domain.NewMockIManagerRepo(ctrl)
	mockManagerRepo.EXPECT().Delete(gomock.Any(), 1, &logrus.Logger{})
	type fields struct {
		managerRepo mock_domain.MockIManagerRepo
		timeout     time.Duration
	}
	fields_ := fields{
		managerRepo: *mockManagerRepo,
		timeout:     time.Second,
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
			ms := services.CreateNewManagerService(mockManagerRepo, time.Second)
			if err := ms.Delete(tt.args.id, tt.args.lg); (err != nil) != tt.wantErr {
				t.Errorf("managerService.Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_managerService_GetByNameSurname(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockManagerRepo := mock_domain.NewMockIManagerRepo(ctrl)
	mngr := domain.Manager{
		ID:         0,
		Name:       "Иван",
		Surname:    "Иванов",
		Department: "Москва Северная",
	}
	mngrs := make([]domain.Manager, 0)
	mngrs = append(mngrs, mngr)
	mockManagerRepo.EXPECT().GetByNameSurname(gomock.Any(), "Иван", "Иванов", &logrus.Logger{}).Return(mngrs, nil)
	type fields struct {
		managerRepo mock_domain.MockIManagerRepo
		timeout     time.Duration
	}
	fields_ := fields{
		managerRepo: *mockManagerRepo,
		timeout:     time.Second,
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
		want    []domain.Manager
		wantErr bool
	}{
		{"normal", fields_, args{"Иван", "Иванов", &logrus.Logger{}}, mngrs, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ms := services.CreateNewManagerService(mockManagerRepo, time.Second)
			got, err := ms.GetByNameSurname(tt.args.name, tt.args.surname, tt.args.lg)
			if (err != nil) != tt.wantErr {
				t.Errorf("managerService.GetByNameSurname() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("managerService.GetByNameSurname() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_managerService_GetByDepartment(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockManagerRepo := mock_domain.NewMockIManagerRepo(ctrl)
	mngr := domain.Manager{
		ID:         0,
		Name:       "Иван",
		Surname:    "Иванов",
		Department: "Москва Северная",
	}
	mngrs := make([]domain.Manager, 0)
	mngrs = append(mngrs, mngr)
	mockManagerRepo.EXPECT().GetByDepartment(gomock.Any(), "Москва Северная", &logrus.Logger{}).Return(mngrs, nil)
	type fields struct {
		managerRepo mock_domain.MockIManagerRepo
		timeout     time.Duration
	}
	fields_ := fields{
		managerRepo: *mockManagerRepo,
		timeout:     time.Second,
	}
	type args struct {
		department string
		lg         *logrus.Logger
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []domain.Manager
		wantErr bool
	}{
		{"normal", fields_, args{"Москва Северная", &logrus.Logger{}}, mngrs, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ms := services.CreateNewManagerService(mockManagerRepo, time.Second)
			got, err := ms.GetByDepartment(tt.args.department, tt.args.lg)
			if (err != nil) != tt.wantErr {
				t.Errorf("managerService.GetByDepartment() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("managerService.GetByDepartment() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_managerService_Update(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockManagerRepo := mock_domain.NewMockIManagerRepo(ctrl)
	mngr := domain.Manager{
		ID:         0,
		Name:       "Иван",
		Surname:    "Иванов",
		Department: "Москва Северная",
	}
	mockManagerRepo.EXPECT().Update(gomock.Any(), 1, &mngr, &logrus.Logger{}).Return(nil)
	type fields struct {
		managerRepo mock_domain.MockIManagerRepo
		timeout     time.Duration
	}
	fields_ := fields{
		managerRepo: *mockManagerRepo,
		timeout:     time.Second,
	}
	type args struct {
		id       int
		newState *domain.Manager
		lg       *logrus.Logger
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{"normal", fields_, args{1, &mngr, &logrus.Logger{}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ms := services.CreateNewManagerService(mockManagerRepo, time.Second)
			if err := ms.Update(tt.args.id, tt.args.newState, tt.args.lg); (err != nil) != tt.wantErr {
				t.Errorf("managerService.Update() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_managerService_GetStatisticsForManagers(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockManagerRepo := mock_domain.NewMockIManagerRepo(ctrl)
	from := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)
	to := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
	mngr := domain.Manager{
		ID:         1,
		Name:       "Иван",
		Surname:    "Иванов",
		Department: "Москва Северная",
	}
	mngrs := make([]domain.Manager, 0)
	mngrs = append(mngrs, mngr)
	req := domain.Request{
		ID:         1,
		TourID:     1,
		ClntID:     1,
		MngrID:     1,
		Status:     "оплачена",
		CreateTime: time.Now(),
		ModifyTime: time.Now(),
		Data:       "",
	}
	reqs := make([]domain.Request, 0)
	reqs = append(reqs, req)
	stat := domain.Statistics{
		ManagerAcc: mngr,
		Efficiency: 100,
		SaleSum:    100000,
	}
	stats := make([]domain.Statistics, 0)
	stats = append(stats, stat)
	mockManagerRepo.EXPECT().GetLimit(gomock.Any(), 10, 10, &logrus.Logger{}).Return(mngrs, nil)
	mockManagerRepo.EXPECT().GetById(gomock.Any(), 1, &logrus.Logger{}).Return(mngr, nil)
	mockManagerRepo.EXPECT().GetAllRequests(gomock.Any(), 1, &logrus.Logger{}).Return(1, nil)
	mockManagerRepo.EXPECT().GetNumberServedRequests(gomock.Any(), 1, &logrus.Logger{}).Return(1, nil)
	mockManagerRepo.EXPECT().GetSumOnPeriod(gomock.Any(), 1, from, to, &logrus.Logger{}).Return(100000, nil)
	mockManagerRepo.EXPECT().GetLimit(gomock.Any(), 1, 1, &logrus.Logger{}).Return(nil, xerrors.New("err"))

	type fields struct {
		managerRepo mock_domain.MockIManagerRepo
		timeout     time.Duration
	}
	fields_ := fields{
		managerRepo: *mockManagerRepo,
		timeout:     time.Second,
	}
	type args struct {
		from   time.Time
		to     time.Time
		offset int
		limit  int
		lg     *logrus.Logger
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []domain.Statistics
		wantErr bool
	}{
		{"normal", fields_, args{from, to, 10, 10, &logrus.Logger{}}, stats, false},
		{"bad get limit", fields_, args{from, to, 1, 1, &logrus.Logger{}}, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ms := services.CreateNewManagerService(mockManagerRepo, time.Second)
			got, err := ms.GetStatisticsForManagers(tt.args.from, tt.args.to, tt.args.offset, tt.args.limit, tt.args.lg)
			if (err != nil) != tt.wantErr {
				t.Errorf("managerService.GetStatisticsForManagers() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("managerService.GetStatisticsForManagers() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_managerService_GetStatisticsForManager(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockManagerRepo := mock_domain.NewMockIManagerRepo(ctrl)
	from := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)
	to := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
	mngr := domain.Manager{
		ID:         1,
		Name:       "Иван",
		Surname:    "Иванов",
		Department: "Москва Северная",
	}
	mngr2 := domain.Manager{
		ID:         3,
		Name:       "Иван",
		Surname:    "Иванов",
		Department: "Москва Северная",
	}
	mngr3 := domain.Manager{
		ID:         4,
		Name:       "Иван",
		Surname:    "Иванов",
		Department: "Москва Северная",
	}
	mngr4 := domain.Manager{
		ID:         5,
		Name:       "Иван",
		Surname:    "Иванов",
		Department: "Москва Северная",
	}
	mngr5 := domain.Manager{
		ID:         6,
		Name:       "Иван",
		Surname:    "Иванов",
		Department: "Москва Северная",
	}
	req := domain.Request{
		ID:         1,
		TourID:     1,
		ClntID:     1,
		MngrID:     1,
		Status:     "оплачена",
		CreateTime: time.Now(),
		ModifyTime: time.Now(),
		Data:       "",
	}
	reqs := make([]domain.Request, 0)
	reqs = append(reqs, req)
	stat := domain.Statistics{
		ManagerAcc: mngr,
		Efficiency: 100,
		SaleSum:    100000,
	}
	mockManagerRepo.EXPECT().GetById(gomock.Any(), 1, &logrus.Logger{}).Return(mngr, nil)
	mockManagerRepo.EXPECT().GetAllRequests(gomock.Any(), 1, &logrus.Logger{}).Return(1, nil)
	mockManagerRepo.EXPECT().GetNumberServedRequests(gomock.Any(), 1, &logrus.Logger{}).Return(1, nil)
	mockManagerRepo.EXPECT().GetSumOnPeriod(gomock.Any(), 1, from, to, &logrus.Logger{}).Return(100000, nil)
	mockManagerRepo.EXPECT().GetById(gomock.Any(), 2, &logrus.Logger{}).Return(domain.Manager{}, xerrors.New("err"))
	mockManagerRepo.EXPECT().GetById(gomock.Any(), 3, &logrus.Logger{}).Return(mngr2, nil)
	mockManagerRepo.EXPECT().GetAllRequests(gomock.Any(), 3, &logrus.Logger{}).Return(-1, xerrors.New("err"))
	mockManagerRepo.EXPECT().GetById(gomock.Any(), 4, &logrus.Logger{}).Return(mngr3, nil)
	mockManagerRepo.EXPECT().GetAllRequests(gomock.Any(), 4, &logrus.Logger{}).Return(1, nil)
	mockManagerRepo.EXPECT().GetNumberServedRequests(gomock.Any(), 4, &logrus.Logger{}).Return(-1, xerrors.New("err"))
	mockManagerRepo.EXPECT().GetById(gomock.Any(), 5, &logrus.Logger{}).Return(mngr4, nil)
	mockManagerRepo.EXPECT().GetAllRequests(gomock.Any(), 5, &logrus.Logger{}).Return(1, nil)
	mockManagerRepo.EXPECT().GetNumberServedRequests(gomock.Any(), 5, &logrus.Logger{}).Return(0, nil)
	mockManagerRepo.EXPECT().GetById(gomock.Any(), 6, &logrus.Logger{}).Return(mngr5, nil)
	mockManagerRepo.EXPECT().GetAllRequests(gomock.Any(), 6, &logrus.Logger{}).Return(1, nil)
	mockManagerRepo.EXPECT().GetNumberServedRequests(gomock.Any(), 6, &logrus.Logger{}).Return(1, nil)
	mockManagerRepo.EXPECT().GetSumOnPeriod(gomock.Any(), 6, from, to, &logrus.Logger{}).Return(-1, xerrors.New("err"))
	type fields struct {
		managerRepo mock_domain.MockIManagerRepo
		timeout     time.Duration
	}
	fields_ := fields{
		managerRepo: *mockManagerRepo,
		timeout:     time.Second,
	}
	type args struct {
		id   int
		from time.Time
		to   time.Time
		lg   *logrus.Logger
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    domain.Statistics
		wantErr bool
	}{
		{"normal", fields_, args{1, from, to, &logrus.Logger{}}, stat, false},
		{"bad get by id", fields_, args{2, from, to, &logrus.Logger{}}, domain.Statistics{}, true},
		{"bad get all requests", fields_, args{3, from, to, &logrus.Logger{}}, domain.Statistics{}, true},
		{"bad get number reqs", fields_, args{4, from, to, &logrus.Logger{}}, domain.Statistics{}, true},
		{"serv reqs = 0", fields_, args{5, from, to, &logrus.Logger{}}, domain.Statistics{}, true},
		{"bad get sum", fields_, args{6, from, to, &logrus.Logger{}}, domain.Statistics{}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ms := services.CreateNewManagerService(mockManagerRepo, time.Second)
			got, err := ms.GetStatisticsForManager(tt.args.id, tt.args.from, tt.args.to, tt.args.lg)
			if (err != nil) != tt.wantErr {
				t.Errorf("managerService.GetStatisticsForManager() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("managerService.GetStatisticsForManager() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_managerService_GetById(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockManagerRepo := mock_domain.NewMockIManagerRepo(ctrl)
	mngr := domain.Manager{
		ID:         0,
		Name:       "Иван",
		Surname:    "Иванов",
		Department: "Москва Северная",
	}
	mockManagerRepo.EXPECT().GetById(gomock.Any(), 1, &logrus.Logger{}).Return(mngr, nil)
	type fields struct {
		managerRepo mock_domain.MockIManagerRepo
		timeout     time.Duration
	}
	fields_ := fields{
		managerRepo: *mockManagerRepo,
		timeout:     time.Second,
	}
	type args struct {
		id int
		lg *logrus.Logger
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    domain.Manager
		wantErr bool
	}{
		{"normal", fields_, args{1, &logrus.Logger{}}, mngr, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ms := services.CreateNewManagerService(mockManagerRepo, time.Second)
			got, err := ms.GetById(tt.args.id, tt.args.lg)
			if (err != nil) != tt.wantErr {
				t.Errorf("managerService.GetById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("managerService.GetById() = %v, want %v", got, tt.want)
			}
		})
	}
}
