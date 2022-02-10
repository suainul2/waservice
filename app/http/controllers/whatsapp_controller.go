package controllers

import (
	"context"
	"strings"
	"time"
	"waservice/app/models"
	"waservice/app/services"
	"waservice/databases"
	"waservice/helpers"
	"waservice/usecases"

	whatsapp "github.com/Rhymen/go-whatsapp"
	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
)

type GenerateController struct {
	GenerateInteractor usecases.GenerateInteractor
	controller
}

// NewUserController returns the resource of users.
func NewGenerateController(Rdb *redis.Client, Ctx context.Context) *GenerateController {
	wac, err := whatsapp.NewConn(5 * time.Second)
	if err != nil {
		panic(err)
	}
	return &GenerateController{
		GenerateInteractor: usecases.GenerateInteractor{
			GenerateRepository: &services.GenerateRepository{
				Rdb: Rdb,
				Ctx: Ctx,
				Wac: wac,
			},
		},
	}
}

func (s *GenerateController) Generate() fiber.Handler {
	return func(c *fiber.Ctx) error {
		user := c.Locals("auth").(*models.User)
		go s.GenerateInteractor.Generate(user.ID)
		return s.ResSuccess(c, "ok")
	}
}

func (s *GenerateController) Live() fiber.Handler {
	return func(c *fiber.Ctx) error {
		user := c.Locals("auth").(*models.User)
		val, err := databases.Rdb.Get(databases.Ctx, "go_red_"+helpers.UintToStr(user.ID)).Result()
		if err != nil {
			panic(err)
		}
		return s.ResSuccess(c, val)
	}
}

func (s *GenerateController) SendText() fiber.Handler {
	return func(c *fiber.Ctx) error {
		user := c.Locals("auth").(*models.User)
		arr := strings.Split(c.Params("*"), "/")
		txt := struct {
			Message string `json:"message"`
		}{}
		if helpers.ParsingBody(c, &txt) {
			return s.ResFail(c, "parsing error")
		}
		_, err := s.GenerateInteractor.Resession(user.ID)
		if err != nil {
			return s.ResFail(c, err)

		}
		if len(arr) > 0 {
			for _, element := range arr {
				s.GenerateInteractor.SendMessage(element, txt.Message, user.ID)
			}

		}
		return s.ResSuccess(c, "ok")
	}
}

func (s *GenerateController) SendImage() fiber.Handler {
	return func(c *fiber.Ctx) error {
		user := c.Locals("auth").(*models.User)
		arr := strings.Split(c.Params("*"), "/")
		txt := struct {
			ImageUrl string `json:"imgUrl"`
			Message  string `json:"message"`
		}{}
		if helpers.ParsingBody(c, &txt) {
			return s.ResFail(c, "parsing error")
		}
		_, err := s.GenerateInteractor.Resession(user.ID)
		if err != nil {
			return s.ResFail(c, err.Error())

		}
		if len(arr) > 0 {
			for _, element := range arr {
				s.GenerateInteractor.SendImage(element, txt.ImageUrl, txt.Message, user.ID)
			}

		}
		return s.ResSuccess(c, "ok")
	}
}
