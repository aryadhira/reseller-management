package handlers

import (
	"github.com/aryadhira/reseller-management/internal/interfaces"
	"github.com/aryadhira/reseller-management/internal/models"
	"github.com/gofiber/fiber/v2"
)

// PaymentRequest represents the request to record a payment
// @Description Payment request information
type PaymentRequest struct {
	// Amount to be paid
	Amount float64 `json:"amount" example:"500.00"`
	// Notes about the payment
	Notes string `json:"notes" example:"Partial payment received"`
}

// CashInRequest represents the request to record a cash in transaction
// @Description Cash in request information
type CashInRequest struct {
	// Amount of the cash in
	Amount float64 `json:"amount" example:"2500.00"`
	// Description of the transaction
	Description string `json:"description" example:"Payment received from client"`
	// Reference ID for the transaction
	ReferenceID *string `json:"reference_id,omitempty" example:"550e8400-e29b-41d4-a716-446655440001"`
}

// CashOutRequest represents the request to record a cash out transaction
// @Description Cash out request information
type CashOutRequest struct {
	// Category of the transaction
	Category models.TransactionCategory `json:"category" example:"SALARY"`
	// Amount of the cash out
	Amount float64 `json:"amount" example:"3000.00"`
	// Description of the transaction
	Description string `json:"description" example:"Monthly salary payment"`
}

// BalanceUpdateRequest represents the request to update the balance
// @Description Balance update request information
type BalanceUpdateRequest struct {
	// New initial balance
	InitialBalance float64 `json:"initial_balance" example:"10000.00"`
	// Notes about the balance update
	Notes string `json:"notes" example:"Adjusted initial balance"`
}

// PaymentHandler handles payment-related requests
type PaymentHandler struct {
	Service interfaces.PaymentService
}

func NewPaymentHandler(service interfaces.PaymentService) *PaymentHandler {
	return &PaymentHandler{Service: service}
}

// GetAllPayments gets all payments
// @Summary Get all payments
// @Description Get a list of all payments linked to their respective orders and resellers
// @Tags Payment Management
// @Produce json
// @Success 200 {array} models.Payment
// @Failure 500 {object} map[string]string
// @Router /payments [get]
func (h *PaymentHandler) GetAllPayments(c *fiber.Ctx) error {
	payments, err := h.Service.GetAllPayments()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(payments)
}

// GetPaymentByOrderID gets payment by order ID
// @Summary Get payment by order ID
// @Description Get payment information for a specific order
// @Tags Payment Management
// @Produce json
// @Param orderID path string true "Order ID"
// @Success 200 {object} models.Payment
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /payments/order/{orderID} [get]
func (h *PaymentHandler) GetPaymentByOrderID(c *fiber.Ctx) error {
	orderID := c.Params("orderID")
	
	payment, err := h.Service.GetPaymentByOrderID(orderID)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Payment not found"})
	}

	return c.JSON(payment)
}

// RecordPayment records a payment
// @Summary Record a payment
// @Description Record a payment (partial or full) against an order
// @Tags Payment Management
// @Accept json
// @Produce json
// @Param orderID path string true "Order ID"
// @Param payment body PaymentRequest true "Payment information"
// @Success 200 {object} models.Payment
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /payments/order/{orderID}/pay [post]
func (h *PaymentHandler) RecordPayment(c *fiber.Ctx) error {
	orderID := c.Params("orderID")
	
	req := new(PaymentRequest)
	if err := c.BodyParser(req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request body"})
	}

	payment, err := h.Service.RecordPayment(orderID, req.Amount, req.Notes)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(payment)
}

// GetAllTransactions gets all transactions
// @Summary Get all transactions
// @Description Get a list of all financial transactions (CASH_IN and CASH_OUT)
// @Tags Financial Management
// @Produce json
// @Success 200 {array} models.Transaction
// @Failure 500 {object} map[string]string
// @Router /transactions [get]
func (h *PaymentHandler) GetAllTransactions(c *fiber.Ctx) error {
	transactions, err := h.Service.GetAllTransactions()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(transactions)
}

// RecordCashIn records a cash in transaction
// @Summary Record a cash in transaction
// @Description Record a CASH_IN transaction for payments received
// @Tags Financial Management
// @Accept json
// @Produce json
// @Param transaction body CashInRequest true "Cash in transaction data"
// @Success 201 {object} models.Transaction
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /transactions/cash-in [post]
func (h *PaymentHandler) RecordCashIn(c *fiber.Ctx) error {
	req := new(CashInRequest)
	if err := c.BodyParser(req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request body"})
	}

	transaction, err := h.Service.RecordCashIn(req.Amount, req.Description, req.ReferenceID)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(201).JSON(transaction)
}

// RecordCashOut records a cash out transaction
// @Summary Record a cash out transaction
// @Description Record a CASH_OUT transaction with category
// @Tags Financial Management
// @Accept json
// @Produce json
// @Param transaction body CashOutRequest true "Cash out transaction data"
// @Success 201 {object} models.Transaction
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /transactions/cash-out [post]
func (h *PaymentHandler) RecordCashOut(c *fiber.Ctx) error {
	req := new(CashOutRequest)
	if err := c.BodyParser(req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request body"})
	}

	transaction, err := h.Service.RecordCashOut(req.Category, req.Amount, req.Description)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(201).JSON(transaction)
}

// UpdateBalance updates the initial balance
// @Summary Update the initial balance
// @Description Update the initial balance and adjust accordingly
// @Tags Financial Management
// @Accept json
// @Produce json
// @Param balance body BalanceUpdateRequest true "Balance information"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /balance [put]
func (h *PaymentHandler) UpdateBalance(c *fiber.Ctx) error {
	req := new(BalanceUpdateRequest)
	if err := c.BodyParser(req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	err := h.Service.UpdateBalance(req.InitialBalance, req.Notes)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "Balance updated successfully"})
}

// GetBalance gets the current balance
// @Summary Get the current balance
// @Description Get the current system balance
// @Tags Financial Management
// @Produce json
// @Success 200 {object} models.Balance
// @Failure 500 {object} map[string]string
// @Router /balance [get]
func (h *PaymentHandler) GetBalance(c *fiber.Ctx) error {
	balance, err := h.Service.GetBalance()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(balance)
}

// GetDashboardData gets dashboard data
// @Summary Get dashboard data
// @Description Get an overview dashboard with current balance, cash flows, and alerts
// @Tags Dashboard
// @Produce json
// @Success 200 {object} interfaces.DashboardData
// @Failure 500 {object} map[string]string
// @Router /dashboard [get]
func (h *PaymentHandler) GetDashboardData(c *fiber.Ctx) error {
	data, err := h.Service.GetDashboardData()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(data)
}