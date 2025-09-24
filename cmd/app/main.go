package main

import (
	"log"

	"github.com/aryadhira/reseller-management/internal/config"
	"github.com/aryadhira/reseller-management/internal/database"
	"github.com/aryadhira/reseller-management/internal/handlers"
	"github.com/aryadhira/reseller-management/internal/repository"
	"github.com/aryadhira/reseller-management/internal/routes"
	"github.com/aryadhira/reseller-management/internal/services"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	// initialize config
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	// database initialization
	dbConfig := &database.DBConfig{
		Host:     cfg.DBHost,
		Port:     cfg.DBPort,
		Username: cfg.DBUsername,
		Password: cfg.DBPassword,
		DBName:   cfg.DBName,
	}

	db, err := database.NewDB(dbConfig)
	if err != nil {
		log.Fatal(err)
	}

	// repository initialization
	userRepo := repository.NewUserRepository(db)

	// service initialization
	authService := services.NewAuthService(userRepo)
	userService := services.NewUserService(userRepo)

	// handler initialization
	authHandler := handlers.NewAuthHandler(authService, cfg.TokenSecret, cfg.TokenExpired)
	userHandler := handlers.NewUserHandler(userService)

	app := fiber.New(fiber.Config{
		AppName: "Reseller Management",
	})

	// middleware
	app.Use(recover.New())
	app.Use(logger.New())

	// Setup Routes
	routes.SetupRoutes(app, authHandler, userHandler, cfg.TokenSecret)

	// Start server
	log.Printf("Server starting on port %s", cfg.AppPort)
	log.Fatal(app.Listen(cfg.AppHost + ":" + cfg.AppPort))

}
