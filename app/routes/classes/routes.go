package classes

import (
	"swadiq-schools/app/config"
	"swadiq-schools/app/database"
	"swadiq-schools/app/models"
	"swadiq-schools/app/routes/auth"

	"github.com/gofiber/fiber/v2"
)

func SetupClassesRoutes(app *fiber.App) {
	classes := app.Group("/classes")
	classes.Use(auth.AuthMiddleware)

	// Routes
	classes.Get("/", ClassesPage)

	// API routes (these were already set up in main.go, but let's make them explicit here too)
	api := app.Group("/api/classes")
	api.Use(auth.AuthMiddleware)
	api.Get("/", GetClassesAPI)
	api.Post("/", CreateClassAPI)
	api.Get("/:id", GetClassAPI)
	api.Put("/:id", UpdateClassAPI)
	api.Delete("/:id", DeleteClassAPI)
}

func ClassesPage(c *fiber.Ctx) error {
	classes, err := database.GetAllClasses(config.GetDB())
	if err != nil {
		// Log the error for debugging
		println("Error getting classes:", err.Error())
		// Initialize empty slice if there's an error
		classes = []*models.Class{}
	}

	// Ensure classes is never nil
	if classes == nil {
		classes = []*models.Class{}
	}

	return c.Render("classes/index", fiber.Map{
		"Title":       "Classes Management - Swadiq Schools",
		"CurrentPage": "classes",
		"classes":     classes,
		"user":        c.Locals("user"),
	})
}
