package service

import (
	"CarCharger-CC/model"
	"CarCharger-CC/repository"
)

type AdminService interface {
	ManageUser(user model.User) error
	ManageStation(station model.Station) error
	ManageSlot(slot model.ChargingSlot) error
	MonitorReservations() ([]model.Reservation, error)
	MonitorChargingSessions() ([]model.ChargingSession, error)
	VerifyPayments() ([]model.Payment, error)
}

type adminService struct {
	userRepo        repository.UserRepository
	stationRepo     repository.StationRepository
	slotRepo        repository.ChargingSlotRepository
	reservationRepo repository.ReservationRepository
	sessionRepo     repository.ChargingSessionRepository
	paymentRepo     repository.PaymentRepository
}

func NewAdminService(
	userRepo repository.UserRepository,
	stationRepo repository.StationRepository,
	slotRepo repository.ChargingSlotRepository,
	reservationRepo repository.ReservationRepository,
	sessionRepo repository.ChargingSessionRepository,
	paymentRepo repository.PaymentRepository,
) AdminService {
	return &adminService{
		userRepo:        userRepo,
		stationRepo:     stationRepo,
		slotRepo:        slotRepo,
		reservationRepo: reservationRepo,
		sessionRepo:     sessionRepo,
		paymentRepo:     paymentRepo,
	}
}

func (s *adminService) ManageUser(user model.User) error {
	return s.userRepo.Save(user)
}

func (s *adminService) ManageStation(station model.Station) error {
	return s.stationRepo.Save(station)
}

func (s *adminService) ManageSlot(slot model.ChargingSlot) error {
	return s.slotRepo.Save(slot)
}

func (s *adminService) MonitorReservations() ([]model.Reservation, error) {
	return s.reservationRepo.FindAll()
}

func (s *adminService) MonitorChargingSessions() ([]model.ChargingSession, error) {
	return s.sessionRepo.FindAll()
}

func (s *adminService) VerifyPayments() ([]model.Payment, error) {
	return s.paymentRepo.FindAll()
}
