package middleware

import (
	"log"
	"waservice/app/models"
	"waservice/databases"
	"waservice/helpers"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

type jwtAuth struct {
}

func (a *jwtAuth) JwtHandle() fiber.Handler {
	return func(c *fiber.Ctx) error {
		user := c.Locals("user").(*jwt.Token)
		claims := user.Claims.(jwt.MapClaims)
		users := &models.User{}
		databases.Db.Where("id = ?", claims["id"].(float64)).First(users)
		dataWaUser, err := databases.Rdb.Get(databases.Ctx, "go_wa_user_"+helpers.UintToStr(users.ID)).Result()
		if err != nil {
			log.Println(err.Error())
		}
		dataWaUserNew := helpers.StrToObWA(dataWaUser)
		dataWaUserNew.UserId = helpers.UintToStr(users.ID)
		c.Locals("auth", users)
		c.Locals("waUser", dataWaUserNew)
		return c.Next()
	}
}

func JwtAuth() *jwtAuth {
	return &jwtAuth{}
}
