package dashboard

import (
	"swadiq-schools/app/models"

	"github.com/gofiber/fiber/v2"
)

// GetDashboard handles dashboard page
func GetDashboard(c *fiber.Ctx) error {
	// Get user from context (set by auth middleware)
	user := c.Locals("user").(*models.User)

	return c.Render("dashboard/index", fiber.Map{
		"Title":       "Dashboard - Swadiq Schools",
		"CurrentPage": "dashboard",
		"FirstName":   user.FirstName,
		"LastName":    user.LastName,
		"Email":       user.Email,
		"user":        user,
	})
}
