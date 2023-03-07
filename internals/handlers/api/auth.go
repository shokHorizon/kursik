package apiHandler

import (
	"fmt"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"github.com/shokHorizon/kursik/database"
	"github.com/shokHorizon/kursik/internals/model"
	"golang.org/x/crypto/bcrypt"
)

func SignUpUser(c *fiber.Ctx) error {
	db := database.DB
	var payload *model.SignUpInput

	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	errors := model.ValidateStruct(payload)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "errors": errors})

	}

	if payload.Password != payload.PasswordConfirm {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": "Passwords do not match"})

	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(payload.Password), bcrypt.DefaultCost)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	newUser := model.User{
		Login:          payload.Login,
		Email:          strings.ToLower(payload.Email),
		HashedPassword: string(hashedPassword),
		AccessLevel:    1,
	}

	result := db.Create(&newUser)

	if result.Error != nil && strings.Contains(result.Error.Error(), "duplicate key value violates unique") {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{"status": "fail", "message": "User with that email already exists"})
	} else if result.Error != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "error", "message": "Something bad happened"})
	}

	// return c.Status(fiber.StatusCreated).JSON(fiber.Map{"status": "success", "data": fiber.Map{"user": model.FilterUserRecord(&newUser)}})
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"status": "success", "message": "User registered", "data": newUser})

}

func SignInUser(c *fiber.Ctx) error {
	db := database.DB
	var payload *model.SignInInput

	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	errors := model.ValidateStruct(payload)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)

	}

	var user model.User
	result := db.First(&user, "email = ?", strings.ToLower(payload.Email))
	if result.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": "Invalid email or Password"})
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.HashedPassword), []byte(payload.Password))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": "Invalid email or Password"})
	}

	secret := "chopka228"

	claims := jwt.MapClaims{
		"id":          user.ID,
		"accessLevel": 1,
		"exp":         time.Now().Add(time.Hour * 72).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(secret))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": fmt.Sprintf("generating JWT Token failed: %v", err)})
	}

	c.Cookie(&fiber.Cookie{
		Name:  "token",
		Value: t,
		Path:  "/",
		// MaxAge:   config.JwtMaxAge * 60,
		Secure:   false,
		HTTPOnly: true,
		Domain:   "localhost",
	})

	return c.JSON(fiber.Map{"token": t})

}

func LogoutUser(c *fiber.Ctx) error {
	expired := time.Now().Add(-time.Hour * 24)
	c.Cookie(&fiber.Cookie{
		Name:    "token",
		Value:   "",
		Expires: expired,
	})
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success"})
}
