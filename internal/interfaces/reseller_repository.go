package interfaces

import (
	"github.com/aryadhira/reseller-management/internal/models"
)

type ResellerRepository interface {
	Create(reseller *models.Reseller) error
	GetAll() ([]models.Reseller, error)
	GetByID(id string) (*models.Reseller, error)
	Update(id string, reseller *models.Reseller) error
	Delete(id string) error
	GetWithOrders(id string) (*models.Reseller, error)
}