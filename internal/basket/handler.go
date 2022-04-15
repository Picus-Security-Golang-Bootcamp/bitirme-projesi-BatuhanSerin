package basket

import (
	"net/http"

	"github.com/BatuhanSerin/final-project/internal/api"
	cfg "github.com/BatuhanSerin/final-project/package/config"
	jwtPackage "github.com/BatuhanSerin/final-project/package/jwt"
	"github.com/BatuhanSerin/final-project/package/middleware"

	"github.com/gin-gonic/gin"
)

type basketHandler struct {
	repo *BasketRepository
}

func NewBasketHandler(r *gin.RouterGroup, repo *BasketRepository, secret string) {
	b := &basketHandler{repo: repo}

	r.Use(middleware.AuthorizationForUser(secret))
	r.POST("/verify", b.VerifyToken)
}

func (b *basketHandler) VerifyToken(c *gin.Context) {

	tokenString := c.GetHeader("Authorization")
	token, err := jwtPackage.ParseToken(tokenString, cfg.GetSecretKey())
	if err != nil {
		c.JSON(401, gin.H{
			"message": "Unauthorized",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Authorized",
	})

	c.JSON(http.StatusOK, token)

	claims := *token
	ID := claims["userID"].(float64)

	basketBody := api.Basket{
		// ProductsIDs: []int64{},
		UserID:     int64(ID),
		TotalPrice: 0,
	}

	c.JSON(http.StatusOK, basketBody)

	b.repo.VerifyToken(c, responseToBasket(&basketBody))
	c.JSON(http.StatusOK, gin.H{
		"message": "Token verified",
	})
}
