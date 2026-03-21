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

		api.GET("/stations/search", stationHandler.SearchStations)
		api.GET("/stations/:stationId", stationHandler.GetStationByID)
		api.GET("/stations/:stationId/slots", stationHandler.GetAvailableSlots)
		api.PUT("/stations/:stationId/status", stationHandler.UpdateStationStatus)

		api.POST("/reservations", reservationHandler.MakeReservation)
		api.POST("/reservations/:reservationId/check-in", reservationHandler.CheckIn)
		api.POST("/walk-in", reservationHandler.WalkIn)
		api.DELETE("/reservations/:reservationId", reservationHandler.CancelReservation)
		api.PUT("/reservations/:reservationId/no-show", reservationHandler.MarkNoShow)
		api.GET("/reservations/:reservationId", reservationHandler.GetReservationByID)
		api.GET("/users/:userId/reservations", reservationHandler.GetReservationsByUserID)

		api.POST("/sessions/start", chargingSessionHandler.StartSession)
		api.PUT("/sessions/:sessionId/status", chargingSessionHandler.UpdateChargingStatus)
		api.POST("/sessions/:sessionId/end", chargingSessionHandler.EndSession)
		api.GET("/sessions/:sessionId", chargingSessionHandler.GetSessionByID)
		api.GET("/users/:userId/active-session", chargingSessionHandler.GetActiveSessionByUserID)

		api.GET("/bills/:billId", billingHandler.GetBillByID)
		api.GET("/users/:userId/bills", billingHandler.GetBillsByUserID)
		api.POST("/bills/:sessionId/generate", billingHandler.GenerateBill)
		api.GET("/sessions/:sessionId/cost", billingHandler.CalculateCost)

		api.POST("/payments", paymentHandler.ProcessPayment)
		api.PUT("/payments/:paymentId/verify", paymentHandler.VerifyPayment)
		api.GET("/payments/:paymentId", paymentHandler.GetPaymentByID)
		api.GET("/bills/:billId/payments", paymentHandler.GetPaymentsByBillID)

		api.POST("/notifications", notificationHandler.SendNotification)
		api.GET("/users/:userId/notifications", notificationHandler.GetNotificationsByUserID)

		admin := api.Group("/admin")
		{
			admin.POST("/users", adminHandler.ManageUser)
			admin.POST("/stations", adminHandler.ManageStation)
			admin.POST("/slots", adminHandler.ManageSlot)
			admin.GET("/reservations", adminHandler.MonitorReservations)
			admin.GET("/sessions", adminHandler.MonitorChargingSessions)
			admin.GET("/payments", adminHandler.VerifyPayments)
		}
	}

	log.Println("EV Charging System started on :8080")
	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
