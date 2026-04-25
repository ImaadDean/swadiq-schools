package reports

import (
	"swadiq-schools/app/config"
	"swadiq-schools/app/database"
	"swadiq-schools/app/models"

	"github.com/gofiber/fiber/v2"
)

func ReportsDashboardHandler(c *fiber.Ctx) error {
	user := c.Locals("user").(*models.User)
	db := config.GetDB()

	// Get dashboard statistics to show some real data on reports page
	stats, err := database.GetDashboardStats(db)
	if err != nil {
		// Fallback if stats fail
		stats = &models.DashboardStats{}
	}

	return c.Render("reports/index", fiber.Map{
		"Title":       "Analytics & Reports",
		"CurrentPage": "reports",
		"FirstName":   user.FirstName,
		"LastName":    user.LastName,
		"Email":       user.Email,
		"user":        user,
		"Stats":       stats,
	})
}
