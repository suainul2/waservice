package controllers

import "github.com/gofiber/fiber/v2"

type controller struct {
}

func (C *controller) ResSuccess(c *fiber.Ctx, data interface{}) error {
	return c.Status(200).JSON(fiber.Map{
		"data": data,
	})
}
func (C *controller) ResFail(c *fiber.Ctx, msg interface{}) error {
	return c.Status(400).JSON(fiber.Map{
		"msg": msg,
	})
}
