package database

import "test_task/models"

type ServiceKind interface {
	Create(serviceKind *models.ServiceKind) (string, error)
	Update(serviceKind *models.ServiceKind) error
	Delete(req models.RequestId) error
	Get(req models.RequestId) (*models.ServiceKind, error)
	GetAll() ([]models.ServiceKind, error)
}

type ServiceClothes interface {
	Create(serviceClothes *models.ServiceClothes) (string, error)
	Update(serviceClothes *models.ServiceClothes) error
	Delete(req models.RequestId) error
	Get(req models.RequestId) (*models.ServiceClothes, error)
	GetAll(req models.GetAllServiceClothesReq) ([]models.ServiceClothes, error)
}

type Order interface {
	Create(order *models.Order) (string, error)
	Update(order *models.Order) error
	Delete(req models.RequestId) error
	Get(req models.RequestId) (*models.Order, error)
	GetAll() ([]models.Order, error)
}

type OrderItem interface {
	Create(orderItem *models.OrderItem) (string, error)
	Update(orderItem *models.OrderItem) error
	Delete(req models.RequestId) error
	Get(req models.RequestId) (*models.OrderItem, error)
	GetAll() ([]models.OrderItem, error)
}
