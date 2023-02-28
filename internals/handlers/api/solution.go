package apiHandler

import (
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

	//If have no user.
	if solution.ID == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No solution present", "data": nil})
	}

	//Return result.
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

	err := c.BodyParser(solution)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}

	err = db.Create(&solution).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not create Solution", "data": err})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Solution created", "data": solution})
}

func PassSolution(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"status": "debug", "message": "PassSolutionHandler", "data": nil})
}

func FailSolution(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"status": "debug", "message": "FailSolutionHandler", "data": nil})
}
