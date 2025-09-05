package auth

import (
	"database/sql"
	"swadiq-schools/app/config"
	"swadiq-schools/app/database"
	"swadiq-schools/app/models"

	"github.com/gofiber/fiber/v2"
)

func SetupAuthRoutes(app *fiber.App) {
	auth := app.Group("/auth")

	// Public routes
	auth.Get("/login", ShowLoginPage)
	auth.Post("/login", LoginAPI)
	auth.Post("/logout", LogoutAPI)
	auth.Get("/forgot-password", ShowForgotPasswordPage)
	auth.Post("/forgot-password", ForgotPasswordAPI)

	// Protected routes
	auth.Use(AuthMiddleware)
	auth.Get("/profile", ShowProfilePage)
	auth.Post("/change-password", ChangePasswordAPI)
}

func ShowLoginPage(c *fiber.Ctx) error {
	// Check if already logged in
	if sessionID := c.Cookies("session_id"); sessionID != "" {
		if _, err := database.GetSessionByID(config.GetDB(), sessionID); err == nil {
			return c.Redirect("/dashboard")
		}
	}

	return c.Render("auth/login", fiber.Map{
		"Title": "Login - Swadiq Schools",
	})
}

func ShowForgotPasswordPage(c *fiber.Ctx) error {
	return c.Render("auth/forgot-password", fiber.Map{
		"Title": "Forgot Password - Swadiq Schools",
	})
}

func ShowProfilePage(c *fiber.Ctx) error {
	return c.Render("auth/profile", fiber.Map{
		"Title":     "Profile - Swadiq Schools",
		"User":      c.Locals("user"),
		"FirstName": c.Locals("user_first_name"),
		"LastName":  c.Locals("user_last_name"),
		"Email":     c.Locals("user_email"),
		"Roles":     c.Locals("user_roles"),
	})
}

// AuthMiddleware validates session and sets user context
func AuthMiddleware(c *fiber.Ctx) error {
	sessionID := c.Cookies("session_id")
	if sessionID == "" {
		return c.Status(401).JSON(fiber.Map{"error": "No session found"})
	}

	session, err := database.GetSessionByID(config.GetDB(), sessionID)
	if err != nil {
		if err == sql.ErrNoRows {
			return c.Status(401).JSON(fiber.Map{"error": "Invalid session"})
		}
		return c.Status(500).JSON(fiber.Map{"error": "Database error"})
	}

	// Get user details by ID
	user, err := database.GetUserByID(config.GetDB(), session.UserID)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "User not found"})
	}

	// Get user roles
	roles, err := database.GetUserRoles(config.GetDB(), session.UserID)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to get user roles"})
	}

	// Set user context
	c.Locals("user_id", user.ID)
	c.Locals("user_email", user.Email)
	c.Locals("user_first_name", user.FirstName)
	c.Locals("user_last_name", user.LastName)
	c.Locals("user_roles", roles)
	c.Locals("user", user)

	return c.Next()
}

// RoleMiddleware checks if user has required role
func RoleMiddleware(allowedRoles ...string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		userRoles := c.Locals("user_roles").([]*models.Role)
		
		for _, userRole := range userRoles {
			for _, allowedRole := range allowedRoles {
				if userRole.Name == allowedRole {
					return c.Next()
				}
			}
		}
		
		return c.Status(403).JSON(fiber.Map{"error": "Insufficient permissions"})
	}
}

