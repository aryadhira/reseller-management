package repository

import (
	"github.com/aryadhira/reseller-management/internal/models"
	"gorm.io/gorm"
)

type orderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) *orderRepository {
	return &orderRepository{db: db}
}

func (r *orderRepository) Create(order *models.Order) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		// Create the order
		if err := tx.Create(order).Error; err != nil {
			return err
		}

		// Create order items
		for i := range order.OrderItems {
			order.OrderItems[i].OrderID = order.ID
			if err := tx.Create(&order.OrderItems[i]).Error; err != nil {
				return err
			}
		}

		return nil
	})
}

func (r *orderRepository) GetAll() ([]models.Order, error) {
	var orders []models.Order
	err := r.db.Preload("Reseller").Preload("OrderItems").Preload("OrderItems.Product").Preload("Payment").Find(&orders).Error
	return orders, err
}

func (r *orderRepository) GetByID(id string) (*models.Order, error) {
	var order models.Order
	err := r.db.Preload("Reseller").Preload("OrderItems").Preload("OrderItems.Product").Preload("Payment").Where("id = ?", id).First(&order).Error
	return &order, err
}

func (r *orderRepository) Update(id string, order *models.Order) error {
	return r.db.Model(&models.Order{}).Where("id = ?", id).Updates(order).Error
}

func (r *orderRepository) Delete(id string) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		// Delete order items first
		if err := tx.Where("order_id = ?", id).Delete(&models.OrderItem{}).Error; err != nil {
			return err
		}

		// Then delete the order
		return tx.Delete(&models.Order{}, "id = ?", id).Error
	})
}

func (r *orderRepository) Cancel(id string) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		// Get the order to access its items
		var order models.Order
		err := tx.Preload("OrderItems").Where("id = ?", id).First(&order).Error
		if err != nil {
			return err
		}

		// Update order status to cancelled
		if err := tx.Model(&models.Order{}).Where("id = ?", id).Update("status", "cancelled").Error; err != nil {
			return err
		}

		// Restore stock for each order item
		for _, item := range order.OrderItems {
			if err := tx.Model(&models.Product{}).Where("id = ?", item.ProductID).
				UpdateColumn("current_stock", gorm.Expr("current_stock + ?", item.Quantity)).Error; err != nil {
				return err
			}
		}

		return nil
	})
}