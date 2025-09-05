package database

import (
	"database/sql"
	"swadiq-schools/app/models"
	"time"
)

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

func GetUserByID(db *sql.DB, userID int) (*models.User, error) {
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

func GetUserRoles(db *sql.DB, userID int) ([]*models.Role, error) {
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

func CreateSession(db *sql.DB, sessionID string, userID int, expiresAt time.Time) error {
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

func UpdateUserPassword(db *sql.DB, userID int, hashedPassword string) error {
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
	query := `SELECT id, name, teacher_id, is_active, created_at, updated_at 
			  FROM classes WHERE is_active = true ORDER BY name`
	
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var classes []*models.Class
	for rows.Next() {
		class := &models.Class{}
		err := rows.Scan(
			&class.ID, &class.Name, &class.TeacherID, 
			&class.IsActive, &class.CreatedAt, &class.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		classes = append(classes, class)
	}
	
	return classes, nil
}

func CreateClass(db *sql.DB, class *models.Class) error {
	query := `INSERT INTO classes (id, name, teacher_id, is_active, created_at, updated_at) 
			  VALUES (uuid_generate_v4(), $1, $2, true, NOW(), NOW()) 
			  RETURNING id, created_at, updated_at`
	
	err := db.QueryRow(query, class.Name, class.TeacherID).Scan(
		&class.ID, &class.CreatedAt, &class.UpdatedAt,
	)
	
	if err != nil {
		return err
	}
	
	class.IsActive = true
	return nil
}