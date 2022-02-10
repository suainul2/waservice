package helpers

import (
	"github.com/gofiber/fiber/v2"
)

func HttpClient(url string, method string, body interface{}) (*fiber.Agent, error) {
	a := fiber.AcquireAgent().ContentType("application/json").JSON(body)
	req := a.Request()
	req.Header.SetMethod(method)
	req.SetRequestURI(url)
	a.ConnectionClose()
	if err := a.Parse(); err != nil {
		return nil, err
	}
	return a, nil
}

func ParsingBody(c *fiber.Ctx, data interface{}) bool {

	if err := c.BodyParser(&data); err != nil {
		return true
	}
	return false
}
