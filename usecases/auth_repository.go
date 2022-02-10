package usecases

import "waservice/app/models"

// A UserRepository belong to the usecases layer.
type AuthRepository interface {
	Login(data *models.User, password string) (*models.User, error)
	Token(data *models.User) (string, error)
	Register(d *models.User) (data *models.User, err error)
}
