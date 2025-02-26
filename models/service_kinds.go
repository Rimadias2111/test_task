package models

import (
	"time"
)

type ServiceKind struct {
	Id        string    `gorm:"type:char(36);primaryKey" json:"id"`
	Name      string    `gorm:"size:100;not null" json:"name"`
	Unit      string    `gorm:"type:enum('item', 'kg');not null" json:"unit"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
}
