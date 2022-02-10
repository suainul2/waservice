package controllers

import (
	"waservice/app/models"
	"waservice/app/services"
	"waservice/helpers"
	"waservice/usecases"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type AuthController struct {
	AuthInteractor usecases.AuthInteractor
	controller
}

// NewUserController returns the resource of users.
func NewAuthController(db *gorm.DB) *AuthController {
	return &AuthController{
		AuthInteractor: usecases.AuthInteractor{
			AuthRepository: &services.AuthRepository{
				Db: db,
			},
		},
	}
}

func (a *AuthController) Login() fiber.Handler {
	return func(c *fiber.Ctx) error {
		user := new(models.User)
		if helpers.ParsingBody(c, user) {
			return a.ResFail(c, "parsing error")
		}
		u, err := a.AuthInteractor.Login(user, user.Password)
		if err != nil {
			return a.ResFail(c, err)

		}
		t, err := a.AuthInteractor.Token(u)
		if err != nil {
			return a.ResFail(c, err)

		}
		return a.ResSuccess(c, t)
	}
}

func (a *AuthController) Register() fiber.Handler {
	return func(c *fiber.Ctx) error {
		user := new(models.User)
		if helpers.ParsingBody(c, user) {
			return a.ResFail(c, "parsing error")
		}
		u, err := a.AuthInteractor.Register(user)
		if err != nil {
			return a.ResFail(c, err)

		}
		return a.ResSuccess(c, u)
	}
}
