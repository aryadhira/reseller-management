package interfaces

import (
	"github.com/aryadhira/reseller-management/internal/models"
)

type OrderRepository interface {
	Create(order *models.Order) error
	GetAll() ([]models.Order, error)
	GetByID(id string) (*models.Order, error)
	Update(id string, order *models.Order) error
	Delete(id string) error
	Cancel(id string) error
}