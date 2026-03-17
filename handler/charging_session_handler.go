package handler

import (
	"net/http"

	"CarCharger-CC/service"

	"github.com/gin-gonic/gin"
)

type ChargingSessionHandler struct {
	service service.ChargingSessionService
}

func NewChargingSessionHandler(service service.ChargingSessionService) *ChargingSessionHandler {
	return &ChargingSessionHandler{service: service}
}

func (h *ChargingSessionHandler) StartSession(c *gin.Context) {
	var req struct {
		UserID        string `json:"userId"`
		ReservationID string `json:"reservationId"`
		SlotID        string `json:"slotId"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	session, err := h.service.StartSession(req.UserID, req.ReservationID, req.SlotID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, session)
}

func (h *ChargingSessionHandler) UpdateChargingStatus(c *gin.Context) {
	var req struct {
		BatteryPercent int     `json:"batteryPercent"`
		EnergyUsed     float64 `json:"energyUsed"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.service.UpdateChargingStatus(c.Param("sessionId"), req.BatteryPercent, req.EnergyUsed); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "charging status updated"})
}

func (h *ChargingSessionHandler) EndSession(c *gin.Context) {
	billing, err := h.service.EndSession(c.Param("sessionId"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, billing)
}

func (h *ChargingSessionHandler) GetSessionByID(c *gin.Context) {
	session, err := h.service.GetSessionByID(c.Param("sessionId"))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, session)
}

func (h *ChargingSessionHandler) GetActiveSessionByUserID(c *gin.Context) {
	session, err := h.service.GetActiveSessionByUserID(c.Param("userId"))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, session)
}
