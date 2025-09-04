package main

import (
	"log"
	"swadiq-schools/app/config"
	"swadiq-schools/app/routes/auth"
	"swadiq-schools/app/routes/dashboard"
	"swadiq-schools/app/routes/students"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/template/html/v2"
)

func main() {
	// Initialize database
	config.InitDB()
	defer config.GetDB().Close()

	// Initialize template engine
	engine := html.New("./app/templates", ".html")

	// Create Fiber app
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	// Middleware
	app.Use(logger.New())
	app.Use(cors.New())

	// Static files
	app.Static("/static", "./static")
	app.Get("/favicon.ico", func(c *fiber.Ctx) error {
		return c.SendFile("./static/favicon.ico")
	})

	// Routes
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Redirect("/auth/login")
	})

	// Setup auth routes
	auth.SetupAuthRoutes(app)

	// Setup dashboard routes
	dashboard.SetupDashboardRoutes(app)

	// Setup students routes
	students.SetupStudentsRoutes(app)

	// Start server
	log.Println("Server starting on :8080")
	log.Fatal(app.Listen(":8080"))
}
