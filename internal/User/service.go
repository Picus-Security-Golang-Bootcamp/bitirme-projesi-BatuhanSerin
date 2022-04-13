package User

import (
	"fmt"

	"github.com/BatuhanSerin/final-project/internal/api"
	"github.com/BatuhanSerin/final-project/internal/models"
)

func UserToResponse(u *models.User) *api.User {
	fmt.Println("user to resp")
	return &api.User{
		Email:     u.Email,
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Phone:     u.Phone,
		Username:  u.Username,
	}
}

func responseToUser(u *api.User) *models.User {
	fmt.Println("res to user")
	return &models.User{
		//Model:     gorm.Model{ID: uint(u.ID)},
		Email:     u.Email,
		Password:  u.Password,
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Phone:     u.Phone,
		Username:  u.Username,
	}
}
