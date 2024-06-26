package user

import (
	"time"
)

type UserRequest struct {
	Name       string    `json:"name" validate:"required"`
	Email      string    `json:"email" validate:"required,email"`
	Password   string    `json:"password" validate:"required"`
	Is_Active  bool      `json:"is_active"`
	Url_Image  string    `json:"url_image"`
	FCMToken   string    `json:"fcm_token" validate:"required"`
	Created_at time.Time `json:"created_at"`
	Updated_at time.Time `json:"updated_at"`
}

type CheckEmailRequest struct {
	Email string `json:"email" validate:"required,email"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
	FCMToken string `json:"fcm_token" validate:"required"`
}

type OTPRequest struct {
	Email string `json:"email" validate:"required,email"`
	OTP   string `json:"otp" validate:"required"`
}

type ResetPasswordRequest struct {
	Email       string `json:"email" validate:"required,email"`
	NewPassword string `json:"new_password" validate:"required"`
}
