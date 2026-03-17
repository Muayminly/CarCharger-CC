package model

type Reservation struct {
	ReservationID     string            `json:"reservationId"`
	UserID            string            `json:"userId"`
	SlotID            string            `json:"slotId"`
	StartTime         string            `json:"startTime"`
	EndTime           string            `json:"endTime"`
	ReservationStatus ReservationStatus `json:"reservationStatus"`
	MaxLateTime       int               `json:"maxLateTime"`
}
