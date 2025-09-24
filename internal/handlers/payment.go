package handlers

import (
	"github.com/aryadhira/reseller-management/internal/interfaces"
	"github.com/aryadhira/reseller-management/internal/models"
	"github.com/gofiber/fiber/v2"
)

type PaymentHandler struct {
	Service interfaces.PaymentService
}

func NewPaymentHandler(service interfaces.PaymentService) *PaymentHandler {
	return &PaymentHandler{Service: service}
}

func (h *PaymentHandler) GetAllPayments(c *fiber.Ctx) error {
	payments, err := h.Service.GetAllPayments()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(payments)
}

func (h *PaymentHandler) GetPaymentByOrderID(c *fiber.Ctx) error {
	orderID := c.Params("orderID")
	
	payment, err := h.Service.GetPaymentByOrderID(orderID)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Payment not found"})
	}

	return c.JSON(payment)
}

func (h *PaymentHandler) RecordPayment(c *fiber.Ctx) error {
	orderID := c.Params("orderID")
	
	var req struct {
		Amount  float64 `json:"amount" validate:"required,gt=0"`
		Notes   string  `json:"notes"`
	}
	
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request body"})
	}

	payment, err := h.Service.RecordPayment(orderID, req.Amount, req.Notes)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(payment)
}

func (h *PaymentHandler) GetAllTransactions(c *fiber.Ctx) error {
	transactions, err := h.Service.GetAllTransactions()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(transactions)
}

func (h *PaymentHandler) RecordCashIn(c *fiber.Ctx) error {
	var req struct {
		Amount      float64  `json:"amount" validate:"required,gt=0"`
		Description string   `json:"description"`
		ReferenceID *string  `json:"reference_id"`
	}
	
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request body"})
	}

	transaction, err := h.Service.RecordCashIn(req.Amount, req.Description, req.ReferenceID)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(201).JSON(transaction)
}

func (h *PaymentHandler) RecordCashOut(c *fiber.Ctx) error {
	var req struct {
		Category    models.TransactionCategory `json:"category" validate:"required,oneof=RENT SALARY EQUIPMENT PAYMENT OTHER"`
		Amount      float64                    `json:"amount" validate:"required,gt=0"`
		Description string                     `json:"description"`
	}
	
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request body"})
	}

	transaction, err := h.Service.RecordCashOut(req.Category, req.Amount, req.Description)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(201).JSON(transaction)
}

func (h *PaymentHandler) UpdateBalance(c *fiber.Ctx) error {
	var req struct {
		InitialBalance float64 `json:"initial_balance"`
		Notes          string  `json:"notes"`
	}
	
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request body"})
	}

	err := h.Service.UpdateBalance(req.InitialBalance, req.Notes)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "Balance updated successfully"})
}

func (h *PaymentHandler) GetBalance(c *fiber.Ctx) error {
	balance, err := h.Service.GetBalance()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(balance)
}

func (h *PaymentHandler) GetDashboardData(c *fiber.Ctx) error {
	data, err := h.Service.GetDashboardData()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(data)
}