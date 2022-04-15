package productInfo

import (
	"github.com/BatuhanSerin/final-project/internal/models"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type ProductInfoRepository struct {
	db *gorm.DB
}

func NewProductInfoRepository(db *gorm.DB) *ProductInfoRepository {
	return &ProductInfoRepository{db: db}
}

func (r *ProductInfoRepository) create(productInfo *models.ProductInfo) (*models.ProductInfo, error) {
	zap.L().Debug("ProductInfoRepository.repo.create", zap.Any("product", productInfo))

	if err := r.db.Where("product_id = ?", productInfo.ProductID).Find(&productInfo).Error; err == nil {
		if err := r.db.Create(&productInfo).Error; err != nil {
			zap.L().Error("productInfo.repo.create Failed", zap.Error(err))
			return nil, err
		}
	}
	// if err := r.db.Model(&productInfo).Where("product_id",&productInfo.ProductID).Create(&productInfo).Error; err != nil {
	// 	zap.L().Error("productInfo.repo.create Failed", zap.Error(err))
	// 	return nil, err
	// }

	return productInfo, nil
}

func (r *ProductInfoRepository) update(productInfo *models.ProductInfo) (*models.ProductInfo, error) {
	zap.L().Debug("ProductInfoRepository.repo.create", zap.Any("product", productInfo))

	if err := r.db.Model(&productInfo).Where("basket_id = ?", &productInfo.BasketID).Where("product_id = ?", &productInfo.ProductID).Update("quantity", &productInfo.Quantity).Error; err != nil {
		zap.L().Error("productInfo.repo.update Failed", zap.Error(err))
		return nil, err
	}

	return productInfo, nil
}
func (r *ProductInfoRepository) Migration() {
	zap.L().Debug("ProductInfo Migration")

	if err := r.db.AutoMigrate(&models.ProductInfo{}); err != nil {
		zap.L().Error("ProductInfo Migration Failed", zap.Error(err))
	}
}
