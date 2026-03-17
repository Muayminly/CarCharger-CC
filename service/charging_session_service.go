package service

import (
	"CarCharger-CC/model"
	"CarCharger-CC/repository"
)

type ChargingSessionService interface {
	StartSession(userID string, reservationID string, slotID string) (*model.ChargingSession, error)
	UpdateChargingStatus(sessionID string, batteryPercent int, energyUsed float64) error
	EndSession(sessionID string) (*model.Billing, error)
	GetSessionByID(sessionID string) (*model.ChargingSession, error)
	GetActiveSessionByUserID(userID string) (*model.ChargingSession, error)
}

type chargingSessionService struct {
	sessionRepo repository.ChargingSessionRepository
	billingRepo repository.BillingRepository
}

func NewChargingSessionService(sessionRepo repository.ChargingSessionRepository, billingRepo repository.BillingRepository) ChargingSessionService {
	return &chargingSessionService{sessionRepo: sessionRepo, billingRepo: billingRepo}
}

func (s *chargingSessionService) StartSession(userID string, reservationID string, slotID string) (*model.ChargingSession, error) {
	session := model.ChargingSession{
		UserID:        userID,
		ReservationID: reservationID,
		SlotID:        slotID,
		Status:        model.ChargingOngoing,
	}
	return s.sessionRepo.Create(session)
}

func (s *chargingSessionService) UpdateChargingStatus(sessionID string, batteryPercent int, energyUsed float64) error {
	session, err := s.sessionRepo.FindByID(sessionID)
	if err != nil {
		return err
	}
	session.CurrentBatteryPercent = batteryPercent
	session.EnergyConsumedKWh = energyUsed
	return s.sessionRepo.Update(*session)
}

func (s *chargingSessionService) EndSession(sessionID string) (*model.Billing, error) {
	session, err := s.sessionRepo.FindByID(sessionID)
	if err != nil {
		return nil, err
	}
	session.Status = model.ChargingCompleted
	_ = s.sessionRepo.Update(*session)
	billing := model.Billing{
		SessionID:     sessionID,
		EnergyUsedKWh: session.EnergyConsumedKWh,
		TotalAmount:   session.EnergyConsumedKWh * 7,
	}
	return s.billingRepo.Create(billing)
}

func (s *chargingSessionService) GetSessionByID(sessionID string) (*model.ChargingSession, error) {
	return s.sessionRepo.FindByID(sessionID)
}

func (s *chargingSessionService) GetActiveSessionByUserID(userID string) (*model.ChargingSession, error) {
	return s.sessionRepo.FindActiveByUserID(userID)
}
