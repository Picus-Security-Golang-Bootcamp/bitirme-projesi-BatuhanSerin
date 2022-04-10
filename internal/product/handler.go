package product

import (
	"net/http"

	"github.com/BatuhanSerin/final-project/internal/api"
	"github.com/BatuhanSerin/final-project/internal/httpErrors"
	"github.com/gin-gonic/gin"
	"github.com/go-openapi/strfmt"
)

type productHandler struct {
	repo *ProductRepository
}

func NewProductHandler(r *gin.RouterGroup, repo *ProductRepository) {
	p := &productHandler{repo: repo}
	r.GET("/", p.getAll)
	r.POST("/create", p.create)
	r.GET("/:id", p.getByID)
}

func (p *productHandler) create(c *gin.Context) {
	productBody := &api.Product{}

	if err := c.Bind(&productBody); err != nil {
		c.JSON(httpErrors.ErrorResponse(httpErrors.CannotBindGivenData))
		return
	}

	if err := productBody.Validate(strfmt.NewFormats()); err != nil {
		c.JSON(httpErrors.ErrorResponse(err))
		return
	}

	product, err := p.repo.create(responseToProduct(productBody))
	if err != nil {
		c.JSON(httpErrors.ErrorResponse(err))
		return
	}

	c.JSON(http.StatusOK, ProductToResponseWithoutCategory(product))

}

func (p *productHandler) getAll(c *gin.Context) {
	products, err := p.repo.getAll()
	if err != nil {
		c.JSON(httpErrors.ErrorResponse(err))
		return
	}
	c.JSON(http.StatusOK, productsToResponseWithoutCategory(products))
}

func (p *productHandler) getByID(c *gin.Context) {
	id := c.Param("id")
	product, err := p.repo.getByID(id)
	if err != nil {
		c.JSON(httpErrors.ErrorResponse(err))
		return
	}
	c.JSON(http.StatusOK, ProductToResponseWithoutCategory(product))
}
