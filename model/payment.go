package model

type Payment struct {
	PaymentID     string        `json:"paymentId"`
	BillID        string        `json:"billId"`
	PaymentMethod PaymentMethod `json:"paymentMethod"`
	PaymentStatus PaymentStatus `json:"paymentStatus"`
	PaymentDate   string        `json:"paymentDate"`
}
