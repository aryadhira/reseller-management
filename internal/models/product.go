package models

type Product struct {
	BaseModel
	Name          string  `json:"name" gorm:"not null"`
	Description   string  `json:"description"`
	SKU           string  `json:"sku" gorm:"unique;not null"`
	Price         float64 `json:"price" gorm:"not null"`
	CurrentStock  int     `json:"current_stock" gorm:"not null;default:0"`
	MinStockAlert int     `json:"min_stock_alert" gorm:"default:10"` // Alert when stock falls below this
	Status        string  `json:"status" gorm:"default:'active"`    // active, inactive
	OrderItems    []OrderItem `json:"order_items" gorm:"foreignKey:ProductID"`
}