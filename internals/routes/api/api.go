package apiRoutes

import (
	"github.com/gofiber/fiber/v2"
	apiHandler "github.com/shokHorizon/kursik/internals/handlers/api"
)

func SetupRoutes(router fiber.Router) {
	api := router.Group("/api")

	setupTaskRoutes(api)
}

func setupTaskRoutes(router fiber.Router) {
	task := router.Group("/task")

	task.Get("/", apiHandler.GetTask)
}
