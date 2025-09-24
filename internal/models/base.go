package models

import (
	"time"

	"gorm.io/gorm"
)

// BaseModel represents the base model with common fields
// @Description Base model with common fields
type BaseModel struct {
	// Unique identifier for the entity
	ID string `gorm:"primaryKey;type:uuid;default:gen_random_uuid()" json:"id" example:"550e8400-e29b-41d4-a716-446655440000"`
	// Creation timestamp
	CreatedAt time.Time `json:"created_at"`
	// Last update timestamp
	UpdatedAt time.Time `json:"updated_at"`
	// Deletion timestamp (for soft deletes)
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"` // For soft deletes, excluded from JSON to avoid swagger issues
}