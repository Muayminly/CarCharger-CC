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
