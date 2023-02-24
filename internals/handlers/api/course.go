package apiHandler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/shokHorizon/kursik/database"
	"github.com/shokHorizon/kursik/internals/model"
)

func GetCourses(c *fiber.Ctx) error {
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

func GetCourse(c *fiber.Ctx) error {
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

func RemoveCourse(c *fiber.Ctx) error {
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

func UpdateCourse(c *fiber.Ctx) error {
	type updateUser struct {
		Login          string `gorm:"unique" json:"login" form:"login"`
		HashedPassword string `json:"hashedPassword" form:"password"`
		AccessLevel    uint16 `json:"accessLevel"`
	}

	db := database.DB
	var user model.User

	// Read the param id.
	id := c.Params("id")

	// Find the user with the given id.
	db.Find(&user, "id = ?", id)

	//If have no user.
	if user.ID == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No user present", "data": nil})
	}

	var updateUserData updateUser
	err := c.BodyParser(&updateUserData)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}

	//Edit user.
	user.Login = updateUserData.Login
	user.HashedPassword = updateUserData.HashedPassword
	user.AccessLevel = updateUserData.AccessLevel

	//Save changes.
	db.Save(&user)

	//Return updated message.
	return c.JSON(fiber.Map{"status": "success", "message": "User updated", "data": user})

}

func CreateCourse(c *fiber.Ctx) error {
	db := database.DB
	course := new(model.Course)

	err := c.BodyParser(course)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}

	err = db.Create(&course).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not create course", "data": err})
	}

	// Return the created task.
	return c.JSON(fiber.Map{"status": "success", "message": "Course created", "data": course})
}
