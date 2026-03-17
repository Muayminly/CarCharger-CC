package service

import (
	"errors"

	"CarCharger-CC/model"
	"CarCharger-CC/repository"
)

type BillingService interface {
	CalculateCost(sessionID string) (float64, error)
	GenerateBill(sessionID string) (*model.Billing, error)
	GetBillByID(billID string) (*model.Billing, error)
	GetBillsByUserID(userID string) ([]model.Billing, error)
}

type billingService struct {
	billingRepo repository.BillingRepository
	sessionRepo repository.ChargingSessionRepository
}

func NewBillingService(billingRepo repository.BillingRepository, sessionRepo repository.ChargingSessionRepository) BillingService {
	return &billingService{billingRepo: billingRepo, sessionRepo: sessionRepo}
}

func (s *billingService) CalculateCost(sessionID string) (float64, error) {
	session, err := s.sessionRepo.FindByID(sessionID)
	if err != nil {
		return 0, err
	}
	return session.EnergyConsumedKWh * 7, nil
}

func (s *billingService) GenerateBill(sessionID string) (*model.Billing, error) {
	session, err := s.sessionRepo.FindByID(sessionID)
	if err != nil {
		return nil, err
	}
	billing := model.Billing{
		SessionID:     sessionID,
		EnergyUsedKWh: session.EnergyConsumedKWh,
		TotalAmount:   session.EnergyConsumedKWh * 7,
	}
	return s.billingRepo.Create(billing)
}

func (s *billingService) GetBillByID(billID string) (*model.Billing, error) {
	return s.billingRepo.FindByID(billID)
}

func (s *billingService) GetBillsByUserID(userID string) ([]model.Billing, error) {
	sessions, err := s.sessionRepo.FindByUserID(userID)
	if err != nil {
		return nil, err
	}
	result := []model.Billing{}
	for _, session := range sessions {
		billing, err := s.billingRepo.FindBySessionID(session.SessionID)
		if err == nil {
			result = append(result, *billing)
		}
	}
	if len(result) == 0 {
		return result, errors.New("no billing found")
	}
	return result, nil
}
