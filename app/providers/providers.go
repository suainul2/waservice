package providers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

var AppService *fiber.App

func Handle(app *fiber.App) {
	if AppService == nil {
		fmt.Println("service berjalan")
		AppService = app
	}
	RedisHandle()
	RouteService("/api").Handle()
}
