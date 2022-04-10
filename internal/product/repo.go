package product

import (
	"github.com/BatuhanSerin/final-project/internal/models"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (r *ProductRepository) create(product *models.Product) (*models.Product, error) {
	zap.L().Debug("product.repo.create", zap.Any("product", product))

	if err := r.db.Create(product).Error; err != nil {
		zap.L().Error("product.repo.create Failed", zap.Error(err))
		return nil, err
	}
	return product, nil
}

func (r *ProductRepository) getAll() (*[]models.Product, error) {
	zap.L().Debug("product.repo.getAll")

	var products = &[]models.Product{}

	if err := r.db.Preload("Category").Find(&products).Error; err != nil {
		zap.L().Error("product.repo.getAll Failed", zap.Error(err))
		return nil, err
	}

	return products, nil
}

func (r *ProductRepository) getByID(id string) (*models.Product, error) {
	zap.L().Debug("product.repo.getByID", zap.Any("id", id))

	var product = &models.Product{}

	if err := r.db.Preload("Category").First(&product, id).Error; err != nil {
		zap.L().Error("product.repo.getByID Failed", zap.Error(err))
		return nil, err
	}

	return product, nil
}

func (r *ProductRepository) update(product *models.Product) (*models.Product, error) {
	zap.L().Debug("product.repo.update", zap.Any("product", product))

	if err := r.db.Save(&product).Error; err != nil {
		zap.L().Error("product.repo.update Failed", zap.Error(err))
		return nil, err
	}

	return product, nil
}

func (r *ProductRepository) delete(product *models.Product) error {
	zap.L().Debug("product.repo.delete", zap.Any("product", product))

	if err := r.db.Delete(&product).Error; err != nil {
		zap.L().Error("product.repo.delete Failed", zap.Error(err))
		return err
	}

	return nil
}

func (r *ProductRepository) Migration() {
	zap.L().Debug("Product Migration")

	if err := r.db.AutoMigrate(&models.Product{}); err != nil {
		zap.L().Error("Product Migration Failed", zap.Error(err))
	}
}
