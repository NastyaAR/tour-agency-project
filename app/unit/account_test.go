package unit

import (
	"app/domain"
	"app/services"
	mock_domain "app/unit/mocks"
	"context"
	"go.uber.org/mock/gomock"
	"golang.org/x/xerrors"
	"testing"
)

//func TestCreateNewAccountService(t *testing.T) {
//	type args struct {
//		AccountRepo    domain.IAccountRepo
//		ClientService  domain.IClientService
//		ManagerService domain.IManagerService
//	}
//	tests := []struct {
//		name string
//		args args
//		want domain.IAccountService
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if got := CreateNewAccountService(tt.args.AccountRepo, tt.args.ClientService, tt.args.ManagerService); !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("CreateNewAccountService() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}

func Test_accountService_Register(t *testing.T) {
	ctrlA := gomock.NewController(t)
	ctrlC := gomock.NewController(t)
	ctrlM := gomock.NewController(t)

	mockAccountRepo := mock_domain.NewMockIAccountRepo(ctrlA)
	mockClientRepo := mock_domain.NewMockIClientRepo(ctrlC)
	mockManagerRepo := mock_domain.NewMockIManagerRepo(ctrlM)

	clientService := services.CreateNewClientService(mockClientRepo)
	managerService := services.CreateNewManagerService(mockManagerRepo)

	acc := domain.Account{
		ID:       1,
		UserID:   1,
		Login:    "login",
		Password: "password",
	}
	acc2 := domain.Account{
		ID:       1,
		UserID:   1,
		Login:    "login",
		Password: "password",
	}
	acc3 := domain.Account{
		ID:       1,
		UserID:   1,
		Login:    "nologin",
		Password: "password",
	}
	acc4 := domain.Account{
		ID:       1,
		UserID:   1,
		Login:    "login",
		Password: "password",
	}
	newAcc := domain.NewAccountDTO{
		ID:         1,
		Name:       "имя",
		Surname:    "фамилия",
		Mail:       "почта",
		Phone:      "8888",
		Department: "Москва Измайлово",
		Role:       "manager",
	}
	newAcc2 := domain.NewAccountDTO{
		ID:         1,
		Name:       "имя",
		Surname:    "фамилия",
		Mail:       "почта",
		Phone:      "8888",
		Department: "Москва Измайлово",
		Role:       "client",
	}
	newAcc3 := domain.NewAccountDTO{
		ID:         1,
		Name:       "имя",
		Surname:    "фамилия",
		Mail:       "почта",
		Phone:      "8888",
		Department: "Москва Измайлово",
		Role:       "no role",
	}
	mngr := domain.Manager{
		ID:         1,
		Name:       newAcc.Name,
		Surname:    newAcc.Surname,
		Department: newAcc.Department,
	}

	clnt := domain.Client{
		ID:      1,
		Name:    "имя",
		Surname: "фамилия",
		Mail:    "почта",
		Phone:   "8888",
	}

	mockAccountRepo.EXPECT().GetByLogin(nil, acc.Login).Return(acc, nil)
	mockManagerRepo.EXPECT().Add(nil, &mngr).Return(mngr, nil)
	mockAccountRepo.EXPECT().Add(nil, &acc).Return(nil)

	mockAccountRepo.EXPECT().GetByLogin(nil, acc2.Login).Return(acc2, nil)
	mockClientRepo.EXPECT().Add(nil, &clnt).Return(clnt, nil)
	mockAccountRepo.EXPECT().Add(nil, &acc2).Return(nil)

	mockAccountRepo.EXPECT().GetByLogin(nil, acc3.Login).Return(domain.Account{}, xerrors.New("ошибка"))

	mockAccountRepo.EXPECT().GetByLogin(nil, acc4.Login).Return(acc4, nil)

	type fields struct {
		AccountRepo    mock_domain.MockIAccountRepo
		ClientService  mock_domain.MockIClientRepo
		ManagerService mock_domain.MockIManagerRepo
	}
	fields_ := fields{
		AccountRepo:    *mockAccountRepo,
		ClientService:  *mockClientRepo,
		ManagerService: *mockManagerRepo,
	}
	type args struct {
		c         context.Context
		acc       *domain.Account
		newAccDTO *domain.NewAccountDTO
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{"normal mngr", fields_, args{nil, &acc, &newAcc}, false},
		{"normal clnt", fields_, args{nil, &acc2, &newAcc2}, false},
		{"no login", fields_, args{nil, &acc3, &newAcc2}, true},
		{"bad role", fields_, args{nil, &acc4, &newAcc3}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			as := services.CreateNewAccountService(mockAccountRepo, clientService, managerService)
			if err := as.Register(tt.args.c, tt.args.acc, tt.args.newAccDTO); (err != nil) != tt.wantErr {
				t.Errorf("accountService.Register() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

//func Test_accountService_Login(t *testing.T) {
//	ctrlA := gomock.NewController(t)
//	ctrlC := gomock.NewController(t)
//	ctrlM := gomock.NewController(t)
//
//	mockAccountRepo := mock_domain.NewMockIAccountRepo(ctrlA)
//	mockClientRepo := mock_domain.NewMockIClientRepo(ctrlC)
//	mockManagerRepo := mock_domain.NewMockIManagerRepo(ctrlM)
//
//	clientService := services.CreateNewClientService(mockClientRepo)
//	managerService := services.CreateNewManagerService(mockManagerRepo)
//
//	acc := domain.Account{
//		ID:       1,
//		UserID:   1,
//		Login:    "login",
//		Password: "password",
//	}
//
//	type fields struct {
//		AccountRepo    mock_domain.MockIAccountRepo
//		ClientService  mock_domain.MockIClientRepo
//		ManagerService mock_domain.MockIManagerRepo
//	}
//	fields_ := fields{
//		AccountRepo:    *mockAccountRepo,
//		ClientService:  *mockClientRepo,
//		ManagerService: *mockManagerRepo,
//	}
//	type args struct {
//		c   context.Context
//		acc *domain.Account
//	}
//	tests := []struct {
//		name    string
//		fields  fields
//		args    args
//		want    string
//		wantErr bool
//	}{
//		{"normal", fields_, args{nil, &acc}, }
//	},
//		for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			as := services.CreateNewAccountService(mockAccountRepo, clientService, managerService)
//			got, err := as.Login(tt.args.c, tt.args.acc)
//			if (err != nil) != tt.wantErr {
//				t.Errorf("accountService.Login() error = %v, wantErr %v", err, tt.wantErr)
//				return
//			}
//			if got != tt.want {
//				t.Errorf("accountService.Login() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}

//func Test_accountService_Update(t *testing.T) {
//	type fields struct {
//		AccountRepo    domain.IAccountRepo
//		ClientService  domain.IClientService
//		ManagerService domain.IManagerService
//	}
//	type args struct {
//		c        context.Context
//		newState *domain.Account
//	}
//	tests := []struct {
//		name    string
//		fields  fields
//		args    args
//		wantErr bool
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			as := &accountService{
//				AccountRepo:    tt.fields.AccountRepo,
//				ClientService:  tt.fields.ClientService,
//				ManagerService: tt.fields.ManagerService,
//			}
//			if err := as.Update(tt.args.c, tt.args.newState); (err != nil) != tt.wantErr {
//				t.Errorf("accountService.Update() error = %v, wantErr %v", err, tt.wantErr)
//			}
//		})
//	}
//}

//func Test_accountService_GetByLogin(t *testing.T) {
//	type fields struct {
//		AccountRepo    domain.IAccountRepo
//		ClientService  domain.IClientService
//		ManagerService domain.IManagerService
//	}
//	type args struct {
//		c     context.Context
//		login string
//	}
//	tests := []struct {
//		name    string
//		fields  fields
//		args    args
//		want    domain.Account
//		wantErr bool
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			as := &accountService{
//				AccountRepo:    tt.fields.AccountRepo,
//				ClientService:  tt.fields.ClientService,
//				ManagerService: tt.fields.ManagerService,
//			}
//			got, err := as.GetByLogin(tt.args.c, tt.args.login)
//			if (err != nil) != tt.wantErr {
//				t.Errorf("accountService.GetByLogin() error = %v, wantErr %v", err, tt.wantErr)
//				return
//			}
//			if !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("accountService.GetByLogin() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}

//func Test_accountService_GetClientByUserId(t *testing.T) {
//	type fields struct {
//		AccountRepo    domain.IAccountRepo
//		ClientService  domain.IClientService
//		ManagerService domain.IManagerService
//	}
//	type args struct {
//		c      context.Context
//		userID int
//	}
//	tests := []struct {
//		name    string
//		fields  fields
//		args    args
//		want    domain.Client
//		wantErr bool
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			as := &accountService{
//				AccountRepo:    tt.fields.AccountRepo,
//				ClientService:  tt.fields.ClientService,
//				ManagerService: tt.fields.ManagerService,
//			}
//			got, err := as.GetClientByUserId(tt.args.c, tt.args.userID)
//			if (err != nil) != tt.wantErr {
//				t.Errorf("accountService.GetClientByUserId() error = %v, wantErr %v", err, tt.wantErr)
//				return
//			}
//			if !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("accountService.GetClientByUserId() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
//
//func Test_accountService_GetManagerByUserId(t *testing.T) {
//	type fields struct {
//		AccountRepo    domain.IAccountRepo
//		ClientService  domain.IClientService
//		ManagerService domain.IManagerService
//	}
//	type args struct {
//		c      context.Context
//		userID int
//	}
//	tests := []struct {
//		name    string
//		fields  fields
//		args    args
//		want    domain.Manager
//		wantErr bool
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			as := &accountService{
//				AccountRepo:    tt.fields.AccountRepo,
//				ClientService:  tt.fields.ClientService,
//				ManagerService: tt.fields.ManagerService,
//			}
//			got, err := as.GetManagerByUserId(tt.args.c, tt.args.userID)
//			if (err != nil) != tt.wantErr {
//				t.Errorf("accountService.GetManagerByUserId() error = %v, wantErr %v", err, tt.wantErr)
//				return
//			}
//			if !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("accountService.GetManagerByUserId() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
//
//func Test_accountService_AddClient(t *testing.T) {
//	type fields struct {
//		AccountRepo    domain.IAccountRepo
//		ClientService  domain.IClientService
//		ManagerService domain.IManagerService
//	}
//	type args struct {
//		c         context.Context
//		acc       *domain.Account
//		newAccDTO *domain.NewAccountDTO
//	}
//	tests := []struct {
//		name    string
//		fields  fields
//		args    args
//		wantErr bool
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			as := &accountService{
//				AccountRepo:    tt.fields.AccountRepo,
//				ClientService:  tt.fields.ClientService,
//				ManagerService: tt.fields.ManagerService,
//			}
//			if err := as.AddClient(tt.args.c, tt.args.acc, tt.args.newAccDTO); (err != nil) != tt.wantErr {
//				t.Errorf("accountService.AddClient() error = %v, wantErr %v", err, tt.wantErr)
//			}
//		})
//	}
//}
//
//func Test_accountService_AddManager(t *testing.T) {
//	type fields struct {
//		AccountRepo    domain.IAccountRepo
//		ClientService  domain.IClientService
//		ManagerService domain.IManagerService
//	}
//	type args struct {
//		c         context.Context
//		acc       *domain.Account
//		newAccDTO *domain.NewAccountDTO
//	}
//	tests := []struct {
//		name    string
//		fields  fields
//		args    args
//		wantErr bool
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			as := &accountService{
//				AccountRepo:    tt.fields.AccountRepo,
//				ClientService:  tt.fields.ClientService,
//				ManagerService: tt.fields.ManagerService,
//			}
//			if err := as.AddManager(tt.args.c, tt.args.acc, tt.args.newAccDTO); (err != nil) != tt.wantErr {
//				t.Errorf("accountService.AddManager() error = %v, wantErr %v", err, tt.wantErr)
//			}
//		})
//	}
//}
