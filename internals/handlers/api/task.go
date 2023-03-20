package apiHandler

import (
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/shokHorizon/kursik/database"
	"github.com/shokHorizon/kursik/internals/model"
)

func GetTasks(c *fiber.Ctx) error {
	db := database.DB
	var tasks []model.Task

	db.Preload("Tags").Order("sequence_id asc").Find(&tasks)

	if len(tasks) == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No tasks present", "data": nil})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Tasks Found", "data": tasks})
}

func GetTask(c *fiber.Ctx) error {
	db := database.DB
	var task model.Task

	id := c.Params("id")
	db.Preload("Tags").Find(&task, "id =?", id)

	if task.ID == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No task present", "data": nil})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Task Found", "data": task})
}

func RemoveTask(c *fiber.Ctx) error {
	db := database.DB
	var task model.Task

	id := c.Params("id")

	db.Find(&task, "id = ?", id)

	if task.ID == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No task present", "data": nil})
	}

	err := db.Delete(&task, "id = ?", id).Error

	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Failed to delete task", "data": nil})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Task removed"})
}

func UpdateTask(c *fiber.Ctx) error {
	db := database.DB
	var task model.Task

	var updateTaskData model.Task
	err := c.BodyParser(&updateTaskData)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}

	db.Find(&task, "id = ?", updateTaskData.ID)

	if task.ID == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No user present", "data": nil})
	}

	task.Description = updateTaskData.Description
	task.Tests = updateTaskData.Tests
	task.Name = updateTaskData.Name
	task.Tags = updateTaskData.Tags
	task.CourseID = updateTaskData.CourseID
	task.SequenceID = updateTaskData.SequenceID

	db.Save(&task)

	return c.JSON(fiber.Map{"status": "success", "message": "Task updated", "data": task})

}

func CreateTask(c *fiber.Ctx) error {
	db := database.DB
	task := new(model.Task)

	err := c.BodyParser(task)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}

	err = db.Create(&task).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not create task", "data": err})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Task created", "data": task})
}

func AddTagsToTask(c *fiber.Ctx) error {
	db := database.DB
	taskID := c.Params("id")

	tagsString := c.Query("tags")
	if tagsString == "" {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "tagsProvided": tagsString})
	}
	tagsNumbers := strings.Split(tagsString, ",")
	tags := make([]model.Tag, 0, len(tagsNumbers))
	for _, v := range tagsNumbers {
		if vInt, err := strconv.ParseUint(v, 10, 64); err != nil {
			return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
		} else {
			tags = append(tags, model.Tag{ID: vInt})
		}
	}

	var task model.Task
	db.Find(&task, "id =?", taskID)

	if task.ID == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No task present", "data": nil})
	}

	db.Model(&task).Association("Tags").Append(tags)
	//db.Save(&task)
	return c.JSON(fiber.Map{"status": "success", "message": "tags added to task", "data": task})

}

// func GetByTags(c *fiber.Ctx) error {
// 	db := database.DB
// 	var tasks []model.Task
// 	// var tags []model.Tag

// 	type Tag struct {
// 		ID uint64 `json:"tag_id"`
// 	}

// 	var tag Tag

// 	err := c.BodyParser(&tag)
// 	if err != nil {
// 		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
// 	}

// 	db.Where("id = (?)", db.Table("tags").Select("task_id").Where("id = ?", tag.ID)).Find(&tasks)
// 	if len(tasks) == 0 {
// 		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No tasks present", "data": nil})
// 	}

//		return c.JSON(fiber.Map{"status": "success", "message": "Tasks Found", "data": tasks})
//	}
func GetTasksByTags(c *fiber.Ctx) error {
	db := database.DB
	var tasks []model.Task

	tagsString := c.Query("tags")
	if tagsString == "" {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "tagsProvided": tagsString})
	}
	tagsNumbers := strings.Split(tagsString, ",")
	tags := make([]model.Tag, 0, len(tagsNumbers))
	for _, v := range tagsNumbers {
		if vInt, err := strconv.ParseUint(v, 10, 64); err != nil {
			return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
		} else {
			tags = append(tags, model.Tag{ID: vInt})
		}
	}

	db.Model(&tags).Association("Tasks").Find(&tasks)

	return c.JSON(fiber.Map{"status": "success", "message": "tags added to task", "data": tasks})
}

func RemoveTagFromTask(c *fiber.Ctx) error {
	db := database.DB
	taskID := c.Params("id")

	tagsString := c.Query("tags")
	if tagsString == "" {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "tagsProvided": tagsString})
	}
	tagsNumbers := strings.Split(tagsString, ",")
	tags := make([]model.Tag, 0, len(tagsNumbers))
	for _, v := range tagsNumbers {
		if vInt, err := strconv.ParseUint(v, 10, 64); err != nil {
			return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
		} else {
			tags = append(tags, model.Tag{ID: vInt})
		}
	}

	var task model.Task
	db.Find(&task, "id =?", taskID)

	if task.ID == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No task present", "data": nil})
	}

	db.Model(&task).Association("Tags").Delete(tags)
	db.Preload("Tags").Find(&task, "id =?", taskID)

	return c.JSON(fiber.Map{"status": "success", "message": "tags deleted form task", "data": task})
}

func ReplaceSequence(c *fiber.Ctx) error {
	db := database.DB
	taskID := c.Params("id")

	sequenceString := c.Query("sequence")
	if sequenceString == "" {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "tagsProvided": sequenceString})
	}
	sequence, err := strconv.ParseUint(sequenceString, 10, 64)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}

	taskOne := new(model.Task)
	taskTwo := new(model.Task)

	db.Find(&taskOne, "id = ?", taskID)
	if taskOne == nil {
		return c.Status(404).JSON(fiber.Map{"status": "not found", "message": "Task not found", "data": taskOne})
	}
	db.Find(&taskTwo, "sequence_id = ?", sequence)
	if taskTwo == nil {
		return c.Status(404).JSON(fiber.Map{"status": "not found", "message": "Task with old sequence not found", "data": taskTwo})
	}

	taskOne.SequenceID, taskTwo.SequenceID = taskTwo.SequenceID, taskOne.SequenceID

	db.Model(&taskOne).UpdateColumn("sequence_id", taskOne.SequenceID)
	db.Model(&taskTwo).UpdateColumn("sequence_id", taskTwo.SequenceID)
	return c.JSON(fiber.Map{"status": "success", "message": "Sequence replaced", "data": taskOne})

}
