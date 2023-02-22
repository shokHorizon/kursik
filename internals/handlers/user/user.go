package userHandler

import (
	"github.com/gofiber/fiber/v2"

	"github.com/shokHorizon/kursik/internals/model"
	"github.com/shokHorizon/kursik/middleware"
)

func GetLogin(c *fiber.Ctx) error {
	return c.Render("auth/login", fiber.Map{})
}

func PostLogin(c *fiber.Ctx) error {
	user := new(model.User)
	if err := c.BodyParser(user); err != nil || !user.LoginUser() {
		return c.Render("auth/login", fiber.Map{})
	}
	middleware.SetAuthenticated(c, user.Login, user.HashedPassword)
	return c.Redirect("../")
}
