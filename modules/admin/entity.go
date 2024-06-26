package admin

import (
	"time"
)

type Admin struct {
	ID         int `gorm:"primaryKey"`
	Name       string
	Email      string
	Password   string
	Url_Image  string
	Created_at time.Time
	Updated_at time.Time
}
