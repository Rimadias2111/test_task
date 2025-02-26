package services

import (
	"log/slog"
	"test_task/models"
	"test_task/storage"
)

type OrderItemService struct {
	store  storage.IStore
	logger *slog.Logger
}

func NewOrderItemService(store storage.IStore, logger *slog.Logger) *OrderItemService {
	return &OrderItemService{
		store:  store,
		logger: logger,
	}
}

func (s *OrderItemService) Create(orderItem *models.OrderItem) (string, error) {
	id, err := s.store.OrderItem().Create(orderItem)
	if err != nil {
		s.logger.Error("Error while creating order item")
		return "", err
	}

	return id, nil
}

func (s *OrderItemService) Update(orderItem *models.OrderItem) error {
	err := s.store.OrderItem().Update(orderItem)
	if err != nil {
		s.logger.Error("Error while updating order item")
		return err
	}

	return nil
}

func (s *OrderItemService) Delete(req models.RequestId) error {
	err := s.store.OrderItem().Delete(req)
	if err != nil {
		s.logger.Error("Error while deleting order item")
		return err
	}

	return nil
}

func (s *OrderItemService) Get(req models.RequestId) (*models.OrderItem, error) {
	orderItem, err := s.store.OrderItem().Get(req)
	if err != nil {
		s.logger.Error("Error while getting order item")
		return nil, err
	}

	return orderItem, nil
}

func (s *OrderItemService) GetAll() ([]models.OrderItem, error) {
	orderItems, err := s.store.OrderItem().GetAll()
	if err != nil {
		s.logger.Error("Error while getting all order items")
		return nil, err
	}

	return orderItems, nil
}
