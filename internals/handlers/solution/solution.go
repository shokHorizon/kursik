package solutionHandler

import (
	"github.com/gofiber/fiber/v2"
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
	return c.JSON(fiber.Map{"status": "debug", "message": "CreateSolutionHandler", "data": nil})
}
