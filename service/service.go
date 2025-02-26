package service

import "test_task/service/services"

type IService interface {
	ServiceKind() *services.ServiceKindService
	ServiceClothes() *services.ServiceClothesService
	Order() *services.OrderService
	OrderItem() *services.OrderItemService
}

type Service struct {
	serviceKindService    services.ServiceKindService
	serviceClothesService services.ServiceClothesService
	orderService          services.OrderService
	orderItemService      services.OrderItemService
}

func (s *Service) ServiceKind() *services.ServiceKindService {
	return &s.serviceKindService
}

func (s *Service) ServiceClothes() *services.ServiceClothesService {
	return &s.serviceClothesService
}

func (s *Service) Order() *services.OrderService {
	return &s.orderService
}

func (s *Service) OrderItem() *services.OrderItemService {
	return &s.orderItemService
}
