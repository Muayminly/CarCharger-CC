// Database Mapping
package repository

import (
	"errors"
	"sync"

	"CarCharger-CC/model"
)

type UserRepository interface {
	CreateDriver(driver model.EVDriver) (*model.EVDriver, error)
	FindByID(userID string) (*model.User, error)
	FindDriverByID(userID string) (*model.EVDriver, error)
	FindByEmailOrPhone(emailOrPhone string) (*model.User, error)
	Update(user model.User) error
	Save(user model.User) error
	ListUsers() ([]model.User, error)
}

type InMemoryUserRepository struct {
	mu      sync.RWMutex
	users   map[string]model.User
	drivers map[string]model.EVDriver
}

func NewInMemoryUserRepository() *InMemoryUserRepository {
	return &InMemoryUserRepository{
		users:   make(map[string]model.User),
		drivers: make(map[string]model.EVDriver),
	}
}

func (r *InMemoryUserRepository) CreateDriver(driver model.EVDriver) (*model.EVDriver, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if driver.UserID == "" {
		driver.UserID = "USR-001"
	}
	r.users[driver.UserID] = driver.User
	r.drivers[driver.UserID] = driver
	return &driver, nil
}

func (r *InMemoryUserRepository) FindByID(userID string) (*model.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	user, ok := r.users[userID]
	if !ok {
		return nil, errors.New("user not found")
	}
	return &user, nil
}

func (r *InMemoryUserRepository) FindDriverByID(userID string) (*model.EVDriver, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	driver, ok := r.drivers[userID]
	if !ok {
		return nil, errors.New("driver not found")
	}
	return &driver, nil
}

func (r *InMemoryUserRepository) FindByEmailOrPhone(emailOrPhone string) (*model.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	for _, user := range r.users {
		if user.Email == emailOrPhone || user.PhoneNumber == emailOrPhone {
			copyUser := user
			return &copyUser, nil
		}
	}
	return nil, errors.New("user not found")
}

func (r *InMemoryUserRepository) Update(user model.User) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.users[user.UserID] = user
	if driver, ok := r.drivers[user.UserID]; ok {
		driver.User = user
		r.drivers[user.UserID] = driver
	}
	return nil
}

func (r *InMemoryUserRepository) Save(user model.User) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.users[user.UserID] = user
	return nil
}

func (r *InMemoryUserRepository) ListUsers() ([]model.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	result := make([]model.User, 0, len(r.users))
	for _, u := range r.users {
		result = append(result, u)
	}
	return result, nil
}
