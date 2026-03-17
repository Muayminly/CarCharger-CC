package model

type User struct {
	UserID      string   `json:"userId"`
	Email       string   `json:"email"`
	PhoneNumber string   `json:"phoneNumber"`
	Password    string   `json:"password"`
	Role        UserRole `json:"role"`
}
