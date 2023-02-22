package route

import (
	"github.com/gofiber/fiber/v2"
	userHandler "github.com/shokHorizon/kursik/internals/handlers/user"
)

func SetupUserRoutes(router fiber.Router) {
	task := router.Group("/auth")

	task.Get("/login", userHandler.GetLogin)
	task.Post("/login", userHandler.PostLogin)
}
