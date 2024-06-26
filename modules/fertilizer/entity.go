package fertilizer

import (
	"time"

	"github.com/OctavianoRyan25/be-agriculture/modules/plant"
)

type Fertilizer struct {
	Id           int         `json:"id" gorm:"primaryKey"`
	Name         string      `json:"name"`
	Compostition string      `json:"compostition"`
	CreateAt     time.Time   `json:"createAt"`
	PlantID int `json:"plant_id"`
	Plant        plant.Plant `gorm:"foreignKey:PlantID;references:ID"`
}
