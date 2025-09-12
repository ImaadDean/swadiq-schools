package students

import (
	"swadiq-schools/app/config"
	"swadiq-schools/app/database"
	"swadiq-schools/app/models"
	"time"

	"github.com/gofiber/fiber/v2"
)

func GetStudentsAPI(c *fiber.Ctx) error {
	students, err := database.GetAllStudents(config.GetDB())
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to fetch students"})
	}

	return c.JSON(fiber.Map{
		"students": students,
		"count":    len(students),
	})
}

func CreateStudentAPI(c *fiber.Ctx) error {
	type CreateStudentRequest struct {
		StudentID     string `json:"student_id"`
		FirstName     string `json:"first_name"`
		LastName      string `json:"last_name"`
		DateOfBirth   string `json:"date_of_birth"`
		Gender        string `json:"gender"`
		Address       string `json:"address"`
		ClassID       string `json:"class_id"`
		ParentID      string `json:"parent_id"`
	}

	var req CreateStudentRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}

	if req.StudentID == "" || req.FirstName == "" || req.LastName == "" {
		return c.Status(400).JSON(fiber.Map{"error": "Student ID, first name, and last name are required"})
	}

	// Create student
	student := &models.Student{
		StudentID: req.StudentID,
		FirstName: req.FirstName,
		LastName:  req.LastName,
	}

	if req.DateOfBirth != "" {
		if parsedDate, err := time.Parse("2006-01-02", req.DateOfBirth); err == nil {
			student.DateOfBirth = &parsedDate
		}
	}
	if req.Gender != "" {
		gender := models.Gender(req.Gender)
		student.Gender = &gender
	}
	if req.Address != "" {
		student.Address = &req.Address
	}
	if req.ClassID != "" {
		student.ClassID = &req.ClassID
	}

	if err := database.CreateStudent(config.GetDB(), student); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to create student"})
	}

	// Link student to parent if provided
	if req.ParentID != "" {
		if err := database.LinkStudentToParent(config.GetDB(), student.ID, req.ParentID, "Parent"); err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "Failed to link student to parent"})
		}
	}

	return c.Status(201).JSON(fiber.Map{
		"message": "Student created successfully",
		"student": student,
	})
}
