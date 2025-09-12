package subjects

import (
	"swadiq-schools/app/config"
	"swadiq-schools/app/database"
	"swadiq-schools/app/models"

	"github.com/gofiber/fiber/v2"
)

func GetSubjectsAPI(c *fiber.Ctx) error {
	departmentID := c.Query("department_id")
	
	var subjects []*models.Subject
	var err error
	
	if departmentID != "" {
		subjects, err = database.GetSubjectsByDepartment(config.GetDB(), departmentID)
	} else {
		subjects, err = database.GetAllSubjects(config.GetDB())
	}
	
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to fetch subjects"})
	}

	return c.JSON(fiber.Map{
		"subjects": subjects,
		"count":    len(subjects),
	})
}

func CreateSubjectAPI(c *fiber.Ctx) error {
	type CreateSubjectRequest struct {
		Name         string `json:"name"`
		Code         string `json:"code"`
		DepartmentID string `json:"department_id"`
	}

	var req CreateSubjectRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}

	if req.Name == "" || req.Code == "" {
		return c.Status(400).JSON(fiber.Map{"error": "Name and code are required"})
	}

	subject := &models.Subject{
		Name: req.Name,
		Code: req.Code,
	}

	if req.DepartmentID != "" {
		subject.DepartmentID = &req.DepartmentID
	}

	if err := database.CreateSubject(config.GetDB(), subject); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error":   "Failed to create subject",
			"details": err.Error(),
		})
	}

	return c.Status(201).JSON(fiber.Map{
		"message": "Subject created successfully",
		"subject": subject,
	})
}
