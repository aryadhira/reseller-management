package routes

import (
	"github.com/aryadhira/reseller-management/internal/handlers"
	"github.com/aryadhira/reseller-management/internal/services"
	"github.com/aryadhira/reseller-management/internal/repository"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func SetupRoutes(app *fiber.App, db *gorm.DB) {
	// Initialize repository
	repo := repository.NewRepository(db)
	
	// Initialize services
	serviceInstance := services.NewService(repo)
	
	// Initialize handlers
	resellerHandler := handlers.NewResellerHandler(serviceInstance.Reseller)
	productHandler := handlers.NewProductHandler(serviceInstance.Product)
	orderHandler := handlers.NewOrderHandler(serviceInstance.Order)
	paymentHandler := handlers.NewPaymentHandler(serviceInstance.Payment)
	
	// API routes
	api := app.Group("/api/v1")
	
	// Reseller routes
	resellers := api.Group("/resellers")
	resellers.Post("/", resellerHandler.CreateReseller)
	resellers.Get("/", resellerHandler.GetAllResellers)
	resellers.Get("/:id", resellerHandler.GetResellerByID)
	resellers.Get("/:id/profile", resellerHandler.GetResellerWithOrders) // Detailed profile with order history
	resellers.Put("/:id", resellerHandler.UpdateReseller)
	resellers.Delete("/:id", resellerHandler.DeleteReseller)
	
	// Product routes
	products := api.Group("/products")
	products.Post("/", productHandler.CreateProduct)
	products.Get("/", productHandler.GetAllProducts)
	products.Get("/:id", productHandler.GetProductByID)
	products.Put("/:id", productHandler.UpdateProduct)
	products.Delete("/:id", productHandler.DeleteProduct)
	products.Post("/:id/restock", productHandler.RestockProduct)
	products.Get("/low-stock", productHandler.GetLowStockProducts)
	
	// Order routes
	orders := api.Group("/orders")
	orders.Post("/", orderHandler.CreateOrder)
	orders.Get("/", orderHandler.GetAllOrders)
	orders.Get("/:id", orderHandler.GetOrderByID)
	orders.Put("/:id", orderHandler.UpdateOrder)
	orders.Delete("/:id", orderHandler.DeleteOrder)
	orders.Patch("/:id/cancel", orderHandler.CancelOrder)
	
	// Payment and financial routes
	payments := api.Group("/payments")
	payments.Get("/", paymentHandler.GetAllPayments)
	payments.Get("/order/:orderID", paymentHandler.GetPaymentByOrderID)
	payments.Post("/order/:orderID/pay", paymentHandler.RecordPayment)
	
	transactions := api.Group("/transactions")
	transactions.Get("/", paymentHandler.GetAllTransactions)
	transactions.Post("/cash-in", paymentHandler.RecordCashIn)
	transactions.Post("/cash-out", paymentHandler.RecordCashOut)
	
	// Balance routes
	balance := api.Group("/balance")
	balance.Put("/", paymentHandler.UpdateBalance)
	balance.Get("/", paymentHandler.GetBalance)
	
	// Dashboard route
	api.Get("/dashboard", paymentHandler.GetDashboardData)
}