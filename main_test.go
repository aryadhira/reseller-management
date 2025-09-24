package main

import (
	"bytes"
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/aryadhira/reseller-management/internal/config"
	"github.com/aryadhira/reseller-management/internal/database"
	"github.com/aryadhira/reseller-management/internal/middleware"
	"github.com/aryadhira/reseller-management/internal/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestAppInitialization(t *testing.T) {
	// Initialize configuration
	cfg := config.LoadConfig()

	// Initialize database (in-memory for testing)
	db, err := database.ConnectDB(cfg)
	assert.NoError(t, err)

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

	// Setup routes
	routes.SetupRoutes(app, db)

	// Test a basic endpoint
	req := httptest.NewRequest("GET", "/api/v1/resellers", nil)
	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode)
}

func TestResellerCRUD(t *testing.T) {
	// Initialize configuration
	cfg := config.LoadConfig()

	// Initialize database (in-memory for testing)
	db, err := database.ConnectDB(cfg)
	assert.NoError(t, err)

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

	// Setup routes
	routes.SetupRoutes(app, db)

	// Test creating a reseller
	resellerData := map[string]interface{}{
		"name":    "Test Reseller",
		"email":   "test@example.com",
		"phone":   "1234567890",
		"address": "123 Test St",
	}

	jsonData, _ := json.Marshal(resellerData)
	req := httptest.NewRequest("POST", "/api/v1/resellers", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, 201, resp.StatusCode)
}
