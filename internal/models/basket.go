package models

import (
	"gorm.io/gorm"
)

type Basket struct {
	gorm.Model
	Products   []ProductInfo
	TotalPrice float64 `gorm:"type:decimal"`
	UserID     int64   `gorm:"type:bigint"`
}
