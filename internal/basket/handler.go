package basket

import (
	"net/http"
	"strconv"

	"github.com/BatuhanSerin/final-project/internal/api"
	"github.com/BatuhanSerin/final-project/internal/httpErrors"
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
	r.POST("/create/:id/:quantity", b.create)
	r.POST("/inc/:id", b.increment)
	r.POST("/dec/:id", b.decrement)
	r.GET("/:id", b.getByID)
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

	basketBody := &api.Basket{
		// ProductsIDs: []int64{},
		UserID: int64(ID),
	}

	//c.JSON(http.StatusOK, basketBody)

	basket, err := b.repo.VerifyToken(c, responseToBasket(basketBody))
	c.JSON(http.StatusOK, gin.H{
		"message": "Token verified",
	})
	c.JSON(http.StatusOK, basketToResponse(basket))
}
func (b *basketHandler) create(c *gin.Context) {

	basketBody := Verify(c)

	id, _ := strconv.ParseInt(c.Param("id"), 10, 0)
	quantity, _ := strconv.ParseInt(c.Param("quantity"), 10, 0)

	basketBody.ProductID = int64(id)
	basketBody.Quantity = int64(quantity)

	basket, err := b.repo.Create(c, responseToBasket(basketBody))
	if err != nil {
		c.JSON(httpErrors.ErrorResponse(err))
		c.JSON(http.StatusOK, gin.H{
			"message": "Stock is NOT enough",
		})
		return
	}
	c.JSON(http.StatusOK, basketToResponse(basket))
}

func (b *basketHandler) increment(c *gin.Context) {
	basketBody := Verify(c)

	id, _ := strconv.ParseInt(c.Param("id"), 10, 0)
	basketBody.ProductID = int64(id)

	basket, err := b.repo.Increment(c, responseToBasket(basketBody))
	if err != nil {
		c.JSON(httpErrors.ErrorResponse(err))
		c.JSON(http.StatusOK, gin.H{
			"message": "Stock is NOT enough",
		})
		return
	}
	c.JSON(http.StatusOK, basketToResponse(basket))
}

func (b *basketHandler) decrement(c *gin.Context) {
	basketBody := Verify(c)

	id, _ := strconv.ParseInt(c.Param("id"), 10, 0)
	basketBody.ProductID = int64(id)

	basket, err := b.repo.Decrement(c, responseToBasket(basketBody))
	if err != nil {
		c.JSON(httpErrors.ErrorResponse(err))
		return
	}
	c.JSON(http.StatusOK, basketToResponse(basket))
}

func (b *basketHandler) getByID(c *gin.Context) {
	id := c.Param("id")
	basket, err := b.repo.GetByID(c, id)
	if err != nil {
		c.JSON(httpErrors.ErrorResponse(err))
		return
	}
	c.JSON(http.StatusOK, basketToResponse(basket))
}
