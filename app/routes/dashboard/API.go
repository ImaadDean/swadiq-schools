package dashboard

import (
	"github.com/gofiber/fiber/v2"
)

// GetDashboard handles dashboard page
func GetDashboard(c *fiber.Ctx) error {
	return c.Render("dashboard/index", fiber.Map{
		"Title":       "Dashboard - Swadiq Schools",
		"CurrentPage": "dashboard",
	})
}
