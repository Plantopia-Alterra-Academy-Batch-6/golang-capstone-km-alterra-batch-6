package notification

import (
	"time"

	"github.com/OctavianoRyan25/be-agriculture/modules/user"
)

type NotificationResponse struct {
	Id        int       `json:"id"`
	Title     string    `json:"title"`
	Body      string    `json:"body"`
	UserId    int       `json:"user_id"`
	PlantId   int       `json:"plant_id"`
	IsRead    bool      `json:"is_read"`
	CreatedAt time.Time `json:"created_at"`
}

type CustomizeWateringReminderResponse struct {
	Id        int           `json:"id"`
	PlantID   int           `json:"plant_id"`
	Plant     PlantResponse `json:"plant"`
	UserID    int           `json:"user_id"`
	User      user.User     `json:"user"`
	Time      string        `json:"time"`
	Recurring bool          `json:"recurring"`
	Type      string        `json:"type"`
	CreatedAt time.Time     `json:"created_at"`
}

type PlantResponse struct {
	ID               int                  `json:"id" gorm:"primaryKey"`
	Name             string               `json:"name"`
	Description      string               `json:"description"`
	IsToxic          bool                 `json:"is_toxic"`
	HarvestDuration  int                  `json:"harvest_duration"`
	Sunlight         string               `json:"sunlight"`
	PlantingTime     string               `json:"planting_time"`
	ClimateCondition string               `json:"climate_condition"`
	PlantImage       []PlantImageResponse `json:"plant_image"`
	CreatedAt        time.Time            `json:"created_at"`
}

type PlantImageResponse struct {
	ID       int    `json:"id"`
	FileName string `json:"file_name"`
}
