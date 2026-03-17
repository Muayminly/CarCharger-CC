package service

import (
	"CarCharger-CC/model"
	"CarCharger-CC/repository"
)

type NotificationService interface {
	SendNotification(userID string, message string, nType model.NotificationType) error
	GetNotificationsByUserID(userID string) ([]model.Notification, error)
}

type notificationService struct {
	notificationRepo repository.NotificationRepository
}

func NewNotificationService(notificationRepo repository.NotificationRepository) NotificationService {
	return &notificationService{notificationRepo: notificationRepo}
}

func (s *notificationService) SendNotification(userID string, message string, nType model.NotificationType) error {
	noti := model.Notification{
		NotificationID:   "NOTI-001",
		UserID:           userID,
		Message:          message,
		NotificationType: nType,
	}
	return s.notificationRepo.Create(noti)
}

func (s *notificationService) GetNotificationsByUserID(userID string) ([]model.Notification, error) {
	return s.notificationRepo.FindByUserID(userID)
}
