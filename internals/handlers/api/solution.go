package apiHandler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/shokHorizon/kursik/database"
	"github.com/shokHorizon/kursik/internals/model"
)

func GetSolutions(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"status": "debug", "message": "GetSolutionsHandler", "data": nil})
}

func GetSolution(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"status": "debug", "message": "GetSolutionIDHandler", "data": nil})
}

func RemoveSolution(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"status": "debug", "message": "RemoveSolutionHandler", "data": nil})
}

func UpdateSolution(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"status": "debug", "message": "UpdateSolutionHandler", "data": nil})
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

	// Return the created solution.
	return c.JSON(fiber.Map{"status": "success", "message": "Solution created", "data": solution})
}

func PassSolution(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"status": "debug", "message": "PassSolutionHandler", "data": nil})
}

func FailSolution(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"status": "debug", "message": "FailSolutionHandler", "data": nil})
}
