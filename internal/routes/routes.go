package routes

import (
	"github.com/aryadhira/reseller-management/internal/handlers"
	"github.com/aryadhira/reseller-management/internal/middleware"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App, authHandler *handlers.AuthHandler, userHandler *handlers.UserHandler, tokenSecret string) {
	// Public routes
	api := app.Group("/api/v1")
	auth := api.Group("/auth")
	auth.Post("/register", authHandler.Register)
	auth.Post("/login", authHandler.Login)

	// Protected routes
	protected := api.Group("/users", middleware.AuthValidation(tokenSecret))
	protected.Get("/profile", userHandler.GetProfile)

	// Health check
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"status": "ok"})
	})
}
