package students

import (
	"database/sql"
	"fmt"
	"strconv"
	"strings"
	"time"
)

// GenerateStudentID generates a unique student ID in the format STU-YYYY-XXX
func GenerateStudentID(db *sql.DB) (string, error) {
	currentYear := time.Now().Year()

	// Get the next sequential number for this year
	nextNumber, err := getNextStudentNumber(db, currentYear)
	if err != nil {
		return "", err
	}

	// Format: STU-2024-001
	return fmt.Sprintf("STU-%d-%03d", currentYear, nextNumber), nil
}

// getNextStudentNumber finds the highest student number for the given year and returns the next one
func getNextStudentNumber(db *sql.DB, year int) (int, error) {
	// Query to find the highest student number for the current year
	// Student IDs are in format STU-YYYY-XXX, so we need to extract the XXX part
	query := `
		SELECT student_id
		FROM students
		WHERE student_id LIKE $1
		AND is_active = true
		ORDER BY student_id DESC
		LIMIT 1
	`

	yearPrefix := fmt.Sprintf("STU-%d-%%", year)

	var lastStudentID string
	err := db.QueryRow(query, yearPrefix).Scan(&lastStudentID)

	if err == sql.ErrNoRows {
		// No students found for this year, start with 001
		return 1, nil
	} else if err != nil {
		return 0, err
	}

	// Extract the number part from the student ID (STU-YYYY-XXX)
	parts := strings.Split(lastStudentID, "-")
	if len(parts) != 3 {
		// Invalid format, start with 001
		return 1, nil
	}

	lastNumber, err := strconv.Atoi(parts[2])
	if err != nil {
		// Invalid number format, start with 001
		return 1, nil
	}

	// Return the next number
	return lastNumber + 1, nil
}
