package handler

import (
	"net/http"

	"CarCharger-CC/service"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	service service.UserService
}

func NewUserHandler(service service.UserService) *UserHandler {
	return &UserHandler{service: service}
}

//viewpayment

func (h *UserHandler) ViewPaymentHistory(c *gin.Context) {
	payments, err := h.service.ViewPaymentHistory(c.Param("userId"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, payments)
}
