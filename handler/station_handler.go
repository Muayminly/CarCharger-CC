package handler

import (
	"net/http"

	"CarCharger-CC/model"
	"CarCharger-CC/service"

	"github.com/gin-gonic/gin"
)

type StationHandler struct {
	service service.StationService
}

func NewStationHandler(service service.StationService) *StationHandler {
	return &StationHandler{service: service}
}

func (h *StationHandler) SearchStations(c *gin.Context) {
	stations, err := h.service.SearchStations(c.Query("location"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, stations)
}

func (h *StationHandler) GetStationByID(c *gin.Context) {
	station, err := h.service.GetStationByID(c.Param("stationId"))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, station)
}

func (h *StationHandler) GetAvailableSlots(c *gin.Context) {
	slots, err := h.service.GetAvailableSlots(c.Param("stationId"), c.Query("dateTime"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, slots)
}

func (h *StationHandler) UpdateStationStatus(c *gin.Context) {
	var req struct {
		Status model.StationStatus `json:"status"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.service.UpdateStationStatus(c.Param("stationId"), req.Status); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "station status updated"})
}
