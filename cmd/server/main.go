package main

import (
	"log"

	"CarCharger-CC/handler"
	"CarCharger-CC/repository"
	"CarCharger-CC/service"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// repositories
	userRepo := repository.NewInMemoryUserRepository()
	stationRepo := repository.NewInMemoryStationRepository()
	slotRepo := repository.NewInMemoryChargingSlotRepository()
	reservationRepo := repository.NewInMemoryReservationRepository()
	sessionRepo := repository.NewInMemoryChargingSessionRepository()
	billingRepo := repository.NewInMemoryBillingRepository()
	paymentRepo := repository.NewInMemoryPaymentRepository()
	notificationRepo := repository.NewInMemoryNotificationRepository()

	// services
	authService := service.NewAuthService(userRepo)
	userService := service.NewUserService(userRepo, sessionRepo, paymentRepo)
	stationService := service.NewStationService(stationRepo, slotRepo)
	reservationService := service.NewReservationService(reservationRepo, slotRepo, sessionRepo)
	chargingSessionService := service.NewChargingSessionService(sessionRepo, billingRepo)
	billingService := service.NewBillingService(billingRepo, sessionRepo)
	paymentService := service.NewPaymentService(paymentRepo, billingRepo)
	notificationService := service.NewNotificationService(notificationRepo)
	adminService := service.NewAdminService(userRepo, stationRepo, slotRepo, reservationRepo, sessionRepo, paymentRepo)

	// handlers
	authHandler := handler.NewAuthHandler(authService)
	userHandler := handler.NewUserHandler(userService)
	stationHandler := handler.NewStationHandler(stationService)
	reservationHandler := handler.NewReservationHandler(reservationService)
	chargingSessionHandler := handler.NewChargingSessionHandler(chargingSessionService)
	billingHandler := handler.NewBillingHandler(billingService)
	paymentHandler := handler.NewPaymentHandler(paymentService)
	notificationHandler := handler.NewNotificationHandler(notificationService)
	adminHandler := handler.NewAdminHandler(adminService)

	api := router.Group("/api")
	{
		auth := api.Group("/auth")
		{
			auth.POST("/register", authHandler.Register)
			auth.POST("/login", authHandler.Login)
			auth.POST("/logout", authHandler.Logout)
			auth.PUT("/profile", authHandler.UpdateProfile)
		}

		api.GET("/users/:userId", userHandler.GetUserByID)
		api.GET("/drivers/:userId", userHandler.GetDriverByID)
		api.GET("/users/:userId/payments", userHandler.ViewPaymentHistory)
		api.GET("/users/:userId/sessions", userHandler.ViewUsageHistory)
		
	}

	log.Println("EV Charging System started on :8080")
	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
