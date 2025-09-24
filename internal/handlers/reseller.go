package handlers

import (
	"github.com/aryadhira/reseller-management/internal/interfaces"
	"github.com/aryadhira/reseller-management/internal/models"
	"github.com/gofiber/fiber/v2"
)

// ResellerHandler handles reseller-related requests
type ResellerHandler struct {
	Service interfaces.ResellerService
}

func NewResellerHandler(service interfaces.ResellerService) *ResellerHandler {
	return &ResellerHandler{Service: service}
}

// CreateReseller creates a new reseller
// @Summary Create a new reseller
// @Description Create a new reseller with provided details
// @Tags Reseller Management
// @Accept json
// @Produce json
// @Param reseller body models.Reseller true "Reseller data"
// @Success 201 {object} models.Reseller
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /resellers [post]
func (h *ResellerHandler) CreateReseller(c *fiber.Ctx) error {
	reseller := new(models.Reseller)
	if err := c.BodyParser(reseller); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request body"})
	}

	createdReseller, err := h.Service.CreateReseller(reseller)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(201).JSON(createdReseller)
}

// GetAllResellers gets all resellers
// @Summary Get all resellers
// @Description Get a list of all resellers
// @Tags Reseller Management
// @Produce json
// @Success 200 {array} models.Reseller
// @Failure 500 {object} map[string]string
// @Router /resellers [get]
func (h *ResellerHandler) GetAllResellers(c *fiber.Ctx) error {
	resellers, err := h.Service.GetAllResellers()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(resellers)
}

// GetResellerByID gets a reseller by ID
// @Summary Get a reseller by ID
// @Description Get a reseller by its unique ID
// @Tags Reseller Management
// @Produce json
// @Param id path string true "Reseller ID"
// @Success 200 {object} models.Reseller
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /resellers/{id} [get]
func (h *ResellerHandler) GetResellerByID(c *fiber.Ctx) error {
	id := c.Params("id")
	
	reseller, err := h.Service.GetResellerByID(id)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Reseller not found"})
	}

	return c.JSON(reseller)
}

// GetResellerWithOrders gets a reseller with order history
// @Summary Get reseller profile with order history
// @Description Get a reseller with their order history and payment status
// @Tags Reseller Management
// @Produce json
// @Param id path string true "Reseller ID"
// @Success 200 {object} models.Reseller
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /resellers/{id}/profile [get]
func (h *ResellerHandler) GetResellerWithOrders(c *fiber.Ctx) error {
	id := c.Params("id")
	
	reseller, err := h.Service.GetResellerWithOrders(id)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Reseller not found"})
	}

	return c.JSON(reseller)
}

// UpdateReseller updates a reseller
// @Summary Update a reseller
// @Description Update an existing reseller's details
// @Tags Reseller Management
// @Accept json
// @Produce json
// @Param id path string true "Reseller ID"
// @Param reseller body models.Reseller true "Updated reseller data"
// @Success 200 {object} models.Reseller
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /resellers/{id} [put]
func (h *ResellerHandler) UpdateReseller(c *fiber.Ctx) error {
	id := c.Params("id")
	
	reseller := new(models.Reseller)
	if err := c.BodyParser(reseller); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request body"})
	}

	updatedReseller, err := h.Service.UpdateReseller(id, reseller)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(updatedReseller)
}

// DeleteReseller deletes a reseller
// @Summary Delete a reseller
// @Description Delete an existing reseller by ID
// @Tags Reseller Management
// @Produce json
// @Param id path string true "Reseller ID"
// @Success 204 {object} nil
// @Failure 500 {object} map[string]string
// @Router /resellers/{id} [delete]
func (h *ResellerHandler) DeleteReseller(c *fiber.Ctx) error {
	id := c.Params("id")
	
	err := h.Service.DeleteReseller(id)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.SendStatus(204)
}