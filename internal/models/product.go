package models

// Product represents a product in the system
// @Description Product information
type Product struct {
	BaseModel
	// Product name
	Name string `json:"name" gorm:"not null" example:"Laptop"`
	// Product description
	Description string `json:"description" example:"High performance laptop"`
	// Product SKU (Stock Keeping Unit)
	SKU string `json:"sku" gorm:"unique;not null" example:"LAP-001"`
	// Product price
	Price float64 `json:"price" gorm:"not null" example:"999.99"`
	// Current stock quantity
	CurrentStock int `json:"current_stock" gorm:"not null;default:0" example:"50"`
	// Minimum stock alert threshold
	MinStockAlert int `json:"min_stock_alert" gorm:"default:10" example:"10"` // Alert when stock falls below this
	// Status of the product
	Status string `json:"status" gorm:"default:'active'" example:"active"` // active, inactive
	// Order items associated with this product (simplified to avoid recursion)
	OrderItems []OrderItem `json:"order_items,omitempty" gorm:"foreignKey:ProductID"`
}