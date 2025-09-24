package models

import "time"

// TransactionType defines the type of transaction
type TransactionType string

const (
	CashIn  TransactionType = "CASH_IN"
	CashOut TransactionType = "CASH_OUT"
)

// TransactionCategory defines the category of transaction
type TransactionCategory string

const (
	Rent       TransactionCategory = "RENT"
	Salary     TransactionCategory = "SALARY"
	Equipment  TransactionCategory = "EQUIPMENT"
	PaymentTran TransactionCategory = "PAYMENT"  // Renamed from "Payment" to "PaymentTran" to avoid conflict with Payment model
	Other      TransactionCategory = "OTHER"
)

// Transaction represents a financial transaction
// @Description Transaction information
type Transaction struct {
	BaseModel
	// Type of the transaction
	Type TransactionType `json:"type" gorm:"not null" example:"CASH_IN"` // CASH_IN, CASH_OUT
	// Category of the transaction
	Category TransactionCategory `json:"category" gorm:"not null" example:"SALARY"`
	// Amount of the transaction
	Amount float64 `json:"amount" gorm:"not null" example:"3000.00"`
	// Description of the transaction
	Description string `json:"description" example:"Monthly salary payment"`
	// Date of the transaction
	Date time.Time `json:"date" gorm:"default:CURRENT_TIMESTAMP"`
	// Reference ID for the transaction (e.g. Order ID, Payment ID)
	ReferenceID *string `json:"reference_id"` // Order ID, Payment ID, etc. // example:"550e8400-e29b-41d4-a716-446655440001"
	// Payment ID if this transaction is related to a payment
	PaymentID *string `json:"payment_id"` // For CASH_IN related to a payment // example:"550e8400-e29b-41d4-a716-446655440002"
}