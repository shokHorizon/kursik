package middleware

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"github.com/shokHorizon/kursik/database"
	"github.com/shokHorizon/kursik/internals/model"
)

func IsAdmin(c *fiber.Ctx) error {
	return DeserializeUser(c, 1)
}

func IsUser(c *fiber.Ctx) error {
	return DeserializeUser(c, 0)
}

func DeserializeUser(c *fiber.Ctx, requireLevel int) error {
	jwt_secret := "chopka228"
	db := database.DB
	var tokenString string
	authorization := c.Get("Authorization")

	if strings.HasPrefix(authorization, "Bearer ") {
		tokenString = strings.TrimPrefix(authorization, "Bearer ")
	} else if c.Cookies("token") != "" {
		tokenString = c.Cookies("token")
	}

	if tokenString == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "fail", "message": "You are not logged in"})
	}

	tokenByte, err := jwt.Parse(tokenString, func(jwtToken *jwt.Token) (interface{}, error) {
		if _, ok := jwtToken.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %s", jwtToken.Header["alg"])
		}

		return []byte(jwt_secret), nil
	})

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "fail", "message": fmt.Sprintf("invalidate token: %v", err)})
	}

	claims, ok := tokenByte.Claims.(jwt.MapClaims)
	if !ok || !tokenByte.Valid {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "fail", "message": "invalid token claim"})

	}

	var user model.User
	//db.First(&user, "id = ?", fmt.Sprint(claims["id"]))
	db.First(&user, "id = ?", fmt.Sprint(claims["id"]))

	if float64(user.ID) != claims["id"] {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"status": "fail", "message": "the user belonging to this token no logger exists", "user": user.ID, "claims": claims["id"]})
	}
	s := fmt.Sprint(claims["accessLevel"])

	res1, _ := strconv.ParseUint(s, 16, 16)
	if uint16(res1) < uint16(requireLevel) {
		return c.Status(fiber.StatusMethodNotAllowed).JSON(fiber.Map{"status": "fail", "message": "Access denied"})
	}
	//c.Locals("user", model.FilterUserRecord(&user))
	ctxUserId := fmt.Sprint(claims["id"])
	c.Set("User_id", ctxUserId)

	return c.Next()

}

func GetUserId(c *fiber.Ctx) (uint64, error) {
	userId := c.Get("User_id")
	stringUserId := fmt.Sprint(userId)
	uintId, err := strconv.ParseUint(stringUserId, 10, 64)
	fmt.Println(uintId)
	if err != nil {
		return 0, c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": "Invalid user id"})
	}
	return uintId, nil
}
