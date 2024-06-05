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

func Test_tourService_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockTourRepo := mock_domain.NewMockITourRepo(ctrl)

	tour := domain.Tour{
		ID:             0,
		ChillPlace:     "Москва",
		FromPlace:      "Санкт-Петербург",
		Date:           time.Date(2024, 10, 10, 10, 0, 0, 0, time.UTC),
		Duration:       10,
		Cost:           750000,
		TouristsNumber: 4,
		ChillType:      "городской",
	}
	mockTourRepo.EXPECT().Add(gomock.Any(), &tour, &logrus.Logger{}).Return(nil)

	type fields struct {
		tourRepository mock_domain.MockITourRepo
		timeout        time.Duration
	}
	fields_ := fields{
		tourRepository: *mockTourRepo,
		timeout:        time.Second,
	}
	type args struct {
		tour *domain.Tour
		lg   *logrus.Logger
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{"normal", fields_, args{&tour, &logrus.Logger{}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ts := services.NewTourService(mockTourRepo, time.Second)
			if err := ts.Create(tt.args.tour, tt.args.lg); (err != nil) != tt.wantErr {
				t.Errorf("tourService.Create() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_tourService_GetById(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockTourRepo := mock_domain.NewMockITourRepo(ctrl)

	tour := domain.Tour{
		ID:             0,
		ChillPlace:     "Москва",
		FromPlace:      "Санкт-Петербург",
		Date:           time.Date(2024, 10, 10, 10, 0, 0, 0, time.UTC),
		Duration:       10,
		Cost:           750000,
		TouristsNumber: 4,
		ChillType:      "городской",
	}
	mockTourRepo.EXPECT().GetById(gomock.Any(), 0, &logrus.Logger{}).Return(tour, nil)

	type fields struct {
		tourRepository mock_domain.MockITourRepo
		timeout        time.Duration
	}
	fields_ := fields{
		tourRepository: *mockTourRepo,
		timeout:        time.Second,
	}
	type args struct {
		id int
		lg *logrus.Logger
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    domain.Tour
		wantErr bool
	}{
		{"normal", fields_, args{0, &logrus.Logger{}}, tour, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ts := services.NewTourService(mockTourRepo, time.Second)
			got, err := ts.GetById(tt.args.id, tt.args.lg)
			if (err != nil) != tt.wantErr {
				t.Errorf("tourService.GetById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("tourService.GetById() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_tourService_Update(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockTourRepo := mock_domain.NewMockITourRepo(ctrl)

	tour := domain.Tour{
		ID:             0,
		ChillPlace:     "Москва",
		FromPlace:      "Санкт-Петербург",
		Date:           time.Date(2024, 10, 10, 10, 0, 0, 0, time.UTC),
		Duration:       10,
		Cost:           750000,
		TouristsNumber: 4,
		ChillType:      "городской",
	}
	mockTourRepo.EXPECT().Update(gomock.Any(), 0, &tour, &logrus.Logger{}).Return(nil)

	type fields struct {
		tourRepository mock_domain.MockITourRepo
		timeout        time.Duration
	}
	fields_ := fields{
		tourRepository: *mockTourRepo,
		timeout:        time.Second,
	}
	type args struct {
		id       int
		newState *domain.Tour
		lg       *logrus.Logger
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{"normal", fields_, args{0, &tour, &logrus.Logger{}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ts := services.NewTourService(mockTourRepo, time.Second)
			if err := ts.Update(tt.args.id, tt.args.newState, tt.args.lg); (err != nil) != tt.wantErr {
				t.Errorf("tourService.Update() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_tourService_SetSale(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockTourRepo := mock_domain.NewMockITourRepo(ctrl)
	sale := domain.Sale{
		ID:          0,
		Name:        "Sale",
		ExpiredTime: time.Date(2024, 10, 10, 10, 0, 0, 0, time.UTC),
		Percent:     5,
	}
	mockTourRepo.EXPECT().UpdateSale(gomock.Any(), 0, &sale, &logrus.Logger{}).Return(nil)

	type fields struct {
		tourRepository mock_domain.MockITourRepo
		timeout        time.Duration
	}
	fields_ := fields{
		tourRepository: *mockTourRepo,
		timeout:        time.Second,
	}
	type args struct {
		id      int
		newSale *domain.Sale
		lg      *logrus.Logger
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{"normal", fields_, args{0, &sale, &logrus.Logger{}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ts := services.NewTourService(mockTourRepo, time.Second)
			if err := ts.SetSale(tt.args.id, tt.args.newSale, tt.args.lg); (err != nil) != tt.wantErr {
				t.Errorf("tourService.SetSale() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_tourService_Delete(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockTourRepo := mock_domain.NewMockITourRepo(ctrl)
	mockTourRepo.EXPECT().Delete(gomock.Any(), 0, &logrus.Logger{}).Return(nil)

	type fields struct {
		tourRepository mock_domain.MockITourRepo
		timeout        time.Duration
	}
	fields_ := fields{
		tourRepository: *mockTourRepo,
		timeout:        time.Second,
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
		{"normal", fields_, args{0, &logrus.Logger{}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ts := services.NewTourService(mockTourRepo, time.Second)
			if err := ts.Delete(tt.args.id, tt.args.lg); (err != nil) != tt.wantErr {
				t.Errorf("tourService.Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_tourService_GetByCriteria(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockTourRepo := mock_domain.NewMockITourRepo(ctrl)
	tour := domain.Tour{
		ID:             0,
		ChillPlace:     "Москва",
		FromPlace:      "Санкт-Петербург",
		Date:           time.Date(2024, 10, 10, 10, 0, 0, 0, time.UTC),
		Duration:       10,
		Cost:           750000,
		TouristsNumber: 4,
		ChillType:      "городской",
	}
	tours := make([]domain.Tour, 0)
	tours = append(tours, tour)

	mockTourRepo.EXPECT().GetByCriteria(gomock.Any(), 10, 10, &tour, &logrus.Logger{}).Return(tours, nil)

	type fields struct {
		tourRepository mock_domain.MockITourRepo
		timeout        time.Duration
	}
	fields_ := fields{
		tourRepository: *mockTourRepo,
		timeout:        time.Second,
	}
	type args struct {
		criteria *domain.Tour
		offset   int
		limit    int
		lg       *logrus.Logger
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []domain.Tour
		wantErr bool
	}{
		{"normal", fields_, args{&tour, 10, 10, &logrus.Logger{}}, tours, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ts := services.NewTourService(mockTourRepo, time.Second)
			got, err := ts.GetByCriteria(tt.args.criteria, tt.args.offset, tt.args.limit, tt.args.lg)
			if (err != nil) != tt.wantErr {
				t.Errorf("tourService.GetByCriteria() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("tourService.GetByCriteria() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_tourService_GetTours(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockTourRepo := mock_domain.NewMockITourRepo(ctrl)
	tour := domain.Tour{
		ID:             0,
		ChillPlace:     "Москва",
		FromPlace:      "Санкт-Петербург",
		Date:           time.Date(2024, 10, 10, 10, 0, 0, 0, time.UTC),
		Duration:       10,
		Cost:           750000,
		TouristsNumber: 4,
		ChillType:      "городской",
	}
	tours := make([]domain.Tour, 0)
	tours = append(tours, tour)
	mockTourRepo.EXPECT().GetLimit(gomock.Any(), 10, 10, &logrus.Logger{}).Return(tours, nil)
	type fields struct {
		tourRepository mock_domain.MockITourRepo
		timeout        time.Duration
	}
	fields_ := fields{
		tourRepository: *mockTourRepo,
		timeout:        time.Second,
	}
	type args struct {
		offset int
		limit  int
		lg     *logrus.Logger
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []domain.Tour
		wantErr bool
	}{
		{"normal", fields_, args{10, 10, &logrus.Logger{}}, tours, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ts := services.NewTourService(mockTourRepo, time.Second)
			got, err := ts.GetTours(tt.args.offset, tt.args.limit, tt.args.lg)
			if (err != nil) != tt.wantErr {
				t.Errorf("tourService.GetTours() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("tourService.GetTours() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_tourService_GetHotTours(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockTourRepo := mock_domain.NewMockITourRepo(ctrl)
	tour := domain.HotTourDto{
		ID:             0,
		ChillPlace:     "Москва",
		FromPlace:      "Санкт-Петербург",
		Date:           time.Date(2024, 10, 10, 10, 0, 0, 0, time.UTC),
		Duration:       10,
		Cost:           750000,
		TouristsNumber: 4,
		ChillType:      "городской",
		SaleName:       "скидка",
		SalePercemt:    10,
	}
	tours := make([]domain.HotTourDto, 0)
	tours = append(tours, tour)
	mockTourRepo.EXPECT().GetHotTours(gomock.Any(), 10, 10, &logrus.Logger{}).Return(tours, nil)
	type fields struct {
		tourRepository mock_domain.MockITourRepo
		timeout        time.Duration
	}
	fields_ := fields{
		tourRepository: *mockTourRepo,
		timeout:        time.Second,
	}
	type args struct {
		offset int
		limit  int
		lg     *logrus.Logger
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []domain.HotTourDto
		wantErr bool
	}{
		{"normal", fields_, args{10, 10, &logrus.Logger{}}, tours, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ts := services.NewTourService(mockTourRepo, time.Second)
			got, err := ts.GetHotTours(tt.args.offset, tt.args.limit, tt.args.lg)
			if (err != nil) != tt.wantErr {
				t.Errorf("tourService.GetHotTours() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("tourService.GetHotTours() = %v, want %v", got, tt.want)
			}
		})
	}
}
