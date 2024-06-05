package unit

import (
	"app/services"
	mock_domain "app/unit/mocks"
	"go.uber.org/mock/gomock"
	"testing"
	"time"

	"github.com/sirupsen/logrus"
)

func Test_authService_GetToken(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockAuthRepo := mock_domain.NewMockIAuthRepo(ctrl)

	mockAuthRepo.EXPECT().GetToken(gomock.Any(), "login", &logrus.Logger{}).Return("token", nil)

	type fields struct {
		authRepo mock_domain.MockIAuthRepo
		timeout  time.Duration
	}
	fields_ := fields{
		authRepo: *mockAuthRepo,
		timeout:  time.Second,
	}
	type args struct {
		login string
		lg    *logrus.Logger
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{"normal", fields_, args{"login", &logrus.Logger{}}, "token", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			as := services.CreateNewAuthService(mockAuthRepo, time.Second)
			got, err := as.GetToken(tt.args.login, tt.args.lg)
			if (err != nil) != tt.wantErr {
				t.Errorf("authService.GetToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("authService.GetToken() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_authService_AddToken(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockAuthRepo := mock_domain.NewMockIAuthRepo(ctrl)
	mockAuthRepo.EXPECT().AddToken(gomock.Any(), "login", "token", &logrus.Logger{}).Return(nil)

	type fields struct {
		authRepo mock_domain.MockIAuthRepo
		timeout  time.Duration
	}
	fields_ := fields{
		authRepo: *mockAuthRepo,
		timeout:  time.Second,
	}
	type args struct {
		login string
		token string
		lg    *logrus.Logger
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{"normal", fields_, args{"login", "token", &logrus.Logger{}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			as := services.CreateNewAuthService(mockAuthRepo, time.Second)
			if err := as.AddToken(tt.args.login, tt.args.token, tt.args.lg); (err != nil) != tt.wantErr {
				t.Errorf("authService.AddToken() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_authService_CheckAccessRights(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockAuthRepo := mock_domain.NewMockIAuthRepo(ctrl)
	type fields struct {
		authRepo mock_domain.MockIAuthRepo
		timeout  time.Duration
	}
	fields_ := fields{
		authRepo: *mockAuthRepo,
		timeout:  time.Second,
	}
	type args struct {
		tokenString   string
		needRoleLevel string
		lg            *logrus.Logger
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int
		wantErr bool
	}{
		{"normal client"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			as := services.CreateNewAuthService(mockAuthRepo, time.Second)
			got, err := as.CheckAccessRights(tt.args.tokenString, tt.args.needRoleLevel, tt.args.lg)
			if (err != nil) != tt.wantErr {
				t.Errorf("authService.CheckAccessRights() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("authService.CheckAccessRights() = %v, want %v", got, tt.want)
			}
		})
	}
}

//func Test_authService_ExtractIdFromToken(t *testing.T) {
//	type fields struct {
//		authRepo domain.IAuthRepo
//		timeout  time.Duration
//	}
//	type args struct {
//		tokenString string
//		lg          *logrus.Logger
//	}
//	tests := []struct {
//		name    string
//		fields  fields
//		args    args
//		want    int
//		wantErr bool
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			a := &authService{
//				authRepo: tt.fields.authRepo,
//				timeout:  tt.fields.timeout,
//			}
//			got, err := a.ExtractIdFromToken(tt.args.tokenString, tt.args.lg)
//			if (err != nil) != tt.wantErr {
//				t.Errorf("authService.ExtractIdFromToken() error = %v, wantErr %v", err, tt.wantErr)
//				return
//			}
//			if got != tt.want {
//				t.Errorf("authService.ExtractIdFromToken() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
