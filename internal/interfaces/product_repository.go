package interfaces

import (
	"github.com/aryadhira/reseller-management/internal/models"
)

type ProductRepository interface {
	Create(product *models.Product) error
	GetAll() ([]models.Product, error)
	GetByID(id string) (*models.Product, error)
	Update(id string, product *models.Product) error
	Delete(id string) error
	Restock(id string, quantity int) error
	GetLowStock() ([]models.Product, error)
}