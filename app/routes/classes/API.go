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

func GetClassAPI(c *fiber.Ctx) error {
	classID := c.Params("id")
	if classID == "" {
		return c.Status(400).JSON(fiber.Map{"error": "Class ID is required"})
	}

	class, err := GetClassByID(config.GetDB(), classID)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Class not found"})
	}

	return c.JSON(fiber.Map{
		"class": class,
	})
}

func UpdateClassAPI(c *fiber.Ctx) error {
	classID := c.Params("id")
	if classID == "" {
		return c.Status(400).JSON(fiber.Map{"error": "Class ID is required"})
	}

	type UpdateClassRequest struct {
		Name      string `json:"name"`
		TeacherID string `json:"teacher_id"`
	}

	var req UpdateClassRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}

	if req.Name == "" {
		return c.Status(400).JSON(fiber.Map{"error": "Class name is required"})
	}

	// Check if class exists
	existingClass, err := GetClassByID(config.GetDB(), classID)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Class not found"})
	}

	// Update class data
	existingClass.Name = req.Name
	if req.TeacherID != "" {
		existingClass.TeacherID = &req.TeacherID
	} else {
		existingClass.TeacherID = nil
	}

	if err := UpdateClass(config.GetDB(), existingClass); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to update class"})
	}

	return c.JSON(fiber.Map{
		"message": "Class updated successfully",
		"class":   existingClass,
	})
}

func DeleteClassAPI(c *fiber.Ctx) error {
	classID := c.Params("id")
	if classID == "" {
		return c.Status(400).JSON(fiber.Map{"error": "Class ID is required"})
	}

	// Check if class exists
	_, err := GetClassByID(config.GetDB(), classID)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Class not found"})
	}

	// TODO: Check if class has students before deleting
	// For now, we'll do a soft delete
	if err := DeleteClass(config.GetDB(), classID); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to delete class"})
	}

	return c.JSON(fiber.Map{
		"message": "Class deleted successfully",
	})
}
