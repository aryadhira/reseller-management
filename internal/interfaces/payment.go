package interfaces

import (
	"github.com/aryadhira/reseller-management/internal/models"
)

// DashboardData represents the dashboard overview data
// @Description Dashboard information
type DashboardData struct {
	// Current balance in the system
	CurrentBalance float64 `json:"current_balance" example:"12500.75"`
	// Total cash in for today
	TodayCashIn float64 `json:"today_cash_in" example:"2500.00"`
	// Total cash in for this month
	ThisMonthCashIn float64 `json:"this_month_cash_in" example:"15000.00"`
	// Total cash in for all time
	AllTimeCashIn float64 `json:"all_time_cash_in" example:"50000.00"`
	// Total cash out for today
	TodayCashOut float64 `json:"today_cash_out" example:"1200.00"`
	// Total cash out for this month
	ThisMonthCashOut float64 `json:"this_month_cash_out" example:"8000.00"`
	// Total cash out for all time
	AllTimeCashOut float64 `json:"all_time_cash_out" example:"25000.00"`
	// Recent transactions
	RecentTransactions []models.Transaction `json:"recent_transactions"`
	// Products with low stock alerts
	LowStockAlerts []models.Product `json:"low_stock_alerts"`
	// Unpaid orders
	UnpaidOrders []models.Order `json:"unpaid_orders"`
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