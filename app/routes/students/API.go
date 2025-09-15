package students

import (
	"swadiq-schools/app/config"
	"swadiq-schools/app/database"
	"swadiq-schools/app/models"
	"time"

	"github.com/gofiber/fiber/v2"
)

func GetStudentsAPI(c *fiber.Ctx) error {
	students, err := database.GetStudentsWithDetails(config.GetDB())
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to fetch students"})
	}

	return c.JSON(fiber.Map{
		"students": students,
		"count":    len(students),
	})
}

// GetStudentsStatsAPI returns students statistics for the students page
func GetStudentsStatsAPI(c *fiber.Ctx) error {
	// Get database connection
	db := config.GetDB()

	// Get students statistics
	stats, err := database.GetStudentsStats(db)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error":   "Failed to fetch students statistics",
			"details": err.Error(),
		})
	}

	// Return statistics as JSON
	return c.JSON(fiber.Map{
		"success": true,
		"data":    stats,
	})
}

// GetStudentsTableAPI returns students formatted for table display with filtering support
func GetStudentsTableAPI(c *fiber.Ctx) error {
	// Get query parameters for filtering
	search := c.Query("search")
	status := c.Query("status")
	classID := c.Query("class_id")
	gender := c.Query("gender")
	dateFrom := c.Query("date_from")
	dateTo := c.Query("date_to")
	sortBy := c.Query("sort_by", "name")      // default to name
	sortOrder := c.Query("sort_order", "asc") // default to ascending

	// Create filter parameters
	filters := database.StudentFilters{
		Search:    search,
		Status:    status,
		ClassID:   classID,
		Gender:    gender,
		DateFrom:  dateFrom,
		DateTo:    dateTo,
		SortBy:    sortBy,
		SortOrder: sortOrder,
	}

	students, err := database.GetStudentsWithFilters(config.GetDB(), filters)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to fetch students"})
	}

	// Format students for table display
	type StudentTableData struct {
		ID          string `json:"id"`
		StudentID   string `json:"student_id"`
		FirstName   string `json:"first_name"`
		LastName    string `json:"last_name"`
		FullName    string `json:"full_name"`
		ClassName   string `json:"class_name,omitempty"`
		ParentName  string `json:"parent_name,omitempty"`
		ParentPhone string `json:"parent_phone,omitempty"`
		ParentEmail string `json:"parent_email,omitempty"`
		Status      string `json:"status"`
		Initials    string `json:"initials"`
		DateOfBirth string `json:"date_of_birth,omitempty"`
		Gender      string `json:"gender,omitempty"`
		Address     string `json:"address,omitempty"`
	}

	var tableData []StudentTableData
	for _, student := range students {
		// Create initials from first and last name
		initials := ""
		if len(student.FirstName) > 0 {
			initials += string(student.FirstName[0])
		}
		if len(student.LastName) > 0 {
			initials += string(student.LastName[0])
		}

		// Get primary parent (first one in the list)
		parentName := ""
		parentPhone := ""
		parentEmail := ""
		if len(student.Parents) > 0 {
			parent := student.Parents[0]
			parentName = parent.FirstName + " " + parent.LastName
			if parent.Phone != nil {
				parentPhone = *parent.Phone
			}
			if parent.Email != nil {
				parentEmail = *parent.Email
			}
		}

		// Get class name
		className := ""
		if student.Class != nil {
			className = student.Class.Name
		}

		// Format date of birth
		dateOfBirth := ""
		if student.DateOfBirth != nil {
			dateOfBirth = student.DateOfBirth.Format("2006-01-02")
		}

		// Format gender
		gender := ""
		if student.Gender != nil {
			gender = string(*student.Gender)
		}

		// Format address
		address := ""
		if student.Address != nil {
			address = *student.Address
		}

		// Determine status
		status := "Active"
		if !student.IsActive {
			status = "Inactive"
		}

		tableData = append(tableData, StudentTableData{
			ID:          student.ID,
			StudentID:   student.StudentID,
			FirstName:   student.FirstName,
			LastName:    student.LastName,
			FullName:    student.FirstName + " " + student.LastName,
			ClassName:   className,
			ParentName:  parentName,
			ParentPhone: parentPhone,
			ParentEmail: parentEmail,
			Status:      status,
			Initials:    initials,
			DateOfBirth: dateOfBirth,
			Gender:      gender,
			Address:     address,
		})
	}

	return c.JSON(fiber.Map{
		"students": tableData,
		"count":    len(tableData),
	})
}

// GetStudentByIDAPI returns a single student by ID
func GetStudentByIDAPI(c *fiber.Ctx) error {
	studentID := c.Params("id")
	if studentID == "" {
		return c.Status(400).JSON(fiber.Map{"error": "Student ID is required"})
	}

	student, err := database.GetStudentByID(config.GetDB(), studentID)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Student not found"})
	}

	// Format response for edit modal
	response := fiber.Map{
		"student": student,
	}

	// Add parent information if available
	if len(student.Parents) > 0 {
		parent := student.Parents[0] // Get primary parent

		// Get relationship information
		relationship, err := database.GetStudentParentRelationship(config.GetDB(), studentID, parent.ID)
		if err != nil {
			relationship = "guardian" // Default if not found
		}

		response["parent"] = fiber.Map{
			"id":           parent.ID,
			"first_name":   parent.FirstName,
			"last_name":    parent.LastName,
			"email":        parent.Email,
			"phone":        parent.Phone,
			"address":      parent.Address,
			"relationship": relationship,
		}
	}

	return c.JSON(response)
}

// GetStudentsByYearAPI returns students for a specific year
func GetStudentsByYearAPI(c *fiber.Ctx) error {
	year := c.QueryInt("year")
	if year == 0 {
		return c.Status(400).JSON(fiber.Map{"error": "Year parameter is required"})
	}

	students, err := database.GetStudentsByYear(config.GetDB(), year)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to fetch students"})
	}

	return c.JSON(fiber.Map{
		"students": students,
		"count":    len(students),
		"year":     year,
	})
}

// GetStudentsByClassAPI returns students for a specific class
func GetStudentsByClassAPI(c *fiber.Ctx) error {
	classID := c.Query("class_id")
	if classID == "" {
		return c.Status(400).JSON(fiber.Map{"error": "Class ID parameter is required"})
	}

	students, err := database.GetStudentsByClass(config.GetDB(), classID)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to fetch students"})
	}

	return c.JSON(fiber.Map{
		"students": students,
		"count":    len(students),
		"class_id": classID,
	})
}

func CreateStudentAPI(c *fiber.Ctx) error {
	type CreateStudentRequest struct {
		StudentID          string `json:"student_id"` // Optional - will be auto-generated if empty
		FirstName          string `json:"first_name"`
		LastName           string `json:"last_name"`
		DateOfBirth        string `json:"date_of_birth"`
		Gender             string `json:"gender"`
		Address            string `json:"address"`
		ClassID            string `json:"class_id"`
		ParentID           string `json:"parent_id"`
		ParentRelationship string `json:"parent_relationship"`
	}

	var req CreateStudentRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}

	// Only first name and last name are required now
	if req.FirstName == "" || req.LastName == "" {
		return c.Status(400).JSON(fiber.Map{"error": "First name and last name are required"})
	}

	// Always auto-generate student ID
	studentID, err := GenerateStudentID(config.GetDB())
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to generate student ID"})
	}

	// Create student
	student := &models.Student{
		StudentID: studentID,
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
		relationship := req.ParentRelationship
		if relationship == "" {
			relationship = "guardian" // Default relationship
		}
		if err := database.LinkStudentToParent(config.GetDB(), student.ID, req.ParentID, relationship); err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "Failed to link student to parent"})
		}
	}

	return c.Status(201).JSON(fiber.Map{
		"message": "Student created successfully",
		"student": student,
	})
}

// UpdateStudentAPI updates an existing student
func UpdateStudentAPI(c *fiber.Ctx) error {
	studentID := c.Params("id")
	if studentID == "" {
		return c.Status(400).JSON(fiber.Map{"error": "Student ID is required"})
	}

	type ParentInfo struct {
		ID           string `json:"id,omitempty"`
		FirstName    string `json:"first_name"`
		LastName     string `json:"last_name"`
		Email        string `json:"email"`
		Phone        string `json:"phone"`
		Address      string `json:"address"`
		Relationship string `json:"relationship"`
	}

	type UpdateStudentRequest struct {
		FirstName          string `json:"first_name"`
		LastName           string `json:"last_name"`
		DateOfBirth        string `json:"date_of_birth"`
		Gender             string `json:"gender"`
		Address            string `json:"address"`
		ClassID            string `json:"class_id"`
		ParentID           string `json:"parent_id"`
		ParentRelationship string `json:"parent_relationship"`
	}

	var req UpdateStudentRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request body"})
	}

	// Get existing student
	student, err := database.GetStudentByID(config.GetDB(), studentID)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Student not found"})
	}

	// Update student fields
	if req.FirstName != "" {
		student.FirstName = req.FirstName
	}
	if req.LastName != "" {
		student.LastName = req.LastName
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

	// Update student in database
	if err := database.UpdateStudent(config.GetDB(), student); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to update student"})
	}

	// Handle parent changes
	if req.ParentID != "" {
		// Change the parent for this student
		relationship := req.ParentRelationship
		if relationship == "" {
			relationship = "guardian" // Default relationship
		}
		if err := database.ChangeStudentParent(config.GetDB(), studentID, req.ParentID, relationship); err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "Failed to update parent relationship"})
		}
	} else if req.ParentID == "" && req.ParentRelationship == "" {
		// Remove parent if both are empty (this handles the remove parent case)
		if err := database.ChangeStudentParent(config.GetDB(), studentID, "", ""); err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "Failed to remove parent"})
		}
	}

	return c.JSON(fiber.Map{
		"message": "Student updated successfully",
		"student": student,
	})
}

// DeleteStudentAPI deletes a student
func DeleteStudentAPI(c *fiber.Ctx) error {
	studentID := c.Params("id")
	if studentID == "" {
		return c.Status(400).JSON(fiber.Map{"error": "Student ID is required"})
	}

	// Check if student exists
	_, err := database.GetStudentByID(config.GetDB(), studentID)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Student not found"})
	}

	// Delete student
	if err := database.DeleteStudent(config.GetDB(), studentID); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to delete student"})
	}

	return c.JSON(fiber.Map{
		"message": "Student deleted successfully",
	})
}

// GetParentsAPI returns all parents for selection
func GetParentsAPI(c *fiber.Ctx) error {
	search := c.Query("search", "")

	parents, err := database.GetParentsForSelection(config.GetDB(), search)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to fetch parents"})
	}

	return c.JSON(fiber.Map{
		"parents": parents,
		"count":   len(parents),
	})
}
