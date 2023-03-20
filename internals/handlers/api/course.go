package apiHandler

import (
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/shokHorizon/kursik/database"
	"github.com/shokHorizon/kursik/internals/model"
)

func GetCourses(c *fiber.Ctx) error {
	db := database.DB
	var courses []model.Course

	db.Find(&courses)

	if len(courses) == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No courses present", "data": nil})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Users Found", "data": courses})
}

func GetCourse(c *fiber.Ctx) error {
	db := database.DB
	var course model.Course

	//Read Id param.
	id := c.Params("id")
	//Find user by id.
	db.Find(&course, "id =?", id)

	//If have no user.
	if course.ID == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No user present", "data": nil})
	}

	//Return result.
	return c.JSON(fiber.Map{"status": "success", "message": "User Found", "data": course})

}

func RemoveCourse(c *fiber.Ctx) error {
	db := database.DB
	var course model.Course

	// Read the param id.
	id := c.Params("id")

	// Find the note with the given id.
	db.Find(&course, "id = ?", id)

	// If no such user present return an error.
	if course.ID == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No user present", "data": nil})
	}

	// Delete the user and return error if encountered.
	err := db.Delete(&course, "id = ?", id).Error

	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Failed to delete user", "data": nil})
	}

	// Return deleted message.
	return c.JSON(fiber.Map{"status": "success", "message": "User removed"})

}

func UpdateCourse(c *fiber.Ctx) error {

	db := database.DB
	var course model.Course

	var updateCourseData model.Course
	err := c.BodyParser(&updateCourseData)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}

	db.Find(&course, "id = ?", updateCourseData.ID)

	if course.ID == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No course present", "data": nil})
	}

	course.Title = updateCourseData.Title
	course.Tasks = updateCourseData.Tasks
	course.OwnerID = updateCourseData.OwnerID

	db.Save(&course)

	return c.JSON(fiber.Map{"status": "success", "message": "Course updated", "data": course})

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

func AddUserToCourse(c *fiber.Ctx) error {
	db := database.DB
	cource := new(model.Course)

	courceID := c.Params("id")
	userIdString := c.Query("user")
	userID, err := strconv.ParseUint(userIdString, 10, 64)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}

	// user := model.User{ID: userID}
	user := make([]model.User, 0, 1)
	user = append(user, model.User{ID: userID})

	db.Find(&cource, "id = ?", courceID)
	if cource.ID == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No cource present", "data": nil})
	}
	db.Model(&cource).Association("Users").Append(user)

	return c.JSON(fiber.Map{"status": "success", "message": "User added to cource", "data": cource})

}

func RemoveUserFromCource(c *fiber.Ctx) error {
	db := database.DB
	courseID := c.Params("id")

	usersString := c.Query("users")
	if usersString == "" {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "tagsProvided": usersString})
	}
	usersNumbers := strings.Split(usersString, ",")
	users := make([]model.User, 0, len(usersNumbers))
	for _, v := range usersNumbers {
		if vInt, err := strconv.ParseUint(v, 10, 64); err != nil {
			return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
		} else {
			users = append(users, model.User{ID: vInt})
		}
	}

	var course model.Course
	db.Find(&course, "id =?", courseID)

	if course.ID == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No cource present", "data": nil})
	}

	db.Model(&course).Association("Users").Delete(users)
	db.Preload("Users").Find(&course, "id =?", courseID)

	return c.JSON(fiber.Map{"status": "success", "message": "users deleted form cource", "data": course})
}

func AddTasksToCource(c *fiber.Ctx) error {
	db := database.DB
	courseID := c.Params("id")

	tasksString := c.Query("tasks")
	if tasksString == "" {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": tasksString})
	}
	tasksNumbers := strings.Split(tasksString, ",")
	tasks := make([]model.Task, 0, len(tasksNumbers))
	for _, v := range tasksNumbers {
		if vInt, err := strconv.ParseUint(v, 10, 64); err != nil {
			return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
		} else {
			tasks = append(tasks, model.Task{ID: vInt})
		}
	}

	var course model.Course
	db.Find(&course, "id =?", courseID)

	if course.ID == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No cource present", "data": nil})
	}

	db.Model(&course).Association("Tasks").Append(tasks)

	return c.JSON(fiber.Map{"status": "success", "message": "tasks added to cource", "data": course})
}

func GetTasksByCource(c *fiber.Ctx) error {
	db := database.DB
	var tasks []model.Task

	courseID := c.Params("id")

	db.Find(&tasks, "course_id = ?", courseID)
	if len(tasks) == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No tasks present", "data": nil})
	}
	return c.JSON(fiber.Map{"status": "success", "message": "Users Found", "data": tasks})
}

func RemoveTaskFromCourse(c *fiber.Ctx) error {
	db := database.DB
	courseID := c.Params("id")

	tasksString := c.Query("tasks")
	if tasksString == "" {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": tasksString})
	}
	tasksNumbers := strings.Split(tasksString, ",")
	tasks := make([]model.Task, 0, len(tasksNumbers))
	for _, v := range tasksNumbers {
		if vInt, err := strconv.ParseUint(v, 10, 64); err != nil {
			return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
		} else {
			tasks = append(tasks, model.Task{ID: vInt})
		}
	}

	var course model.Course
	db.Find(&course, "id =?", courseID)

	if course.ID == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No course present", "data": nil})
	}

	db.Model(&course).Association("Tasks").Delete(tasks)

	return c.JSON(fiber.Map{"status": "success", "message": "tasks deleted form course", "data": course})
}
