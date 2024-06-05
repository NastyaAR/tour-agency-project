package unit

import (
	"app/domain"
	"app/services"
	mock_domain "app/unit/mocks"
	"context"
	"go.uber.org/mock/gomock"
	"golang.org/x/xerrors"
	"reflect"
	"testing"
)

//func TestCreateNewRequestService(t *testing.T) {
//	type args struct {
//		RequestRepo domain.IRequestRepo
//		PayAdapter  domain.IPayAdapter
//	}
//	tests := []struct {
//		name string
//		args args
//		want domain.IRequestService
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if got := CreateNewRequestService(tt.args.RequestRepo, tt.args.PayAdapter); !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("CreateNewRequestService() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}

func Test_requestService_Create(t *testing.T) {
	ctrlR := gomock.NewController(t)
	ctrlP := gomock.NewController(t)
	mockRequestRepo := mock_domain.NewMockIRequestRepo(ctrlR)
	mockPayAdapter := mock_domain.NewMockIPayAdapter(ctrlP)

	req := domain.Request{
		ID:         1,
		TourID:     1,
		ClntID:     1,
		MngrID:     1,
		Status:     "принята",
		CreateTime: "2024.03.12",
		ModifyTime: "",
		Data:       "",
	}

	mockRequestRepo.EXPECT().Add(nil, &req).Return(nil)

	type fields struct {
		RequestRepo mock_domain.MockIRequestRepo
		PayAdapter  mock_domain.MockIPayAdapter
	}
	fields_ := fields{
		RequestRepo: *mockRequestRepo,
		PayAdapter:  *mockPayAdapter,
	}
	type args struct {
		c   context.Context
		req *domain.Request
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{"normal", fields_, args{nil, &req}, false},
		{"bad tour id", fields_, args{nil, &domain.Request{TourID: -1}}, true},
		{"bad mngr id", fields_, args{nil, &domain.Request{MngrID: -1}}, true},
		{"bad clnt id", fields_, args{nil, &domain.Request{ClntID: -1}}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rs := services.CreateNewRequestService(mockRequestRepo, mockPayAdapter)
			if err := rs.Create(tt.args.c, tt.args.req); (err != nil) != tt.wantErr {
				t.Errorf("requestService.Create() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_requestService_GetById(t *testing.T) {
	ctrlR := gomock.NewController(t)
	ctrlP := gomock.NewController(t)
	mockRequestRepo := mock_domain.NewMockIRequestRepo(ctrlR)
	mockPayAdapter := mock_domain.NewMockIPayAdapter(ctrlP)

	req := domain.Request{
		ID:         1,
		TourID:     1,
		ClntID:     1,
		MngrID:     1,
		Status:     "принята",
		CreateTime: "2024.03.12",
		ModifyTime: "",
		Data:       "",
	}

	mockRequestRepo.EXPECT().GetById(nil, 1).Return(req, nil)

	type fields struct {
		RequestRepo mock_domain.MockIRequestRepo
		PayAdapter  mock_domain.MockIPayAdapter
	}
	fields_ := fields{
		RequestRepo: *mockRequestRepo,
		PayAdapter:  *mockPayAdapter,
	}
	type args struct {
		c  context.Context
		id int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    domain.Request
		wantErr bool
	}{
		{"normal", fields_, args{nil, 1}, req, false},
		{"bad id", fields_, args{nil, -1}, domain.Request{}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rs := services.CreateNewRequestService(mockRequestRepo, mockPayAdapter)
			got, err := rs.GetById(tt.args.c, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("requestService.GetById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("requestService.GetById() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_requestService_Update(t *testing.T) {
	ctrlR := gomock.NewController(t)
	ctrlP := gomock.NewController(t)
	mockRequestRepo := mock_domain.NewMockIRequestRepo(ctrlR)
	mockPayAdapter := mock_domain.NewMockIPayAdapter(ctrlP)

	req := domain.Request{
		ID:         1,
		TourID:     1,
		ClntID:     1,
		MngrID:     1,
		Status:     "принята",
		CreateTime: "2024.03.12",
		ModifyTime: "",
		Data:       "",
	}

	mockRequestRepo.EXPECT().Update(nil, 1, &req).Return(nil)

	type fields struct {
		RequestRepo mock_domain.MockIRequestRepo
		PayAdapter  mock_domain.MockIPayAdapter
	}
	fields_ := fields{
		RequestRepo: *mockRequestRepo,
		PayAdapter:  *mockPayAdapter,
	}
	type args struct {
		c        context.Context
		id       int
		newState *domain.Request
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{"normal", fields_, args{nil, 1, &req}, false},
		{"bad id", fields_, args{nil, -1, &req}, true},
		{"empty", fields_, args{nil, 1, nil}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rs := services.CreateNewRequestService(mockRequestRepo, mockPayAdapter)
			if err := rs.Update(tt.args.c, tt.args.id, tt.args.newState); (err != nil) != tt.wantErr {
				t.Errorf("requestService.Update() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_requestService_GetByStatus(t *testing.T) {
	ctrlR := gomock.NewController(t)
	ctrlP := gomock.NewController(t)
	mockRequestRepo := mock_domain.NewMockIRequestRepo(ctrlR)
	mockPayAdapter := mock_domain.NewMockIPayAdapter(ctrlP)

	reqs := []domain.Request{
		{Status: "принята"},
		{Status: "обрабатывается"},
		{Status: "отклонена"},
		{Status: "оплачена"},
		{Status: "подтверждена"},
	}

	mockRequestRepo.EXPECT().GetByStatus(nil, "принята").Return([]domain.Request{reqs[0]}, nil)
	mockRequestRepo.EXPECT().GetByStatus(nil, "обрабатывается").Return([]domain.Request{reqs[1]}, nil)
	mockRequestRepo.EXPECT().GetByStatus(nil, "отклонена").Return([]domain.Request{reqs[2]}, nil)
	mockRequestRepo.EXPECT().GetByStatus(nil, "оплачена").Return([]domain.Request{reqs[3]}, nil)
	mockRequestRepo.EXPECT().GetByStatus(nil, "подтверждена").Return([]domain.Request{reqs[4]}, nil)

	type fields struct {
		RequestRepo mock_domain.MockIRequestRepo
		PayAdapter  mock_domain.MockIPayAdapter
	}
	fields_ := fields{
		RequestRepo: *mockRequestRepo,
		PayAdapter:  *mockPayAdapter,
	}
	type args struct {
		c      context.Context
		status string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []domain.Request
		wantErr bool
	}{
		{"normal accepted", fields_, args{nil, "принята"}, []domain.Request{reqs[0]}, false},
		{"normal handle", fields_, args{nil, "обрабатывается"}, []domain.Request{reqs[1]}, false},
		{"normal reject", fields_, args{nil, "отклонена"}, []domain.Request{reqs[2]}, false},
		{"normal payed", fields_, args{nil, "оплачена"}, []domain.Request{reqs[3]}, false},
		{"normal approved", fields_, args{nil, "подтверждена"}, []domain.Request{reqs[4]}, false},
		{"bad status", fields_, args{nil, "неизвестно"}, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rs := services.CreateNewRequestService(mockRequestRepo, mockPayAdapter)
			got, err := rs.GetByStatus(tt.args.c, tt.args.status)
			if (err != nil) != tt.wantErr {
				t.Errorf("requestService.GetByStatus() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("requestService.GetByStatus() = %v, want %v", got, tt.want)
			}
		})
	}
}

//func Test_requestService_GetByTime(t *testing.T) {
//	type fields struct {
//		RequestRepo domain.IRequestRepo
//		PayAdapter  domain.IPayAdapter
//	}
//	type args struct {
//		c    context.Context
//		from string
//		to   string
//	}
//	tests := []struct {
//		name    string
//		fields  fields
//		args    args
//		want    []domain.Request
//		wantErr bool
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			rs := services.CreateNewRequestService(mockRequestRepo, mockPayAdapter)
//			got, err := rs.GetByTime(tt.args.c, tt.args.from, tt.args.to)
//			if (err != nil) != tt.wantErr {
//				t.Errorf("requestService.GetByTime() error = %v, wantErr %v", err, tt.wantErr)
//				return
//			}
//			if !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("requestService.GetByTime() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}

func Test_requestService_Pay(t *testing.T) {
	ctrlR := gomock.NewController(t)
	ctrlP := gomock.NewController(t)
	mockRequestRepo := mock_domain.NewMockIRequestRepo(ctrlR)
	mockPayRepo := mock_domain.NewMockIPayRepo(ctrlP)

	req := domain.Request{
		ID:         1,
		TourID:     1,
		ClntID:     1,
		MngrID:     1,
		Status:     "подтверждена",
		CreateTime: "2024.03.12",
		ModifyTime: "",
		Data:       "",
	}
	badReq := domain.Request{
		ID:         1,
		TourID:     1,
		ClntID:     1,
		MngrID:     1,
		Status:     "принята",
		CreateTime: "2024.03.12",
		ModifyTime: "",
		Data:       "",
	}
	req2 := domain.Request{
		ID:         3,
		TourID:     1,
		ClntID:     1,
		MngrID:     1,
		Status:     "подтверждена",
		CreateTime: "2024.03.12",
		ModifyTime: "",
		Data:       "",
	}
	req3 := domain.Request{
		ID:         4,
		TourID:     1,
		ClntID:     1,
		MngrID:     1,
		Status:     "подтверждена",
		CreateTime: "2024.03.12",
		ModifyTime: "",
		Data:       "",
	}
	req4 := domain.Request{
		ID:         5,
		TourID:     1,
		ClntID:     1,
		MngrID:     1,
		Status:     "подтверждена",
		CreateTime: "2024.03.12",
		ModifyTime: "",
		Data:       "",
	}
	mockRequestRepo.EXPECT().GetById(nil, 1).Return(req, nil)
	mockRequestRepo.EXPECT().CountFinalCost(nil, 1).Return(float32(100.0), nil)
	mockPayRepo.EXPECT().Pay(nil, gomock.Any()).Return(1, nil)
	mockRequestRepo.EXPECT().Pay(nil, &req).Return(nil)
	mockRequestRepo.EXPECT().GetById(nil, 2).Return(badReq, nil)
	mockRequestRepo.EXPECT().GetById(nil, 3).Return(req2, nil)
	mockRequestRepo.EXPECT().CountFinalCost(nil, 3).Return(float32(0), xerrors.New("ошибка"))
	mockRequestRepo.EXPECT().GetById(nil, 4).Return(req3, nil)
	mockRequestRepo.EXPECT().CountFinalCost(nil, 4).Return(float32(10), nil)
	mockPayRepo.EXPECT().Pay(nil, gomock.Any()).Return(4, xerrors.New("ошибка"))
	mockRequestRepo.EXPECT().GetById(nil, 5).Return(req4, nil)
	mockRequestRepo.EXPECT().CountFinalCost(nil, 5).Return(float32(10), nil)
	mockPayRepo.EXPECT().Pay(nil, gomock.Any()).Return(5, nil)
	mockRequestRepo.EXPECT().Pay(nil, &req4).Return(xerrors.New("ошибка"))

	type fields struct {
		RequestRepo mock_domain.MockIRequestRepo
		PayRepo     mock_domain.MockIPayRepo
	}
	fields_ := fields{
		RequestRepo: *mockRequestRepo,
		PayRepo:     *mockPayRepo,
	}
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
		{"bad status", fields_, args{nil, 2}, true},
		{"bad cost", fields_, args{nil, 3}, true},
		{"bad pay", fields_, args{nil, 4}, true},
		{"bad request pay", fields_, args{nil, 5}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rs := services.CreateNewRequestService(mockRequestRepo, mockPayRepo)
			if err := rs.Pay(tt.args.c, tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("requestService.Pay() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_requestService_Reject(t *testing.T) {
	ctrlR := gomock.NewController(t)
	ctrlP := gomock.NewController(t)
	mockRequestRepo := mock_domain.NewMockIRequestRepo(ctrlR)
	mockPayRepo := mock_domain.NewMockIPayRepo(ctrlP)

	mockRequestRepo.EXPECT().Reject(nil, 1)

	type fields struct {
		RequestRepo mock_domain.MockIRequestRepo
		PayRepo     mock_domain.MockIPayRepo
	}
	fields_ := fields{
		RequestRepo: *mockRequestRepo,
		PayRepo:     *mockPayRepo,
	}
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
		{"bad id", fields_, args{nil, -1}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rs := services.CreateNewRequestService(mockRequestRepo, mockPayRepo)
			if err := rs.Reject(tt.args.c, tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("requestService.Reject() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_requestService_Approve(t *testing.T) {
	ctrlR := gomock.NewController(t)
	ctrlP := gomock.NewController(t)
	mockRequestRepo := mock_domain.NewMockIRequestRepo(ctrlR)
	mockPayRepo := mock_domain.NewMockIPayRepo(ctrlP)

	mockRequestRepo.EXPECT().Approve(nil, 1)
	type fields struct {
		RequestRepo mock_domain.MockIRequestRepo
		PayRepo     mock_domain.MockIPayRepo
	}
	fields_ := fields{
		RequestRepo: *mockRequestRepo,
		PayRepo:     *mockPayRepo,
	}
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
		{"bad id", fields_, args{nil, -1}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rs := services.CreateNewRequestService(mockRequestRepo, mockPayRepo)
			if err := rs.Approve(tt.args.c, tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("requestService.Approve() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
