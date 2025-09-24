package models

// Reseller represents a reseller in the system
// @Description Reseller information
type Reseller struct {
	BaseModel
	// Reseller name
	Name string `json:"name" gorm:"not null" example:"John Doe"`
	// Reseller email
	Email string `json:"email" gorm:"unique;not null" example:"john@example.com"`
	// Reseller phone number
	Phone string `json:"phone" example:"+1-234-567-8900"`
	// Reseller address
	Address string `json:"address" example:"123 Main St, City, Country"`
	// Status of the reseller
	Status string `json:"status" gorm:"default:'active'" example:"active"` // active, inactive
	// Orders associated with this reseller (simplified to avoid recursion)
	Orders []Order `json:"orders,omitempty" gorm:"foreignKey:ResellerID"`
}