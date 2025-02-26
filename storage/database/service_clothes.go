package database

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"test_task/models"
)

type ServiceClothesRepo struct {
	db *gorm.DB
}

func NewServiceClothesRepo(db *gorm.DB) ServiceClothes {
	return &ServiceClothesRepo{
		db: db,
	}
}

func (s ServiceClothesRepo) Create(serviceClothes *models.ServiceClothes) (string, error) {
	id := uuid.New()
	serviceClothes.Id = id.String()

	err := s.db.Create(serviceClothes).Error
	if err != nil {
		return "", err
	}

	return serviceClothes.Id, nil
}

func (s ServiceClothesRepo) Update(serviceClothes *models.ServiceClothes) error {
	err := s.db.Save(serviceClothes).Error
	if err != nil {
		return err
	}

	return nil
}

func (s ServiceClothesRepo) Delete(req models.RequestId) error {
	err := s.db.Delete(&models.ServiceClothes{}, req.Id).Error
	if err != nil {
		return err
	}

	return nil
}

func (s ServiceClothesRepo) Get(req models.RequestId) (*models.ServiceClothes, error) {
	var serviceClothes models.ServiceClothes
	err := s.db.First(&serviceClothes, req.Id).Error
	if err != nil {
		return nil, err
	}

	return &serviceClothes, nil
}

func (s ServiceClothesRepo) GetAll(req models.GetAllServiceClothesReq) ([]models.ServiceClothes, error) {
	var serviceClothes []models.ServiceClothes
	query := s.db

	if req.ServiceKindId != "" {
		query = query.Where("service_kind_id = ?", req.ServiceKindId)
	}

	err := query.Find(&serviceClothes).Error
	if err != nil {
		return nil, err
	}

	return serviceClothes, nil
}
