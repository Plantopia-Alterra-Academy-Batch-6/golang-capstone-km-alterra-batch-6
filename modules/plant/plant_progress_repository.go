package plant

import "gorm.io/gorm"

type PlantProgressRepository interface {
	FindByUserIDAndPlantID(userID int, plantID int) ([]PlantProgress, error)
	FindByID(id int) (PlantProgress, error)
	Create(progress PlantProgress) (PlantProgress, error)
	Delete(progress PlantProgress) error
}

type plantProgressRepository struct {
	db *gorm.DB
}

func NewPlantProgressRepository(db *gorm.DB) PlantProgressRepository {
	return &plantProgressRepository{db}
}

func (r *plantProgressRepository) FindByUserIDAndPlantID(userID int, plantID int) ([]PlantProgress, error) {
	var progresses []PlantProgress
	err := r.db.Where("user_id = ? AND plant_id = ?", userID, plantID).Find(&progresses).Error
	return progresses, err
}

func (r *plantProgressRepository) FindByID(id int) (PlantProgress, error) {
	var progress PlantProgress
	err := r.db.First(&progress, id).Error
	return progress, err
}

func (r *plantProgressRepository) Create(progress PlantProgress) (PlantProgress, error) {
	err := r.db.Create(&progress).Error
	return progress, err
}

func (r *plantProgressRepository) Delete(progress PlantProgress) error {
	err := r.db.Delete(&progress).Error
	return err
}
