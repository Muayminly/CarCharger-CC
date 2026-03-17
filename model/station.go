package model

type Station struct {
	StationID     string        `json:"stationId"`
	StationName   string        `json:"stationName"`
	Location      string        `json:"location"`
	TotalSlots    int           `json:"totalSlots"`
	StationStatus StationStatus `json:"stationStatus"`
}
