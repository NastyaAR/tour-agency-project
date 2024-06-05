package unit

import (
	"app/domain"
	"app/services"
	mock_domain "app/unit/mocks"
	"go.uber.org/mock/gomock"
	"testing"
)

func TestCreateNormal(t *testing.T) {
	ctrl := gomock.NewController(t)

	mockTourRepo := mock_domain.NewMockITourRepo(ctrl)

	tour := domain.Tour{
		ID:             1,
		ChillPlace:     "Крым",
		FromPlace:      "Москва",
		Date:           "2024.06.30",
		Duration:       7,
		Cost:           100000,
		TouristsNumber: 2,
		ChillType:      "пляжный",
		TransferOn:     true,
	}
	mockTourRepo.EXPECT().Add(
		nil,
		&tour,
	).Return(nil)

	tourService := services.NewTourService(mockTourRepo)
	ret := tourService.Create(nil, &tour)
	if ret != nil {
		t.Fatalf("Возникла ошибка при создании тура: %s", ret)
	}
}

func TestCreateBadCost(t *testing.T) {
	ctrl := gomock.NewController(t)

	mockTourRepo := mock_domain.NewMockITourRepo(ctrl)

	tour := domain.Tour{
		ID:             1,
		ChillPlace:     "Крым",
		FromPlace:      "Москва",
		Date:           "2024.06.30",
		Duration:       7,
		Cost:           -100000,
		TouristsNumber: 2,
		ChillType:      "пляжный",
		TransferOn:     true,
	}

	tourService := services.NewTourService(mockTourRepo)
	ret := tourService.Create(nil, &tour)
	if ret == nil {
		t.Fail()
	}
}

func TestCreateBadTouristsNumber(t *testing.T) {
	ctrl := gomock.NewController(t)

	mockTourRepo := mock_domain.NewMockITourRepo(ctrl)

	tour := domain.Tour{
		ID:             1,
		ChillPlace:     "Крым",
		FromPlace:      "Москва",
		Date:           "2024.06.30",
		Duration:       7,
		Cost:           100000,
		TouristsNumber: 0,
		ChillType:      "пляжный",
		TransferOn:     true,
	}

	tourService := services.NewTourService(mockTourRepo)
	ret := tourService.Create(nil, &tour)
	if ret == nil {
		t.Fail()
	}
}

func TestGetByIdNormal(t *testing.T) {
	ctrl := gomock.NewController(t)

	mockTourRepo := mock_domain.NewMockITourRepo(ctrl)

	tour := domain.Tour{
		ID:             1,
		ChillPlace:     "Крым",
		FromPlace:      "Москва",
		Date:           "2024.06.30",
		Duration:       7,
		Cost:           100000,
		TouristsNumber: 2,
		ChillType:      "пляжный",
		TransferOn:     true,
	}
	mockTourRepo.EXPECT().GetById(
		nil,
		1,
	).Return(tour, nil)

	tourService := services.NewTourService(mockTourRepo)
	_, err := tourService.GetById(nil, 1)
	if err != nil {
		t.Fatalf("Возникла ошибка при поиске тура по id: %s", err)
	}
}

func TestGetByIdBadId(t *testing.T) {
	ctrl := gomock.NewController(t)

	mockTourRepo := mock_domain.NewMockITourRepo(ctrl)

	tourService := services.NewTourService(mockTourRepo)
	_, err := tourService.GetById(nil, -1)
	if err == nil {
		t.Fatalf("Возникла ошибка при поиске тура по id: %s", err)
	}
}

func TestUpdateNormal(t *testing.T) {
	ctrl := gomock.NewController(t)

	mockTourRepo := mock_domain.NewMockITourRepo(ctrl)
	tour := domain.Tour{
		ID:             1,
		ChillPlace:     "Крым",
		FromPlace:      "Москва",
		Date:           "2024.06.30",
		Duration:       7,
		Cost:           100000,
		TouristsNumber: 2,
		ChillType:      "пляжный",
		TransferOn:     true,
	}
	mockTourRepo.EXPECT().Update(nil, 1, &tour).Return(nil)
	tourService := services.NewTourService(mockTourRepo)
	err := tourService.Update(nil, 1, &tour)
	if err != nil {
		t.Fatalf("Возникла ошибка при обновлении тура по id: %s", err)
	}
}

func TestUpdateNilNew(t *testing.T) {
	ctrl := gomock.NewController(t)

	mockTourRepo := mock_domain.NewMockITourRepo(ctrl)
	tourService := services.NewTourService(mockTourRepo)
	err := tourService.Update(nil, 1, nil)
	if err == nil {
		t.Fatalf("Возникла ошибка при обновлении тура по id: %s", err)
	}
}

func TestSetSaleNormal(t *testing.T) {
	ctrl := gomock.NewController(t)

	mockTourRepo := mock_domain.NewMockITourRepo(ctrl)
	newSale := domain.Sale{
		ID:          1,
		Name:        "Скидка-скидка",
		ExpiredTime: "2024.05.05",
		Percent:     10,
	}
	mockTourRepo.EXPECT().UpdateSale(nil, 1, &newSale)
	tourService := services.NewTourService(mockTourRepo)
	err := tourService.SetSale(nil, 1, &newSale)
	if err != nil {
		t.Fatalf("Возникла ошибка при добавлении скидки на тур: %s", err)
	}
}

func TestSetSaleBadPercentSmall(t *testing.T) {
	ctrl := gomock.NewController(t)

	mockTourRepo := mock_domain.NewMockITourRepo(ctrl)
	newSale := domain.Sale{
		ID:          1,
		Name:        "Скидка-скидка",
		ExpiredTime: "2024.05.05",
		Percent:     -10,
	}
	tourService := services.NewTourService(mockTourRepo)
	err := tourService.SetSale(nil, 1, &newSale)
	if err == nil {
		t.Fatalf("Возникла ошибка при добавлении скидки на тур: %s", err)
	}
}

func TestSetSaleBadPercentBig(t *testing.T) {
	ctrl := gomock.NewController(t)

	mockTourRepo := mock_domain.NewMockITourRepo(ctrl)
	newSale := domain.Sale{
		ID:          1,
		Name:        "Скидка-скидка",
		ExpiredTime: "2024.05.05",
		Percent:     101,
	}
	tourService := services.NewTourService(mockTourRepo)
	err := tourService.SetSale(nil, 1, &newSale)
	if err == nil {
		t.Fatalf("Возникла ошибка при добавлении скидки на тур: %s", err)
	}
}

func TestDeleteNormal(t *testing.T) {
	ctrl := gomock.NewController(t)

	mockTourRepo := mock_domain.NewMockITourRepo(ctrl)
	mockTourRepo.EXPECT().Delete(nil, 1).Return(nil)
	tourService := services.NewTourService(mockTourRepo)
	err := tourService.Delete(nil, 1)
	if err != nil {
		t.Fatalf("Возникла ошибка при удалении тура: %s", err)
	}
}

func TestGetByCriteriaNormal(t *testing.T) {
	ctrl := gomock.NewController(t)

	mockTourRepo := mock_domain.NewMockITourRepo(ctrl)
	tour := domain.Tour{
		ID:             1,
		ChillPlace:     "Крым",
		FromPlace:      "Москва",
		Date:           "2024.06.30",
		Duration:       7,
		Cost:           100000,
		TouristsNumber: 2,
		ChillType:      "",
		TransferOn:     true,
	}
	tours := make([]domain.Tour, 0)
	tours = append(tours, tour)
	mockTourRepo.EXPECT().GetByCriteria(nil, &tour).Return(tours, nil)
	tourService := services.NewTourService(mockTourRepo)
	_, err := tourService.GetByCriteria(nil, &tour)
	if err != nil {
		t.Fatalf("Возникла ошибка при удалении тура: %s", err)
	}
}
