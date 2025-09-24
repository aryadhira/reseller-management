package models

import "time"

// Payment represents a payment in the system
// @Description Payment information
type Payment struct {
	BaseModel
	// ID of the order this payment is for
	OrderID string `json:"order_id" gorm:"not null;unique" example:"550e8400-e29b-41d4-a716-446655440001"`
	// Total amount of the order
	TotalAmount float64 `json:"total_amount" gorm:"not null" example:"1999.98"`
	// Amount that has been paid
	AmountPaid float64 `json:"amount_paid" gorm:"default:0" example:"500.00"`
	// Status of the payment
	Status string `json:"status" gorm:"default:'unpaid'" example:"partially_paid"` // unpaid, partially_paid, paid, overdue
	// Date when the payment was made
	PaymentDate time.Time `json:"payment_date" gorm:"default:CURRENT_TIMESTAMP"`
	// Additional notes about the payment
	Notes string `json:"notes" example:"Partial payment received"`
	// Cash in transactions associated with this payment (simplified to avoid recursion)
	CashInTransactions []Transaction `json:"cash_in_transactions,omitempty" gorm:"foreignKey:PaymentID"`
	// Order this payment is for (simplified to avoid recursion)
	Order Order `json:"order,omitempty" gorm:"foreignKey:OrderID"`
}