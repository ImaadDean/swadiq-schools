package classes

import (
	"swadiq-schools/app/config"
	"swadiq-schools/app/database"
	"swadiq-schools/app/models"

	"github.com/gofiber/fiber/v2"
)

func GetClassesAPI(c *fiber.Ctx) error {
	classes, err := database.GetAllClasses(config.GetDB())
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to fetch classes"})
	}

	return c.JSON(fiber.Map{
		"classes": classes,
		"count":   len(classes),
	})
}

func CreateClassAPI(c *fiber.Ctx) error {
	type CreateClassRequest struct {
		Name      string `json:"name"`
		TeacherID string `json:"teacher_id"`
	}

	var req CreateClassRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}

	if req.Name == "" {
		return c.Status(400).JSON(fiber.Map{"error": "Class name is required"})
	}

	class := &models.Class{
		Name: req.Name,
	}

	if req.TeacherID != "" {
		class.TeacherID = &req.TeacherID
	}

	if err := database.CreateClass(config.GetDB(), class); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to create class"})
	}

	return c.Status(201).JSON(fiber.Map{
		"message": "Class created successfully",
		"class":   class,
	})
}
