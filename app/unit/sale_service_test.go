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

func Test_saleService_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockSaleRepo := mock_domain.NewMockISaleRepo(ctrl)

	sale := domain.Sale{
		ID:          0,
		Name:        "весенняя скидка",
		ExpiredTime: time.Date(2025, 10, 10, 0, 0, 0, 0, time.UTC),
		Percent:     0,
	}
	mockSaleRepo.EXPECT().Add(gomock.Any(), &sale, &logrus.Logger{}).Return(nil)

	type fields struct {
		saleRepo mock_domain.MockISaleRepo
		timeout  time.Duration
	}
	fields_ := fields{
		saleRepo: *mockSaleRepo,
		timeout:  time.Second,
	}
	type args struct {
		sale *domain.Sale
		lg   *logrus.Logger
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{"normal", fields_, args{&sale, &logrus.Logger{}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sr := services.NewSaleService(mockSaleRepo, time.Second)
			if err := sr.Create(tt.args.sale, tt.args.lg); (err != nil) != tt.wantErr {
				t.Errorf("saleService.Create() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_saleService_GetById(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockSaleRepo := mock_domain.NewMockISaleRepo(ctrl)

	sale := domain.Sale{
		ID:          1,
		Name:        "весенняя скидка",
		ExpiredTime: time.Date(2025, 10, 10, 0, 0, 0, 0, time.UTC),
		Percent:     0,
	}
	mockSaleRepo.EXPECT().GetById(gomock.Any(), 1, &logrus.Logger{}).Return(sale, nil)

	type fields struct {
		saleRepo mock_domain.MockISaleRepo
		timeout  time.Duration
	}
	fields_ := fields{
		saleRepo: *mockSaleRepo,
		timeout:  time.Second,
	}
	type args struct {
		id int
		lg *logrus.Logger
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    domain.Sale
		wantErr bool
	}{
		{"normal", fields_, args{1, &logrus.Logger{}}, sale, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sr := services.NewSaleService(mockSaleRepo, time.Second)
			got, err := sr.GetById(tt.args.id, tt.args.lg)
			if (err != nil) != tt.wantErr {
				t.Errorf("saleService.GetById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("saleService.GetById() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_saleService_Update(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockSaleRepo := mock_domain.NewMockISaleRepo(ctrl)
	sale := domain.Sale{
		ID:          1,
		Name:        "весенняя скидка",
		ExpiredTime: time.Date(2025, 10, 10, 0, 0, 0, 0, time.UTC),
		Percent:     0,
	}

	mockSaleRepo.EXPECT().Update(gomock.Any(), 1, &sale, &logrus.Logger{}).Return(nil)

	type fields struct {
		saleRepo mock_domain.MockISaleRepo
		timeout  time.Duration
	}
	fields_ := fields{
		saleRepo: *mockSaleRepo,
		timeout:  time.Second,
	}
	type args struct {
		id       int
		newState *domain.Sale
		lg       *logrus.Logger
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{"normal", fields_, args{1, &sale, &logrus.Logger{}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sr := services.NewSaleService(mockSaleRepo, time.Second)
			if err := sr.Update(tt.args.id, tt.args.newState, tt.args.lg); (err != nil) != tt.wantErr {
				t.Errorf("saleService.Update() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_saleService_Delete(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockSaleRepo := mock_domain.NewMockISaleRepo(ctrl)
	mockSaleRepo.EXPECT().Delete(gomock.Any(), 1, &logrus.Logger{}).Return(nil)

	type fields struct {
		saleRepo mock_domain.MockISaleRepo
		timeout  time.Duration
	}
	fields_ := fields{
		saleRepo: *mockSaleRepo,
		timeout:  time.Second,
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
			sr := services.NewSaleService(mockSaleRepo, time.Second)
			if err := sr.Delete(tt.args.id, tt.args.lg); (err != nil) != tt.wantErr {
				t.Errorf("saleService.Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_saleService_GetByCriteria(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockSaleRepo := mock_domain.NewMockISaleRepo(ctrl)
	sale := domain.Sale{
		ID:          1,
		Name:        "весенняя скидка",
		ExpiredTime: time.Date(2025, 10, 10, 0, 0, 0, 0, time.UTC),
		Percent:     0,
	}
	sales := make([]domain.Sale, 0)
	sales = append(sales, sale)

	mockSaleRepo.EXPECT().GetByCriteria(gomock.Any(), 10, 10, &sale, &logrus.Logger{}).Return(sales, nil)

	type fields struct {
		saleRepo mock_domain.MockISaleRepo
		timeout  time.Duration
	}
	fields_ := fields{
		saleRepo: *mockSaleRepo,
		timeout:  time.Second,
	}
	type args struct {
		saleCriteria *domain.Sale
		offset       int
		limit        int
		lg           *logrus.Logger
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []domain.Sale
		wantErr bool
	}{
		{"normal", fields_, args{&sale, 10, 10, &logrus.Logger{}}, sales, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sr := services.NewSaleService(mockSaleRepo, time.Second)
			got, err := sr.GetByCriteria(tt.args.saleCriteria, tt.args.offset, tt.args.limit, tt.args.lg)
			if (err != nil) != tt.wantErr {
				t.Errorf("saleService.GetByCriteria() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("saleService.GetByCriteria() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_saleService_GetSales(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockSaleRepo := mock_domain.NewMockISaleRepo(ctrl)
	sale := domain.Sale{
		ID:          1,
		Name:        "весенняя скидка",
		ExpiredTime: time.Date(2025, 10, 10, 0, 0, 0, 0, time.UTC),
		Percent:     0,
	}
	sales := make([]domain.Sale, 0)
	sales = append(sales, sale)

	mockSaleRepo.EXPECT().GetLimit(gomock.Any(), 10, 10, &logrus.Logger{}).Return(sales, nil)

	type fields struct {
		saleRepo mock_domain.MockISaleRepo
		timeout  time.Duration
	}
	fields_ := fields{
		saleRepo: *mockSaleRepo,
		timeout:  time.Second,
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
		want    []domain.Sale
		wantErr bool
	}{
		{"normal", fields_, args{10, 10, &logrus.Logger{}}, sales, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sr := services.NewSaleService(mockSaleRepo, time.Second)
			got, err := sr.GetSales(tt.args.offset, tt.args.limit, tt.args.lg)
			if (err != nil) != tt.wantErr {
				t.Errorf("saleService.GetSales() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("saleService.GetSales() = %v, want %v", got, tt.want)
			}
		})
	}
}
