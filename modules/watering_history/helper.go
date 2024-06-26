package wateringhistory

import "github.com/OctavianoRyan25/be-agriculture/modules/plant"

func MapPlantToPlantResponse(plant *plant.Plant) *PlantResponse {
	return &PlantResponse{
		ID:               plant.ID,
		Name:             plant.Name,
		Description:      plant.Description,
		IsToxic:          plant.IsToxic,
		HarvestDuration:  plant.HarvestDuration,
		Sunlight:         plant.Sunlight,
		PlantingTime:     plant.PlantingTime,
		ClimateCondition: plant.ClimateCondition,
		PlantImage:       MapPlantImagesToPlantImageResponses(plant.PlantImages),
		CreatedAt:        plant.CreatedAt,
	}
}

func MapPlantImagesToPlantImageResponses(images []plant.PlantImage) []PlantImageResponse {
	var plantImageResponses []PlantImageResponse

	for _, image := range images {
		if image.IsPrimary == 1 {
			plantImageResponses = append(plantImageResponses, PlantImageResponse{
				ID:       image.ID,
				FileName: image.FileName,
			})
		}
	}

	return plantImageResponses
}
