package notification

type CustomizeWateringReminderRequest struct {
	PlantID   int    `json:"plant_id" validate:"required"`
	Time      string `json:"time" validate:"required"`
	Recurring bool   `json:"recurring" validate:"required"`
	Type      string `json:"type" validate:"required"`
}
