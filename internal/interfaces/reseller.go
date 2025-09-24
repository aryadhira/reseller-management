package interfaces

import (
	"github.com/aryadhira/reseller-management/internal/models"
)

type ResellerService interface {
	CreateReseller(reseller *models.Reseller) (*models.Reseller, error)
	GetAllResellers() ([]models.Reseller, error)
	GetResellerByID(id string) (*models.Reseller, error)
	UpdateReseller(id string, reseller *models.Reseller) (*models.Reseller, error)
	DeleteReseller(id string) error
	GetResellerWithOrders(id string) (*models.Reseller, error)
}