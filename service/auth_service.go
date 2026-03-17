package service

import (
	"errors"

	"CarCharger-CC/model"
	"CarCharger-CC/repository"
)

type AuthService interface {
	Register(user model.EVDriver) (*model.EVDriver, error)
	Login(emailOrPhone string, password string) (*model.User, error)
	Logout(userID string) error
	UpdateProfile(user model.User) error
}

type authService struct {
	userRepo repository.UserRepository
}

func NewAuthService(userRepo repository.UserRepository) AuthService {
	return &authService{userRepo: userRepo}
}

func (s *authService) Register(user model.EVDriver) (*model.EVDriver, error) {
	if user.UserID == "" {
		user.UserID = "USR-001"
	}
	user.Role = model.RoleEVDriver
	return s.userRepo.CreateDriver(user)
}

func (s *authService) Login(emailOrPhone string, password string) (*model.User, error) {
	user, err := s.userRepo.FindByEmailOrPhone(emailOrPhone)
	if err != nil {
		return nil, err
	}
	if user.Password != password {
		return nil, errors.New("invalid credentials")
	}
	return user, nil
}

func (s *authService) Logout(userID string) error {
	_, err := s.userRepo.FindByID(userID)
	return err
}

func (s *authService) UpdateProfile(user model.User) error {
	return s.userRepo.Update(user)
}
