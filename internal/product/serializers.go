package product

import (
	"gorm.io/gorm"

	"github.com/BatuhanSerin/final-project/internal/api"
	"github.com/BatuhanSerin/final-project/internal/models"
)

func ProductToResponse(p *models.Product) *api.Product {

	return &api.Product{
		Category: &api.CategoryWithoutRequired{
			ID:   int64(p.Category.ID),
			Name: *p.Category.Name,
		},
		ID:    int64(p.ID),
		Name:  p.Name,
		Price: &p.Price,
		Stock: &p.Stock,
	}
}

func ProductToResponseWithoutCategory(p *models.Product) *api.Product {

	return &api.Product{
		ID:    int64(p.ID),
		Name:  p.Name,
		Price: &p.Price,
		Stock: &p.Stock,
	}
}

func productsToResponseWithoutCategory(ps *[]models.Product) []*api.Product {
	products := make([]*api.Product, 0)
	for _, p := range *ps {
		products = append(products, ProductToResponseWithoutCategory(&p))
	}
	return products
}

func productsToResponse(ps *[]models.Product) []*api.Product {
	products := make([]*api.Product, 0)
	for _, p := range *ps {
		products = append(products, ProductToResponse(&p))
	}
	return products
}

func responseToProduct(p *api.Product) *models.Product {
	return &models.Product{
		Model:      gorm.Model{ID: uint(p.ID)},
		Name:       p.Name,
		Price:      *p.Price,
		Stock:      *p.Stock,
		CategoryID: uint(p.Category.ID),
	}
}
