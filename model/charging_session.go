package model

type ChargingSession struct {
	SessionID             string         `json:"sessionId"`
	UserID                string         `json:"userId"`
	ReservationID         string         `json:"reservationId"`
	SlotID                string         `json:"slotId"`
	ActualStartTime       string         `json:"actualStartTime"`
	ActualEndTime         string         `json:"actualEndTime"`
	EnergyConsumedKWh     float64        `json:"energyConsumedKWh"`
	CurrentBatteryPercent int            `json:"currentBatteryPercent"`
	Status                ChargingStatus `json:"status"`
}
