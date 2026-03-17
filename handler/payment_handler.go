package handler

import (
	"net/http"

	"CarCharger-CC/model"
	"CarCharger-CC/service"

	"github.com/gin-gonic/gin"
)

type PaymentHandler struct {
	service service.PaymentService
}

func NewPaymentHandler(service service.PaymentService) *PaymentHandler {
	return &PaymentHandler{service: service}
}

func (h *PaymentHandler) ProcessPayment(c *gin.Context) {
	var req struct {
		BillID string              `json:"billId"`
		Method model.PaymentMethod `json:"paymentMethod"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	payment, err := h.service.ProcessPayment(req.BillID, req.Method)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, payment)
}

func (h *PaymentHandler) VerifyPayment(c *gin.Context) {
	if err := h.service.VerifyPayment(c.Param("paymentId")); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "payment verified"})
}

func (h *PaymentHandler) GetPaymentByID(c *gin.Context) {
	payment, err := h.service.GetPaymentByID(c.Param("paymentId"))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, payment)
}

func (h *PaymentHandler) GetPaymentsByBillID(c *gin.Context) {
	payments, err := h.service.GetPaymentsByBillID(c.Param("billId"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, payments)
}
