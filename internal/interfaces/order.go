package interfaces

import (
	"github.com/aryadhira/reseller-management/internal/models"
)

type OrderService interface {
	CreateOrder(order *models.Order) (*models.Order, error)
	GetAllOrders() ([]models.Order, error)
	GetOrderByID(id string) (*models.Order, error)
	UpdateOrder(id string, order *models.Order) (*models.Order, error)
	DeleteOrder(id string) error
	CancelOrder(id string) error
}