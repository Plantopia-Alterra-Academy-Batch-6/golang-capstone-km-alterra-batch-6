package admin

import (
	"time"
)

type AdminResponse struct {
	ID         int       `json:"id"`
	Name       string    `json:"name"`
	Email      string    `json:"email"`
	Url_Image  string    `json:"url_image"`
	Created_at time.Time `json:"created_at"`
}

type LoginResponse struct {
	Token string `json:"token"`
}
