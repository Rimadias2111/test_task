package storage

import (
	"gorm.io/gorm"
	"test_task/storage/database"
)

func New(db *gorm.DB) IStore {
	return &Store{
		db:             db,
		serviceKind:    database.NewServiceKindRepo(db),
		serviceClothes: database.NewServiceClothesRepo(db),
		order:          database.NewOrderRepo(db),
		orderItem:      database.NewOrderItemRepo(db),
	}
}
