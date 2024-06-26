package fertilizer

import (
	"time"
)

type FertilizerResponse struct {
	Id int `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
	Compostition string `json:"compostition"`
	CreateAt time.Time `json:"createAt"`
}

func NewFertilizerResponse(fertilizer Fertilizer) FertilizerResponse {
	return FertilizerResponse{
		Id:       fertilizer.Id,
		Name:     fertilizer.Name,
		Compostition: fertilizer.Compostition,
		CreateAt: fertilizer.CreateAt,
	}
}