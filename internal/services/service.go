package services

import (
	"github.com/aryadhira/reseller-management/internal/interfaces"
	"github.com/aryadhira/reseller-management/internal/repository"
)

type Service struct {
	Reseller interfaces.ResellerService
	Product  interfaces.ProductService
	Order    interfaces.OrderService
	Payment  interfaces.PaymentService
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Reseller: NewResellerService(repo),
		Product:  NewProductService(repo),
		Order:    NewOrderService(repo),
		Payment:  NewPaymentService(repo),
	}
}