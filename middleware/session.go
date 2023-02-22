package middleware

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

var store session.Store

func init() {
	store = *session.New()
}

func IsAuthenticated(c *fiber.Ctx) error {
	_, err := store.Get(c)
	if err != nil {
		return c.RedirectToRoute("login", fiber.Map{})
	}

	return c.Next()
}

func SetAuthenticated(c *fiber.Ctx, login, password string) error {
	sess, err := store.Get(c)
	if err != nil {
		fmt.Println(err)
		return c.RedirectToRoute("home", fiber.Map{})
	}
	sess.Set(login, password)
	if err := sess.Save(); err != nil {
		return c.RedirectToRoute("auth/login", fiber.Map{})
	}
	return c.RedirectToRoute("home", fiber.Map{})
}
