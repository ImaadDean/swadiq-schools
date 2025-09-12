package database

import (
	"database/sql"
	"fmt"
	"swadiq-schools/app/models"
	"time"

	"golang.org/x/crypto/bcrypt"
)

// hashPassword hashes a password using bcrypt
func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func GetUserByEmail(db *sql.DB, email string) (*models.User, error) {
	user := &models.User{}
	query := `SELECT id, email, password, first_name, last_name, is_active, created_at, updated_at 
			  FROM users WHERE email = $1 AND is_active = true`
	
	err := db.QueryRow(query, email).Scan(
		&user.ID, &user.Email, &user.Password, &user.FirstName, 
		&user.LastName, &user.IsActive, &user.CreatedAt, &user.UpdatedAt,
	)
	
	if err != nil {
		return nil, err
	}
	return user, nil
}

func GetUserByID(db *sql.DB, userID string) (*models.User, error) {
	user := &models.User{}
	query := `SELECT id, email, password, first_name, last_name, is_active, created_at, updated_at
			  FROM users WHERE id = $1 AND is_active = true`

	err := db.QueryRow(query, userID).Scan(
		&user.ID, &user.Email, &user.Password, &user.FirstName,
		&user.LastName, &user.IsActive, &user.CreatedAt, &user.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}
	return user, nil
}

func GetUserRoles(db *sql.DB, userID string) ([]*models.Role, error) {
	query := `
		SELECT r.id, r.name
		FROM roles r
		JOIN user_roles ur ON r.id = ur.role_id
		WHERE ur.user_id = $1
	`
	rows, err := db.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var roles []*models.Role
	for rows.Next() {
		var role models.Role
		if err := rows.Scan(&role.ID, &role.Name); err != nil {
			return nil, err
		}
		roles = append(roles, &role)
	}
	return roles, nil
}

func CreateSession(db *sql.DB, sessionID string, userID string, expiresAt time.Time) error {
	query := `INSERT INTO sessions (id, user_id, expires_at, created_at) VALUES ($1, $2, $3, $4)`
	_, err := db.Exec(query, sessionID, userID, expiresAt, time.Now())
	return err
}

func GetSessionByID(db *sql.DB, sessionID string) (*models.Session, error) {
	session := &models.Session{}
	query := `SELECT id, user_id, expires_at, created_at FROM sessions WHERE id = $1 AND expires_at > NOW()`
	
	err := db.QueryRow(query, sessionID).Scan(
		&session.ID, &session.UserID, &session.ExpiresAt, &session.CreatedAt,
	)
	
	if err != nil {
		return nil, err
	}
	return session, nil
}

func DeleteSession(db *sql.DB, sessionID string) error {
	query := `DELETE FROM sessions WHERE id = $1`
	_, err := db.Exec(query, sessionID)
	return err
}

func UpdateUserPassword(db *sql.DB, userID string, hashedPassword string) error {
	query := `UPDATE users SET password = $1, updated_at = NOW() WHERE id = $2`
	_, err := db.Exec(query, hashedPassword, userID)
	return err
}

func GetAllStudents(db *sql.DB) ([]models.Student, error) {
	// Simple query first to check if table exists
	query := `SELECT s.id, s.student_id, s.first_name, s.last_name, s.date_of_birth, 
			  s.gender, s.address, s.class_id, s.is_active, s.created_at, s.updated_at
			  FROM students s 
			  WHERE s.is_active = true ORDER BY s.created_at DESC`
	
	rows, err := db.Query(query)
	if err != nil {
		// Return empty slice if table doesn't exist
		return []models.Student{}, nil
	}
	defer rows.Close()

	var students []models.Student
	for rows.Next() {
		var student models.Student
		
		err := rows.Scan(
			&student.ID, &student.StudentID, &student.FirstName, &student.LastName,
			&student.DateOfBirth, &student.Gender, &student.Address,
			&student.ClassID, &student.IsActive, &student.CreatedAt, &student.UpdatedAt,
		)
		if err != nil {
			continue
		}
		
		students = append(students, student)
	}
	return students, nil
}

func CreateStudent(db *sql.DB, student *models.Student) error {
	query := `INSERT INTO students (student_id, first_name, last_name, date_of_birth, 
			  gender, address, class_id) 
			  VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id, created_at, updated_at`
	
	err := db.QueryRow(query, student.StudentID, student.FirstName, student.LastName,
		student.DateOfBirth, student.Gender, student.Address,
		student.ClassID).Scan(&student.ID, &student.CreatedAt, &student.UpdatedAt)
	
	return err
}

func LinkStudentToParent(db *sql.DB, studentID string, parentID string, relationship string) error {
	query := `INSERT INTO student_parents (student_id, parent_id, relationship) VALUES ($1, $2, $3)`
	_, err := db.Exec(query, studentID, parentID, relationship)
	return err
}

func GetAllParents(db *sql.DB) ([]*models.Parent, error) {
	query := `SELECT id, first_name, last_name, phone, email, address, is_active, created_at, updated_at 
			  FROM parents WHERE is_active = true ORDER BY first_name, last_name`
	
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var parents []*models.Parent
	for rows.Next() {
		parent := &models.Parent{}
		err := rows.Scan(
			&parent.ID, &parent.FirstName, &parent.LastName, &parent.Phone, 
			&parent.Email, &parent.Address, &parent.IsActive, &parent.CreatedAt, &parent.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		parents = append(parents, parent)
	}
	
	return parents, nil
}

func CreateParent(db *sql.DB, parent *models.Parent) error {
	query := `INSERT INTO parents (id, first_name, last_name, phone, email, address, is_active, created_at, updated_at) 
			  VALUES (uuid_generate_v4(), $1, $2, $3, $4, $5, true, NOW(), NOW()) 
			  RETURNING id, created_at, updated_at`
	
	err := db.QueryRow(query, parent.FirstName, parent.LastName, parent.Phone, parent.Email, parent.Address).Scan(
		&parent.ID, &parent.CreatedAt, &parent.UpdatedAt,
	)
	
	if err != nil {
		return err
	}
	
	parent.IsActive = true
	return nil
}

func GetAllClasses(db *sql.DB) ([]*models.Class, error) {
	query := `SELECT c.id, c.name, c.teacher_id, c.is_active, c.created_at, c.updated_at,
			  u.first_name, u.last_name, u.email
			  FROM classes c
			  LEFT JOIN users u ON c.teacher_id = u.id
			  WHERE c.is_active = true ORDER BY c.name`

	rows, err := db.Query(query)
	if err != nil {
		return []*models.Class{}, nil
	}
	defer rows.Close()

	var classes []*models.Class
	for rows.Next() {
		class := &models.Class{}
		var teacherID *int
		var teacherFirstName, teacherLastName, teacherEmail *string

		err := rows.Scan(
			&class.ID, &class.Name, &teacherID,
			&class.IsActive, &class.CreatedAt, &class.UpdatedAt,
			&teacherFirstName, &teacherLastName, &teacherEmail,
		)
		if err != nil {
			continue
		}

		// Convert teacher ID to string if exists
		if teacherID != nil {
			teacherIDStr := fmt.Sprintf("%d", *teacherID)
			class.TeacherID = &teacherIDStr

			// Set teacher info if exists
			if teacherFirstName != nil && teacherLastName != nil {
				class.Teacher = &models.User{
					ID:        teacherIDStr,
					FirstName: *teacherFirstName,
					LastName:  *teacherLastName,
					Email:     *teacherEmail,
				}
			}
		}

		classes = append(classes, class)
	}

	if classes == nil {
		classes = []*models.Class{}
	}

	return classes, nil
}

func CreateClass(db *sql.DB, class *models.Class) error {
	var teacherID *int
	if class.TeacherID != nil && *class.TeacherID != "" {
		// Convert string teacher ID to integer
		var tid int
		if err := db.QueryRow("SELECT id FROM users WHERE id = $1", *class.TeacherID).Scan(&tid); err != nil {
			return err
		}
		teacherID = &tid
	}

	query := `INSERT INTO classes (name, teacher_id, is_active, created_at, updated_at)
			  VALUES ($1, $2, true, NOW(), NOW())
			  RETURNING id, created_at, updated_at`

	err := db.QueryRow(query, class.Name, teacherID).Scan(
		&class.ID, &class.CreatedAt, &class.UpdatedAt,
	)

	if err != nil {
		return err
	}

	class.IsActive = true
	return nil
}

// Teacher-related functions
func GetAllTeachers(db *sql.DB) ([]*models.User, error) {
	query := `SELECT u.id, u.email, u.first_name, u.last_name, u.is_active, u.created_at, u.updated_at
			  FROM users u
			  INNER JOIN user_roles ur ON u.id = ur.user_id
			  INNER JOIN roles r ON ur.role_id = r.id
			  WHERE r.name IN ('class_teacher', 'subject_teacher')
			  AND u.is_active = true
			  ORDER BY u.first_name, u.last_name`

	rows, err := db.Query(query)
	if err != nil {
		// Return empty slice instead of error for better UX
		return []*models.User{}, nil
	}
	defer rows.Close()

	var teachers []*models.User
	for rows.Next() {
		teacher := &models.User{}
		err := rows.Scan(
			&teacher.ID, &teacher.Email, &teacher.FirstName, &teacher.LastName,
			&teacher.IsActive, &teacher.CreatedAt, &teacher.UpdatedAt,
		)
		if err != nil {
			continue // Skip invalid rows instead of failing
		}
		teachers = append(teachers, teacher)
	}

	// Ensure we always return a valid slice
	if teachers == nil {
		teachers = []*models.User{}
	}

	return teachers, nil
}

func CreateTeacher(db *sql.DB, user *models.User) error {
	// Hash password before storing
	hashedPassword, err := hashPassword(user.Password)
	if err != nil {
		return err
	}

	// Create user account
	query := `INSERT INTO users (email, password, first_name, last_name, phone, is_active, created_at, updated_at)
			  VALUES ($1, $2, $3, $4, $5, true, NOW(), NOW())
			  RETURNING id, created_at, updated_at`

	var userID int
	err = db.QueryRow(query, user.Email, hashedPassword, user.FirstName, user.LastName, user.Phone).Scan(
		&userID, &user.CreatedAt, &user.UpdatedAt,
	)

	if err != nil {
		return err
	}

	// Convert integer ID to string for the model
	user.ID = fmt.Sprintf("%d", userID)

	if err != nil {
		return err
	}

	// Assign class_teacher role by default
	roleQuery := `INSERT INTO user_roles (user_id, role_id, created_at)
				  SELECT $1, r.id, NOW()
				  FROM roles r
				  WHERE r.name = 'class_teacher'`

	_, err = db.Exec(roleQuery, userID)
	if err != nil {
		return err
	}

	user.IsActive = true
	return nil
}

// Department-related functions
func GetAllDepartments(db *sql.DB) ([]*models.Department, error) {
	query := `SELECT d.id, d.name, d.code, d.description, d.head_of_department_id, d.assistant_head_id,
			  d.is_active, d.created_at, d.updated_at,
			  h.first_name as head_first_name, h.last_name as head_last_name, h.email as head_email,
			  a.first_name as assistant_first_name, a.last_name as assistant_last_name, a.email as assistant_email
			  FROM departments d
			  LEFT JOIN users h ON d.head_of_department_id = h.id
			  LEFT JOIN users a ON d.assistant_head_id = a.id
			  WHERE d.is_active = true ORDER BY d.name`

	rows, err := db.Query(query)
	if err != nil {
		return []*models.Department{}, nil
	}
	defer rows.Close()

	var departments []*models.Department
	for rows.Next() {
		department := &models.Department{}
		var headFirstName, headLastName, headEmail *string
		var assistantFirstName, assistantLastName, assistantEmail *string

		err := rows.Scan(
			&department.ID, &department.Name, &department.Code, &department.Description,
			&department.HeadOfDepartmentID, &department.AssistantHeadID,
			&department.IsActive, &department.CreatedAt, &department.UpdatedAt,
			&headFirstName, &headLastName, &headEmail,
			&assistantFirstName, &assistantLastName, &assistantEmail,
		)
		if err != nil {
			continue
		}

		// Set head of department if exists
		if headFirstName != nil && headLastName != nil && department.HeadOfDepartmentID != nil {
			department.HeadOfDepartment = &models.User{
				ID:        *department.HeadOfDepartmentID,
				FirstName: *headFirstName,
				LastName:  *headLastName,
				Email:     *headEmail,
			}
		}

		// Set assistant head if exists
		if assistantFirstName != nil && assistantLastName != nil && department.AssistantHeadID != nil {
			department.AssistantHead = &models.User{
				ID:        *department.AssistantHeadID,
				FirstName: *assistantFirstName,
				LastName:  *assistantLastName,
				Email:     *assistantEmail,
			}
		}

		departments = append(departments, department)
	}

	if departments == nil {
		departments = []*models.Department{}
	}

	return departments, nil
}

// Subject-related functions
func GetAllSubjects(db *sql.DB) ([]*models.Subject, error) {
	query := `SELECT s.id, s.name, s.code, s.department_id, s.is_active, s.created_at, s.updated_at,
			  d.name as department_name
			  FROM subjects s
			  LEFT JOIN departments d ON s.department_id = d.id
			  WHERE s.is_active = true ORDER BY s.name`

	rows, err := db.Query(query)
	if err != nil {
		return []*models.Subject{}, nil
	}
	defer rows.Close()

	var subjects []*models.Subject
	for rows.Next() {
		subject := &models.Subject{}
		var departmentName *string
		err := rows.Scan(
			&subject.ID, &subject.Name, &subject.Code, &subject.DepartmentID,
			&subject.IsActive, &subject.CreatedAt, &subject.UpdatedAt, &departmentName,
		)
		if err != nil {
			continue
		}

		// Set department if exists
		if departmentName != nil && subject.DepartmentID != nil {
			subject.Department = &models.Department{
				ID:   *subject.DepartmentID,
				Name: *departmentName,
			}
		}

		subjects = append(subjects, subject)
	}

	if subjects == nil {
		subjects = []*models.Subject{}
	}

	return subjects, nil
}

// Attendance-related functions
func GetAttendanceByClassAndDate(db *sql.DB, classID string, date time.Time) ([]*models.Attendance, error) {
	query := `SELECT a.id, a.student_id, a.class_id, a.date, a.status, a.created_at, a.updated_at,
			  s.student_id as student_number, s.first_name, s.last_name
			  FROM attendance a
			  INNER JOIN students s ON a.student_id = s.id
			  WHERE a.class_id = $1 AND a.date = $2
			  ORDER BY s.first_name, s.last_name`

	rows, err := db.Query(query, classID, date)
	if err != nil {
		return []*models.Attendance{}, nil
	}
	defer rows.Close()

	var attendanceRecords []*models.Attendance
	for rows.Next() {
		attendance := &models.Attendance{
			Student: &models.Student{},
		}
		err := rows.Scan(
			&attendance.ID, &attendance.StudentID, &attendance.ClassID, &attendance.Date, &attendance.Status,
			&attendance.CreatedAt, &attendance.UpdatedAt,
			&attendance.Student.StudentID, &attendance.Student.FirstName, &attendance.Student.LastName,
		)
		if err != nil {
			return nil, err
		}
		attendanceRecords = append(attendanceRecords, attendance)
	}

	return attendanceRecords, nil
}

func GetStudentsByClass(db *sql.DB, classID string) ([]*models.Student, error) {
	query := `SELECT id, student_id, first_name, last_name, date_of_birth, gender, address, class_id, is_active, created_at, updated_at
			  FROM students
			  WHERE class_id = $1 AND is_active = true
			  ORDER BY first_name, last_name`

	rows, err := db.Query(query, classID)
	if err != nil {
		return []*models.Student{}, nil
	}
	defer rows.Close()

	var students []*models.Student
	for rows.Next() {
		student := &models.Student{}
		err := rows.Scan(
			&student.ID, &student.StudentID, &student.FirstName, &student.LastName,
			&student.DateOfBirth, &student.Gender, &student.Address, &student.ClassID,
			&student.IsActive, &student.CreatedAt, &student.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		students = append(students, student)
	}

	return students, nil
}

func CreateOrUpdateAttendance(db *sql.DB, attendance *models.Attendance) error {
	// First check if attendance record exists for this student, class, and date
	var existingID string
	checkQuery := `SELECT id FROM attendance WHERE student_id = $1 AND class_id = $2 AND date = $3`
	err := db.QueryRow(checkQuery, attendance.StudentID, attendance.ClassID, attendance.Date).Scan(&existingID)

	if err == sql.ErrNoRows {
		// Create new record
		insertQuery := `INSERT INTO attendance (student_id, class_id, date, status, created_at, updated_at)
						VALUES ($1, $2, $3, $4, NOW(), NOW())
						RETURNING id, created_at, updated_at`
		err = db.QueryRow(insertQuery, attendance.StudentID, attendance.ClassID, attendance.Date, attendance.Status).Scan(
			&attendance.ID, &attendance.CreatedAt, &attendance.UpdatedAt,
		)
		return err
	} else if err != nil {
		return err
	} else {
		// Update existing record
		updateQuery := `UPDATE attendance SET status = $1, updated_at = NOW() WHERE id = $2`
		_, err = db.Exec(updateQuery, attendance.Status, existingID)
		attendance.ID = existingID
		return err
	}
}

func GetAttendanceStats(db *sql.DB, classID string, startDate, endDate time.Time) (map[string]interface{}, error) {
	query := `SELECT
				COUNT(*) as total_records,
				COUNT(CASE WHEN status = 'present' THEN 1 END) as present_count,
				COUNT(CASE WHEN status = 'absent' THEN 1 END) as absent_count,
				COUNT(CASE WHEN status = 'late' THEN 1 END) as late_count
			  FROM attendance
			  WHERE class_id = $1 AND date BETWEEN $2 AND $3`

	var total, present, absent, late int
	err := db.QueryRow(query, classID, startDate, endDate).Scan(&total, &present, &absent, &late)
	if err != nil {
		return nil, err
	}

	stats := map[string]interface{}{
		"total":   total,
		"present": present,
		"absent":  absent,
		"late":    late,
	}

	if total > 0 {
		stats["present_percentage"] = float64(present) / float64(total) * 100
		stats["absent_percentage"] = float64(absent) / float64(total) * 100
		stats["late_percentage"] = float64(late) / float64(total) * 100
	}

	return stats, nil
}

// Exam-related functions
func GetAllExams(db *sql.DB) ([]*models.Exam, error) {
	query := `SELECT e.id, e.name, e.class_id, e.start_date, e.end_date, e.is_active, e.created_at, e.updated_at,
			  c.name as class_name
			  FROM exams e
			  LEFT JOIN classes c ON e.class_id = c.id
			  WHERE e.is_active = true ORDER BY e.start_date DESC`

	rows, err := db.Query(query)
	if err != nil {
		return []*models.Exam{}, nil
	}
	defer rows.Close()

	var exams []*models.Exam
	for rows.Next() {
		exam := &models.Exam{
			Class: &models.Class{},
		}
		err := rows.Scan(
			&exam.ID, &exam.Name, &exam.ClassID, &exam.StartDate, &exam.EndDate,
			&exam.IsActive, &exam.CreatedAt, &exam.UpdatedAt,
			&exam.Class.Name,
		)
		if err != nil {
			return nil, err
		}
		exams = append(exams, exam)
	}

	return exams, nil
}

func GetExamsByClass(db *sql.DB, classID string) ([]*models.Exam, error) {
	query := `SELECT e.id, e.name, e.class_id, e.start_date, e.end_date, e.is_active, e.created_at, e.updated_at,
			  c.name as class_name
			  FROM exams e
			  LEFT JOIN classes c ON e.class_id = c.id
			  WHERE e.class_id = $1 AND e.is_active = true ORDER BY e.start_date DESC`

	rows, err := db.Query(query, classID)
	if err != nil {
		return []*models.Exam{}, nil
	}
	defer rows.Close()

	var exams []*models.Exam
	for rows.Next() {
		exam := &models.Exam{
			Class: &models.Class{},
		}
		err := rows.Scan(
			&exam.ID, &exam.Name, &exam.ClassID, &exam.StartDate, &exam.EndDate,
			&exam.IsActive, &exam.CreatedAt, &exam.UpdatedAt,
			&exam.Class.Name,
		)
		if err != nil {
			return nil, err
		}
		exams = append(exams, exam)
	}

	return exams, nil
}

func CreateExam(db *sql.DB, exam *models.Exam) error {
	query := `INSERT INTO exams (name, class_id, start_date, end_date, is_active, created_at, updated_at)
			  VALUES ($1, $2, $3, $4, true, NOW(), NOW())
			  RETURNING id, created_at, updated_at`

	err := db.QueryRow(query, exam.Name, exam.ClassID, exam.StartDate, exam.EndDate).Scan(
		&exam.ID, &exam.CreatedAt, &exam.UpdatedAt,
	)

	if err != nil {
		return err
	}

	exam.IsActive = true
	return nil
}

func GetExamByID(db *sql.DB, examID string) (*models.Exam, error) {
	exam := &models.Exam{
		Class: &models.Class{},
	}
	query := `SELECT e.id, e.name, e.class_id, e.start_date, e.end_date, e.is_active, e.created_at, e.updated_at,
			  c.name as class_name
			  FROM exams e
			  LEFT JOIN classes c ON e.class_id = c.id
			  WHERE e.id = $1 AND e.is_active = true`

	err := db.QueryRow(query, examID).Scan(
		&exam.ID, &exam.Name, &exam.ClassID, &exam.StartDate, &exam.EndDate,
		&exam.IsActive, &exam.CreatedAt, &exam.UpdatedAt,
		&exam.Class.Name,
	)

	if err != nil {
		return nil, err
	}

	return exam, nil
}

// Paper-related functions
func GetAllPapers(db *sql.DB) ([]*models.Paper, error) {
	query := `SELECT p.id, p.subject_id, p.teacher_id, p.name, p.code, p.is_active, p.created_at, p.updated_at,
			  s.name as subject_name, s.code as subject_code,
			  u.first_name, u.last_name, u.email
			  FROM papers p
			  LEFT JOIN subjects s ON p.subject_id = s.id
			  LEFT JOIN users u ON p.teacher_id = u.id
			  WHERE p.is_active = true ORDER BY s.name, p.name`

	rows, err := db.Query(query)
	if err != nil {
		return []*models.Paper{}, nil
	}
	defer rows.Close()

	var papers []*models.Paper
	for rows.Next() {
		paper := &models.Paper{
			Subject: &models.Subject{},
			Teacher: &models.User{},
		}
		var teacherFirstName, teacherLastName, teacherEmail sql.NullString

		err := rows.Scan(
			&paper.ID, &paper.SubjectID, &paper.TeacherID, &paper.Name, &paper.Code,
			&paper.IsActive, &paper.CreatedAt, &paper.UpdatedAt,
			&paper.Subject.Name, &paper.Subject.Code,
			&teacherFirstName, &teacherLastName, &teacherEmail,
		)
		if err != nil {
			return nil, err
		}

		if teacherFirstName.Valid {
			paper.Teacher.FirstName = teacherFirstName.String
			paper.Teacher.LastName = teacherLastName.String
			paper.Teacher.Email = teacherEmail.String
		} else {
			paper.Teacher = nil
		}

		papers = append(papers, paper)
	}

	return papers, nil
}

func CreatePaper(db *sql.DB, paper *models.Paper) error {
	query := `INSERT INTO papers (subject_id, teacher_id, name, code, is_active, created_at, updated_at)
			  VALUES ($1, $2, $3, $4, true, NOW(), NOW())
			  RETURNING id, created_at, updated_at`

	err := db.QueryRow(query, paper.SubjectID, paper.TeacherID, paper.Name, paper.Code).Scan(
		&paper.ID, &paper.CreatedAt, &paper.UpdatedAt,
	)

	if err != nil {
		return err
	}

	paper.IsActive = true
	return nil
}

func CreateSubject(db *sql.DB, subject *models.Subject) error {
	query := `INSERT INTO subjects (name, code, department_id, is_active, created_at, updated_at)
			  VALUES ($1, $2, $3, true, NOW(), NOW())
			  RETURNING id, created_at, updated_at`

	err := db.QueryRow(query, subject.Name, subject.Code, subject.DepartmentID).Scan(
		&subject.ID, &subject.CreatedAt, &subject.UpdatedAt,
	)

	if err != nil {
		return err
	}

	subject.IsActive = true
	return nil
}

func GetSubjectsByDepartment(db *sql.DB, departmentID string) ([]*models.Subject, error) {
	query := `SELECT id, name, code, department_id, is_active, created_at, updated_at
			  FROM subjects WHERE department_id = $1 AND is_active = true ORDER BY name`

	rows, err := db.Query(query, departmentID)
	if err != nil {
		return []*models.Subject{}, nil
	}
	defer rows.Close()

	var subjects []*models.Subject
	for rows.Next() {
		subject := &models.Subject{}
		err := rows.Scan(
			&subject.ID, &subject.Name, &subject.Code, &subject.DepartmentID,
			&subject.IsActive, &subject.CreatedAt, &subject.UpdatedAt,
		)
		if err != nil {
			continue
		}
		subjects = append(subjects, subject)
	}

	if subjects == nil {
		subjects = []*models.Subject{}
	}

	return subjects, nil
}