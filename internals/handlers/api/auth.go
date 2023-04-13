package apiHandler

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"github.com/segmentio/kafka-go"
	"github.com/shokHorizon/kursik/database"
	"github.com/shokHorizon/kursik/internals/model"
	"github.com/shokHorizon/kursik/producer"
	"github.com/thanhpk/randstr"
	"golang.org/x/crypto/bcrypt"
)

const jwt_secret = "chopka228"

var emailProducer *kafka.Writer = producer.EmailSending()

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

	code := randstr.String(20)
	timeout := time.Now().Add(time.Minute * 15)

	newUser := model.User{
		Login:          payload.Login,
		Email:          strings.ToLower(payload.Email),
		HashedPassword: string(hashedPassword),
		AccessLevel:    0,
		VerHash:        code,
		Timeout:        timeout,
		IsActivate:     false,
	}

	result := db.Create(&newUser)

	if result.Error != nil && strings.Contains(result.Error.Error(), "duplicate key value violates unique") {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{"status": "fail", "message": "User with that email already exists"})
	} else if result.Error != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "error", "message": "Something bad happened"})
	}

	type Content struct {
		Email   string `json:"email"`
		VerHash string `json:"ver_hash"`
	}

	content := Content{Email: newUser.Email, VerHash: newUser.VerHash}
	bytes, r := json.Marshal(content)
	if r != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "error", "message": "Something bad happened"})
	}

	emailProducer.WriteMessages(context.Background(), kafka.Message{
		Value: []byte(bytes),
	})

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"status": "success", "message": "Confirm registration in email"})

}

func SignUpVerification(c *fiber.Ctx) error {
	db := database.DB
	var user model.User

	hashedCode := c.Params("ver_hash")

	db.Find(&user, "ver_hash = ?", hashedCode)
	if user.ID == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No user present", "data": nil})
	}

	db.Model(&user).UpdateColumn("is_activate", true)

	return c.JSON(fiber.Map{"status": "success", "message": "User registered!!!"})
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

	if !user.IsActivate {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": "Confirm registration in email"})
	}

	claims := jwt.MapClaims{
		"id":          user.ID,
		"accessLevel": 0,
		"exp":         time.Now().Add(time.Hour * 12).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(jwt_secret))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": fmt.Sprintf("generating JWT Token failed: %v", err)})
	}

	c.Cookie(&fiber.Cookie{
		Name:     "token",
		Value:    t,
		Path:     "/",
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
