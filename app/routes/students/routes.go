package students

import (
	"swadiq-schools/app/config"
	"swadiq-schools/app/database"
	"swadiq-schools/app/routes/auth"

	"github.com/gofiber/fiber/v2"
)

func SetupStudentsRoutes(app *fiber.App) {
	students := app.Group("/students")
	students.Use(auth.AuthMiddleware)

	// Routes
	students.Get("/", StudentsPage)
	students.Get("/add", AddStudentPage)

	// API routes
	api := app.Group("/api/students")
	api.Use(auth.AuthMiddleware)
	api.Get("/", GetStudentsAPI)
	api.Post("/", CreateStudentAPI)
}

func StudentsPage(c *fiber.Ctx) error {
	students, err := database.GetAllStudents(config.GetDB())
	if err != nil {
		return c.Status(500).Render("error", fiber.Map{
			"Title":        "Error - Swadiq Schools",
			"CurrentPage":  "students",
			"ErrorCode":    "500",
			"ErrorTitle":   "Database Error",
			"ErrorMessage": "Failed to load students. Please try again later.",
			"ShowRetry":    true,
			"user":         c.Locals("user"),
		})
	}

	return c.Render("students/index", fiber.Map{
		"Title":       "Students - Swadiq Schools",
		"CurrentPage": "students",
		"students":    students,
		"user":        c.Locals("user"),
	})
}

func AddStudentPage(c *fiber.Ctx) error {
	return c.Render("students/add", fiber.Map{
		"Title":       "Add Student - Swadiq Schools",
		"CurrentPage": "students",
		"user":        c.Locals("user"),
	})
}
