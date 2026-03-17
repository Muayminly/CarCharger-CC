package handler

import (
	"net/http"

	"CarCharger-CC/service"

	"github.com/gin-gonic/gin"
)

type BillingHandler struct {
	service service.BillingService
}

func NewBillingHandler(service service.BillingService) *BillingHandler {
	return &BillingHandler{service: service}
}

func (h *BillingHandler) CalculateCost(c *gin.Context) {
	cost, err := h.service.CalculateCost(c.Param("sessionId"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"totalAmount": cost})
}

func (h *BillingHandler) GenerateBill(c *gin.Context) {
	bill, err := h.service.GenerateBill(c.Param("sessionId"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, bill)
}

func (h *BillingHandler) GetBillByID(c *gin.Context) {
	bill, err := h.service.GetBillByID(c.Param("billId"))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, bill)
}

func (h *BillingHandler) GetBillsByUserID(c *gin.Context) {
	bills, err := h.service.GetBillsByUserID(c.Param("userId"))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, bills)
}
