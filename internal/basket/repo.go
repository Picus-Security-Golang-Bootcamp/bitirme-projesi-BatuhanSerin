package basket

import (
	"fmt"
	"strconv"

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

func (b *BasketRepository) Increment(c *gin.Context, basket *models.Basket) (*models.Basket, error) {

	zap.L().Debug("basket.repo.Increment")

	if err := b.db.Model(&basket).Where(fmt.Sprintf("user_id = %v", basket.UserID)).Where(fmt.Sprintf("product_id = %v", basket.ProductID)).First(basket).Error; err != nil {
		basket.Quantity = 1
		if err := b.db.Create(basket).Error; err != nil {
			zap.L().Error("basket.repo.Increment Failed", zap.Error(err))
			return nil, err
		}
	} else {

		basket.Quantity++
		quantity := strconv.FormatUint(uint64(basket.Quantity), 10)

		if err := CheckStock(b, c, basket); err != nil {
			return nil, err
		}

		if err := b.db.Model(&basket).Where(fmt.Sprintf("user_id = %v", basket.UserID)).Where(fmt.Sprintf("product_id = %v", basket.ProductID)).Update("quantity", quantity).Error; err != nil {
			zap.L().Error("basket.repo.Increment Failed", zap.Error(err))
			return nil, err
		}
	}
	return basket, nil

}
func (b *BasketRepository) Decrement(c *gin.Context, basket *models.Basket) (*models.Basket, error) {

	zap.L().Debug("basket.repo.Decrement")

	if err := b.db.Model(&basket).Where(fmt.Sprintf("user_id = %v", basket.UserID)).Where(fmt.Sprintf("product_id = %v", basket.ProductID)).First(basket).Error; err == nil {
		basket.Quantity--
		if basket.Quantity > 0 {
			quantity := strconv.FormatUint(uint64(basket.Quantity), 10)
			if err := b.db.Model(&basket).Where(fmt.Sprintf("user_id = %v", basket.UserID)).Where(fmt.Sprintf("product_id = %v", basket.ProductID)).Update("quantity", quantity).Error; err != nil {
				zap.L().Error("basket.repo.Decrement Failed", zap.Error(err))
				return nil, err
			}
		} else {
			if err := b.db.Delete(basket).Error; err != nil {
				zap.L().Error("basket.repo.Decrement Failed", zap.Error(err))
				return nil, err
			}
			return nil, err
		}
	}

	return basket, nil

}

func (b *BasketRepository) Create(c *gin.Context, basket *models.Basket) (*models.Basket, error) {

	if err := CheckStock(b, c, basket); err != nil {
		return nil, err
	}

	zap.L().Debug("basket.repo.Create")
	quantity := strconv.FormatUint(uint64(basket.Quantity), 10)
	if err := b.db.Model(&basket).Where(fmt.Sprintf("user_id = %v", basket.UserID)).Where(fmt.Sprintf("product_id = %v", basket.ProductID)).Update("quantity", quantity).First(basket).Error; err != nil {
		if err := b.db.Create(basket).Error; err != nil {
			zap.L().Error("basket.repo.Create Failed", zap.Error(err))
			return nil, err
		}
	}
	return basket, nil
}

func (b *BasketRepository) GetByID(c *gin.Context, id string) (*models.Basket, error) {

	zap.L().Debug("basket.repo.GetByID")

	basket := &models.Basket{}

	if err := b.db.Preload("Products").First(basket, id).Error; err != nil {
		zap.L().Error("basket.repo.GetByID Failed", zap.Error(err))
		return nil, err
	}

	return basket, nil

}
