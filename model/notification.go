package model

type Notification struct {
	NotificationID   string           `json:"notificationId"`
	UserID           string           `json:"userId"`
	Message          string           `json:"message"`
	NotificationType NotificationType `json:"notificationType"`
	SentTime         string           `json:"sentTime"`
}
