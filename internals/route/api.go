package route

import (
	"github.com/gofiber/fiber/v2"
	apiHandler "github.com/shokHorizon/kursik/internals/handlers/api"
)

func SetupApiRoutes(router fiber.Router) {
	api := router.Group("/api")

	setupTaskRoutes(api)
	setupSolutionRoutes(api)
	setupUserRoutes(api)
	setupCourseRoutes(api)
}

func setupTaskRoutes(router fiber.Router) {
	task := router.Group("/task")

	task.Get("/", apiHandler.GetTasks)
	task.Get("/:id", apiHandler.GetTask)
	task.Post("/", apiHandler.CreateTask)
	task.Put("/:id", apiHandler.UpdateTask)
	task.Delete("/:id", apiHandler.RemoveTask)
}

func setupSolutionRoutes(router fiber.Router) {
	solution := router.Group("/solution")

	solution.Get("/", apiHandler.GetSolutions)
	solution.Get("/:id", apiHandler.GetSolution)
	solution.Post("/", apiHandler.CreateSolution)
	solution.Put("/:id", apiHandler.UpdateSolution)
	solution.Delete("/:id", apiHandler.RemoveSolution)

	solution.Post("/pased", apiHandler.PassSolution)
	solution.Post("/failed", apiHandler.FailSolution)

}

func setupUserRoutes(router fiber.Router) {
	user := router.Group("/user")

	user.Get("/", apiHandler.GetUsers)
	user.Get("/:id", apiHandler.GetUser)
	user.Post("/", apiHandler.CreateUser)
	user.Put("/:id", apiHandler.UpdateUser)
	user.Delete("/:id", apiHandler.RemoveUser)

}

func setupCourseRoutes(router fiber.Router) {
	course := router.Group("/course")

	course.Get("/", apiHandler.GetCourses)
	course.Get("/:id", apiHandler.GetCourse)
	course.Post("/", apiHandler.CreateCourse)
	course.Put("/:id", apiHandler.UpdateUser)
	course.Delete("/:id", apiHandler.RemoveCourse)

}
