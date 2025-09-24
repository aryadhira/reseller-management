package handlers

import (
	"github.com/aryadhira/reseller-management/internal/interfaces"
	"github.com/aryadhira/reseller-management/internal/models"
	"github.com/gofiber/fiber/v2"
)

// RestockRequest represents the request to restock a product
// @Description Restock request information
type RestockRequest struct {
	// Quantity to add to the product's stock
	Quantity int `json:"quantity" example:"20"`
}

// ProductHandler handles product-related requests
type ProductHandler struct {
	Service interfaces.ProductService
}

func NewProductHandler(service interfaces.ProductService) *ProductHandler {
	return &ProductHandler{Service: service}
}

// CreateProduct creates a new product
// @Summary Create a new product
// @Description Create a new product with provided details
// @Tags Product Management
// @Accept json
// @Produce json
// @Param product body models.Product true "Product data"
// @Success 201 {object} models.Product
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /products [post]
func (h *ProductHandler) CreateProduct(c *fiber.Ctx) error {
	product := new(models.Product)
	if err := c.BodyParser(product); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request body"})
	}

	createdProduct, err := h.Service.CreateProduct(product)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(201).JSON(createdProduct)
}

// GetAllProducts gets all products
// @Summary Get all products
// @Description Get a list of all products with current quantities
// @Tags Product Management
// @Produce json
// @Success 200 {array} models.Product
// @Failure 500 {object} map[string]string
// @Router /products [get]
func (h *ProductHandler) GetAllProducts(c *fiber.Ctx) error {
	products, err := h.Service.GetAllProducts()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(products)
}

// GetProductByID gets a product by ID
// @Summary Get a product by ID
// @Description Get a product by its unique ID
// @Tags Product Management
// @Produce json
// @Param id path string true "Product ID"
// @Success 200 {object} models.Product
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /products/{id} [get]
func (h *ProductHandler) GetProductByID(c *fiber.Ctx) error {
	id := c.Params("id")
	
	product, err := h.Service.GetProductByID(id)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Product not found"})
	}

	return c.JSON(product)
}

// UpdateProduct updates a product
// @Summary Update a product
// @Description Update an existing product's details (name, price, etc.)
// @Tags Product Management
// @Accept json
// @Produce json
// @Param id path string true "Product ID"
// @Param product body models.Product true "Updated product data"
// @Success 200 {object} models.Product
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /products/{id} [put]
func (h *ProductHandler) UpdateProduct(c *fiber.Ctx) error {
	id := c.Params("id")
	
	product := new(models.Product)
	if err := c.BodyParser(product); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request body"})
	}

	updatedProduct, err := h.Service.UpdateProduct(id, product)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(updatedProduct)
}

// DeleteProduct deletes a product
// @Summary Delete a product
// @Description Delete an existing product by ID (with checks to ensure it's not used in any orders)
// @Tags Product Management
// @Produce json
// @Param id path string true "Product ID"
// @Success 204 {object} nil
// @Failure 500 {object} map[string]string
// @Router /products/{id} [delete]
func (h *ProductHandler) DeleteProduct(c *fiber.Ctx) error {
	id := c.Params("id")
	
	err := h.Service.DeleteProduct(id)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.SendStatus(204)
}

// RestockProduct restocks a product
// @Summary Restock a product
// @Description Add quantity to an existing product's stock
// @Tags Product Management
// @Accept json
// @Produce json
// @Param id path string true "Product ID"
// @Param quantity body RestockRequest true "Quantity to add"
// @Success 200 {object} models.Product
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /products/{id}/restock [post]
func (h *ProductHandler) RestockProduct(c *fiber.Ctx) error {
	id := c.Params("id")
	
	req := new(RestockRequest)
	if err := c.BodyParser(req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request body"})
	}

	updatedProduct, err := h.Service.RestockProduct(id, req.Quantity)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(updatedProduct)
}

// GetLowStockProducts gets low stock products
// @Summary Get products with low stock
// @Description Get a list of products with stock levels below the minimum alert threshold
// @Tags Product Management
// @Produce json
// @Success 200 {array} models.Product
// @Failure 500 {object} map[string]string
// @Router /products/low-stock [get]
func (h *ProductHandler) GetLowStockProducts(c *fiber.Ctx) error {
	products, err := h.Service.GetLowStockProducts()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(products)
}