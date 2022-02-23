package main

import (
	"waservice/app/providers"
	"waservice/configs"
	"waservice/helpers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	app := fiber.New()
	app.Use(cors.New())
	app.Use(recover.New())
	app.Use(logger.New())
	app.Static("/", helpers.Pwd+"/public")
	configs.Handle()
	providers.Handle(app)
	app.Listen(configs.Server.Port)
}
