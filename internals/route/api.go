package route

import (
	"github.com/gofiber/fiber/v2"
	apiHandler "github.com/shokHorizon/kursik/internals/handlers/api"
)

func SetupApiRoutes(router fiber.Router) {
	api := router.Group("/api")

	setupTaskRoutes(api)
	setupSolutionRoutes(api)
}

func setupTaskRoutes(router fiber.Router) {
	task := router.Group("/task")

	task.Get("/", apiHandler.GetTasks)
	task.Get("/:id<int>", apiHandler.GetTask)
	task.Post("/create", apiHandler.CreateTask)
	task.Put("/update", apiHandler.UpdateTask)
	task.Delete("/remove", apiHandler.RemoveTask)
}

func setupSolutionRoutes(router fiber.Router) {
	solution := router.Group("/solution")

	solution.Get("/", apiHandler.GetSolutions)
	solution.Get("/:id<int>", apiHandler.GetSolution)
	solution.Post("/create", apiHandler.CreateSolution)
	solution.Put("/update", apiHandler.UpdateSolution)
	solution.Delete("/remove", apiHandler.RemoveSolution)

	solution.Post("/pased", apiHandler.PassSolution)
	solution.Post("/failed", apiHandler.FailSolution)

}
