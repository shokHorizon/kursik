package route

import (
	"github.com/gofiber/fiber/v2"
	taskHandler "github.com/shokHorizon/kursik/internals/handlers/task"
)

func SetupTaskRoutes(router fiber.Router) {
	task := router.Group("/task")

	task.Get("/", taskHandler.GetTasks)
	task.Get("/:id<int>", taskHandler.GetTask)
	task.Post("/create", taskHandler.CreateTask)
	task.Put("/update", taskHandler.UpdateTask)
	task.Delete("/remove", taskHandler.RemoveTask)
}
