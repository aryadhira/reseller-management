package services

import (
	"errors"

	"github.com/aryadhira/reseller-management/internal/models"
	"github.com/aryadhira/reseller-management/internal/repository"
	"github.com/google/uuid"
)

type productService struct {
	repo *repository.Repository
}

func NewProductService(repo *repository.Repository) *productService {
	return &productService{repo: repo}
}

func (s *productService) CreateProduct(product *models.Product) (*models.Product, error) {
	product.BaseModel = models.BaseModel{ID: uuid.NewString()}
	
	err := s.repo.Product.Create(product)
	if err != nil {
		return nil, err
	}
	
	return product, nil
}

func (s *productService) GetAllProducts() ([]models.Product, error) {
	return s.repo.Product.GetAll()
}

func (s *productService) GetProductByID(id string) (*models.Product, error) {
	return s.repo.Product.GetByID(id)
}

func (s *productService) UpdateProduct(id string, product *models.Product) (*models.Product, error) {
	existing, err := s.repo.Product.GetByID(id)
	if err != nil {
		return nil, errors.New("product not found")
	}
	
	product.ID = existing.ID // Preserve the ID
	
	err = s.repo.Product.Update(id, product)
	if err != nil {
		return nil, err
	}
	
	return product, nil
}

func (s *productService) DeleteProduct(id string) error {
	// Check if the product is used in any orders
	orders, err := s.getOrdersContainingProduct(id)
	if err != nil {
		return err
	}
	
	if len(orders) > 0 {
		return errors.New("cannot delete product that is used in existing orders")
	}
	
	return s.repo.Product.Delete(id)
}

func (s *productService) RestockProduct(id string, quantity int) (*models.Product, error) {
	_, err := s.repo.Product.GetByID(id)
	if err != nil {
		return nil, errors.New("product not found")
	}
	
	err = s.repo.Product.Restock(id, quantity)
	if err != nil {
		return nil, err
	}
	
	updatedProduct, err := s.repo.Product.GetByID(id)
	if err != nil {
		return nil, err
	}
	
	return updatedProduct, nil
}

func (s *productService) GetLowStockProducts() ([]models.Product, error) {
	return s.repo.Product.GetLowStock()
}

func (s *productService) getOrdersContainingProduct(productID string) ([]models.Order, error) {
	// This would need to be implemented to check if any orders contain this product
	// For now, we'll return an empty list
	// In a real implementation, we'd need to query orders that contain this product
	return []models.Order{}, nil
}