package repository

import (
	"errors"
	"sync"

	"CarCharger-CC/model"
)

type ChargingSessionRepository interface {
	Create(session model.ChargingSession) (*model.ChargingSession, error)
	FindByID(sessionID string) (*model.ChargingSession, error)
	FindActiveByUserID(userID string) (*model.ChargingSession, error)
	FindByUserID(userID string) ([]model.ChargingSession, error)
	FindAll() ([]model.ChargingSession, error)
	Update(session model.ChargingSession) error
}

type InMemoryChargingSessionRepository struct {
	mu       sync.RWMutex
	sessions map[string]model.ChargingSession
}

func NewInMemoryChargingSessionRepository() *InMemoryChargingSessionRepository {
	return &InMemoryChargingSessionRepository{sessions: make(map[string]model.ChargingSession)}
}

func (r *InMemoryChargingSessionRepository) Create(session model.ChargingSession) (*model.ChargingSession, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if session.SessionID == "" {
		session.SessionID = "SES-001"
	}
	r.sessions[session.SessionID] = session
	return &session, nil
}

func (r *InMemoryChargingSessionRepository) FindByID(sessionID string) (*model.ChargingSession, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	session, ok := r.sessions[sessionID]
	if !ok {
		return nil, errors.New("session not found")
	}
	return &session, nil
}

func (r *InMemoryChargingSessionRepository) FindActiveByUserID(userID string) (*model.ChargingSession, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	for _, session := range r.sessions {
		if session.UserID == userID && session.Status == model.ChargingOngoing {
			copySession := session
			return &copySession, nil
		}
	}
	return nil, errors.New("active session not found")
}

func (r *InMemoryChargingSessionRepository) FindByUserID(userID string) ([]model.ChargingSession, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	result := []model.ChargingSession{}
	for _, session := range r.sessions {
		if session.UserID == userID {
			result = append(result, session)
		}
	}
	return result, nil
}

func (r *InMemoryChargingSessionRepository) FindAll() ([]model.ChargingSession, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	result := make([]model.ChargingSession, 0, len(r.sessions))
	for _, session := range r.sessions {
		result = append(result, session)
	}
	return result, nil
}

func (r *InMemoryChargingSessionRepository) Update(session model.ChargingSession) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.sessions[session.SessionID] = session
	return nil
}
