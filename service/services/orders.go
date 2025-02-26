package services

import (
	"log/slog"
	"test_task/models"
	"test_task/storage"
)

type OrderService struct {
	store  storage.IStore
	logger *slog.Logger
}

func NewOrderService(store storage.IStore, logger *slog.Logger) *OrderService {
	return &OrderService{
		store:  store,
		logger: logger,
	}
}

func (s *OrderService) Create(order *models.Order) (string, error) {
	id, err := s.store.Order().Create(order)
	if err != nil {
		s.logger.Error("Error while creating order")
		return "", err
	}

	return id, nil
}

func (s *OrderService) Update(order *models.Order) error {
	err := s.store.Order().Update(order)
	if err != nil {
		s.logger.Error("Error while updating order")
		return err
	}

	return nil
}

func (s *OrderService) Delete(req models.RequestId) error {
	err := s.store.Order().Delete(req)
	if err != nil {
		s.logger.Error("Error while deleting order")
		return err
	}

	return nil
}

func (s *OrderService) Get(req models.RequestId) (*models.Order, error) {
	order, err := s.store.Order().Get(req)
	if err != nil {
		s.logger.Error("Error while getting order")
		return nil, err
	}

	return order, nil
}

func (s *OrderService) GetAll() ([]models.Order, error) {
	orders, err := s.store.Order().GetAll()
	if err != nil {
		s.logger.Error("Error while getting all orders")
		return nil, err
	}

	return orders, nil
}
