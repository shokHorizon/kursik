package route

import (
	"github.com/gofiber/fiber/v2"
	apiHandler "github.com/shokHorizon/kursik/internals/handlers/api"
	"github.com/shokHorizon/kursik/middleware"
)

func SetupApiRoutes(router fiber.Router) {
	api := router.Group("/api")

	setupTaskRoutes(api)
	setupSolutionRoutes(api)
	setupUserRoutes(api)
	setupCourseRoutes(api)
	setupTagRoutes(api)
	SetupAuthRoutes(api)
}

func setupTaskRoutes(router fiber.Router) {
	task := router.Group("/task")

	task.Get("/getAll", apiHandler.GetTasks)
	task.Get("/get/:id", apiHandler.GetTask)
	task.Post("/create", apiHandler.CreateTask)
	task.Put("/update", apiHandler.UpdateTask)
	task.Delete("/delete/:id", apiHandler.RemoveTask)
}

func setupSolutionRoutes(router fiber.Router) {
	solution := router.Group("/solution")

	solution.Get("/getAll", apiHandler.GetSolutions)
	solution.Get("/get/:id", apiHandler.GetSolution)
	solution.Post("/create", apiHandler.CreateSolution)
	solution.Put("/update", apiHandler.UpdateSolution)
	solution.Delete("/delete/:id", apiHandler.RemoveSolution)

	solution.Post("/pased", apiHandler.PassSolution)
	solution.Post("/failed", apiHandler.FailSolution)

}

func setupUserRoutes(router fiber.Router) {
	user := router.Group("/user")

	user.Get("/getAll", middleware.DeserializeUser, apiHandler.GetUsers)
	user.Get("/get/:id", apiHandler.GetUser)
	user.Post("/create", apiHandler.CreateUser)
	user.Put("/update", apiHandler.UpdateUser)
	user.Delete("/delete/:id", apiHandler.RemoveUser)

}

func setupCourseRoutes(router fiber.Router) {
	course := router.Group("/course")

	course.Get("/getAll", apiHandler.GetCourses)
	course.Get("/get/:id", apiHandler.GetCourse)
	course.Post("/create", apiHandler.CreateCourse)
	course.Put("/update", apiHandler.UpdateCourse)
	course.Delete("/delete/:id", apiHandler.RemoveCourse)

}

func setupTagRoutes(router fiber.Router) {
	course := router.Group("/tag")

	course.Get("/getAll", apiHandler.GetTags)
	course.Get("/get/:id", apiHandler.GetTag)
	course.Post("/create", apiHandler.CreateTag)
	course.Put("/update", apiHandler.UpdateTag)
	course.Delete("/delete/:id", apiHandler.RemoveTag)

}

func SetupAuthRoutes(router fiber.Router) {
	auth := router.Group("/auth")

	auth.Post("/signup", apiHandler.SignUpUser)
	auth.Post("/signin", apiHandler.SignInUser)
	auth.Get("/logout", apiHandler.LogoutUser)
}
