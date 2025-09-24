package models

type Reseller struct {
	BaseModel
	Name        string `json:"name" gorm:"not null"`
	Email       string `json:"email" gorm:"unique;not null"`
	Phone       string `json:"phone"`
	Address     string `json:"address"`
	Status      string `json:"status" gorm:"default:'active"` // active, inactive
	Orders      []Order `json:"orders" gorm:"foreignKey:ResellerID"`
}