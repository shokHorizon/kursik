package router

import (
	"github.com/gofiber/fiber/v2"

	"github.com/shokHorizon/kursik/internals/route"
)

func SetupRoutes(app *fiber.App) {
	route.SetupApiRoutes(app)
	route.SetupUserRoutes(app)
	route.SetupHomeRoutes(app)
	route.SetupTaskRoutes(app)
	route.SetupSolutionRoutes(app)
}
