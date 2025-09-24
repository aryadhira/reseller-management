package repository

import (
	"github.com/aryadhira/reseller-management/internal/models"
	"gorm.io/gorm"
)

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *productRepository {
	return &productRepository{db: db}
}

func (r *productRepository) Create(product *models.Product) error {
	return r.db.Create(product).Error
}

func (r *productRepository) GetAll() ([]models.Product, error) {
	var products []models.Product
	err := r.db.Find(&products).Error
	return products, err
}

func (r *productRepository) GetByID(id string) (*models.Product, error) {
	var product models.Product
	err := r.db.Where("id = ?", id).First(&product).Error
	return &product, err
}

func (r *productRepository) Update(id string, product *models.Product) error {
	return r.db.Model(&models.Product{}).Where("id = ?", id).Updates(product).Error
}

func (r *productRepository) Delete(id string) error {
	return r.db.Delete(&models.Product{}, "id = ?", id).Error
}

func (r *productRepository) Restock(id string, quantity int) error {
	return r.db.Model(&models.Product{}).Where("id = ?", id).UpdateColumn("current_stock", gorm.Expr("current_stock + ?", quantity)).Error
}

func (r *productRepository) GetLowStock() ([]models.Product, error) {
	var products []models.Product
	err := r.db.Where("current_stock <= min_stock_alert AND status = 'active'").Find(&products).Error
	return products, err
}