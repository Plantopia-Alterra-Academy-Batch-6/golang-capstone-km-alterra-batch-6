package user

import (
	"time"
)

type User struct {
	ID         int `gorm:"primaryKey"`
	Name       string
	Email      string
	Password   string
	Is_Active  bool
	OTP        string
	FCMToken   string
	Url_Image  string
	Created_at time.Time
	Updated_at time.Time
}
