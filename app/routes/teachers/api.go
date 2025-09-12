package teachers

import (
	"swadiq-schools/app/config"
	"swadiq-schools/app/database"
	"swadiq-schools/app/models"

	"github.com/gofiber/fiber/v2"
)

func GetTeachersAPI(c *fiber.Ctx) error {
	teachers, err := database.GetAllTeachers(config.GetDB())
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to fetch teachers"})
	}

	return c.JSON(fiber.Map{
		"teachers": teachers,
		"count":    len(teachers),
	})
}

func CreateTeacherAPI(c *fiber.Ctx) error {
	type CreateTeacherRequest struct {
		FirstName    string   `json:"first_name"`
		LastName     string   `json:"last_name"`
		Email        string   `json:"email"`
		Password     string   `json:"password"`
		Phone        string   `json:"phone"`
		DepartmentID string   `json:"department_id"`
		SubjectIDs   []string `json:"subject_ids"`
	}

	var req CreateTeacherRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}

	if req.FirstName == "" || req.LastName == "" || req.Email == "" || req.Password == "" {
		return c.Status(400).JSON(fiber.Map{"error": "First name, last name, email, and password are required"})
	}

	// Create user account for teacher
	user := &models.User{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Email:     req.Email,
		Password:  req.Password, // This will be hashed in the database function
		Phone:     req.Phone,
	}

	if err := database.CreateTeacher(config.GetDB(), user); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Failed to create teacher",
			"details": err.Error(),
		})
	}

	return c.Status(201).JSON(fiber.Map{
		"message": "Teacher created successfully",
		"teacher": user,
		"department_id": req.DepartmentID,
		"subject_ids": req.SubjectIDs,
	})
}

func GetDepartmentsAPI(c *fiber.Ctx) error {
	departments, err := database.GetAllDepartments(config.GetDB())
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to fetch departments"})
	}

	return c.JSON(fiber.Map{
		"departments": departments,
		"count":       len(departments),
	})
}

func SearchTeachersAPI(c *fiber.Ctx) error {
	query := c.Query("q", "")
	limit := c.QueryInt("limit", 10)

	teachers, err := SearchTeachers(config.GetDB(), query, limit)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to search teachers"})
	}

	return c.JSON(fiber.Map{
		"teachers": teachers,
		"count":    len(teachers),
		"query":    query,
	})
}

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
