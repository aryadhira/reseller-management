package models

// Balance represents the financial balance
// @Description Balance information
type Balance struct {
	BaseModel
	// Initial balance amount
	InitialBalance float64 `json:"initial_balance" gorm:"default:0" example:"10000.00"`
	// Current balance amount
	CurrentBalance float64 `json:"current_balance" gorm:"default:0" example:"12500.75"`
	// Notes about the balance
	Notes string `json:"notes" example:"Adjusted initial balance"`
}