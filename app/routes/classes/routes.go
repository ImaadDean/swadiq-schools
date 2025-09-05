package classes

import (
	"swadiq-schools/app/routes/auth"

	"github.com/gofiber/fiber/v2"
)

func SetupClassesRoutes(app *fiber.App) {
	classes := app.Group("/classes")
	classes.Use(auth.AuthMiddleware)

	// Routes
	classes.Get("/", ClassesPage)
}

func ClassesPage(c *fiber.Ctx) error {
	return c.Render("classes/index", fiber.Map{
		"Title":       "Classes Management - Swadiq Schools",
		"CurrentPage": "classes",
		"user":        c.Locals("user"),
	})
}
