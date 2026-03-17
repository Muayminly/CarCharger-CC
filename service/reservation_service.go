package service

import (
	"CarCharger-CC/model"
	"CarCharger-CC/repository"
)

type ReservationService interface {
	MakeReservation(userID string, slotID string, startTime string, endTime string) (*model.Reservation, error)
	CheckIn(reservationID string) (*model.ChargingSession, error)
	WalkIn(userID string, slotID string) (*model.ChargingSession, error)
	CancelReservation(reservationID string) error
	MarkNoShow(reservationID string) error
	GetReservationByID(reservationID string) (*model.Reservation, error)
	GetReservationsByUserID(userID string) ([]model.Reservation, error)
}

type reservationService struct {
	reservationRepo repository.ReservationRepository
	slotRepo        repository.ChargingSlotRepository
	sessionRepo     repository.ChargingSessionRepository
}

func NewReservationService(reservationRepo repository.ReservationRepository, slotRepo repository.ChargingSlotRepository, sessionRepo repository.ChargingSessionRepository) ReservationService {
	return &reservationService{reservationRepo: reservationRepo, slotRepo: slotRepo, sessionRepo: sessionRepo}
}

func (s *reservationService) MakeReservation(userID string, slotID string, startTime string, endTime string) (*model.Reservation, error) {
	_ = s.slotRepo.UpdateStatus(slotID, model.SlotReserved)
	reservation := model.Reservation{
		UserID:            userID,
		SlotID:            slotID,
		StartTime:         startTime,
		EndTime:           endTime,
		ReservationStatus: model.ReservationReserved,
		MaxLateTime:       15,
	}
	return s.reservationRepo.Create(reservation)
}

func (s *reservationService) CheckIn(reservationID string) (*model.ChargingSession, error) {
	reservation, err := s.reservationRepo.FindByID(reservationID)
	if err != nil {
		return nil, err
	}
	_ = s.reservationRepo.UpdateStatus(reservationID, model.ReservationCheckedIn)
	_ = s.slotRepo.UpdateStatus(reservation.SlotID, model.SlotCharging)
	session := model.ChargingSession{
		UserID:        reservation.UserID,
		ReservationID: reservation.ReservationID,
		SlotID:        reservation.SlotID,
		Status:        model.ChargingOngoing,
	}
	return s.sessionRepo.Create(session)
}

func (s *reservationService) WalkIn(userID string, slotID string) (*model.ChargingSession, error) {
	_ = s.slotRepo.UpdateStatus(slotID, model.SlotCharging)
	session := model.ChargingSession{
		UserID: userID,
		SlotID: slotID,
		Status: model.ChargingOngoing,
	}
	return s.sessionRepo.Create(session)
}

func (s *reservationService) CancelReservation(reservationID string) error {
	reservation, err := s.reservationRepo.FindByID(reservationID)
	if err != nil {
		return err
	}
	_ = s.slotRepo.UpdateStatus(reservation.SlotID, model.SlotAvailable)
	return s.reservationRepo.UpdateStatus(reservationID, model.ReservationCancelled)
}

func (s *reservationService) MarkNoShow(reservationID string) error {
	reservation, err := s.reservationRepo.FindByID(reservationID)
	if err != nil {
		return err
	}
	_ = s.slotRepo.UpdateStatus(reservation.SlotID, model.SlotAvailable)
	return s.reservationRepo.UpdateStatus(reservationID, model.ReservationNoShow)
}

func (s *reservationService) GetReservationByID(reservationID string) (*model.Reservation, error) {
	return s.reservationRepo.FindByID(reservationID)
}

func (s *reservationService) GetReservationsByUserID(userID string) ([]model.Reservation, error) {
	return s.reservationRepo.FindByUserID(userID)
}
