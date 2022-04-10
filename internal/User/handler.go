package User

import (
	"net/http"

	"github.com/BatuhanSerin/final-project/internal/api"
	//"github.com/BatuhanSerin/final-project/package/config"
	"github.com/gin-gonic/gin"
)

type userHandler struct {
	repo *UserRepository
}

func NewUserHandler(r *gin.RouterGroup, repo *UserRepository) {
	u := &userHandler{repo: repo}

	r.POST("/login", u.login)
}

func (u *userHandler) login(c *gin.Context) {
	var user *api.Login

	if err := c.Bind(&user); err != nil {
		c.JSON(400, gin.H{
			"message": "Bad Request",
		})
	}

	a, _ := u.repo.getUser(user.Email, user.Password)
	if a == nil {
		c.JSON(401, gin.H{
			"message": "Unauthorized",
		})
	}
	token, err := u.repo.Login(a.Email, a.Password)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "Internal Server Error",
		})
	}
	c.JSON(http.StatusOK, token)

}
