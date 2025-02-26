package services

import (
	"log/slog"
	"test_task/models"
	"test_task/storage"
)

type ServiceKindService struct {
	store  storage.IStore
	logger *slog.Logger
}

func NewServiceKindService(store storage.IStore, logger *slog.Logger) *ServiceKindService {
	return &ServiceKindService{
		store:  store,
		logger: logger,
	}
}

func (s *ServiceKindService) Create(serviceKind *models.ServiceKind) (string, error) {
	id, err := s.store.ServiceKind().Create(serviceKind)
	if err != nil {
		s.logger.Error("Error while creating service kind")
		return "", err
	}

	return id, nil
}

func (s *ServiceKindService) Update(serviceKind *models.ServiceKind) error {
	err := s.store.ServiceKind().Update(serviceKind)
	if err != nil {
		s.logger.Error("Error while updating service kind")
		return err
	}

	return nil
}

func (s *ServiceKindService) Delete(req models.RequestId) error {
	err := s.store.ServiceKind().Delete(req)
	if err != nil {
		s.logger.Error("Error while deleting service kind")
		return err
	}

	return nil
}

func (s *ServiceKindService) Get(req models.RequestId) (*models.ServiceKind, error) {
	serviceKind, err := s.store.ServiceKind().Get(req)
	if err != nil {
		s.logger.Error("Error while getting service kind")
		return nil, err
	}

	return serviceKind, nil
}

func (s *ServiceKindService) GetAll() ([]models.ServiceKind, error) {
	serviceKinds, err := s.store.ServiceKind().GetAll()
	if err != nil {
		s.logger.Error("Error while getting all service kinds")
		return nil, err
	}

	return serviceKinds, nil
}
