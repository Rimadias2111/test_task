package services

import (
	"log/slog"
	"test_task/models"
	"test_task/storage"
)

type ServiceClothesService struct {
	store  storage.IStore
	logger *slog.Logger
}

func NewServiceClothesService(store storage.IStore, logger *slog.Logger) *ServiceClothesService {
	return &ServiceClothesService{
		store:  store,
		logger: logger,
	}
}

func (s *ServiceClothesService) Create(serviceClothes *models.ServiceClothes) (string, error) {
	id, err := s.store.ServiceClothes().Create(serviceClothes)
	if err != nil {
		s.logger.Error("Error while creating service clothes")
		return "", err
	}

	return id, nil
}

func (s *ServiceClothesService) Update(serviceClothes *models.ServiceClothes) error {
	err := s.store.ServiceClothes().Update(serviceClothes)
	if err != nil {
		s.logger.Error("Error while updating service clothes")
		return err
	}

	return nil
}

func (s *ServiceClothesService) Delete(req models.RequestId) error {
	err := s.store.ServiceClothes().Delete(req)
	if err != nil {
		s.logger.Error("Error while deleting service clothes")
		return err
	}

	return nil
}

func (s *ServiceClothesService) Get(req models.RequestId) (*models.ServiceClothes, error) {
	serviceClothes, err := s.store.ServiceClothes().Get(req)
	if err != nil {
		s.logger.Error("Error while getting service clothes")
		return nil, err
	}

	return serviceClothes, nil
}

func (s *ServiceClothesService) GetAll(req models.GetAllServiceClothesReq) ([]models.ServiceClothes, error) {
	serviceClothes, err := s.store.ServiceClothes().GetAll(req)
	if err != nil {
		s.logger.Error("Error while getting all service clothes")
		return nil, err
	}

	return serviceClothes, nil
}
