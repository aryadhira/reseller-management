package services

import (
	"errors"

	"github.com/aryadhira/reseller-management/internal/models"
	"github.com/aryadhira/reseller-management/internal/repository"
	"github.com/google/uuid"
)

type resellerService struct {
	repo *repository.Repository
}

func NewResellerService(repo *repository.Repository) *resellerService {
	return &resellerService{repo: repo}
}

func (s *resellerService) CreateReseller(reseller *models.Reseller) (*models.Reseller, error) {
	reseller.BaseModel = models.BaseModel{ID: uuid.NewString()}
	
	err := s.repo.Reseller.Create(reseller)
	if err != nil {
		return nil, err
	}
	
	return reseller, nil
}

func (s *resellerService) GetAllResellers() ([]models.Reseller, error) {
	return s.repo.Reseller.GetAll()
}

func (s *resellerService) GetResellerByID(id string) (*models.Reseller, error) {
	return s.repo.Reseller.GetByID(id)
}

func (s *resellerService) UpdateReseller(id string, reseller *models.Reseller) (*models.Reseller, error) {
	existing, err := s.repo.Reseller.GetByID(id)
	if err != nil {
		return nil, errors.New("reseller not found")
	}
	
	reseller.BaseModel = models.BaseModel{ID: existing.ID} // Preserve the ID
	
	err = s.repo.Reseller.Update(id, reseller)
	if err != nil {
		return nil, err
	}
	
	return reseller, nil
}

func (s *resellerService) DeleteReseller(id string) error {
	return s.repo.Reseller.Delete(id)
}

func (s *resellerService) GetResellerWithOrders(id string) (*models.Reseller, error) {
	return s.repo.Reseller.GetWithOrders(id)
}