package search

type PlantSearchParams struct {
	Name            string `json:"name" query:"name"`
	PlantCategory   string `json:"plant_category" query:"plant_category"`
	DifficultyLevel string `json:"difficulty_level" query:"difficulty_level"`
	Sunlight        string `json:"sunlight" query:"sunlight"`
	HarvestDuration string `json:"harvest_duration" query:"harvest_duration"`
	IsToxic         *bool  `json:"is_toxic" query:"is_toxic"`
}
