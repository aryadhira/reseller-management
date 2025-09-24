package repository

import (
	"github.com/aryadhira/reseller-management/internal/models"
	"gorm.io/gorm"
)

type resellerRepository struct {
	db *gorm.DB
}

func NewResellerRepository(db *gorm.DB) *resellerRepository {
	return &resellerRepository{db: db}
}

func (r *resellerRepository) Create(reseller *models.Reseller) error {
	return r.db.Create(reseller).Error
}

func (r *resellerRepository) GetAll() ([]models.Reseller, error) {
	var resellers []models.Reseller
	err := r.db.Find(&resellers).Error
	return resellers, err
}

func (r *resellerRepository) GetByID(id string) (*models.Reseller, error) {
	var reseller models.Reseller
	err := r.db.Where("id = ?", id).First(&reseller).Error
	return &reseller, err
}

func (r *resellerRepository) Update(id string, reseller *models.Reseller) error {
	return r.db.Model(&models.Reseller{}).Where("id = ?", id).Updates(reseller).Error
}

func (r *resellerRepository) Delete(id string) error {
	return r.db.Delete(&models.Reseller{}, "id = ?", id).Error
}

func (r *resellerRepository) GetWithOrders(id string) (*models.Reseller, error) {
	var reseller models.Reseller
	err := r.db.Preload("Orders").Preload("Orders.OrderItems").Preload("Orders.Payment").Where("id = ?", id).First(&reseller).Error
	return &reseller, err
}