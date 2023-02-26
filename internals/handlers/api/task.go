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

	if len(tasks) == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No tasks present", "data": nil})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Tasks Found", "data": tasks})
}

func GetTask(c *fiber.Ctx) error {
	db := database.DB
	var task model.Task

	id := c.Params("id")
	db.Find(&task, "id =?", id)

	if task.ID == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No task present", "data": nil})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Task Found", "data": task})
}

func RemoveTask(c *fiber.Ctx) error {
	db := database.DB
	var task model.Task

	id := c.Params("id")

	db.Find(&task, "id = ?", id)

	if task.ID == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No task present", "data": nil})
	}

	err := db.Delete(&task, "id = ?", id).Error

	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Failed to delete task", "data": nil})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Task removed"})
}

func UpdateTask(c *fiber.Ctx) error {
	type updateTask struct {
		Name           string         `json:"title"`
		Description    string         `json:"description"`
		Tests          string         `json:"tests"`
		AuthorSolution model.Solution `gorm:"foreignKey:ID" json:"solution_id"`
		Tags           []model.Tag    `gorm:"foreignKey:ID" json:"tag_id"`
	}

	db := database.DB
	var task model.Task

	id := c.Params("id")

	db.Find(&task, "id = ?", id)

	if task.ID == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No task present", "data": nil})
	}

	var updateTaskData updateTask
	err := c.BodyParser(&updateTaskData)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}

	task.Name = updateTaskData.Name
	task.Description = updateTaskData.Description
	task.Tests = updateTaskData.Tests
	task.Tags = updateTaskData.Tags

	db.Save(&task)

	return c.JSON(fiber.Map{"status": "success", "message": "Task updated", "data": task})
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

	return c.JSON(fiber.Map{"status": "success", "message": "User created", "data": task})
}
