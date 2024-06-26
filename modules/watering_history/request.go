package wateringhistory

type WateringHistoryRequest struct {
	PlantID int `json:"plant_id" validate:"required"`
}
