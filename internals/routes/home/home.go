package homeRoutes

import (
	"github.com/gofiber/fiber/v2"
	homeHandler "github.com/shokHorizon/kursik/internals/handlers/home"
)

func SetupRoutes(router fiber.Router) {
	router.Get("/", homeHandler.GetHome)
}
