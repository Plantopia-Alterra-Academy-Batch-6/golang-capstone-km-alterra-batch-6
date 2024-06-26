package user

import (
	"time"
)

type UserResponse struct {
	ID         int       `json:"id"`
	Name       string    `json:"name"`
	Email      string    `json:"email"`
	Is_Active  bool      `json:"is_active"`
	OTP        string    `json:"otp"`
	Url_Image  string    `json:"url_image"`
	FCMToken   string    `json:"fcm_token"`
	Created_at time.Time `json:"created_at"`
}

type LoginResponse struct {
	Token string `json:"token"`
}
