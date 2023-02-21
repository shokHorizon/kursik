package router

import (
	"github.com/gofiber/fiber/v2"
	_ "github.com/gofiber/fiber/v2/middleware/logger"

	apiRoutes "github.com/shokHorizon/kursik/internals/routes/api"
	homeRoutes "github.com/shokHorizon/kursik/internals/routes/home"
	solutionRoutes "github.com/shokHorizon/kursik/internals/routes/solution"
	taskRoutes "github.com/shokHorizon/kursik/internals/routes/task"
)

func SetupRoutes(app *fiber.App) {
	// Group api calls with param '/api'
	// api := app.Group("/api", logger.New())

	// Setup note routes, can use same syntax to add routes for more models

	apiRoutes.SetupRoutes(app)
	taskRoutes.SetupRoutes(app)
	solutionRoutes.SetupRoutes(app)
	homeRoutes.SetupRoutes(app)
}
