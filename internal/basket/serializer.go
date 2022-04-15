package basket

import (
	"github.com/BatuhanSerin/final-project/internal/api"
	"github.com/BatuhanSerin/final-project/internal/models"
	"gorm.io/gorm"
)

func responseToBasket(b *api.Basket) *models.Basket {

	return &models.Basket{
		Model:      gorm.Model{ID: uint(b.UserID)},
		UserID:     b.UserID,
		TotalPrice: b.TotalPrice,
	}
}
