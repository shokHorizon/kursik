package apiHandler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/shokHorizon/kursik/database"
	"github.com/shokHorizon/kursik/internals/model"
)

func GetUsers(c *fiber.Ctx) error {
	db := database.DB
	var users []model.User

	db.Find(&users)

	//If have no users.
	if len(users) == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No users present[]", "data": nil})
	}

	//Return result.
	return c.JSON(fiber.Map{"status": "success", "message": "Users Found", "data": users})
}

func GetUser(c *fiber.Ctx) error {
	db := database.DB
	var user model.User

	//Read Id param.
	id := c.Params("id")
	//Find user by id.
	db.Find(&user, "id =?", id)

	//If have no user.
	if user.ID == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No user present", "data": nil})
	}

	//Return result.
	return c.JSON(fiber.Map{"status": "success", "message": "User Found", "data": user})

}

func RemoveUser(c *fiber.Ctx) error {
	db := database.DB
	var user model.User

	// Read the param id.
	id := c.Params("id")

	// Find the note with the given id.
	db.Find(&user, "id = ?", id)

	// If no such user present return an error.
	if user.ID == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No user present", "data": nil})
	}

	// Delete the user and return error if encountered.
	err := db.Delete(&user, "id = ?", id).Error

	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Failed to delete user", "data": nil})
	}

	// Return deleted message.
	return c.JSON(fiber.Map{"status": "success", "message": "User removed"})

}

func UpdateUser(c *fiber.Ctx) error {
	db := database.DB
	var user model.User

	var updateUserData model.User
	err := c.BodyParser(&updateUserData)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}

	db.Find(&user, "id = ?", updateUserData.ID)

	if user.ID == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No user present", "data": nil})
	}

	user.Login = updateUserData.Login
	user.HashedPassword = updateUserData.HashedPassword
	user.AccessLevel = updateUserData.AccessLevel
	user.Tasks = updateUserData.Tasks
	user.Solutions = updateUserData.Solutions
	user.Courses = updateUserData.Courses

	db.Save(&user)

	return c.JSON(fiber.Map{"status": "success", "message": "User updated", "data": user})

}

func CreateUser(c *fiber.Ctx) error {
	db := database.DB
	user := new(model.User)

	err := c.BodyParser(user)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}

	err = db.Create(&user).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not create task", "data": err})
	}

	// Return the created task.
	return c.JSON(fiber.Map{"status": "success", "message": "User created", "data": user})
}
