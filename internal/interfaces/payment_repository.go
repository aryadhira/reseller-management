package interfaces

import (
	"time"

	"github.com/aryadhira/reseller-management/internal/models"
)

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