package handlers

import (
	"github.com/aryadhira/reseller-management/internal/interfaces"
	"github.com/aryadhira/reseller-management/internal/models"
	"github.com/gofiber/fiber/v2"
)

type ResellerHandler struct {
	Service interfaces.ResellerService
}

func NewResellerHandler(service interfaces.ResellerService) *ResellerHandler {
	return &ResellerHandler{Service: service}
}

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

func (h *ResellerHandler) GetAllResellers(c *fiber.Ctx) error {
	resellers, err := h.Service.GetAllResellers()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(resellers)
}

func (h *ResellerHandler) GetResellerByID(c *fiber.Ctx) error {
	id := c.Params("id")
	
	reseller, err := h.Service.GetResellerByID(id)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Reseller not found"})
	}

	return c.JSON(reseller)
}

func (h *ResellerHandler) GetResellerWithOrders(c *fiber.Ctx) error {
	id := c.Params("id")
	
	reseller, err := h.Service.GetResellerWithOrders(id)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Reseller not found"})
	}

	return c.JSON(reseller)
}

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

func (h *ResellerHandler) DeleteReseller(c *fiber.Ctx) error {
	id := c.Params("id")
	
	err := h.Service.DeleteReseller(id)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.SendStatus(204)
}