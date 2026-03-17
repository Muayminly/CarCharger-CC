package repository

import (
	"errors"
	"sync"

	"CarCharger-CC/model"
)

type PaymentRepository interface {
	Create(payment model.Payment) (*model.Payment, error)
	FindByID(paymentID string) (*model.Payment, error)
	FindByBillID(billID string) ([]model.Payment, error)
	FindByUserID(userID string) ([]model.Payment, error)
	FindAll() ([]model.Payment, error)
	UpdateStatus(paymentID string, status model.PaymentStatus) error
}

type InMemoryPaymentRepository struct {
	mu       sync.RWMutex
	payments map[string]model.Payment
}

func NewInMemoryPaymentRepository() *InMemoryPaymentRepository {
	return &InMemoryPaymentRepository{payments: make(map[string]model.Payment)}
}

func (r *InMemoryPaymentRepository) Create(payment model.Payment) (*model.Payment, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if payment.PaymentID == "" {
		payment.PaymentID = "PAY-001"
	}
	r.payments[payment.PaymentID] = payment
	return &payment, nil
}

func (r *InMemoryPaymentRepository) FindByID(paymentID string) (*model.Payment, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	payment, ok := r.payments[paymentID]
	if !ok {
		return nil, errors.New("payment not found")
	}
	return &payment, nil
}

func (r *InMemoryPaymentRepository) FindByBillID(billID string) ([]model.Payment, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	result := []model.Payment{}
	for _, payment := range r.payments {
		if payment.BillID == billID {
			result = append(result, payment)
		}
	}
	return result, nil
}

func (r *InMemoryPaymentRepository) FindByUserID(userID string) ([]model.Payment, error) {
	// Skeleton only: no direct userID field on payment, so return all or adapt later.
	return r.FindAll()
}

func (r *InMemoryPaymentRepository) FindAll() ([]model.Payment, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	result := make([]model.Payment, 0, len(r.payments))
	for _, payment := range r.payments {
		result = append(result, payment)
	}
	return result, nil
}

func (r *InMemoryPaymentRepository) UpdateStatus(paymentID string, status model.PaymentStatus) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	payment, ok := r.payments[paymentID]
	if !ok {
		return errors.New("payment not found")
	}
	payment.PaymentStatus = status
	r.payments[paymentID] = payment
	return nil
}
