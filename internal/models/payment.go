package models

import "time"

type Payment struct {
	BaseModel
	OrderID       string    `json:"order_id" gorm:"not null;unique"`
	TotalAmount   float64   `json:"total_amount" gorm:"not null"`
	AmountPaid    float64   `json:"amount_paid" gorm:"default:0"`
	Status        string    `json:"status" gorm:"default:'unpaid"` // unpaid, partially_paid, paid, overdue
	PaymentDate   time.Time `json:"payment_date" gorm:"default:CURRENT_TIMESTAMP"`
	Notes         string    `json:"notes"`
	CashInTransactions []Transaction `json:"cash_in_transactions" gorm:"foreignKey:PaymentID"`
	Order         Order     `json:"order" gorm:"foreignKey:OrderID"`
}