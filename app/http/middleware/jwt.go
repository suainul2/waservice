package middleware

import (
	"waservice/app/models"
	"waservice/databases"

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
		c.Locals("auth", users)
		return c.Next()
	}
}

func JwtAuth() *jwtAuth {
	return &jwtAuth{}
}
