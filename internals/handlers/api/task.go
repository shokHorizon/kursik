package apiHandler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/shokHorizon/kursik/database"
	"github.com/shokHorizon/kursik/internals/model"
)

func GetTasks(c *fiber.Ctx) error {
	db := database.DB
	var tasks []model.Task

	db.Find(&tasks)

	//If have no tasks.
	if len(tasks) == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No tasks present", "data": nil})
	}

	//Return result.
	return c.JSON(fiber.Map{"status": "success", "message": "Tasks Found", "data": tasks})
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
	db := database.DB
	task := new(model.Task)

	err := c.BodyParser(task)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}

	err = db.Create(&task).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not create task", "data": err})
	}

	// Return the created task.
	return c.JSON(fiber.Map{"status": "success", "message": "User created", "data": task})
}
