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
	apiRoutes.SetupRoutes(app)
	taskRoutes.SetupRoutes(app)
	solutionRoutes.SetupRoutes(app)
	homeRoutes.SetupRoutes(app)
}
