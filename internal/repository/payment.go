package repository

import (
	"time"

	"github.com/aryadhira/reseller-management/internal/models"
	"gorm.io/gorm"
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

type paymentRepository struct {
	db *gorm.DB
}

func NewPaymentRepository(db *gorm.DB) *paymentRepository {
	return &paymentRepository{db: db}
}

func (r *paymentRepository) GetAll() ([]models.Payment, error) {
	var payments []models.Payment
	err := r.db.Preload("Order").Preload("Order.Reseller").Find(&payments).Error
	return payments, err
}

func (r *paymentRepository) GetByOrderID(orderID string) (*models.Payment, error) {
	var payment models.Payment
	err := r.db.Preload("Order").Preload("Order.Reseller").Where("order_id = ?", orderID).First(&payment).Error
	return &payment, err
}

func (r *paymentRepository) Create(payment *models.Payment) error {
	return r.db.Create(payment).Error
}

func (r *paymentRepository) Update(payment *models.Payment) error {
	return r.db.Save(payment).Error
}

func (r *paymentRepository) GetAllTransactions() ([]models.Transaction, error) {
	var transactions []models.Transaction
	err := r.db.Order("created_at DESC").Find(&transactions).Error
	return transactions, err
}

func (r *paymentRepository) CreateTransaction(transaction *models.Transaction) error {
	return r.db.Create(transaction).Error
}

func (r *paymentRepository) UpdateBalance(initialBalance float64) error {
	var balance models.Balance
	err := r.db.First(&balance).Error
	if err != nil {
		// If no balance record exists, create one
		if err == gorm.ErrRecordNotFound {
			balance = models.Balance{
				InitialBalance: initialBalance,
				CurrentBalance: initialBalance,
			}
			return r.db.Create(&balance).Error
		}
		return err
	}

	// Update the initial balance
	balance.InitialBalance = initialBalance
	balance.CurrentBalance = initialBalance // Reset current balance to initial

	// Recalculate current balance from all transactions
	cashInTotal := r.calculateCashInTotal()
	cashOutTotal := r.calculateCashOutTotal()
	balance.CurrentBalance = balance.InitialBalance + cashInTotal - cashOutTotal

	return r.db.Save(&balance).Error
}

func (r *paymentRepository) GetBalance() (*models.Balance, error) {
	var balance models.Balance
	err := r.db.First(&balance).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			// If no balance record exists, create one with default values
			balance = models.Balance{
				InitialBalance: 0,
				CurrentBalance: 0,
			}
			err = r.db.Create(&balance).Error
			if err != nil {
				return nil, err
			}
		} else {
			return nil, err
		}
	} else {
		// Recalculate current balance from all transactions
		cashInTotal := r.calculateCashInTotal()
		cashOutTotal := r.calculateCashOutTotal()
		balance.CurrentBalance = balance.InitialBalance + cashInTotal - cashOutTotal
		
		// Save the updated balance
		r.db.Save(&balance)
	}

	return &balance, nil
}

func (r *paymentRepository) GetDashboardData() (*DashboardData, error) {
	today := time.Now().Truncate(24 * time.Hour)
	thisMonth := time.Date(time.Now().Year(), time.Now().Month(), 1, 0, 0, 0, 0, time.Now().Location())

	balance, err := r.GetBalance()
	if err != nil {
		return nil, err
	}

	todayCashIn, err := r.GetCashInByDateRange(today, today.Add(24*time.Hour))
	if err != nil {
		return nil, err
	}

	thisMonthCashIn, err := r.GetCashInByDateRange(thisMonth, thisMonth.AddDate(0, 1, 0))
	if err != nil {
		return nil, err
	}

	allTimeCashIn, err := r.GetCashInByDateRange(time.Time{}, time.Now().Add(24*time.Hour))
	if err != nil {
		return nil, err
	}

	todayCashOut, err := r.GetCashOutByDateRange(today, today.Add(24*time.Hour))
	if err != nil {
		return nil, err
	}

	thisMonthCashOut, err := r.GetCashOutByDateRange(thisMonth, thisMonth.AddDate(0, 1, 0))
	if err != nil {
		return nil, err
	}

	allTimeCashOut, err := r.GetCashOutByDateRange(time.Time{}, time.Now().Add(24*time.Hour))
	if err != nil {
		return nil, err
	}

	recentTransactions, err := r.GetRecentTransactions(10)
	if err != nil {
		return nil, err
	}

	unpaidOrders, err := r.GetUnpaidOrders()
	if err != nil {
		return nil, err
	}

	// Get low stock products from product repository
	// We'll need to access the product repository for this
	// For now, let's return the dashboard data without low stock alerts
	// Later we'll handle the integration properly

	data := &DashboardData{
		CurrentBalance:     balance.CurrentBalance,
		TodayCashIn:        todayCashIn,
		ThisMonthCashIn:    thisMonthCashIn,
		AllTimeCashIn:      allTimeCashIn,
		TodayCashOut:       todayCashOut,
		ThisMonthCashOut:   thisMonthCashOut,
		AllTimeCashOut:     allTimeCashOut,
		RecentTransactions: recentTransactions,
		LowStockAlerts:     []models.Product{}, // Will be populated from product repo later
		UnpaidOrders:       unpaidOrders,
	}

	return data, nil
}

func (r *paymentRepository) GetRecentTransactions(limit int) ([]models.Transaction, error) {
	var transactions []models.Transaction
	err := r.db.Order("created_at DESC").Limit(limit).Find(&transactions).Error
	return transactions, err
}

func (r *paymentRepository) GetCashInByDateRange(start, end time.Time) (float64, error) {
	var total float64
	err := r.db.Model(&models.Transaction{}).
		Where("type = ? AND date >= ? AND date < ?", "CASH_IN", start, end).
		Select("COALESCE(SUM(amount), 0)").
		Scan(&total).Error
	return total, err
}

func (r *paymentRepository) GetCashOutByDateRange(start, end time.Time) (float64, error) {
	var total float64
	err := r.db.Model(&models.Transaction{}).
		Where("type = ? AND date >= ? AND date < ?", "CASH_OUT", start, end).
		Select("COALESCE(SUM(amount), 0)").
		Scan(&total).Error
	return total, err
}

func (r *paymentRepository) GetUnpaidOrders() ([]models.Order, error) {
	var orders []models.Order
	err := r.db.Preload("Reseller").Preload("OrderItems").Where("payment_status IN ?", []string{"unpaid", "partially_paid"}).Find(&orders).Error
	return orders, err
}

func (r *paymentRepository) calculateCashInTotal() float64 {
	var total float64
	r.db.Model(&models.Transaction{}).
		Where("type = ?", "CASH_IN").
		Select("COALESCE(SUM(amount), 0)").
		Scan(&total)
	return total
}

func (r *paymentRepository) calculateCashOutTotal() float64 {
	var total float64
	r.db.Model(&models.Transaction{}).
		Where("type = ?", "CASH_OUT").
		Select("COALESCE(SUM(amount), 0)").
		Scan(&total)
	return total
}