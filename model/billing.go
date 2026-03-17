package model

type Billing struct {
	BillID        string  `json:"billId"`
	SessionID     string  `json:"sessionId"`
	EnergyUsedKWh float64 `json:"energyUsedKWh"`
	TotalAmount   float64 `json:"totalAmount"`
	BillingDate   string  `json:"billingDate"`
}
