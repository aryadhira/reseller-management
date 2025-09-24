package handlers

import (
	"github.com/aryadhira/reseller-management/internal/interfaces"
	"github.com/aryadhira/reseller-management/internal/models"
	"github.com/gofiber/fiber/v2"
)

// OrderHandler handles order-related requests
type OrderHandler struct {
	Service interfaces.OrderService
}

func NewOrderHandler(service interfaces.OrderService) *OrderHandler {
	return &OrderHandler{Service: service}
}

// CreateOrder creates a new order
// @Summary Create a new order
// @Description Create a new order for a selected reseller with multiple items
// @Tags Order Management
// @Accept json
// @Produce json
// @Param order body models.Order true "Order data"
// @Success 201 {object} models.Order
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /orders [post]
func (h *OrderHandler) CreateOrder(c *fiber.Ctx) error {
	order := new(models.Order)
	if err := c.BodyParser(order); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request body"})
	}

	createdOrder, err := h.Service.CreateOrder(order)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(201).JSON(createdOrder)
}

// GetAllOrders gets all orders
// @Summary Get all orders
// @Description Get a list of all orders
// @Tags Order Management
// @Produce json
// @Success 200 {array} models.Order
// @Failure 500 {object} map[string]string
// @Router /orders [get]
func (h *OrderHandler) GetAllOrders(c *fiber.Ctx) error {
	orders, err := h.Service.GetAllOrders()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(orders)
}

// GetOrderByID gets an order by ID
// @Summary Get an order by ID
// @Description Get an order by its unique ID
// @Tags Order Management
// @Produce json
// @Param id path string true "Order ID"
// @Success 200 {object} models.Order
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /orders/{id} [get]
func (h *OrderHandler) GetOrderByID(c *fiber.Ctx) error {
	id := c.Params("id")
	
	order, err := h.Service.GetOrderByID(id)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Order not found"})
	}

	return c.JSON(order)
}

// UpdateOrder updates an order
// @Summary Update an order
// @Description Update an existing order
// @Tags Order Management
// @Accept json
// @Produce json
// @Param id path string true "Order ID"
// @Param order body models.Order true "Updated order data"
// @Success 200 {object} models.Order
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /orders/{id} [put]
func (h *OrderHandler) UpdateOrder(c *fiber.Ctx) error {
	id := c.Params("id")
	
	order := new(models.Order)
	if err := c.BodyParser(order); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request body"})
	}

	updatedOrder, err := h.Service.UpdateOrder(id, order)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(updatedOrder)
}

// DeleteOrder deletes an order
// @Summary Delete an order
// @Description Delete an existing order by ID
// @Tags Order Management
// @Produce json
// @Param id path string true "Order ID"
// @Success 204 {object} nil
// @Failure 500 {object} map[string]string
// @Router /orders/{id} [delete]
func (h *OrderHandler) DeleteOrder(c *fiber.Ctx) error {
	id := c.Params("id")
	
	err := h.Service.DeleteOrder(id)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.SendStatus(204)
}

// CancelOrder cancels an order
// @Summary Cancel an order
// @Description Cancel an existing order and restore stock quantities
// @Tags Order Management
// @Produce json
// @Param id path string true "Order ID"
// @Success 200 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /orders/{id}/cancel [patch]
func (h *OrderHandler) CancelOrder(c *fiber.Ctx) error {
	id := c.Params("id")
	
	err := h.Service.CancelOrder(id)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "Order cancelled successfully"})
}