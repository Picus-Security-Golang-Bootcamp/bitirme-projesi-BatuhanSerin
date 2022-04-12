package product

import (
	"encoding/csv"
	"log"
	"net/http"
	"strconv"

	"github.com/BatuhanSerin/final-project/internal/api"
	"github.com/BatuhanSerin/final-project/internal/httpErrors"
	"github.com/BatuhanSerin/final-project/package/middleware"
	"github.com/gin-gonic/gin"
	"github.com/go-openapi/strfmt"
)

type productHandler struct {
	repo *ProductRepository
}

func NewProductHandler(r *gin.RouterGroup, repo *ProductRepository, secret string) {
	p := &productHandler{repo: repo}
	r.GET("/", p.getAll)
	r.GET("/:id", p.getByID)
	r.Use(middleware.Authorization(secret))
	r.POST("/create", p.create)
	r.POST("/createBulk", p.createBulk)
}

func (p *productHandler) createBulk(c *gin.Context) {

	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(httpErrors.ErrorResponse(err))
	}

	csvFile, err := file.Open()
	if err != nil {
		c.JSON(httpErrors.ErrorResponse(err))
		log.Println("Open")
	}

	defer csvFile.Close()

	csvLines, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		c.JSON(httpErrors.ErrorResponse(err))
		log.Println("NewReader")
	}
	for _, line := range csvLines[1:] {

		CategoryId, _ := strconv.ParseInt(line[4], 10, 0)
		Id, _ := strconv.ParseInt(line[0], 10, 0)
		Price, _ := strconv.ParseFloat(line[2], 32)
		Stock, _ := strconv.ParseInt(line[3], 10, 0)

		productBody := &api.Product{
			Category: &api.CategoryWithoutRequired{
				ID: CategoryId,
			},
			ID:    Id,
			Name:  &line[1],
			Price: &Price,
			Stock: &Stock,
		}

		if err := productBody.Validate(strfmt.NewFormats()); err != nil {
			c.JSON(httpErrors.ErrorResponse(err))
			return
		}
		product, err := p.repo.create(responseToProduct(productBody))
		if err != nil {
			c.JSON(httpErrors.ErrorResponse(err))
		}
		c.JSON(http.StatusOK, ProductToResponseWithoutCategory(product))
	}

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
