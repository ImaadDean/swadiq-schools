package teachers

import (
	"swadiq-schools/app/config"
	"swadiq-schools/app/database"
	"swadiq-schools/app/models"
	"swadiq-schools/app/routes/auth"

	"github.com/gofiber/fiber/v2"
)

func SetupTeachersRoutes(app *fiber.App) {
	teachers := app.Group("/teachers")
	teachers.Use(auth.AuthMiddleware)

	// Routes
	teachers.Get("/", TeachersPage)

	// API routes
	api := app.Group("/api/teachers")
	api.Use(auth.AuthMiddleware)
	api.Get("/", GetTeachersAPI)
	api.Get("/search", SearchTeachersAPI)
	api.Post("/", CreateTeacherAPI)

	// Additional API routes for departments and subjects
	departmentsAPI := app.Group("/api/departments")
	departmentsAPI.Use(auth.AuthMiddleware)
	departmentsAPI.Get("/", GetDepartmentsAPI)

	subjectsAPI := app.Group("/api/subjects")
	subjectsAPI.Use(auth.AuthMiddleware)
	subjectsAPI.Get("/", GetSubjectsAPI)
}

func TeachersPage(c *fiber.Ctx) error {
	teachers, err := database.GetAllTeachers(config.GetDB())
	if err != nil {
		// Log the error for debugging
		println("Error getting teachers:", err.Error())
		// Initialize empty slice if there's an error
		teachers = []*models.User{}
	}

	// Ensure teachers is never nil
	if teachers == nil {
		teachers = []*models.User{}
	}

	return c.Render("teachers/index", fiber.Map{
		"Title":       "Teachers - Swadiq Schools",
		"CurrentPage": "teachers",
		"teachers":    teachers,
		"user":        c.Locals("user"),
	})
}


