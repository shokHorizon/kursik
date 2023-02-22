package route

import (
	"github.com/gofiber/fiber/v2"
	homeHandler "github.com/shokHorizon/kursik/internals/handlers/home"
)

func SetupHomeRoutes(router fiber.Router) {
	router.Get("/", homeHandler.GetHome)
}
