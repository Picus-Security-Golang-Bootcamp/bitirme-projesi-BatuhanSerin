package basket

import (
	"github.com/BatuhanSerin/final-project/internal/models"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type BasketRepository struct {
	db *gorm.DB
}

func NewBasketRepository(db *gorm.DB) *BasketRepository {
	return &BasketRepository{db: db}
}

func (b *BasketRepository) Migration() {
	zap.L().Debug("basket Migration")

	if err := b.db.AutoMigrate(&models.Basket{}); err != nil {
		zap.L().Error("basket Migration Failed", zap.Error(err))
	}
}
func (b *BasketRepository) VerifyToken(c *gin.Context, basket *models.Basket) (*models.Basket, error) {

	zap.L().Debug("basket.repo.VerifyToken")

	if err := b.db.FirstOrCreate(basket).Error; err != nil {
		zap.L().Error("basket.repo.VerifyToken Failed", zap.Error(err))
		return nil, err
	}

	return basket, nil

}
