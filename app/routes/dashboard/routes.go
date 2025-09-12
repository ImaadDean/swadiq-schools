package dashboard

import (
	"swadiq-schools/app/routes/auth"

	"github.com/gofiber/fiber/v2"
)

// SetupDashboardRoutes sets up dashboard routes
func SetupDashboardRoutes(app *fiber.App) {
	app.Get("/dashboard", auth.AuthMiddleware, GetDashboard)
}
