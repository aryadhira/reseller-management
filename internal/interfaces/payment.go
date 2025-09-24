package interfaces

import (
	"github.com/aryadhira/reseller-management/internal/models"
)

type DashboardData struct {
	CurrentBalance       float64                    `json:"current_balance"`
	TodayCashIn          float64                    `json:"today_cash_in"`
	ThisMonthCashIn      float64                    `json:"this_month_cash_in"`
	AllTimeCashIn        float64                    `json:"all_time_cash_in"`
	TodayCashOut         float64                    `json:"today_cash_out"`
	ThisMonthCashOut     float64                    `json:"this_month_cash_out"`
	AllTimeCashOut       float64                    `json:"all_time_cash_out"`
	RecentTransactions   []models.Transaction       `json:"recent_transactions"`
	LowStockAlerts       []models.Product           `json:"low_stock_alerts"`
	UnpaidOrders         []models.Order             `json:"unpaid_orders"`
}

type PaymentService interface {
	GetAllPayments() ([]models.Payment, error)
	GetPaymentByOrderID(orderID string) (*models.Payment, error)
	RecordPayment(orderID string, amount float64, notes string) (*models.Payment, error)
	GetAllTransactions() ([]models.Transaction, error)
	RecordCashIn(amount float64, description string, referenceID *string) (*models.Transaction, error)
	RecordCashOut(category models.TransactionCategory, amount float64, description string) (*models.Transaction, error)
	UpdateBalance(initialBalance float64, notes string) error
	GetBalance() (*models.Balance, error)
	GetDashboardData() (*DashboardData, error)
}