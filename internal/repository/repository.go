package repository

import (
	"time"

	"github.com/aryadhira/reseller-management/internal/models"
	"gorm.io/gorm"
)

type Repository struct {
	Reseller ResellerRepository
	Product  ProductRepository
	Order    OrderRepository
	Payment  PaymentRepository
}

type ResellerRepository interface {
	Create(reseller *models.Reseller) error
	GetAll() ([]models.Reseller, error)
	GetByID(id string) (*models.Reseller, error)
	Update(id string, reseller *models.Reseller) error
	Delete(id string) error
	GetWithOrders(id string) (*models.Reseller, error)
}

type ProductRepository interface {
	Create(product *models.Product) error
	GetAll() ([]models.Product, error)
	GetByID(id string) (*models.Product, error)
	Update(id string, product *models.Product) error
	Delete(id string) error
	Restock(id string, quantity int) error
	GetLowStock() ([]models.Product, error)
}

type OrderRepository interface {
	Create(order *models.Order) error
	GetAll() ([]models.Order, error)
	GetByID(id string) (*models.Order, error)
	Update(id string, order *models.Order) error
	Delete(id string) error
	Cancel(id string) error
}

type PaymentRepository interface {
	GetAll() ([]models.Payment, error)
	GetByOrderID(orderID string) (*models.Payment, error)
	Create(payment *models.Payment) error
	Update(payment *models.Payment) error
	GetAllTransactions() ([]models.Transaction, error)
	CreateTransaction(transaction *models.Transaction) error
	UpdateBalance(initialBalance float64) error
	GetBalance() (*models.Balance, error)
	GetDashboardData() (*DashboardData, error)
	GetRecentTransactions(limit int) ([]models.Transaction, error)
	GetCashInByDateRange(start, end time.Time) (float64, error)
	GetCashOutByDateRange(start, end time.Time) (float64, error)
	GetUnpaidOrders() ([]models.Order, error)
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		Reseller: NewResellerRepository(db),
		Product:  NewProductRepository(db),
		Order:    NewOrderRepository(db),
		Payment:  NewPaymentRepository(db),
	}
}