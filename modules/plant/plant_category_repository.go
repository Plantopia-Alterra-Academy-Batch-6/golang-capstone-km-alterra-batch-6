package plant

import "gorm.io/gorm"

type PlantCategoryRepository interface {
	FindAll() ([]PlantCategory, error)
	FindByID(id int) (PlantCategory, error)
	Create(category PlantCategory) (PlantCategory, error)
	Update(category PlantCategory) (PlantCategory, error)
	Delete(category PlantCategory) error
}

type plantCategoryRepository struct {
	db *gorm.DB
}

func NewPlantCategoryRepository(db *gorm.DB) PlantCategoryRepository {
	return &plantCategoryRepository{db}
}

func (r *plantCategoryRepository) FindAll() ([]PlantCategory, error) {
	var categories []PlantCategory
	err := r.db.Find(&categories).Error
	return categories, err
}

func (r *plantCategoryRepository) FindByID(id int) (PlantCategory, error) {
	var category PlantCategory
	err := r.db.First(&category, id).Error
	return category, err
}

func (r *plantCategoryRepository) Create(category PlantCategory) (PlantCategory, error) {
	err := r.db.Create(&category).Error
	return category, err
}

func (r *plantCategoryRepository) Update(category PlantCategory) (PlantCategory, error) {
	err := r.db.Save(&category).Error
	return category, err
}

func (r *plantCategoryRepository) Delete(category PlantCategory) error {
	err := r.db.Delete(&category).Error
	return err
}
