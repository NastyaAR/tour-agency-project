package unit

import (
	"app/domain"
	"app/services"
	mock_domain "app/unit/mocks"
	"context"
	"go.uber.org/mock/gomock"
	"reflect"
	"testing"
)

func Test_saleService_Create(t *testing.T) {
	sale := domain.Sale{
		ID:          1,
		Name:        "скидка-скидка",
		ExpiredTime: "2024.04.40",
		Percent:     10,
	}
	sale2 := domain.Sale{
		ID:          1,
		Name:        "скидка-скидка",
		ExpiredTime: "2024.04.40",
		Percent:     -10,
	}
	sale3 := domain.Sale{
		ID:          1,
		Name:        "скидка-скидка",
		ExpiredTime: "2024.04.40",
		Percent:     101,
	}

	ctrl := gomock.NewController(t)
	mockSaleRepo := mock_domain.NewMockISaleRepo(ctrl)
	mockSaleRepo.EXPECT().Add(
		nil,
		&sale,
	).Return(nil)

	type fields struct {
		SaleRepo mock_domain.MockISaleRepo
	}
	type args struct {
		c    context.Context
		sale *domain.Sale
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{"normal", fields{*mockSaleRepo}, args{nil, &sale}, false},
		{"bad_percent_small", fields{*mockSaleRepo}, args{nil, &sale2}, true},
		{"bad_percent_big", fields{*mockSaleRepo}, args{nil, &sale3}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sr := services.NewSaleService(mockSaleRepo)
			if err := sr.Create(tt.args.c, tt.args.sale); (err != nil) != tt.wantErr {
				t.Errorf("saleService.Create() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_saleService_GetById(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockSaleRepo := mock_domain.NewMockISaleRepo(ctrl)
	type fields struct {
		SaleRepo mock_domain.MockISaleRepo
	}
	sale := domain.Sale{
		ID:          1,
		Name:        "скидка-скидка",
		ExpiredTime: "2024.04.40",
		Percent:     10,
	}

	mockSaleRepo.EXPECT().GetById(nil, 1).Return(sale, nil)
	fields_ := fields{SaleRepo: *mockSaleRepo}
	type args struct {
		c  context.Context
		id int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    domain.Sale
		wantErr bool
	}{
		{"normal", fields_, args{nil, 1}, sale, false},
		{"bad id", fields_, args{nil, -1}, domain.Sale{}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sr := services.NewSaleService(mockSaleRepo)
			got, err := sr.GetById(tt.args.c, tt.args.id)
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

	type fields struct {
		SaleRepo mock_domain.MockISaleRepo
	}
	sale := domain.Sale{
		ID:          1,
		Name:        "скидка-скидка",
		ExpiredTime: "2024.04.40",
		Percent:     10,
	}
	badSale := domain.Sale{
		ID:          1,
		Name:        "скидка-скидка",
		ExpiredTime: "2024.04.40",
		Percent:     -10,
	}
	badSale2 := domain.Sale{
		ID:          1,
		Name:        "скидка-скидка",
		ExpiredTime: "2024.04.40",
		Percent:     101,
	}
	mockSaleRepo.EXPECT().Update(nil, 1, &sale).Return(nil)

	fields_ := fields{SaleRepo: *mockSaleRepo}
	type args struct {
		c        context.Context
		id       int
		newState *domain.Sale
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{"normal", fields_, args{nil, 1, &sale}, false},
		{"bad percent", fields_, args{nil, 1, &badSale}, true},
		{"bad percent", fields_, args{nil, 1, &badSale2}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sr := services.NewSaleService(mockSaleRepo)
			if err := sr.Update(tt.args.c, tt.args.id, tt.args.newState); (err != nil) != tt.wantErr {
				t.Errorf("saleService.Update() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_saleService_Delete(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockSaleRepo := mock_domain.NewMockISaleRepo(ctrl)

	mockSaleRepo.EXPECT().Delete(nil, 1).Return(nil)

	type fields struct {
		SaleRepo mock_domain.MockISaleRepo
	}
	fields_ := fields{SaleRepo: *mockSaleRepo}
	type args struct {
		c  context.Context
		id int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{"normal", fields_, args{nil, 1}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sr := services.NewSaleService(mockSaleRepo)
			if err := sr.Delete(tt.args.c, tt.args.id); (err != nil) != tt.wantErr {
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
		Name:        "скидка-скидка",
		ExpiredTime: "2024.04.40",
		Percent:     10,
	}
	sales := make([]domain.Sale, 0)
	sales = append(sales, sale)

	mockSaleRepo.EXPECT().GetByCriteria(nil, &sale).Return(sales, nil)

	type fields struct {
		SaleRepo mock_domain.MockISaleRepo
	}
	fields_ := fields{SaleRepo: *mockSaleRepo}
	type args struct {
		c            context.Context
		saleCriteria *domain.Sale
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []domain.Sale
		wantErr bool
	}{
		{"normal", fields_, args{nil, &sale}, sales, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sr := services.NewSaleService(mockSaleRepo)
			got, err := sr.GetByCriteria(tt.args.c, tt.args.saleCriteria)
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
