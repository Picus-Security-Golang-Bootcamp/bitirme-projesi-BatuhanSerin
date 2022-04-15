package models

import "gorm.io/gorm"

type ProductInfo struct {
	gorm.Model
	ProductID *int64 `gorm:"type:bigint"`
	Quantity  int64  `gorm:"type:bigint"`
	BasketID  uint
	Basket    Basket
}
