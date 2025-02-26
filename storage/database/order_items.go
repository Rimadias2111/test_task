package database

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"test_task/models"
)

type OrderItemRepo struct {
	db *gorm.DB
}

func NewOrderItemRepo(db *gorm.DB) OrderItem {
	return &OrderItemRepo{
		db: db,
	}
}

func (o OrderItemRepo) Create(orderItem *models.OrderItem) (string, error) {
	if orderItem.Id == "" {
		id := uuid.New()
		orderItem.Id = id.String()
	}

	err := o.db.Create(orderItem).Error
	if err != nil {
		return "", err
	}

	return orderItem.Id, nil
}

func (o OrderItemRepo) Update(orderItem *models.OrderItem) error {
	err := o.db.Save(orderItem).Error
	if err != nil {
		return err
	}

	return nil
}

func (o OrderItemRepo) Delete(req models.RequestId) error {
	err := o.db.Delete(&models.OrderItem{}, req.Id).Error
	if err != nil {
		return err
	}

	return nil
}

func (o OrderItemRepo) Get(req models.RequestId) (*models.OrderItem, error) {
	var orderItem models.OrderItem
	err := o.db.First(&orderItem, req.Id).Error
	if err != nil {
		return nil, err
	}

	return &orderItem, nil
}

func (o OrderItemRepo) GetAll() ([]models.OrderItem, error) {
	var orderItems []models.OrderItem
	err := o.db.Find(&orderItems).Error
	if err != nil {
		return nil, err
	}

	return orderItems, nil
}
