package User

import (
	"log"
	"time"

	"github.com/BatuhanSerin/final-project/internal/models"
	jwtPackage "github.com/BatuhanSerin/final-project/package/jwt"
	"github.com/golang-jwt/jwt"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (u *UserRepository) Migration() {
	zap.L().Debug("category Migration")

	if err := u.db.AutoMigrate(&models.User{}); err != nil {
		zap.L().Error("category Migration Failed", zap.Error(err))
	}
}

// func (u *UserRepository) getUsers() *[]models.User {
// 	zap.L().Debug("user.repo.getUsers")

// 	var users = &[]models.User{}

// 	if err := u.db.Find(&users).Error; err != nil {
// 		zap.L().Error("user.repo.getUsers Failed", zap.Error(err))
// 		return nil
// 	}

// 	return users
// }
func (u *UserRepository) getUser(email, password *string) (*models.User, error) {
	users := []*models.User{}
	//users2 := &models.User{}
	pass := *password
	emailValue := *email

	if err := u.db.Find(&users).Error; err != nil {
		zap.L().Error("user.repo.getUsers Failed", zap.Error(err))
		return nil, err
	}
	zap.L().Debug("user.repo.getUser", zap.Any("email", email), zap.Any("password", password))
	//fmt.Printf("%s 11111111111 %s ", emailValue, pass)
	for _, v := range users {
		//fmt.Printf("%s %s11 %s %s ", v.Email, v.Password, emailValue, pass)
		if v.Email == emailValue && v.Password == pass {
			//fmt.Print("%s %s    3333311 %s %s ", v.Email, v.Password, emailValue, pass)
			// fmt.Println("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa")
			// fmt.Print(v.Email, v.Password)
			return v, nil
		}
	}
	return nil, nil
}

func (r *UserRepository) Login(email, password string) (string, error) {
	var user *models.User
	if err := r.db.Where("email = ? AND password = ?", email, password).First(&user).Error; err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID": user.ID,
		"email":  user.Email,
		"role":   user.IsAdmin,
		"exp":    time.Now().Add(time.Hour * 72).Unix(),
	})

	tokenString, _ := jwtPackage.GenerateToken(token, "secret")
	log.Println(tokenString)
	return tokenString, nil
}
