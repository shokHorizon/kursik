package apiHandler

import (
	"github.com/gofiber/fiber/v2"
)

func GetTasks(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"status": "debug", "message": "GetTasksHandler", "data": nil})
}

func GetTask(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"status": "debug", "message": "GetTaskIDHandler", "data": nil})
}

func RemoveTask(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"status": "debug", "message": "RemoveTaskHandler", "data": nil})
}

func UpdateTask(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"status": "debug", "message": "UpdateTaskHandler", "data": nil})
}

func CreateTask(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"status": "debug", "message": "CreateTaskHandler", "data": nil})
}
