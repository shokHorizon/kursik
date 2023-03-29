package apiHandler

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/shokHorizon/kursik/database"
	"github.com/shokHorizon/kursik/internals/model"
)

func GetSolutions(c *fiber.Ctx) error {
	db := database.DB
	var solutions []model.Solution

	db.Find(&solutions)

	if len(solutions) == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No solution present", "data": nil})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Solutions Found", "data": solutions})
}

func GetSolution(c *fiber.Ctx) error {
	db := database.DB
	var solution model.Solution

	id := c.Params("id")
	db.Find(&solution, "id =?", id)

	if solution.ID == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No solution present", "data": nil})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Solution Found", "data": solution})
}

func RemoveSolution(c *fiber.Ctx) error {
	db := database.DB
	var solution model.Solution

	id := c.Params("id")

	db.Find(&solution, "id = ?", id)

	if solution.ID == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No solution present", "data": nil})
	}

	err := db.Delete(&solution, "id = ?", id).Error

	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Failed to delete solution", "data": nil})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Solution removed"})
}

func UpdateSolution(c *fiber.Ctx) error {
	db := database.DB
	var solution model.Solution

	var updateSolutionData model.Solution
	err := c.BodyParser(&updateSolutionData)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}

	db.Find(&solution, "id = ?", updateSolutionData.ID)

	if solution.ID == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No solution present", "data": nil})
	}

	solution.UserID = updateSolutionData.UserID
	solution.TaskID = updateSolutionData.TaskID
	solution.Code = updateSolutionData.Code
	solution.Status = updateSolutionData.Status

	db.Save(&solution)

	return c.JSON(fiber.Map{"status": "success", "message": "Solution updated", "data": solution})

}

func CreateSolution(c *fiber.Ctx) error {
	db := database.DB
	solution := new(model.Solution)
	stringTaskID := c.Params("id")
	taskID, e := strconv.ParseUint(stringTaskID, 10, 64)
	if e != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Invalid task id"})
	}
	type CreateSolution struct {
		Code   string `json:"code"`
		Status bool   `gorm:"default:null" json:"status"`
	}
	createSolution := new(CreateSolution)
	err := c.BodyParser(createSolution)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}

	// userID, er := middleware.GetUserId(c)
	// if er != nil {
	// 	return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": er})
	// }

	solution.TaskID = taskID
	solution.UserID = 1
	solution.Code = createSolution.Code
	solution.Status = createSolution.Status

	err = db.Create(&solution).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not create Solution", "data": err})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Solution created", "data": solution})
}
