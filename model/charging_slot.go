package model

type ChargingSlot struct {
	SlotID     string     `json:"slotId"`
	StationID  string     `json:"stationId"`
	SlotType   SlotType   `json:"slotType"`
	SlotStatus SlotStatus `json:"slotStatus"`
}
