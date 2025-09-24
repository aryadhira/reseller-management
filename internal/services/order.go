package services

import (
	"errors"
	"fmt"

	"github.com/aryadhira/reseller-management/internal/models"
	"github.com/aryadhira/reseller-management/internal/repository"
	"github.com/google/uuid"
)

type orderService struct {
	repo *repository.Repository
}

func NewOrderService(repo *repository.Repository) *orderService {
	return &orderService{repo: repo}
}

func (s *orderService) CreateOrder(order *models.Order) (*models.Order, error) {
	order.BaseModel = models.BaseModel{ID: uuid.NewString()}
	
	// Validate that the reseller exists
	_, err := s.repo.Reseller.GetByID(order.ResellerID)
	if err != nil {
		return nil, errors.New("reseller not found")
	}
	
	// Validate products and check stock availability
	totalAmount := 0.0
	for i := range order.OrderItems {
		product, err := s.repo.Product.GetByID(order.OrderItems[i].ProductID)
		if err != nil {
			return nil, fmt.Errorf("product with ID %s not found", order.OrderItems[i].ProductID)
		}
		
		if product.CurrentStock < order.OrderItems[i].Quantity {
			return nil, fmt.Errorf("insufficient stock for product %s. Available: %d, Requested: %d", 
				product.Name, product.CurrentStock, order.OrderItems[i].Quantity)
		}
		
		// Set the price at the time of order
		order.OrderItems[i].Price = product.Price
		order.OrderItems[i].Subtotal = float64(order.OrderItems[i].Quantity) * product.Price
		totalAmount += order.OrderItems[i].Subtotal
	}
	
	order.TotalAmount = totalAmount
	
	// Deduct stock from products
	for _, item := range order.OrderItems {
		err := s.repo.Product.Restock(item.ProductID, -item.Quantity) // Negative quantity to reduce stock
		if err != nil {
			return nil, err
		}
	}
	
	// Create the order
	err = s.repo.Order.Create(order)
	if err != nil {
		// If order creation fails, restore stock
		for _, item := range order.OrderItems {
			s.repo.Product.Restock(item.ProductID, item.Quantity) // Positive quantity to restore
		}
		return nil, err
	}
	
	// Create payment record with status 'UNPAID'
	payment := &models.Payment{
		BaseModel:   models.BaseModel{ID: uuid.NewString()},
		OrderID:     order.ID,
		TotalAmount: order.TotalAmount,
		AmountPaid:  0,
		Status:      "unpaid",
	}
	
	err = s.repo.Payment.Create(payment)
	if err != nil {
		return nil, err
	}
	
	return order, nil
}

func (s *orderService) GetAllOrders() ([]models.Order, error) {
	return s.repo.Order.GetAll()
}

func (s *orderService) GetOrderByID(id string) (*models.Order, error) {
	return s.repo.Order.GetByID(id)
}

func (s *orderService) UpdateOrder(id string, order *models.Order) (*models.Order, error) {
	existing, err := s.repo.Order.GetByID(id)
	if err != nil {
		return nil, errors.New("order not found")
	}
	
	order.BaseModel = models.BaseModel{ID: existing.ID} // Preserve the ID
	
	err = s.repo.Order.Update(id, order)
	if err != nil {
		return nil, err
	}
	
	return order, nil
}

func (s *orderService) DeleteOrder(id string) error {
	return s.repo.Order.Delete(id)
}

func (s *orderService) CancelOrder(id string) error {
	order, err := s.repo.Order.GetByID(id)
	if err != nil {
		return errors.New("order not found")
	}
	
	if order.Status == "cancelled" {
		return errors.New("order is already cancelled")
	}
	
	// Cancel the order (which will restore stock)
	err = s.repo.Order.Cancel(id)
	if err != nil {
		return err
	}
	
	// Delete the associated payment record
	payment, err := s.repo.Payment.GetByOrderID(id)
	if err == nil {
		// If payment exists, delete it
		// In a real implementation, we might want to soft delete or mark as void
		// For now, let's just update its status to cancelled
		payment.Status = "cancelled"
		err = s.repo.Payment.Update(payment)
		if err != nil {
			return err
		}
	}
	
	// Delete any associated CASH_IN transactions for this order
	// In a real implementation, we would query and delete transactions linked to this order
	
	return nil
}