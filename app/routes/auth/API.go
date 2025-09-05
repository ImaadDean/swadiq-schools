package auth

import (
	"database/sql"
	"strconv"
	"swadiq-schools/app/config"
	"swadiq-schools/app/database"
	"time"

	"github.com/gofiber/fiber/v2"
)

func LoginAPI(c *fiber.Ctx) error {
	type LoginRequest struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	var req LoginRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}

	user, err := database.GetUserByEmail(config.GetDB(), req.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			return c.Status(401).JSON(fiber.Map{"error": "Invalid credentials"})
		}
		return c.Status(500).JSON(fiber.Map{"error": "Database error"})
	}

	if !CheckPasswordHash(req.Password, user.Password) {
		return c.Status(401).JSON(fiber.Map{"error": "Invalid credentials"})
	}

	userIDInt, err := strconv.Atoi(user.ID)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Invalid user ID"})
	}

	roles, err := database.GetUserRoles(config.GetDB(), userIDInt)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to get user roles"})
	}
	user.Roles = roles

	sessionID := GenerateSessionID()
	expiresAt := GetSessionExpiry()

	if err := database.CreateSession(config.GetDB(), sessionID, userIDInt, expiresAt); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to create session"})
	}

	c.Cookie(&fiber.Cookie{
		Name:     "session_id",
		Value:    sessionID,
		Expires:  expiresAt,
		HTTPOnly: true,
		Secure:   false, // Set to true in production with HTTPS
		SameSite: "Lax",
	})

	return c.JSON(fiber.Map{
		"message": "Login successful",
		"user":    user,
	})
}

func LogoutAPI(c *fiber.Ctx) error {
	sessionID := c.Cookies("session_id")
	if sessionID != "" {
		database.DeleteSession(config.GetDB(), sessionID)
	}

	c.Cookie(&fiber.Cookie{
		Name:     "session_id",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
	})

	return c.Redirect("/auth/login")
}

func ChangePasswordAPI(c *fiber.Ctx) error {
	type ChangePasswordRequest struct {
		CurrentPassword string `json:"current_password"`
		NewPassword     string `json:"new_password"`
	}

	var req ChangePasswordRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}

	userID := c.Locals("user_id").(int)

	// Get current user to verify current password
	user, err := database.GetUserByEmail(config.GetDB(), c.Locals("user_email").(string))
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Database error"})
	}

	if !CheckPasswordHash(req.CurrentPassword, user.Password) {
		return c.Status(400).JSON(fiber.Map{"error": "Current password is incorrect"})
	}

	hashedPassword, err := HashPassword(req.NewPassword)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to hash password"})
	}

	if err := database.UpdateUserPassword(config.GetDB(), userID, hashedPassword); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to update password"})
	}

	return c.JSON(fiber.Map{"message": "Password changed successfully"})
}

func ForgotPasswordAPI(c *fiber.Ctx) error {
	type ForgotPasswordRequest struct {
		Email       string `json:"email"`
		NewPassword string `json:"new_password,omitempty"`
	}

	var req ForgotPasswordRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}

	// Check if user exists
	user, err := database.GetUserByEmail(config.GetDB(), req.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			return c.Status(404).JSON(fiber.Map{"error": "Email not found"})
		}
		return c.Status(500).JSON(fiber.Map{"error": "Database error"})
	}

	// If no new password provided, just verify email exists
	if req.NewPassword == "" {
		return c.JSON(fiber.Map{
			"message": "Email verified",
			"user_found": true,
		})
	}

	// Hash new password
	hashedPassword, err := HashPassword(req.NewPassword)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to hash password"})
	}

	// Convert user ID to int
	userIDInt, err := strconv.Atoi(user.ID)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Invalid user ID"})
	}

	// Update password
	if err := database.UpdateUserPassword(config.GetDB(), userIDInt, hashedPassword); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to update password"})
	}

	return c.JSON(fiber.Map{"message": "Password reset successfully"})
}
