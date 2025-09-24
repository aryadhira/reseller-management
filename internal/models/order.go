package models

import "time"

type Order struct {
	BaseModel
	ResellerID    string      `json:"reseller_id" gorm:"not null"`
	OrderItems    []OrderItem `json:"order_items" gorm:"foreignKey:OrderID"`
	TotalAmount   float64     `json:"total_amount" gorm:"not null"`
	Status        string      `json:"status" gorm:"default:'pending"` // pending, confirmed, cancelled, completed
	PaymentStatus string      `json:"payment_status" gorm:"default:'unpaid"` // unpaid, partially_paid, paid, overdue
	OrderDate     time.Time   `json:"order_date" gorm:"default:CURRENT_TIMESTAMP"`
	Notes         string      `json:"notes"`
	Reseller      Reseller    `json:"reseller" gorm:"foreignKey:ResellerID"`
	Payment       *Payment    `json:"payment" gorm:"foreignKey:OrderID"`
}