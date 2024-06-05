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
		CreateTime: time.Date(2024, 10, 10, 1, 0, 0, 0, time.UTC),
		ModifyTime: time.Date(2024, 10, 10, 1, 0, 0, 0, time.UTC),
		Data:       "",
	}
	mockRequestRepo.EXPECT().GetNonSendEvents(gomock.Any(), 10, &logrus.Logger{}).AnyTimes()
	mockRequestRepo.EXPECT().Add(gomock.Any(), &req, &logrus.Logger{}).Return(nil)

	type fields struct {
		RequestRepo     mock_domain.MockIRequestRepo
		PayAdapter      mock_domain.MockIPayAdapter
		timeout         time.Duration
		payingDone      chan bool
		payingLimit     int
		payingFrequency time.Duration
	}
	var done chan bool
	fields_ := fields{
		RequestRepo:     *mockRequestRepo,
		PayAdapter:      *mockPayAdapter,
		timeout:         time.Second,
		payingDone:      done,
		payingLimit:     10,
		payingFrequency: time.Second,
	}
	type args struct {
		req *domain.Request
		lg  *logrus.Logger
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{"normal", fields_, args{req: &req, lg: &logrus.Logger{}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rs := services.CreateNewRequestService(mockRequestRepo, mockPayAdapter,
				time.Second, done, 10, time.Second, &logrus.Logger{})
			if err := rs.Create(tt.args.req, tt.args.lg); (err != nil) != tt.wantErr {
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
		CreateTime: time.Date(2024, 10, 10, 1, 0, 0, 0, time.UTC),
		ModifyTime: time.Date(2024, 10, 10, 1, 0, 0, 0, time.UTC),
		Data:       "",
	}
	mockRequestRepo.EXPECT().GetNonSendEvents(gomock.Any(), 10, &logrus.Logger{}).AnyTimes()
	mockRequestRepo.EXPECT().GetById(gomock.Any(), 1, &logrus.Logger{}).Return(req, nil)

	type fields struct {
		RequestRepo     mock_domain.MockIRequestRepo
		PayAdapter      mock_domain.MockIPayAdapter
		timeout         time.Duration
		payingDone      chan bool
		payingLimit     int
		payingFrequency time.Duration
	}
	var done chan bool
	fields_ := fields{
		RequestRepo:     *mockRequestRepo,
		PayAdapter:      *mockPayAdapter,
		timeout:         time.Second,
		payingDone:      done,
		payingLimit:     10,
		payingFrequency: time.Second,
	}
	type args struct {
		id int
		lg *logrus.Logger
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    domain.Request
		wantErr bool
	}{
		{"normal", fields_, args{1, &logrus.Logger{}}, req, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rs := services.CreateNewRequestService(mockRequestRepo, mockPayAdapter,
				time.Second, done, 10, time.Second, &logrus.Logger{})
			got, err := rs.GetById(tt.args.id, tt.args.lg)
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
		CreateTime: time.Date(2024, 10, 10, 1, 0, 0, 0, time.UTC),
		ModifyTime: time.Date(2024, 10, 10, 1, 0, 0, 0, time.UTC),
		Data:       "",
	}
	mockRequestRepo.EXPECT().GetNonSendEvents(gomock.Any(), 10, &logrus.Logger{}).AnyTimes()
	mockRequestRepo.EXPECT().Update(gomock.Any(), 1, &req, &logrus.Logger{})

	type fields struct {
		RequestRepo     mock_domain.MockIRequestRepo
		PayAdapter      mock_domain.MockIPayAdapter
		timeout         time.Duration
		payingDone      chan bool
		payingLimit     int
		payingFrequency time.Duration
	}
	var done chan bool
	fields_ := fields{
		RequestRepo:     *mockRequestRepo,
		PayAdapter:      *mockPayAdapter,
		timeout:         time.Second,
		payingDone:      done,
		payingLimit:     10,
		payingFrequency: time.Second,
	}
	type args struct {
		id       int
		newState *domain.Request
		lg       *logrus.Logger
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{"normal", fields_, args{1, &req, &logrus.Logger{}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rs := services.CreateNewRequestService(mockRequestRepo, mockPayAdapter,
				time.Second, done, 10, time.Second, &logrus.Logger{})
			if err := rs.Update(tt.args.id, tt.args.newState, tt.args.lg); (err != nil) != tt.wantErr {
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

	req := domain.Request{
		ID:         1,
		TourID:     1,
		ClntID:     1,
		MngrID:     1,
		Status:     "принята",
		CreateTime: time.Date(2024, 10, 10, 1, 0, 0, 0, time.UTC),
		ModifyTime: time.Date(2024, 10, 10, 1, 0, 0, 0, time.UTC),
		Data:       "",
	}
	reqs := make([]domain.Request, 0)
	reqs = append(reqs, req)
	mockRequestRepo.EXPECT().GetNonSendEvents(gomock.Any(), 10, &logrus.Logger{}).AnyTimes()
	mockRequestRepo.EXPECT().GetByStatus(gomock.Any(), "принята", 10, 10, &logrus.Logger{}).Return(reqs, nil)

	type fields struct {
		RequestRepo     mock_domain.MockIRequestRepo
		PayAdapter      mock_domain.MockIPayAdapter
		timeout         time.Duration
		payingDone      chan bool
		payingLimit     int
		payingFrequency time.Duration
	}
	var done chan bool
	fields_ := fields{
		RequestRepo:     *mockRequestRepo,
		PayAdapter:      *mockPayAdapter,
		timeout:         time.Second,
		payingDone:      done,
		payingLimit:     10,
		payingFrequency: time.Second,
	}
	type args struct {
		status string
		offset int
		limit  int
		lg     *logrus.Logger
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []domain.Request
		wantErr bool
	}{
		{"normal", fields_, args{"принята", 10, 10, &logrus.Logger{}}, reqs, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rs := services.CreateNewRequestService(mockRequestRepo, mockPayAdapter,
				time.Second, done, 10, time.Second, &logrus.Logger{})
			got, err := rs.GetByStatus(tt.args.status, tt.args.offset, tt.args.limit, tt.args.lg)
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

func Test_requestService_Pay(t *testing.T) {
	ctrlR := gomock.NewController(t)
	ctrlP := gomock.NewController(t)
	mockRequestRepo := mock_domain.NewMockIRequestRepo(ctrlR)
	mockPayAdapter := mock_domain.NewMockIPayAdapter(ctrlP)

	req := domain.Request{
		ID:         1,
		TourID:     1,
		ClntID:     1,
		MngrID:     1,
		Status:     "подтверждена",
		CreateTime: time.Date(2024, 10, 10, 1, 0, 0, 0, time.UTC),
		ModifyTime: time.Date(2024, 10, 10, 1, 0, 0, 0, time.UTC),
		Data:       "",
	}

	badStatReq := domain.Request{
		ID:         3,
		TourID:     1,
		ClntID:     1,
		MngrID:     1,
		Status:     "принята",
		CreateTime: time.Date(2024, 10, 10, 1, 0, 0, 0, time.UTC),
		ModifyTime: time.Date(2024, 10, 10, 1, 0, 0, 0, time.UTC),
		Data:       "",
	}

	badCntFCReq := domain.Request{
		ID:         4,
		TourID:     1,
		ClntID:     4,
		MngrID:     1,
		Status:     "подтверждена",
		CreateTime: time.Date(2024, 10, 10, 1, 0, 0, 0, time.UTC),
		ModifyTime: time.Date(2024, 10, 10, 1, 0, 0, 0, time.UTC),
		Data:       "",
	}

	badAtomPayReq := domain.Request{
		ID:         5,
		TourID:     1,
		ClntID:     1,
		MngrID:     1,
		Status:     "подтверждена",
		CreateTime: time.Date(2024, 10, 10, 1, 0, 0, 0, time.UTC),
		ModifyTime: time.Date(2024, 10, 10, 1, 0, 0, 0, time.UTC),
		Data:       "",
	}

	mockRequestRepo.EXPECT().GetNonSendEvents(gomock.Any(), 10, &logrus.Logger{}).AnyTimes()
	mockRequestRepo.EXPECT().GetById(gomock.Any(), 1, &logrus.Logger{}).Return(req, nil)
	mockRequestRepo.EXPECT().CountFinalCost(gomock.Any(), 1, &logrus.Logger{}).Return(100, nil)
	mockRequestRepo.EXPECT().AtomicPay(gomock.Any(), 100, &req, &logrus.Logger{}).Return(nil)
	mockRequestRepo.EXPECT().GetById(gomock.Any(), 2, &logrus.Logger{}).Return(domain.Request{}, xerrors.New("err"))
	mockRequestRepo.EXPECT().GetNonSendEvents(gomock.Any(), 10, &logrus.Logger{}).AnyTimes()
	mockRequestRepo.EXPECT().GetById(gomock.Any(), 3, &logrus.Logger{}).Return(badStatReq, nil)
	mockRequestRepo.EXPECT().GetById(gomock.Any(), 4, &logrus.Logger{}).Return(badCntFCReq, nil)
	mockRequestRepo.EXPECT().CountFinalCost(gomock.Any(), 1, &logrus.Logger{}).Return(-1, xerrors.New("err"))
	mockRequestRepo.EXPECT().GetById(gomock.Any(), 5, &logrus.Logger{}).Return(badAtomPayReq, nil)

	type fields struct {
		RequestRepo     mock_domain.MockIRequestRepo
		PayAdapter      mock_domain.MockIPayAdapter
		timeout         time.Duration
		payingDone      chan bool
		payingLimit     int
		payingFrequency time.Duration
	}
	var done chan bool
	fields_ := fields{
		RequestRepo:     *mockRequestRepo,
		PayAdapter:      *mockPayAdapter,
		timeout:         time.Second,
		payingDone:      done,
		payingLimit:     10,
		payingFrequency: time.Second,
	}
	type args struct {
		id      int
		clnt_id int
		lg      *logrus.Logger
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{"normal", fields_, args{1, 1, &logrus.Logger{}}, false},
		{"bad id", fields_, args{-1, 1, &logrus.Logger{}}, true},
		{"bad get by id", fields_, args{2, 1, &logrus.Logger{}}, true},
		{"bad status", fields_, args{3, 1, &logrus.Logger{}}, true},
		{"bad count final cost", fields_, args{4, 4, &logrus.Logger{}}, true},
		{"bad atomic pay", fields_, args{5, 5, &logrus.Logger{}}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rs := services.CreateNewRequestService(mockRequestRepo, mockPayAdapter,
				time.Second, done, 10, time.Second, &logrus.Logger{})
			if err := rs.Pay(tt.args.id, tt.args.clnt_id, tt.args.lg); (err != nil) != tt.wantErr {
				t.Errorf("requestService.Pay() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_requestService_Paying(t *testing.T) {
	ctrlR := gomock.NewController(t)
	ctrlP := gomock.NewController(t)
	mockRequestRepo := mock_domain.NewMockIRequestRepo(ctrlR)
	mockPayAdapter := mock_domain.NewMockIPayAdapter(ctrlP)

	req := domain.PayEvent{
		Id:    0,
		ReqID: 1,
		Sum:   100,
		State: "non-send",
	}

	req2 := domain.PayEvent{
		Id:    1,
		ReqID: 2,
		Sum:   120,
		State: "non-send",
	}

	reqs := make([]domain.PayEvent, 0)
	reqs = append(reqs, req, req2)

	mockRequestRepo.EXPECT().GetNonSendEvents(gomock.Any(), 2, &logrus.Logger{}).Return(reqs, nil).AnyTimes()

	mockPayAdapter.EXPECT().SendPaymentRequest(&req, &logrus.Logger{}).Return(nil).AnyTimes()
	mockPayAdapter.EXPECT().SendPaymentRequest(&req2, &logrus.Logger{}).Return(nil).AnyTimes()
	mockRequestRepo.EXPECT().UpdateOutbox(gomock.Any(), req.ReqID, &logrus.Logger{}).Return(nil).AnyTimes()
	mockRequestRepo.EXPECT().UpdateOutbox(gomock.Any(), req2.ReqID, &logrus.Logger{}).Return(nil).AnyTimes()

	type fields struct {
		RequestRepo     mock_domain.MockIRequestRepo
		PayAdapter      mock_domain.MockIPayAdapter
		timeout         time.Duration
		payingDone      chan bool
		payingLimit     int
		payingFrequency time.Duration
	}
	done := make(chan bool, 1)
	fields_ := fields{
		RequestRepo:     *mockRequestRepo,
		PayAdapter:      *mockPayAdapter,
		timeout:         time.Second,
		payingDone:      done,
		payingLimit:     2,
		payingFrequency: time.Second,
	}
	type args struct {
		done      chan bool
		frequency time.Duration
		lg        *logrus.Logger
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{"normal", fields_, args{done, time.Second, &logrus.Logger{}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_ = services.CreateNewRequestService(mockRequestRepo, mockPayAdapter,
				time.Second, done, 10, time.Second, &logrus.Logger{})
			done <- true
		})
	}
}

func Test_requestService_Reject(t *testing.T) {
	ctrlR := gomock.NewController(t)
	ctrlP := gomock.NewController(t)
	mockRequestRepo := mock_domain.NewMockIRequestRepo(ctrlR)
	mockPayAdapter := mock_domain.NewMockIPayAdapter(ctrlP)

	req := domain.Request{
		ID:         1,
		TourID:     1,
		ClntID:     1,
		MngrID:     1,
		Status:     domain.Approved,
		CreateTime: time.Date(2024, 10, 10, 1, 0, 0, 0, time.UTC),
		ModifyTime: time.Date(2024, 10, 10, 1, 0, 0, 0, time.UTC),
		Data:       "",
	}

	payedReq := domain.Request{
		ID:         2,
		TourID:     1,
		ClntID:     1,
		MngrID:     1,
		Status:     domain.Paid,
		CreateTime: time.Date(2024, 10, 10, 1, 0, 0, 0, time.UTC),
		ModifyTime: time.Date(2024, 10, 10, 1, 0, 0, 0, time.UTC),
		Data:       "",
	}

	badMngrIdReq := domain.Request{
		ID:         3,
		TourID:     1,
		ClntID:     1,
		MngrID:     1,
		Status:     domain.Paid,
		CreateTime: time.Date(2024, 10, 10, 1, 0, 0, 0, time.UTC),
		ModifyTime: time.Date(2024, 10, 10, 1, 0, 0, 0, time.UTC),
		Data:       "",
	}

	rejectedReq := domain.Request{
		ID:         4,
		TourID:     1,
		ClntID:     1,
		MngrID:     1,
		Status:     domain.Rejected,
		CreateTime: time.Date(2024, 10, 10, 1, 0, 0, 0, time.UTC),
		ModifyTime: time.Date(2024, 10, 10, 1, 0, 0, 0, time.UTC),
		Data:       "",
	}

	mockRequestRepo.EXPECT().GetById(gomock.Any(), 1, &logrus.Logger{}).Return(req, nil)
	mockRequestRepo.EXPECT().Reject(gomock.Any(), 1, &logrus.Logger{}).Return(nil)
	mockRequestRepo.EXPECT().GetNonSendEvents(gomock.Any(), 10, &logrus.Logger{}).AnyTimes()
	mockRequestRepo.EXPECT().GetById(gomock.Any(), 2, &logrus.Logger{}).Return(payedReq, nil)
	mockRequestRepo.EXPECT().GetById(gomock.Any(), 3, &logrus.Logger{}).Return(badMngrIdReq, nil)
	mockRequestRepo.EXPECT().GetById(gomock.Any(), 4, &logrus.Logger{}).Return(rejectedReq, nil)

	type fields struct {
		RequestRepo     mock_domain.MockIRequestRepo
		PayAdapter      mock_domain.MockIPayAdapter
		timeout         time.Duration
		payingDone      chan bool
		payingLimit     int
		payingFrequency time.Duration
	}
	var done chan bool
	fields_ := fields{
		RequestRepo:     *mockRequestRepo,
		PayAdapter:      *mockPayAdapter,
		timeout:         time.Second,
		payingDone:      done,
		payingLimit:     10,
		payingFrequency: time.Second,
	}
	type args struct {
		id      int
		mngr_id int
		lg      *logrus.Logger
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{"normal", fields_, args{1, 1, &logrus.Logger{}}, false},
		{"almost payed", fields_, args{2, 1, &logrus.Logger{}}, true},
		{"bad manager id", fields_, args{3, 2, &logrus.Logger{}}, true},
		{"almost rejected", fields_, args{4, 1, &logrus.Logger{}}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rs := services.CreateNewRequestService(mockRequestRepo, mockPayAdapter,
				time.Second, done, 10, time.Second, &logrus.Logger{})
			if err := rs.Reject(tt.args.id, tt.args.mngr_id, tt.args.lg); (err != nil) != tt.wantErr {
				t.Errorf("requestService.Reject() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_requestService_Approve(t *testing.T) {
	ctrlR := gomock.NewController(t)
	ctrlP := gomock.NewController(t)
	mockRequestRepo := mock_domain.NewMockIRequestRepo(ctrlR)
	mockPayAdapter := mock_domain.NewMockIPayAdapter(ctrlP)

	req := domain.Request{
		ID:         1,
		TourID:     1,
		ClntID:     1,
		MngrID:     1,
		Status:     domain.Accepted,
		CreateTime: time.Time{},
		ModifyTime: time.Time{},
		Data:       "",
	}

	noTourReq := domain.Request{
		ID:         2,
		TourID:     domain.DefaultEmptyValue,
		ClntID:     1,
		MngrID:     1,
		Status:     domain.Accepted,
		CreateTime: time.Time{},
		ModifyTime: time.Time{},
		Data:       "",
	}

	badMngrIdReq := domain.Request{
		ID:         3,
		TourID:     domain.DefaultEmptyValue,
		ClntID:     1,
		MngrID:     1,
		Status:     domain.Accepted,
		CreateTime: time.Time{},
		ModifyTime: time.Time{},
		Data:       "",
	}

	paidReq := domain.Request{
		ID:         4,
		TourID:     1,
		ClntID:     1,
		MngrID:     1,
		Status:     domain.Paid,
		CreateTime: time.Time{},
		ModifyTime: time.Time{},
		Data:       "",
	}

	rejectedReq := domain.Request{
		ID:         5,
		TourID:     1,
		ClntID:     1,
		MngrID:     1,
		Status:     domain.Rejected,
		CreateTime: time.Time{},
		ModifyTime: time.Time{},
		Data:       "",
	}

	mockRequestRepo.EXPECT().GetNonSendEvents(gomock.Any(), 10, &logrus.Logger{}).AnyTimes()
	mockRequestRepo.EXPECT().GetById(gomock.Any(), 1, &logrus.Logger{}).Return(req, nil)
	mockRequestRepo.EXPECT().Approve(gomock.Any(), 1, &logrus.Logger{}).Return(nil)
	mockRequestRepo.EXPECT().GetById(gomock.Any(), 2, &logrus.Logger{}).Return(noTourReq, nil)
	mockRequestRepo.EXPECT().GetById(gomock.Any(), 3, &logrus.Logger{}).Return(badMngrIdReq, nil)
	mockRequestRepo.EXPECT().GetById(gomock.Any(), 4, &logrus.Logger{}).Return(paidReq, nil)
	mockRequestRepo.EXPECT().GetById(gomock.Any(), 5, &logrus.Logger{}).Return(rejectedReq, nil)

	type fields struct {
		RequestRepo     mock_domain.MockIRequestRepo
		PayAdapter      mock_domain.MockIPayAdapter
		timeout         time.Duration
		payingDone      chan bool
		payingLimit     int
		payingFrequency time.Duration
	}
	var done chan bool
	fields_ := fields{
		RequestRepo:     *mockRequestRepo,
		PayAdapter:      *mockPayAdapter,
		timeout:         time.Second,
		payingDone:      done,
		payingLimit:     10,
		payingFrequency: time.Second,
	}
	type args struct {
		id      int
		mngr_id int
		lg      *logrus.Logger
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{"normal", fields_, args{1, 1, &logrus.Logger{}}, false},
		{"bad tour id", fields_, args{2, 1, &logrus.Logger{}}, true},
		{"bad manager id", fields_, args{3, 2, &logrus.Logger{}}, true},
		{"almost paid", fields_, args{4, 1, &logrus.Logger{}}, true},
		{"rejected request", fields_, args{5, 1, &logrus.Logger{}}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rs := services.CreateNewRequestService(mockRequestRepo, mockPayAdapter,
				time.Second, done, 10, time.Second, &logrus.Logger{})
			if err := rs.Approve(tt.args.id, tt.args.mngr_id, tt.args.lg); (err != nil) != tt.wantErr {
				t.Errorf("requestService.Approve() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_requestService_GetRequests(t *testing.T) {
	ctrlR := gomock.NewController(t)
	ctrlP := gomock.NewController(t)
	mockRequestRepo := mock_domain.NewMockIRequestRepo(ctrlR)
	mockPayAdapter := mock_domain.NewMockIPayAdapter(ctrlP)

	req := domain.Request{
		ID:         1,
		TourID:     1,
		ClntID:     1,
		MngrID:     1,
		Status:     "подтверждена",
		CreateTime: time.Date(2024, 10, 10, 1, 0, 0, 0, time.UTC),
		ModifyTime: time.Date(2024, 10, 10, 1, 0, 0, 0, time.UTC),
		Data:       "",
	}
	reqs := make([]domain.Request, 0)
	reqs = append(reqs, req)
	mockRequestRepo.EXPECT().GetNonSendEvents(gomock.Any(), 10, &logrus.Logger{}).AnyTimes()
	mockRequestRepo.EXPECT().GetLimit(gomock.Any(), 10, 10, &logrus.Logger{}).Return(reqs, nil)

	type fields struct {
		RequestRepo     mock_domain.MockIRequestRepo
		PayAdapter      mock_domain.MockIPayAdapter
		timeout         time.Duration
		payingDone      chan bool
		payingLimit     int
		payingFrequency time.Duration
	}
	var done chan bool
	fields_ := fields{
		RequestRepo:     *mockRequestRepo,
		PayAdapter:      *mockPayAdapter,
		timeout:         time.Second,
		payingDone:      done,
		payingLimit:     10,
		payingFrequency: time.Second,
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
		want    []domain.Request
		wantErr bool
	}{
		{"normal", fields_, args{10, 10, &logrus.Logger{}}, reqs, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rs := services.CreateNewRequestService(mockRequestRepo, mockPayAdapter,
				time.Second, done, 10, time.Second, &logrus.Logger{})
			got, err := rs.GetRequests(tt.args.offset, tt.args.limit, tt.args.lg)
			if (err != nil) != tt.wantErr {
				t.Errorf("requestService.GetRequests() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("requestService.GetRequests() = %v, want %v", got, tt.want)
			}
		})
	}
}
