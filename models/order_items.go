package models

import (
	"time"
)

type OrderItem struct {
	Id               string    `gorm:"type:char(36);primaryKey" json:"id"`
	OrderId          string    `gorm:"type:char(36);not null;index" json:"order_id"`
	ServiceKindId    string    `gorm:"type:char(36);not null;index" json:"service_kind_id"`
	ServiceClothesId string    `gorm:"type:char(36);not null;index" json:"service_clothes_id"`
	Quantity         *int      `json:"quantity,omitempty"`
	Weight           *float64  `json:"weight,omitempty"`
	Price            float64   `gorm:"not null" json:"price"`
	CreatedAt        time.Time `gorm:"autoCreateTime" json:"created_at"`

	Order          Order          `gorm:"foreignKey:OrderId" json:"order"`
	ServiceKind    ServiceKind    `gorm:"foreignKey:ServiceKindId" json:"service_kind"`
	ServiceClothes ServiceClothes `gorm:"foreignKey:ServiceClothesId" json:"service_clothes"`

	ServiceKindName    string `json:"service_kind_name" gorm:"-"`
	ServiceClothesName string `json:"service_clothes_name" gorm:"-"`
}
