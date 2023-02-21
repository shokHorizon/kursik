package apiHandler

import (
	"github.com/gofiber/fiber/v2"
)

func GetTask(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"status": "debug", "message": "[API] GetTask", "data": nil})
}
