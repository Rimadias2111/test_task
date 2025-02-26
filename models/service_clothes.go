package models

import (
	"time"
)

type ServiceClothes struct {
	Id            string    `gorm:"type:char(36);primaryKey" json:"id"`
	ServiceKindId string    `gorm:"type:char(36);not null;index" json:"service_kind_id"`
	Name          string    `gorm:"size:100;not null" json:"name"`
	Price         float64   `gorm:"not null" json:"price"`
	CreatedAt     time.Time `gorm:"autoCreateTime" json:"created_at"`

	ServiceKind ServiceKind `gorm:"foreignKey:ServiceKindId" json:"service_kind"`
}

type GetAllServiceClothesReq struct {
	ServiceKindId string
}
