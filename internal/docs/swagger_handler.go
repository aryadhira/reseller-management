package docs

import (
	"github.com/gofiber/swagger"
	"github.com/gofiber/fiber/v2"
)

// RegisterSwaggerRoutes adds the swagger route to the application
func RegisterSwaggerRoutes(app *fiber.App) {
	app.Get("/swagger/*", swagger.HandlerDefault) 
}