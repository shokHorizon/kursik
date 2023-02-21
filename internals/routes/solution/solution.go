package solutionRoutes

import (
	"github.com/gofiber/fiber/v2"
	solutionHandler "github.com/shokHorizon/kursik/internals/handlers/solution"
)

func SetupRoutes(router fiber.Router) {
	task := router.Group("/solution")

	task.Get("/", solutionHandler.GetSolutions)
	task.Get("/:id<int>", solutionHandler.GetSolution)
	task.Post("/create", solutionHandler.CreateSolution)
	task.Put("/update", solutionHandler.UpdateSolution)
	task.Delete("/remove", solutionHandler.RemoveSolution)
}
