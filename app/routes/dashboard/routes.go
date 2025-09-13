package dashboard

import (
	"swadiq-schools/app/routes/auth"

	"github.com/gofiber/fiber/v2"
)

// SetupDashboardRoutes sets up dashboard routes
func SetupDashboardRoutes(app *fiber.App) {
	// Dashboard page
	app.Get("/dashboard", auth.AuthMiddleware, GetDashboard)

	// Dashboard API routes
	app.Get("/api/dashboard/stats", auth.AuthMiddleware, GetDashboardStatsAPI)
}
