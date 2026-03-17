package handler

import (
	"net/http"

	"CarCharger-CC/model"
	"CarCharger-CC/service"

	"github.com/gin-gonic/gin"
)

type NotificationHandler struct {
	service service.NotificationService
}

func NewNotificationHandler(service service.NotificationService) *NotificationHandler {
	return &NotificationHandler{service: service}
}

func (h *NotificationHandler) SendNotification(c *gin.Context) {
	var req struct {
		UserID  string                 `json:"userId"`
		Message string                 `json:"message"`
		Type    model.NotificationType `json:"notificationType"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.service.SendNotification(req.UserID, req.Message, req.Type); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "notification sent"})
}

func (h *NotificationHandler) GetNotificationsByUserID(c *gin.Context) {
	items, err := h.service.GetNotificationsByUserID(c.Param("userId"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, items)
}
