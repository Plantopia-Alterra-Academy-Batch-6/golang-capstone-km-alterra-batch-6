package plant

import (
	"gorm.io/gorm"
)

type UserPlantRepository interface {
	AddUserPlant(userPlant UserPlant) (UserPlant, error)
	GetUserPlantsByUserID(userID int, limit int, offset int) ([]UserPlant, error)
	DeleteUserPlantByID(userPlantID int) error
	GetUserPlantByID(userPlantID int) (UserPlant, error)
	CountByUserID(userID int, count *int64) error

	Create(userPlantHistory UserPlantHistory) (UserPlantHistory, error)
	GetPlantByID(plantID int) (Plant, error)
	GetPlantCategoryByID(categoryID int) (PlantCategory, error)
	GetPrimaryPlantImageByPlantID(plantID int) (PlantImage, error)
	GetUserPlantHistoryByUserID(userID int) ([]UserPlantHistory, error)
	CheckPlantExists(plantID int) (bool, error)
	UpdateCustomizeName(userPlantID int, customizeName string) error
	CheckUserPlantExists(userPlantID int) (bool, error)
	CheckUserPlantExistsForAdd(userID, plantID int) (bool, error)
	UpdateInstructionCategory(userPlantID int, fieldToUpdate string) error
	GetUserPlantByUserIDAndPlantID(userID int, plantID int) (UserPlant, error)
}

type userPlantRepository struct {
	db *gorm.DB
}

func NewUserPlantRepository(db *gorm.DB) UserPlantRepository {
	return &userPlantRepository{db}
}

func (r *userPlantRepository) GetUserPlantByUserIDAndPlantID(userID int, plantID int) (UserPlant, error) {
	var userPlant UserPlant
	err := r.db.Preload("Plant").
	Preload("Plant.PlantCategory").
	Preload("Plant.PlantCharacteristic").
	Preload("Plant.WateringSchedule").
	Preload("Plant.PlantInstructions").
	Preload("Plant.PlantInstructions.InstructionCategory").
	Preload("Plant.PlantFAQs").
	Preload("Plant.PlantImages").Where("user_id = ? AND plant_id = ?", userID, plantID).First(&userPlant).Error
	if err != nil {
		return userPlant, err
	}
	return userPlant, nil
}

func (r *userPlantRepository) UpdateInstructionCategory(userPlantID int, fieldToUpdate string) error {
	updateQuery := map[string]interface{}{fieldToUpdate: 1}

	err := r.db.Model(&UserPlant{}).
		Where("id = ?", userPlantID).
		Updates(updateQuery).Error

	return err
}

func (r *userPlantRepository) CheckUserPlantExists(userPlantID int) (bool, error) {
	var count int64
	if err := r.db.Model(&UserPlant{}).Where("id = ?", userPlantID).Count(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}

func (r *userPlantRepository) UpdateCustomizeName(userPlantID int, customizeName string) error {
	if err := r.db.Model(&UserPlant{}).Where("id = ?", userPlantID).Update("customize_name", customizeName).Error; err != nil {
		return err
	}
	return nil
}

func (r *userPlantRepository) CheckPlantExists(plantID int) (bool, error) {
	var count int64
	if err := r.db.Model(&Plant{}).Where("id = ?", plantID).Count(&count).Error; err != nil {
			return false, err
	}
	return count > 0, nil
}

func (r *userPlantRepository) GetUserPlantHistoryByUserID(userID int) ([]UserPlantHistory, error) {
	var histories []UserPlantHistory
	if err := r.db.Where("user_id = ?", userID).Find(&histories).Error; err != nil {
		return nil, err
	}
	return histories, nil
}

func (r *userPlantRepository) GetPlantCategoryByID(categoryID int) (PlantCategory, error) {
	var category PlantCategory
	if err := r.db.Where("id = ?", categoryID).First(&category).Error; err != nil {
		return category, err
	}
	return category, nil
}

func (r *userPlantRepository) GetPrimaryPlantImageByPlantID(plantID int) (PlantImage, error) {
	var image PlantImage
	err := r.db.Where("plant_id = ? AND is_primary = 1", plantID).First(&image).Error
	if err != nil {
		err = r.db.Where("plant_id = ?", plantID).First(&image).Error
		if err != nil {
			image.FileName = "https://example.com/images/placeholder.png"
			return image, nil
		}
	}
	return image, nil
}

func (r *userPlantRepository) CheckUserPlantExistsForAdd(userID, plantID int) (bool, error) {
	var count int64
	result := r.db.Model(&UserPlant{}).Where("user_id = ? AND plant_id = ?", userID, plantID).Count(&count)
	if result.Error != nil {
			return false, result.Error
	}
	return count > 0, nil
}

func (r *userPlantRepository) Create(userPlantHistory UserPlantHistory) (UserPlantHistory, error) {
	if err := r.db.Create(&userPlantHistory).Error; err != nil {
		return userPlantHistory, err
	}
	return userPlantHistory, nil
}

func (r *userPlantRepository) GetPlantByID(plantID int) (Plant, error) {
	var plant Plant
	if err := r.db.Where("id = ?", plantID).First(&plant).Error; err != nil {
		return plant, err
	}
	return plant, nil
}

func (r *userPlantRepository) AddUserPlant(userPlant UserPlant) (UserPlant, error) {
	err := r.db.Create(&userPlant).Error
	if err != nil {
		return userPlant, err
	}

	err = r.db.Preload("Plant").
		Preload("Plant.PlantCategory").
		Preload("Plant.PlantCharacteristic").
		Preload("Plant.WateringSchedule").
		Preload("Plant.PlantInstructions").
		Preload("Plant.PlantFAQs").
		Preload("Plant.PlantImages").
		First(&userPlant, userPlant.ID).Error

	return userPlant, err
}

func (r *userPlantRepository) GetUserPlantsByUserID(userID int, limit int, offset int) ([]UserPlant, error) {
	var userPlants []UserPlant
	query := r.db.Preload("Plant").
		Preload("Plant.PlantCategory").
		Preload("Plant.PlantCharacteristic").
		Preload("Plant.WateringSchedule").
		Preload("Plant.PlantInstructions").
		Preload("Plant.PlantInstructions.InstructionCategory").
		Preload("Plant.PlantFAQs").
		Preload("Plant.PlantImages").
		Where("user_id = ?", userID)

	if limit > 0 {
		query = query.Limit(limit)
	}

	if offset >= 0 {
		query = query.Offset(offset)
	}

	err := query.Find(&userPlants).Error
	return userPlants, err
}

func (r *userPlantRepository) DeleteUserPlantByID(userPlantID int) error {
	return r.db.Where("id = ?", userPlantID).Delete(&UserPlant{}).Error
}

func (r *userPlantRepository) GetUserPlantByID(userPlantID int) (UserPlant, error) {
	var userPlant UserPlant
	err := r.db.Preload("Plant").
			Preload("Plant.PlantCategory").
			Preload("Plant.PlantCharacteristic").
			Preload("Plant.WateringSchedule").
			Preload("Plant.PlantInstructions").
			Preload("Plant.PlantFAQs").
			Preload("Plant.PlantImages").
			Where("id = ?", userPlantID).
			First(&userPlant).Error
	if err != nil {
			return UserPlant{}, err
	}
	return userPlant, nil
}

func (r *userPlantRepository) CountByUserID(userID int, count *int64) error {
	return r.db.Model(&UserPlant{}).Where("user_id = ?", userID).Count(count).Error
}
