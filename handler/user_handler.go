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

func (h *UserHandler) GetUserByID(c *gin.Context) {
	user, err := h.service.GetUserByID(c.Param("userId"))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}

func (h *UserHandler) GetDriverByID(c *gin.Context) {
	driver, err := h.service.GetDriverByID(c.Param("userId"))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, driver)
}

func (h *UserHandler) ViewPaymentHistory(c *gin.Context) {
	payments, err := h.service.ViewPaymentHistory(c.Param("userId"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, payments)
}

func (h *UserHandler) ViewUsageHistory(c *gin.Context) {
	sessions, err := h.service.ViewUsageHistory(c.Param("userId"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, sessions)
}
