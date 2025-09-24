package models

type OrderItem struct {
	BaseModel
	OrderID   string `json:"order_id" gorm:"not null"`
	ProductID string `json:"product_id" gorm:"not null"`
	Quantity  int    `json:"quantity" gorm:"not null"`
	Price     float64 `json:"price" gorm:"not null"` // Price at the time of order
	Subtotal  float64 `json:"subtotal" gorm:"not null"`
	Order     Order   `json:"order" gorm:"foreignKey:OrderID"`
	Product   Product `json:"product" gorm:"foreignKey:ProductID"`
}