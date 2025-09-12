package attendance

import (
	"swadiq-schools/app/config"
	"swadiq-schools/app/database"
	"swadiq-schools/app/models"
	"time"

	"github.com/gofiber/fiber/v2"
)

func GetAttendanceByClassAPI(c *fiber.Ctx) error {
	classID := c.Params("classId")
	if classID == "" {
		return c.Status(400).JSON(fiber.Map{"error": "Class ID is required"})
	}

	students, err := database.GetStudentsByClass(config.GetDB(), classID)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to fetch students"})
	}

	return c.JSON(fiber.Map{
		"students": students,
		"count":    len(students),
	})
}

func GetAttendanceByClassAndDateAPI(c *fiber.Ctx) error {
	classID := c.Params("classId")
	dateStr := c.Params("date")

	if classID == "" || dateStr == "" {
		return c.Status(400).JSON(fiber.Map{"error": "Class ID and date are required"})
	}

	// Parse date
	date, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid date format. Use YYYY-MM-DD"})
	}

	attendanceRecords, err := database.GetAttendanceByClassAndDate(config.GetDB(), classID, date)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to fetch attendance records"})
	}

	return c.JSON(fiber.Map{
		"attendance": attendanceRecords,
		"count":      len(attendanceRecords),
		"date":       dateStr,
		"class_id":   classID,
	})
}

func CreateOrUpdateAttendanceAPI(c *fiber.Ctx) error {
	type AttendanceRequest struct {
		StudentID string `json:"student_id" validate:"required,uuid"`
		ClassID   string `json:"class_id" validate:"required,uuid"`
		Date      string `json:"date" validate:"required"`
		Status    string `json:"status" validate:"required,oneof=present absent late"`
	}

	var req AttendanceRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request body"})
	}

	// Parse date
	date, err := time.Parse("2006-01-02", req.Date)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid date format. Use YYYY-MM-DD"})
	}

	// Validate status
	var status models.AttendanceStatus
	switch req.Status {
	case "present":
		status = models.Present
	case "absent":
		status = models.Absent
	case "late":
		status = models.Late
	default:
		return c.Status(400).JSON(fiber.Map{"error": "Invalid status. Must be present, absent, or late"})
	}

	attendance := &models.Attendance{
		StudentID: req.StudentID,
		ClassID:   req.ClassID,
		Date:      date,
		Status:    status,
	}

	if err := database.CreateOrUpdateAttendance(config.GetDB(), attendance); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to save attendance record"})
	}

	return c.JSON(fiber.Map{
		"message":    "Attendance record saved successfully",
		"attendance": attendance,
	})
}

func GetAttendanceStatsAPI(c *fiber.Ctx) error {
	classID := c.Params("classId")
	if classID == "" {
		return c.Status(400).JSON(fiber.Map{"error": "Class ID is required"})
	}

	// Get date range from query parameters (default to current month)
	startDateStr := c.Query("start_date")
	endDateStr := c.Query("end_date")

	var startDate, endDate time.Time
	var err error

	if startDateStr == "" {
		// Default to start of current month
		now := time.Now()
		startDate = time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, now.Location())
	} else {
		startDate, err = time.Parse("2006-01-02", startDateStr)
		if err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "Invalid start_date format. Use YYYY-MM-DD"})
		}
	}

	if endDateStr == "" {
		// Default to end of current month
		now := time.Now()
		endDate = time.Date(now.Year(), now.Month()+1, 0, 23, 59, 59, 0, now.Location())
	} else {
		endDate, err = time.Parse("2006-01-02", endDateStr)
		if err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "Invalid end_date format. Use YYYY-MM-DD"})
		}
	}

	stats, err := database.GetAttendanceStats(config.GetDB(), classID, startDate, endDate)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to fetch attendance statistics"})
	}

	return c.JSON(fiber.Map{
		"stats":      stats,
		"class_id":   classID,
		"start_date": startDate.Format("2006-01-02"),
		"end_date":   endDate.Format("2006-01-02"),
	})
}

// Batch attendance update for multiple students
func BatchUpdateAttendanceAPI(c *fiber.Ctx) error {
	type BatchAttendanceRequest struct {
		ClassID string `json:"class_id" validate:"required,uuid"`
		Date    string `json:"date" validate:"required"`
		Records []struct {
			StudentID string `json:"student_id" validate:"required,uuid"`
			Status    string `json:"status" validate:"required,oneof=present absent late"`
		} `json:"records" validate:"required,min=1"`
	}

	var req BatchAttendanceRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request body"})
	}

	// Parse date
	date, err := time.Parse("2006-01-02", req.Date)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid date format. Use YYYY-MM-DD"})
	}

	var successCount int
	var errors []string

	for _, record := range req.Records {
		// Validate status
		var status models.AttendanceStatus
		switch record.Status {
		case "present":
			status = models.Present
		case "absent":
			status = models.Absent
		case "late":
			status = models.Late
		default:
			errors = append(errors, "Invalid status for student "+record.StudentID)
			continue
		}

		attendance := &models.Attendance{
			StudentID: record.StudentID,
			ClassID:   req.ClassID,
			Date:      date,
			Status:    status,
		}

		if err := database.CreateOrUpdateAttendance(config.GetDB(), attendance); err != nil {
			errors = append(errors, "Failed to save attendance for student "+record.StudentID)
		} else {
			successCount++
		}
	}

	response := fiber.Map{
		"message":       "Batch attendance update completed",
		"success_count": successCount,
		"total_records": len(req.Records),
	}

	if len(errors) > 0 {
		response["errors"] = errors
		response["error_count"] = len(errors)
	}

	return c.JSON(response)
}
