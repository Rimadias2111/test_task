package models

import (
	"time"
)

type Order struct {
	Id         string    `gorm:"type:char(36);primaryKey" json:"id"`
	TotalPrice float64   `gorm:"not null" json:"total_price"`
	Discount   float64   `gorm:"default:0" json:"discount"`
	Express    bool      `gorm:"default:false" json:"express"`
	CreatedAt  time.Time `gorm:"autoCreateTime" json:"created_at"`

	OrderItems []OrderItem `gorm:"foreignKey:OrderId" json:"order_items"`
}
