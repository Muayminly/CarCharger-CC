package model

type UserRole string
type StationStatus string
type SlotType string
type SlotStatus string
type ReservationStatus string
type ChargingStatus string
type PaymentMethod string
type PaymentStatus string
type NotificationType string

const (
	RoleEVDriver UserRole = "EV_DRIVER"
	RoleAdmin    UserRole = "ADMIN"
)

const (
	StationAvailable    StationStatus = "AVAILABLE"
	StationOutOfService StationStatus = "OUT_OF_SERVICE"
)

const (
	SlotAC SlotType = "AC"
	SlotDC SlotType = "DC"
)

const (
	SlotAvailable    SlotStatus = "AVAILABLE"
	SlotReserved     SlotStatus = "RESERVED"
	SlotCharging     SlotStatus = "CHARGING"
	SlotOutOfService SlotStatus = "OUT_OF_SERVICE"
)

const (
	ReservationReserved  ReservationStatus = "RESERVED"
	ReservationCheckedIn ReservationStatus = "CHECKED_IN"
	ReservationCancelled ReservationStatus = "CANCELLED"
	ReservationNoShow    ReservationStatus = "NO_SHOW"
)

const (
	ChargingOngoing   ChargingStatus = "CHARGING"
	ChargingCompleted ChargingStatus = "COMPLETED"
)

const (
	MethodQR   PaymentMethod = "QR"
	MethodCard PaymentMethod = "CARD"
)

const (
	PaymentPending PaymentStatus = "PENDING"
	PaymentSuccess PaymentStatus = "SUCCESS"
	PaymentFailed  PaymentStatus = "FAILED"
)

const (
	NotiReservationReminder NotificationType = "RESERVATION_REMINDER"
	NotiChargingCompleted   NotificationType = "CHARGING_COMPLETED"
	NotiReceipt             NotificationType = "RECEIPT"
	NotiGeneral             NotificationType = "GENERAL"
)
