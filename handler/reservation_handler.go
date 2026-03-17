package handler

import (
	"net/http"

	"CarCharger-CC/service"

	"github.com/gin-gonic/gin"
)

type ReservationHandler struct {
	service service.ReservationService
}

func NewReservationHandler(service service.ReservationService) *ReservationHandler {
	return &ReservationHandler{service: service}
}

func (h *ReservationHandler) MakeReservation(c *gin.Context) {
	var req struct {
		UserID    string `json:"userId"`
		SlotID    string `json:"slotId"`
		StartTime string `json:"startTime"`
		EndTime   string `json:"endTime"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	reservation, err := h.service.MakeReservation(req.UserID, req.SlotID, req.StartTime, req.EndTime)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, reservation)
}

func (h *ReservationHandler) CheckIn(c *gin.Context) {
	session, err := h.service.CheckIn(c.Param("reservationId"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, session)
}

func (h *ReservationHandler) WalkIn(c *gin.Context) {
	var req struct {
		UserID string `json:"userId"`
		SlotID string `json:"slotId"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	session, err := h.service.WalkIn(req.UserID, req.SlotID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, session)
}

func (h *ReservationHandler) CancelReservation(c *gin.Context) {
	if err := h.service.CancelReservation(c.Param("reservationId")); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "reservation cancelled"})
}

func (h *ReservationHandler) MarkNoShow(c *gin.Context) {
	if err := h.service.MarkNoShow(c.Param("reservationId")); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "reservation marked as no-show"})
}

func (h *ReservationHandler) GetReservationByID(c *gin.Context) {
	reservation, err := h.service.GetReservationByID(c.Param("reservationId"))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, reservation)
}

func (h *ReservationHandler) GetReservationsByUserID(c *gin.Context) {
	reservations, err := h.service.GetReservationsByUserID(c.Param("userId"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, reservations)
}
