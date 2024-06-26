package plant

import "gorm.io/gorm"

type PlantInstructionCategoryRepository interface {
	FindAll() ([]PlantInstructionCategory, error)
	FindByID(id int) (PlantInstructionCategory, error)
	FindInstructionByCategoryID(plantID int, instructionCategoryID int) ([]PlantInstruction, error)
	Create(category PlantInstructionCategory) (PlantInstructionCategory, error)
	Update(category PlantInstructionCategory) (PlantInstructionCategory, error)
	Delete(category PlantInstructionCategory) error
}

type plantInstructionCategoryRepository struct {
	db *gorm.DB
}

func NewPlantInstructionCategoryRepository(db *gorm.DB) PlantInstructionCategoryRepository {
	return &plantInstructionCategoryRepository{db}
}

func (r *plantInstructionCategoryRepository) FindAll() ([]PlantInstructionCategory, error) {
	var categories []PlantInstructionCategory
	err := r.db.Find(&categories).Error
	return categories, err
}

func (r *plantInstructionCategoryRepository) FindByID(id int) (PlantInstructionCategory, error) {
	var category PlantInstructionCategory
	err := r.db.First(&category, id).Error
	return category, err
}

func (r *plantInstructionCategoryRepository) FindInstructionByCategoryID(plantID int, instructionCategoryID int) ([]PlantInstruction, error) {
	var instructions []PlantInstruction
	err := r.db.Preload("InstructionCategory").Where("plant_id = ? AND instruction_category_id = ?", plantID ,instructionCategoryID).Find(&instructions).Error
	return instructions, err
}

func (r *plantInstructionCategoryRepository) Create(category PlantInstructionCategory) (PlantInstructionCategory, error) {
	err := r.db.Create(&category).Error
	return category, err
}

func (r *plantInstructionCategoryRepository) Update(category PlantInstructionCategory) (PlantInstructionCategory, error) {
	err := r.db.Save(&category).Error
	return category, err
}

func (r *plantInstructionCategoryRepository) Delete(category PlantInstructionCategory) error {
	err := r.db.Delete(&category).Error
	return err
}
