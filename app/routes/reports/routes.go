package reports

import (
	"swadiq-schools/app/routes/auth"

	"github.com/gofiber/fiber/v2"
)

func SetupReportsRoutes(app *fiber.App) {
	reports := app.Group("/reports")
	reports.Use(auth.AuthMiddleware)

	reports.Get("/", ReportsDashboardHandler)
}
