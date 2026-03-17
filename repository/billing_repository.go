package repository

import (
	"errors"
	"sync"

	"CarCharger-CC/model"
)

type BillingRepository interface {
	Create(billing model.Billing) (*model.Billing, error)
	FindByID(billID string) (*model.Billing, error)
	FindBySessionID(sessionID string) (*model.Billing, error)
	FindAll() ([]model.Billing, error)
}

type InMemoryBillingRepository struct {
	mu       sync.RWMutex
	billings map[string]model.Billing
}

func NewInMemoryBillingRepository() *InMemoryBillingRepository {
	return &InMemoryBillingRepository{billings: make(map[string]model.Billing)}
}

func (r *InMemoryBillingRepository) Create(billing model.Billing) (*model.Billing, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if billing.BillID == "" {
		billing.BillID = "BIL-001"
	}
	r.billings[billing.BillID] = billing
	return &billing, nil
}

func (r *InMemoryBillingRepository) FindByID(billID string) (*model.Billing, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	billing, ok := r.billings[billID]
	if !ok {
		return nil, errors.New("billing not found")
	}
	return &billing, nil
}

func (r *InMemoryBillingRepository) FindBySessionID(sessionID string) (*model.Billing, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	for _, billing := range r.billings {
		if billing.SessionID == sessionID {
			copyBilling := billing
			return &copyBilling, nil
		}
	}
	return nil, errors.New("billing not found")
}

func (r *InMemoryBillingRepository) FindAll() ([]model.Billing, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	result := make([]model.Billing, 0, len(r.billings))
	for _, billing := range r.billings {
		result = append(result, billing)
	}
	return result, nil
}
