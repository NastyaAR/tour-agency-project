package unit

//
//import (
//	"app/domain"
//	"context"
//	"reflect"
//	"testing"
//)
//
//func TestCreateNewManagerService(t *testing.T) {
//	type args struct {
//		ManagerRepo domain.IManagerRepo
//	}
//	tests := []struct {
//		name string
//		args args
//		want domain.IManagerService
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if got := CreateNewManagerService(tt.args.ManagerRepo); !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("CreateNewManagerService() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
//
//func Test_managerService_Create(t *testing.T) {
//	type fields struct {
//		ManagerRepo domain.IManagerRepo
//	}
//	type args struct {
//		c    context.Context
//		mngr *domain.Manager
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
//			ms := &managerService{
//				ManagerRepo: tt.fields.ManagerRepo,
//			}
//			got, err := ms.Create(tt.args.c, tt.args.mngr)
//			if (err != nil) != tt.wantErr {
//				t.Errorf("managerService.Create() error = %v, wantErr %v", err, tt.wantErr)
//				return
//			}
//			if !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("managerService.Create() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
//
//func Test_managerService_Delete(t *testing.T) {
//	type fields struct {
//		ManagerRepo domain.IManagerRepo
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
//			ms := &managerService{
//				ManagerRepo: tt.fields.ManagerRepo,
//			}
//			if err := ms.Delete(tt.args.c, tt.args.id); (err != nil) != tt.wantErr {
//				t.Errorf("managerService.Delete() error = %v, wantErr %v", err, tt.wantErr)
//			}
//		})
//	}
//}
//
//func Test_managerService_GetByNameSurname(t *testing.T) {
//	type fields struct {
//		ManagerRepo domain.IManagerRepo
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
//		want    []domain.Manager
//		wantErr bool
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			ms := &managerService{
//				ManagerRepo: tt.fields.ManagerRepo,
//			}
//			got, err := ms.GetByNameSurname(tt.args.c, tt.args.name, tt.args.surname)
//			if (err != nil) != tt.wantErr {
//				t.Errorf("managerService.GetByNameSurname() error = %v, wantErr %v", err, tt.wantErr)
//				return
//			}
//			if !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("managerService.GetByNameSurname() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
//
//func Test_managerService_GetByDepartment(t *testing.T) {
//	type fields struct {
//		ManagerRepo domain.IManagerRepo
//	}
//	type args struct {
//		c          context.Context
//		department string
//	}
//	tests := []struct {
//		name    string
//		fields  fields
//		args    args
//		want    []domain.Manager
//		wantErr bool
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			ms := &managerService{
//				ManagerRepo: tt.fields.ManagerRepo,
//			}
//			got, err := ms.GetByDepartment(tt.args.c, tt.args.department)
//			if (err != nil) != tt.wantErr {
//				t.Errorf("managerService.GetByDepartment() error = %v, wantErr %v", err, tt.wantErr)
//				return
//			}
//			if !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("managerService.GetByDepartment() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
//
//func Test_managerService_Update(t *testing.T) {
//	type fields struct {
//		ManagerRepo domain.IManagerRepo
//	}
//	type args struct {
//		c        context.Context
//		newState *domain.Manager
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
//			ms := &managerService{
//				ManagerRepo: tt.fields.ManagerRepo,
//			}
//			if err := ms.Update(tt.args.c, tt.args.newState); (err != nil) != tt.wantErr {
//				t.Errorf("managerService.Update() error = %v, wantErr %v", err, tt.wantErr)
//			}
//		})
//	}
//}
//
//func Test_managerService_GetStatisticsForManagers(t *testing.T) {
//	type fields struct {
//		ManagerRepo domain.IManagerRepo
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
//		want    []domain.Statistics
//		wantErr bool
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			ms := &managerService{
//				ManagerRepo: tt.fields.ManagerRepo,
//			}
//			got, err := ms.GetStatisticsForManagers(tt.args.c, tt.args.from, tt.args.to)
//			if (err != nil) != tt.wantErr {
//				t.Errorf("managerService.GetStatisticsForManagers() error = %v, wantErr %v", err, tt.wantErr)
//				return
//			}
//			if !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("managerService.GetStatisticsForManagers() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
//
//func Test_managerService_GetStatisticsForManager(t *testing.T) {
//	type fields struct {
//		ManagerRepo domain.IManagerRepo
//	}
//	type args struct {
//		c    context.Context
//		id   int
//		from string
//		to   string
//	}
//	tests := []struct {
//		name    string
//		fields  fields
//		args    args
//		want    domain.Statistics
//		wantErr bool
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			ms := &managerService{
//				ManagerRepo: tt.fields.ManagerRepo,
//			}
//			got, err := ms.GetStatisticsForManager(tt.args.c, tt.args.id, tt.args.from, tt.args.to)
//			if (err != nil) != tt.wantErr {
//				t.Errorf("managerService.GetStatisticsForManager() error = %v, wantErr %v", err, tt.wantErr)
//				return
//			}
//			if !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("managerService.GetStatisticsForManager() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
