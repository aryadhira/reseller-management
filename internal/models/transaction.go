package models

import "time"

type TransactionType string

const (
	CashIn  TransactionType = "CASH_IN"
	CashOut TransactionType = "CASH_OUT"
)

type TransactionCategory string

const (
	Rent       TransactionCategory = "RENT"
	Salary     TransactionCategory = "SALARY"
	Equipment  TransactionCategory = "EQUIPMENT"
	PaymentTran TransactionCategory = "PAYMENT"  // Renamed from "Payment" to "PaymentTran" to avoid conflict with Payment model
	Other      TransactionCategory = "OTHER"
)

type Transaction struct {
	BaseModel
	Type        TransactionType     `json:"type" gorm:"not null"`
	Category    TransactionCategory `json:"category" gorm:"not null"`
	Amount      float64             `json:"amount" gorm:"not null"`
	Description string              `json:"description"`
	Date        time.Time           `json:"date" gorm:"default:CURRENT_TIMESTAMP"`
	ReferenceID *string             `json:"reference_id"` // Order ID, Payment ID, etc.
	PaymentID   *string             `json:"payment_id"`   // For CASH_IN related to a payment
}