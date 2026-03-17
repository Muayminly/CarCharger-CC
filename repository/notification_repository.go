package repository

import (
	"sync"

	"CarCharger-CC/model"
)

type NotificationRepository interface {
	Create(notification model.Notification) error
	FindByUserID(userID string) ([]model.Notification, error)
}

type InMemoryNotificationRepository struct {
	mu            sync.RWMutex
	notifications []model.Notification
}

func NewInMemoryNotificationRepository() *InMemoryNotificationRepository {
	return &InMemoryNotificationRepository{notifications: []model.Notification{}}
}

func (r *InMemoryNotificationRepository) Create(notification model.Notification) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.notifications = append(r.notifications, notification)
	return nil
}

func (r *InMemoryNotificationRepository) FindByUserID(userID string) ([]model.Notification, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	result := []model.Notification{}
	for _, notification := range r.notifications {
		if notification.UserID == userID {
			result = append(result, notification)
		}
	}
	return result, nil
}
