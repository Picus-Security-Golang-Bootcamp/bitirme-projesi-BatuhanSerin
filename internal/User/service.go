package User

import (
	"github.com/BatuhanSerin/final-project/internal/api"
	"github.com/BatuhanSerin/final-project/internal/models"
)

func UserToResponse(u *models.User) *api.User {

	return &api.User{
		Email:     u.Email,
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Phone:     u.Phone,
		Username:  u.Username,
	}
}

func responseToUser(u *api.User) *models.User {

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
