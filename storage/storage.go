package storage

import (
	"gorm.io/gorm"
	"test_task/storage/database"
)

type IStore interface {
	DB() *gorm.DB
	ServiceKind() database.ServiceKind
	ServiceClothes() database.ServiceClothes
	Order() database.Order
	OrderItem() database.OrderItem
}

type Store struct {
	db             *gorm.DB
	serviceKind    database.ServiceKind
	serviceClothes database.ServiceClothes
	order          database.Order
	orderItem      database.OrderItem
}

func (s *Store) DB() *gorm.DB { return s.db }

func (s *Store) ServiceKind() database.ServiceKind { return s.serviceKind }

func (s *Store) ServiceClothes() database.ServiceClothes { return s.serviceClothes }

func (s *Store) Order() database.Order { return s.order }

func (s *Store) OrderItem() database.OrderItem { return s.orderItem }
