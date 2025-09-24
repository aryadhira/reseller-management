package main

import (
	"log"

	"github.com/aryadhira/reseller-management/internal/docs"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"

	"github.com/aryadhira/reseller-management/internal/config"
	"github.com/aryadhira/reseller-management/internal/database"
	"github.com/aryadhira/reseller-management/internal/middleware"
	"github.com/aryadhira/reseller-management/internal/routes"
)

func main() {
	// Initialize configuration
	cfg := config.LoadConfig()

	// Initialize database
	db, err := database.ConnectDB(cfg)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Initialize Fiber app
	app := fiber.New(fiber.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		},
	})

	// Setup middleware
	middleware.SetupMiddleware(app)

	// Additional middlewares
	app.Use(logger.New())
	app.Use(cors.New())

	// Setup routes
	routes.SetupRoutes(app, db)

	// Setup Swagger documentation
	docs.RegisterSwaggerRoutes(app)

	// Start server
	port := cfg.AppPort
	if port == "" {
		port = "3000"
	}

	log.Printf("Server starting on port %s", port)
	log.Fatal(app.Listen(":" + port))
}
