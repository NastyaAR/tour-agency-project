package unit

//
//import (
//	"app/domain"
//	"context"
//	"reflect"
//	"testing"
//)
//
//func TestCreateNewClientService(t *testing.T) {
//	type args struct {
//		ClientRepo domain.IClientRepo
//	}
//	tests := []struct {
//		name string
//		args args
//		want domain.IClientService
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if got := CreateNewClientService(tt.args.ClientRepo); !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("CreateNewClientService() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
//
//func Test_clientService_Create(t *testing.T) {
//	type fields struct {
//		ClientRepo domain.IClientRepo
//	}
//	type args struct {
//		c    context.Context
//		clnt *domain.Client
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
//			cs := &clientService{
//				ClientRepo: tt.fields.ClientRepo,
//			}
//			got, err := cs.Create(tt.args.c, tt.args.clnt)
//			if (err != nil) != tt.wantErr {
//				t.Errorf("clientService.Create() error = %v, wantErr %v", err, tt.wantErr)
//				return
//			}
//			if !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("clientService.Create() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
//
//func Test_clientService_Delete(t *testing.T) {
//	type fields struct {
//		ClientRepo domain.IClientRepo
//	}
//	type args struct {
//		c  context.Context
//		id int
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
//			cs := &clientService{
//				ClientRepo: tt.fields.ClientRepo,
//			}
//			if err := cs.Delete(tt.args.c, tt.args.id); (err != nil) != tt.wantErr {
//				t.Errorf("clientService.Delete() error = %v, wantErr %v", err, tt.wantErr)
//			}
//		})
//	}
//}
//
//func Test_clientService_GetByNameSurname(t *testing.T) {
//	type fields struct {
//		ClientRepo domain.IClientRepo
//	}
//	type args struct {
//		c       context.Context
//		name    string
//		surname string
//	}
//	tests := []struct {
//		name    string
//		fields  fields
//		args    args
//		want    []domain.Client
//		wantErr bool
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			cs := &clientService{
//				ClientRepo: tt.fields.ClientRepo,
//			}
//			got, err := cs.GetByNameSurname(tt.args.c, tt.args.name, tt.args.surname)
//			if (err != nil) != tt.wantErr {
//				t.Errorf("clientService.GetByNameSurname() error = %v, wantErr %v", err, tt.wantErr)
//				return
//			}
//			if !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("clientService.GetByNameSurname() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
//
//func Test_clientService_GetByPhone(t *testing.T) {
//	type fields struct {
//		ClientRepo domain.IClientRepo
//	}
//	type args struct {
//		c     context.Context
//		phone string
//	}
//	tests := []struct {
//		name    string
//		fields  fields
//		args    args
//		want    []domain.Client
//		wantErr bool
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			cs := &clientService{
//				ClientRepo: tt.fields.ClientRepo,
//			}
//			got, err := cs.GetByPhone(tt.args.c, tt.args.phone)
//			if (err != nil) != tt.wantErr {
//				t.Errorf("clientService.GetByPhone() error = %v, wantErr %v", err, tt.wantErr)
//				return
//			}
//			if !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("clientService.GetByPhone() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
//
//func Test_clientService_Update(t *testing.T) {
//	type fields struct {
//		ClientRepo domain.IClientRepo
//	}
//	type args struct {
//		c        context.Context
//		newState *domain.Client
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
//			cs := &clientService{
//				ClientRepo: tt.fields.ClientRepo,
//			}
//			if err := cs.Update(tt.args.c, tt.args.newState); (err != nil) != tt.wantErr {
//				t.Errorf("clientService.Update() error = %v, wantErr %v", err, tt.wantErr)
//			}
//		})
//	}
//}
//
//func Test_clientService_GetStoryOfRequests(t *testing.T) {
//	type fields struct {
//		ClientRepo domain.IClientRepo
//	}
//	type args struct {
//		c  context.Context
//		id int
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
//			cs := &clientService{
//				ClientRepo: tt.fields.ClientRepo,
//			}
//			got, err := cs.GetStoryOfRequests(tt.args.c, tt.args.id)
//			if (err != nil) != tt.wantErr {
//				t.Errorf("clientService.GetStoryOfRequests() error = %v, wantErr %v", err, tt.wantErr)
//				return
//			}
//			if !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("clientService.GetStoryOfRequests() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
