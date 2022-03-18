package routes

import (
	"fmt"
	"waservice/app/http/controllers"
	"waservice/app/http/middleware"
	"waservice/configs"
	"waservice/databases"
	"waservice/helpers"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
)

func Api(r fiber.Router) {
	generateController := controllers.NewGenerateController(databases.Rdb, databases.Ctx)
	authController := controllers.NewAuthController(databases.Db)

	r.Post("/login", authController.Login())
	r.Post("/regis", authController.Register())
	r.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte(configs.Server.Key),
	}))
	r.Use(middleware.JwtAuth().JwtHandle())
	r.Get("/connect", generateController.Generate())
	r.Get("/live", generateController.Live())
	r.Post("/send/text", generateController.SendText())
	r.Post("/send/image", generateController.SendImage())
	r.Post("/send/images/:x", generateController.SendImageX())
	r.Post("/send/texts/:x", generateController.SendTextX())
	err := databases.Rdb.Set(databases.Ctx, "go_wa_", string(helpers.ObToSTring(map[string]interface{}{"suainul": 2})), 0).Err()
	if err != nil {
		fmt.Println(err)
	}
}