// repository.go
package plant

import "gorm.io/gorm"

type PlantEarliestWateringRepository interface {
	GetEarliestWatering() ([]PlantReminder, error)
}

type plantEarliestWateringRepository struct {
	db *gorm.DB
}

func NewPlantEarliestWateringRepository(db *gorm.DB) PlantEarliestWateringRepository {
	return &plantEarliestWateringRepository{db}
}

// GetEarliestWatering implements PlantEarliestWateringRepository.
func (r *plantEarliestWateringRepository) GetEarliestWatering() ([]PlantReminder, error) {
	var schedule []PlantReminder
	err := r.db.Find(&schedule).Error
	return schedule, err
}