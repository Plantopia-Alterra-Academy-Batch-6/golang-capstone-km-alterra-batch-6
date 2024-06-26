package search

import (
	"errors"

	"github.com/OctavianoRyan25/be-agriculture/modules/plant"
	"gorm.io/gorm"
)

type Repository interface {
	Search(PlantSearchParams) ([]plant.Plant, error)
}

type searchRepo struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *searchRepo {
	return &searchRepo{
		db: db,
	}
}

func (r *searchRepo) Search(params PlantSearchParams) ([]plant.Plant, error) {
	var plants []plant.Plant
	query := r.db.Preload("PlantCategory").Preload("PlantCharacteristic").Preload("WateringSchedule").
		Preload("PlantInstructions").
		Preload("PlantFAQs").
		Preload("PlantImages")

	if params.Name != "" {
		query = query.Where("name LIKE ?", "%"+params.Name+"%")
	}
	if params.PlantCategory != "" {
		query = query.Where("plant_category_id = ?", params.PlantCategory)
	}
	if params.Sunlight != "" {
		query = query.Where("sunlight = ?", params.Sunlight)
	}
	if params.HarvestDuration != "" {
		switch params.HarvestDuration {
		case "less than 1 month":
			query = query.Where("harvest_duration < ?", 1)
		case "1-3 months":
			query = query.Where("harvest_duration >= ? AND harvest_duration <= ?", 1, 3)
		case "3-6 months":
			query = query.Where("harvest_duration >= ? AND harvest_duration <= ?", 3, 6)
		case "more than 6 months":
			query = query.Where("harvest_duration > ?", 6)
		default:
			return nil, errors.New("invalid harvest duration")
		}
	}
	if params.IsToxic != nil {
		query = query.Where("is_toxic = ?", params.IsToxic)
	}

	err := query.Find(&plants).Error
	if err != nil {
		return nil, err
	}
	return plants, nil
}
