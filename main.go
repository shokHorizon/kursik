package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"

	"github.com/shokHorizon/kursik/database"
	"github.com/shokHorizon/kursik/router"
)

func main() {
	engine := html.New("./views", ".tmpl")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	// Connect to the Database
	database.ConnectDB()

	// Send a string back for GET calls to the endpoint "/"
	router.SetupRoutes(app)

	// Listen on PORT 3000
	app.Listen(":3000")
}
