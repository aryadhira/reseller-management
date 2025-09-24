package services

import (
	"errors"
	"time"

	"github.com/aryadhira/reseller-management/internal/interfaces"
	"github.com/aryadhira/reseller-management/internal/models"
	"github.com/aryadhira/reseller-management/internal/repository"
	"github.com/google/uuid"
)

type paymentService struct {
	repo *repository.Repository
}

func NewPaymentService(repo *repository.Repository) *paymentService {
	return &paymentService{repo: repo}
}

func (s *paymentService) GetAllPayments() ([]models.Payment, error) {
	return s.repo.Payment.GetAll()
}

func (s *paymentService) GetPaymentByOrderID(orderID string) (*models.Payment, error) {
	return s.repo.Payment.GetByOrderID(orderID)
}

func (s *paymentService) RecordPayment(orderID string, amount float64, notes string) (*models.Payment, error) {
	// Get the order and its payment
	order, err := s.repo.Order.GetByID(orderID)
	if err != nil {
		return nil, errors.New("order not found")
	}
	
	payment, err := s.repo.Payment.GetByOrderID(orderID)
	if err != nil {
		return nil, errors.New("payment record not found for order")
	}
	
	// Check if the payment amount is valid
	if amount <= 0 {
		return nil, errors.New("payment amount must be greater than zero")
	}
	
	if amount > (payment.TotalAmount - payment.AmountPaid) {
		return nil, errors.New("payment amount exceeds remaining balance")
	}
	
	// Update the payment record
	payment.AmountPaid += amount
	
	// Update the payment status
	if payment.AmountPaid >= payment.TotalAmount {
		payment.Status = "paid"
		order.PaymentStatus = "paid"
	} else if payment.AmountPaid > 0 {
		payment.Status = "partially_paid"
		order.PaymentStatus = "partially_paid"
	}
	
	// Update the order payment status
	err = s.repo.Order.Update(order.ID, order)
	if err != nil {
		return nil, err
	}
	
	// Update the payment record
	err = s.repo.Payment.Update(payment)
	if err != nil {
		return nil, err
	}
	
	// Create a CASH_IN transaction record
	transaction := &models.Transaction{
		BaseModel:   models.BaseModel{ID: uuid.NewString()},
		Type:        models.CashIn,
		Amount:      amount,
		Description: notes,
		Date:        time.Now(),
		ReferenceID: &orderID,
		PaymentID:   &payment.ID,
	}
	
	err = s.repo.Payment.CreateTransaction(transaction)
	if err != nil {
		return nil, err
	}
	
	// Update the balance
	balance, err := s.repo.Payment.GetBalance()
	if err != nil {
		return nil, err
	}
	
	balance.CurrentBalance += amount
	err = s.repo.Payment.UpdateBalance(balance.InitialBalance)
	if err != nil {
		return nil, err
	}
	
	return payment, nil
}

func (s *paymentService) GetAllTransactions() ([]models.Transaction, error) {
	return s.repo.Payment.GetAllTransactions()
}

func (s *paymentService) RecordCashIn(amount float64, description string, referenceID *string) (*models.Transaction, error) {
	if amount <= 0 {
		return nil, errors.New("amount must be greater than zero")
	}
	
	transaction := &models.Transaction{
		BaseModel:   models.BaseModel{ID: uuid.NewString()},
		Type:        models.CashIn,
		Category:    models.Other, // Default category, could be enhanced to accept category as parameter
		Amount:      amount,
		Description: description,
		Date:        time.Now(),
		ReferenceID: referenceID,
	}
	
	err := s.repo.Payment.CreateTransaction(transaction)
	if err != nil {
		return nil, err
	}
	
	// Update the balance
	balance, err := s.repo.Payment.GetBalance()
	if err != nil {
		return nil, err
	}
	
	balance.CurrentBalance += amount
	err = s.repo.Payment.UpdateBalance(balance.InitialBalance)
	if err != nil {
		return nil, err
	}
	
	return transaction, nil
}

func (s *paymentService) RecordCashOut(category models.TransactionCategory, amount float64, description string) (*models.Transaction, error) {
	if amount <= 0 {
		return nil, errors.New("amount must be greater than zero")
	}
	
	// Check if there's enough balance for the cash out
	balance, err := s.repo.Payment.GetBalance()
	if err != nil {
		return nil, err
	}
	
	if balance.CurrentBalance < amount {
		return nil, errors.New("insufficient balance for cash out transaction")
	}
	
	transaction := &models.Transaction{
		BaseModel:   models.BaseModel{ID: uuid.NewString()},
		Type:        models.CashOut,
		Category:    category,
		Amount:      amount,
		Description: description,
		Date:        time.Now(),
	}
	
	err = s.repo.Payment.CreateTransaction(transaction)
	if err != nil {
		return nil, err
	}
	
	// Update the balance
	balance.CurrentBalance -= amount
	err = s.repo.Payment.UpdateBalance(balance.InitialBalance)
	if err != nil {
		return nil, err
	}
	
	return transaction, nil
}

func (s *paymentService) UpdateBalance(initialBalance float64, notes string) error {
	return s.repo.Payment.UpdateBalance(initialBalance)
}

func (s *paymentService) GetBalance() (*models.Balance, error) {
	return s.repo.Payment.GetBalance()
}

func (s *paymentService) GetDashboardData() (*interfaces.DashboardData, error) {
	data, err := s.repo.Payment.GetDashboardData()
	if err != nil {
		return nil, err
	}
	
	// Get low stock products from product repository
	lowStockProducts, err := s.repo.Product.GetLowStock()
	if err != nil {
		// If there's an error getting low stock products, continue with an empty list
		lowStockProducts = []models.Product{}
	}
	
	data.LowStockAlerts = lowStockProducts
	
	return &interfaces.DashboardData{
		CurrentBalance:     data.CurrentBalance,
		TodayCashIn:        data.TodayCashIn,
		ThisMonthCashIn:    data.ThisMonthCashIn,
		AllTimeCashIn:      data.AllTimeCashIn,
		TodayCashOut:       data.TodayCashOut,
		ThisMonthCashOut:   data.ThisMonthCashOut,
		AllTimeCashOut:     data.AllTimeCashOut,
		RecentTransactions: data.RecentTransactions,
		LowStockAlerts:     data.LowStockAlerts,
		UnpaidOrders:       data.UnpaidOrders,
	}, nil
}