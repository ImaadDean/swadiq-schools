package dashboard

import (
	"github.com/gofiber/fiber/v2"
)

// SetupDashboardRoutes sets up dashboard routes
func SetupDashboardRoutes(app *fiber.App) {
	app.Get("/dashboard", GetDashboard)
}
