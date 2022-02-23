package services

import (
	"errors"
	"waservice/app/models"
	"waservice/configs"
	"waservice/helpers"

	"github.com/golang-jwt/jwt/v4"
	"gorm.io/gorm"
)

type AuthRepository struct {
	Db *gorm.DB
}

func (s *AuthRepository) Login(data *models.User, password string) (*models.User, error) {
	s.Db.Where("email = ?", data.Email).First(data)
	if !helpers.PassCompare(data.Password, password) {
		return nil, errors.New("password not valid")
	}
	return data, nil
}

func (s *AuthRepository) Token(data *models.User) (string, error) {
	claims := jwt.MapClaims{
		"name":  data.Name,
		"id":    data.ID,
		"admin": data.Role == data.GetRole()["admin"],
	}
	// "exp":   time.Now().Add(time.Hour * 12).Unix(),

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(configs.Server.Key))
	if err != nil {
		return "", err
	}
	return t, err
}
func (s *AuthRepository) Register(data *models.User) (*models.User, error) {
	data.Role = 0
	data.Password = helpers.Bcrypt(data.Password)
	s.Db.Create(data)
	return data, nil
}
