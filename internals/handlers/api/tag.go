package apiHandler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/shokHorizon/kursik/database"
	"github.com/shokHorizon/kursik/internals/model"
)

func GetTags(c *fiber.Ctx) error {
	db := database.DB
	var tags []model.Tag

	db.Find(&tags)

	if len(tags) == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No tags present", "data": nil})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Tags Found", "data": tags})
}

func GetTag(c *fiber.Ctx) error {
	db := database.DB
	var tag model.Tag

	id := c.Params("id")
	db.Find(&tag, "id =?", id)

	if tag.ID == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No tag present", "data": nil})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Tag Found", "data": tag})

}

func RemoveTag(c *fiber.Ctx) error {
	db := database.DB
	var tag model.Tag

	id := c.Params("id")

	db.Find(&tag, "id = ?", id)

	if tag.ID == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No tag present", "data": nil})
	}

	err := db.Delete(&tag, "id = ?", id).Error

	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Failed to delete tag", "data": nil})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Tag removed"})

}

func UpdateTag(c *fiber.Ctx) error {
	db := database.DB
	var tag model.Tag

	var updateTagData model.Tag
	err := c.BodyParser(&updateTagData)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}

	db.Find(&tag, "id = ?", updateTagData.ID)

	if tag.ID == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No tag present", "data": nil})
	}

	tag.Name = updateTagData.Name
	// tag.Tasks = updateTagData.Tasks

	db.Save(&tag)

	return c.JSON(fiber.Map{"status": "success", "message": "Tag updated", "data": tag})

}

func CreateTag(c *fiber.Ctx) error {
	db := database.DB
	tag := new(model.Tag)

	err := c.BodyParser(tag)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}

	err = db.Create(&tag).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not create tag", "data": err})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Tag created", "data": tag})
}

func FindTagByName(c *fiber.Ctx) error {
	db := database.DB
	tagString := c.Query("name")
	tag := new(model.Tag)

	db.Find(&tag, "name = ?", tagString)
	if tag == nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No tag present", "data": nil})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Tag found", "data": tag})
}
