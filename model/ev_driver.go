package model

type EVDriver struct {
	User
	DriverName   string `json:"driverName"`
	VehiclePlate string `json:"vehiclePlate"`
}
