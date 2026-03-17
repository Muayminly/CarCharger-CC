package model

type Admin struct {
	User
	AdminID string `json:"adminId"`
}
