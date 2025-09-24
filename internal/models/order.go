package models

import "time"

// Order represents an order in the system
// @Description Order information
type Order struct {
	BaseModel
	// ID of the reseller who placed the order
	ResellerID string `json:"reseller_id" gorm:"not null" example:"550e8400-e29b-41d4-a716-446655440000"`
	// Items in the order (simplified to avoid recursion)
	OrderItems []OrderItem `json:"order_items,omitempty" gorm:"foreignKey:OrderID"`
	// Total amount of the order
	TotalAmount float64 `json:"total_amount" gorm:"not null" example:"1999.98"`
	// Status of the order
	Status string `json:"status" gorm:"default:'pending'" example:"pending"` // pending, confirmed, cancelled, completed
	// Payment status of the order
	PaymentStatus string `json:"payment_status" gorm:"default:'unpaid'" example:"unpaid"` // unpaid, partially_paid, paid, overdue
	// Date when the order was placed
	OrderDate time.Time `json:"order_date" gorm:"default:CURRENT_TIMESTAMP"`
	// Additional notes about the order
	Notes string `json:"notes" example:"Special delivery instructions"`
	// Reseller who placed the order (simplified to avoid recursion)
	Reseller Reseller `json:"reseller,omitempty" gorm:"foreignKey:ResellerID"`
	// Payment information for the order (simplified to avoid recursion)
	Payment *Payment `json:"payment,omitempty" gorm:"foreignKey:OrderID"`
}