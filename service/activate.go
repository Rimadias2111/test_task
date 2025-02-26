package service

import (
	"log/slog"
	"test_task/service/services"
	"test_task/storage"
)

func New(store storage.IStore, logger *slog.Logger) IService {
	return &Service{
		serviceKindService:    *services.NewServiceKindService(store, logger),
		serviceClothesService: *services.NewServiceClothesService(store, logger),
		orderService:          *services.NewOrderService(store, logger),
		orderItemService:      *services.NewOrderItemService(store, logger),
	}
}
