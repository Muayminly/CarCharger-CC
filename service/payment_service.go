package service

import (
	"CarCharger-CC/model"
	"CarCharger-CC/repository"
)

type PaymentService interface {
	ProcessPayment(billID string, method model.PaymentMethod) (*model.Payment, error)
	VerifyPayment(paymentID string) error
	GetPaymentByID(paymentID string) (*model.Payment, error)
	GetPaymentsByBillID(billID string) ([]model.Payment, error)
}

type paymentService struct {
	paymentRepo repository.PaymentRepository
	billingRepo repository.BillingRepository
}

func NewPaymentService(paymentRepo repository.PaymentRepository, billingRepo repository.BillingRepository) PaymentService {
	return &paymentService{paymentRepo: paymentRepo, billingRepo: billingRepo}
}

func (s *paymentService) ProcessPayment(billID string, method model.PaymentMethod) (*model.Payment, error) {
	_, err := s.billingRepo.FindByID(billID)
	if err != nil {
		return nil, err
	}
	payment := model.Payment{
		BillID:        billID,
		PaymentMethod: method,
		PaymentStatus: model.PaymentPending,
	}
	return s.paymentRepo.Create(payment)
}

func (s *paymentService) VerifyPayment(paymentID string) error {
	return s.paymentRepo.UpdateStatus(paymentID, model.PaymentSuccess)
}

func (s *paymentService) GetPaymentByID(paymentID string) (*model.Payment, error) {
	return s.paymentRepo.FindByID(paymentID)
}

func (s *paymentService) GetPaymentsByBillID(billID string) ([]model.Payment, error) {
	return s.paymentRepo.FindByBillID(billID)
}
