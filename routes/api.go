package routes

import (
	"waservice/app/http/controllers"
	"waservice/app/http/middleware"
	"waservice/configs"
	"waservice/databases"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
)

func Api(r fiber.Router) {
	generateController := controllers.NewGenerateController(databases.Rdb, databases.Ctx)
	authController := controllers.NewAuthController(databases.Db)

	r.Post("/login", authController.Login())

	r.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte(configs.Server.Key),
	}))
	r.Use(middleware.JwtAuth().JwtHandle())
	r.Get("/connect", generateController.Generate())
	r.Get("/live", generateController.Live())
	r.Post("/send/text/*", generateController.SendText())
	r.Post("/send/image/*", generateController.SendImage())

}
