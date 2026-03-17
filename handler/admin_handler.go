package handler

import (
	"net/http"

	"CarCharger-CC/model"
	"CarCharger-CC/service"

	"github.com/gin-gonic/gin"
)

type AdminHandler struct {
	service service.AdminService
}

func NewAdminHandler(service service.AdminService) *AdminHandler {
	return &AdminHandler{service: service}
}

func (h *AdminHandler) ManageUser(c *gin.Context) {
	var req model.User
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.service.ManageUser(req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "user saved"})
}

func (h *AdminHandler) ManageStation(c *gin.Context) {
	var req model.Station
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.service.ManageStation(req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "station saved"})
}

func (h *AdminHandler) ManageSlot(c *gin.Context) {
	var req model.ChargingSlot
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.service.ManageSlot(req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "slot saved"})
}

func (h *AdminHandler) MonitorReservations(c *gin.Context) {
	items, err := h.service.MonitorReservations()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, items)
}

func (h *AdminHandler) MonitorChargingSessions(c *gin.Context) {
	items, err := h.service.MonitorChargingSessions()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, items)
}

func (h *AdminHandler) VerifyPayments(c *gin.Context) {
	items, err := h.service.VerifyPayments()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, items)
}
