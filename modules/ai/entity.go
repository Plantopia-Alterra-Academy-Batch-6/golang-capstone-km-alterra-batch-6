package ai

type FertilizerRecommendation struct {
	PlantName      string `json:"plant_name"`
	Recommendation string `json:"recommendation"`
}

type PlantingRecommendation struct {
	PlantName      string `json:"plant_name"`
	Recommendation string `json:"recommendation"`
}

type PlantNameForRecommendation struct {
	Name string `json:"name"`
}
