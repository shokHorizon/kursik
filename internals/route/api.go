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
	task.Use(middleware.IsUser)
	// task.Get("/recommended", apiHandler.GetRecommended)
	//task.Post("/:id/solution", apiHandler.CreateSolution)
	task.Get("/bytags", apiHandler.GetTasksByTags)
	// task.Use(middleware.IsAdmin)

	task.Put("/addtags/:id", apiHandler.AddTagsToTask)
	task.Post("/create", apiHandler.CreateTask)
	task.Put("/update", apiHandler.UpdateTask)
	task.Delete("/delete/:id", apiHandler.RemoveTask)

	task.Delete("/removetag/:id", apiHandler.RemoveTagFromTask)
	task.Patch("/sequence/:id", apiHandler.ReplaceSequence)
}

func setupSolutionRoutes(router fiber.Router) {
	solution := router.Group("/solution")

	solution.Use(middleware.IsUser)
	solution.Get("/get/:id", apiHandler.GetSolution)
	solution.Post("/:id/create", apiHandler.CreateSolution)
	solution.Use(middleware.IsAdmin)
	solution.Put("/update", apiHandler.UpdateSolution)
	solution.Delete("/delete/:id", apiHandler.RemoveSolution)
	solution.Get("/getAll", apiHandler.GetSolutions)

}

func setupUserRoutes(router fiber.Router) {
	user := router.Group("/user")

	// user.Use(middleware.IsAdmin)
	user.Post("/create", apiHandler.CreateUser)
	user.Put("/update", apiHandler.UpdateUser)
	user.Delete("/delete/:id", apiHandler.RemoveUser)
	user.Get("/getAll", apiHandler.GetUsers)
	user.Get("/get/:id", apiHandler.GetUser)
}

func setupCourseRoutes(router fiber.Router) {
	course := router.Group("/course")

	course.Get("/getAll", apiHandler.GetCourses)
	course.Get("/get/:id", apiHandler.GetCourse)
	course.Get("/get-tasks/:id", apiHandler.GetTasksByCource)
	course.Use(middleware.IsAdmin)
	course.Post("/create", apiHandler.CreateCourse)
	course.Put("/update", apiHandler.UpdateCourse)
	course.Delete("/delete/:id", apiHandler.RemoveCourse)

	course.Put("adduser/:id", apiHandler.AddUserToCourse)
	course.Delete("/delete-users/:id", apiHandler.RemoveUserFromCource)
	course.Put("/add-tasks/:id", apiHandler.AddTasksToCource)
	course.Delete("delete-tasks/:id", apiHandler.RemoveTaskFromCourse)

}

func setupTagRoutes(router fiber.Router) {
	tag := router.Group("/tag")

	tag.Get("/byname", apiHandler.FindTagByName)
	tag.Get("/getAll", apiHandler.GetTags)
	tag.Get("/get/:id", apiHandler.GetTag)
	// tag.Use(middleware.IsAdmin)
	tag.Post("/create", apiHandler.CreateTag)
	tag.Put("/update", apiHandler.UpdateTag)
	tag.Delete("/delete/:id", apiHandler.RemoveTag)

}

func SetupAuthRoutes(router fiber.Router) {
	auth := router.Group("/auth")

	auth.Post("/signup", apiHandler.SignUpUser)
	auth.Post("/signin", apiHandler.SignInUser)
	auth.Get("/logout", apiHandler.LogoutUser)
	// запрос на подтверждение по почте /confirm
	// запрос на смену пароля
}
