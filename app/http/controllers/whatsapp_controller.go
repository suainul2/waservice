package controllers

import (
	"context"
	"log"
	"time"
	"waservice/app/services"
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
		log.Println(err)
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
		user := c.Locals("waUser").(*helpers.DataWAUser)
		go s.GenerateInteractor.Generate(user)
		return s.ResSuccess(c, "ok")
	}
}

func (s *GenerateController) Live() fiber.Handler {
	return func(c *fiber.Ctx) error {
		user := c.Locals("waUser").(*helpers.DataWAUser)
		return s.ResSuccess(c, user)
	}
}

func (s *GenerateController) SendText() fiber.Handler {
	return func(c *fiber.Ctx) error {
		user := c.Locals("waUser").(*helpers.DataWAUser)

		txt := struct {
			Message string   `json:"message"`
			Target  []string `json:"target"`
		}{}
		if helpers.ParsingBody(c, &txt) {
			return s.ResFail(c, "parsing error")
		}
		arr := txt.Target
		_, err := s.GenerateInteractor.Resession(user)
		if err != nil {
			return s.ResFail(c, user)
		}
		if len(arr) > 0 {
			for _, element := range arr {
				s.GenerateInteractor.SendMessage(element, txt.Message)
			}

		}
		return s.ResSuccess(c, "ok")
	}
}

func (s *GenerateController) SendImage() fiber.Handler {
	return func(c *fiber.Ctx) error {
		user := c.Locals("waUser").(*helpers.DataWAUser)

		txt := struct {
			ImageUrl string   `json:"imgUrl"`
			Message  string   `json:"message"`
			Target   []string `json:"target"`
		}{}
		if helpers.ParsingBody(c, &txt) {
			return s.ResFail(c, "parsing error")
		}
		arr := txt.Target
		_, err := s.GenerateInteractor.Resession(user)
		if err != nil {
			return s.ResFail(c, user)

		}
		if len(arr) > 0 {
			for _, element := range arr {
				s.GenerateInteractor.SendImage(element, txt.ImageUrl, txt.Message)
			}

		}
		return s.ResSuccess(c, "ok")
	}
}

func (s *GenerateController) SendImageX() fiber.Handler {
	return func(c *fiber.Ctx) error {
		user := c.Locals("waUser").(*helpers.DataWAUser)

		txt := struct {
			ImageUrl string   `json:"imgUrl"`
			Message  string   `json:"message"`
			Target   []string `json:"target"`
		}{}
		if helpers.ParsingBody(c, &txt) {
			return s.ResFail(c, "parsing error")
		}
		arr := txt.Target
		_, err := s.GenerateInteractor.Resession(user)
		if err != nil {
			return s.ResFail(c, user)

		}
		if len(arr) > 0 {
			u := helpers.StrToInt(c.Params("x"))
			for i := 0; i < u; i++ {
				for _, element := range arr {
					s.GenerateInteractor.SendImage(element, txt.ImageUrl, txt.Message)
				}
			}
		}
		return s.ResSuccess(c, "ok")
	}
}

func (s *GenerateController) SendTextX() fiber.Handler {
	return func(c *fiber.Ctx) error {
		user := c.Locals("waUser").(*helpers.DataWAUser)

		txt := struct {
			Message string   `json:"message"`
			Target  []string `json:"target"`
		}{}
		if helpers.ParsingBody(c, &txt) {
			return s.ResFail(c, "parsing error")
		}
		arr := txt.Target
		_, err := s.GenerateInteractor.Resession(user)
		if err != nil {
			return s.ResFail(c, user)

		}
		if len(arr) > 0 {
			u := helpers.StrToInt(c.Params("x"))
			for i := 0; i < u; i++ {
				for _, element := range arr {
					s.GenerateInteractor.SendMessage(element, txt.Message)
				}
			}
		}
		return s.ResSuccess(c, "ok")
	}
}
