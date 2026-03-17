package service

import (
	"CarCharger-CC/model"
	"CarCharger-CC/repository"
)

type UserService interface {
	GetUserByID(userID string) (*model.User, error)
	GetDriverByID(userID string) (*model.EVDriver, error)
	ViewPaymentHistory(userID string) ([]model.Payment, error)
	ViewUsageHistory(userID string) ([]model.ChargingSession, error)
}

type userService struct {
	userRepo    repository.UserRepository
	sessionRepo repository.ChargingSessionRepository
	paymentRepo repository.PaymentRepository
}

func NewUserService(userRepo repository.UserRepository, sessionRepo repository.ChargingSessionRepository, paymentRepo repository.PaymentRepository) UserService {
	return &userService{userRepo: userRepo, sessionRepo: sessionRepo, paymentRepo: paymentRepo}
}

func (s *userService) GetUserByID(userID string) (*model.User, error) {
	return s.userRepo.FindByID(userID)
}

func (s *userService) GetDriverByID(userID string) (*model.EVDriver, error) {
	return s.userRepo.FindDriverByID(userID)
}

func (s *userService) ViewPaymentHistory(userID string) ([]model.Payment, error) {
	return s.paymentRepo.FindByUserID(userID)
}

func (s *userService) ViewUsageHistory(userID string) ([]model.ChargingSession, error) {
	return s.sessionRepo.FindByUserID(userID)
}
