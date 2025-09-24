package models

type Balance struct {
	BaseModel
	InitialBalance float64 `json:"initial_balance" gorm:"default:0"`
	CurrentBalance float64 `json:"current_balance" gorm:"default:0"`
	Notes          string  `json:"notes"`
}