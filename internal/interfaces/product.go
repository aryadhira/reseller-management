package interfaces

import (
	"github.com/aryadhira/reseller-management/internal/models"
)

type ProductService interface {
	CreateProduct(product *models.Product) (*models.Product, error)
	GetAllProducts() ([]models.Product, error)
	GetProductByID(id string) (*models.Product, error)
	UpdateProduct(id string, product *models.Product) (*models.Product, error)
	DeleteProduct(id string) error
	RestockProduct(id string, quantity int) (*models.Product, error)
	GetLowStockProducts() ([]models.Product, error)
}