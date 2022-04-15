package productInfo

import (
	"github.com/BatuhanSerin/final-project/internal/api"
	"github.com/BatuhanSerin/final-project/internal/models"
)

func productInfoToResponse(p *models.ProductInfo) *api.ProductInfo {
	return &api.ProductInfo{
		ProductsIDs: p.ProductID,
		Quantity:    p.Quantity,
	}
}

func responseToProductInfo(p *api.ProductInfo) *models.ProductInfo {
	return &models.ProductInfo{
		ProductID: p.ProductsIDs,
		Quantity:  p.Quantity,
		BasketID:  uint(p.Basket.UserID),
	}
}
