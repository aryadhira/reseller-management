package models

// OrderItem represents an item in an order
// @Description Order item information
type OrderItem struct {
	BaseModel
	// ID of the order this item belongs to
	OrderID string `json:"order_id" gorm:"not null" example:"550e8400-e29b-41d4-a716-446655440001"`
	// ID of the product in this order item
	ProductID string `json:"product_id" gorm:"not null" example:"550e8400-e29b-41d4-a716-446655440002"`
	// Quantity of the product ordered
	Quantity int `json:"quantity" gorm:"not null" example:"2"`
	// Price of the product at the time of order
	Price float64 `json:"price" gorm:"not null" example:"999.99"` // Price at the time of order
	// Subtotal for this item (quantity * price)
	Subtotal float64 `json:"subtotal" gorm:"not null" example:"1999.98"`
	// Order this item belongs to (simplified to avoid recursion)
	Order Order `json:"order,omitempty" gorm:"foreignKey:OrderID"`
	// Product being ordered (simplified to avoid recursion)
	Product Product `json:"product,omitempty" gorm:"foreignKey:ProductID"`
}