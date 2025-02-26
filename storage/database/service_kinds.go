package database

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"test_task/models"
)

type ServiceKindRepo struct {
	db *gorm.DB
}

func NewServiceKindRepo(db *gorm.DB) ServiceKind {
	return &ServiceKindRepo{
		db: db,
	}
}

func (s ServiceKindRepo) Create(serviceKind *models.ServiceKind) (string, error) {
	id := uuid.New()
	serviceKind.Id = id.String()

	err := s.db.Create(serviceKind).Error
	if err != nil {
		return "", err
	}

	return serviceKind.Id, nil
}

func (s ServiceKindRepo) Update(serviceKind *models.ServiceKind) error {
	err := s.db.Save(serviceKind).Error
	if err != nil {
		return err
	}

	return nil
}

func (s ServiceKindRepo) Delete(req models.RequestId) error {
	err := s.db.Delete(&models.ServiceKind{}, req.Id).Error
	if err != nil {
		return err
	}

	return nil
}

func (s ServiceKindRepo) Get(req models.RequestId) (*models.ServiceKind, error) {
	var serviceKind models.ServiceKind
	err := s.db.First(&serviceKind, req.Id).Error
	if err != nil {
		return nil, err
	}

	return &serviceKind, nil
}

func (s ServiceKindRepo) GetAll() ([]models.ServiceKind, error) {
	var serviceKinds []models.ServiceKind
	err := s.db.Find(&serviceKinds).Error
	if err != nil {
		return nil, err
	}

	return serviceKinds, nil
}
