package repository

import (
	"errors"
	"sync"

	"CarCharger-CC/model"
)

type ReservationRepository interface {
	Create(reservation model.Reservation) (*model.Reservation, error)
	FindByID(reservationID string) (*model.Reservation, error)
	FindByUserID(userID string) ([]model.Reservation, error)
	FindAll() ([]model.Reservation, error)
	UpdateStatus(reservationID string, status model.ReservationStatus) error
}

type InMemoryReservationRepository struct {
	mu           sync.RWMutex
	reservations map[string]model.Reservation
}

func NewInMemoryReservationRepository() *InMemoryReservationRepository {
	return &InMemoryReservationRepository{reservations: make(map[string]model.Reservation)}
}

func (r *InMemoryReservationRepository) Create(reservation model.Reservation) (*model.Reservation, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if reservation.ReservationID == "" {
		reservation.ReservationID = "RES-001"
	}
	r.reservations[reservation.ReservationID] = reservation
	return &reservation, nil
}

func (r *InMemoryReservationRepository) FindByID(reservationID string) (*model.Reservation, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	reservation, ok := r.reservations[reservationID]
	if !ok {
		return nil, errors.New("reservation not found")
	}
	return &reservation, nil
}

func (r *InMemoryReservationRepository) FindByUserID(userID string) ([]model.Reservation, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	result := []model.Reservation{}
	for _, reservation := range r.reservations {
		if reservation.UserID == userID {
			result = append(result, reservation)
		}
	}
	return result, nil
}

func (r *InMemoryReservationRepository) FindAll() ([]model.Reservation, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	result := make([]model.Reservation, 0, len(r.reservations))
	for _, reservation := range r.reservations {
		result = append(result, reservation)
	}
	return result, nil
}

func (r *InMemoryReservationRepository) UpdateStatus(reservationID string, status model.ReservationStatus) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	reservation, ok := r.reservations[reservationID]
	if !ok {
		return errors.New("reservation not found")
	}
	reservation.ReservationStatus = status
	r.reservations[reservationID] = reservation
	return nil
}
