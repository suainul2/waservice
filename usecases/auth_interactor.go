package usecases

import "waservice/app/models"

type AuthInteractor struct {
	AuthRepository AuthRepository
}

// Index is display a listing of the resource.
func (gi *AuthInteractor) Login(d *models.User, p string) (data *models.User, err error) {
	data, err = gi.AuthRepository.Login(d, p)
	return
}
func (gi *AuthInteractor) Token(d *models.User) (data string, err error) {
	data, err = gi.AuthRepository.Token(d)
	return
}

func (gi *AuthInteractor) Register(d *models.User) (data *models.User, err error) {
	data, err = gi.AuthRepository.Register(d)
	return
}
