package plant

import "gorm.io/gorm"

type PlantRepository interface {
	FindAll(page, limit int) ([]Plant, error)
	FindByID(id int) (Plant, error)
	Create(plant Plant) (Plant, error)
	Update(plant Plant) (Plant, error)
	Delete(id int) error
	FindByIDWithRelations(id int) (Plant, error)
	CountAll(count *int64) error
	SearchPlantsByName(name string, limit, offset int) ([]Plant, error)
	FindEarliestWateringTime(name string, limit, offset int) ([]Plant, error)
	CountPlantsByName(name string) (int64, error)
	ClearPlantInstructions(plantID int) error
	ClearPlantFAQs(plantID int) error
	ClearPlantImages(plantID int) error
	FindRecommendations(userID int, plants *[]Plant) *gorm.DB
	FindByCategoryID(categoryID int, plants *[]Plant) *gorm.DB
	CategoryExists(categoryID int) (bool, error)
}

type plantRepository struct {
	db *gorm.DB
}

func NewPlantRepository(db *gorm.DB) PlantRepository {
	return &plantRepository{db}
}

func (r *plantRepository) FindRecommendations(userID int, plants *[]Plant) *gorm.DB {
	subQuery := r.db.Table("user_plants").Select("plant_id").Where("user_id = ?", userID)
	return r.db.Preload("PlantCategory").
			Preload("PlantCharacteristic").
			Preload("WateringSchedule").
			Preload("PlantInstructions").
			Preload("PlantInstructions.InstructionCategory").
			Preload("PlantFAQs").
			Preload("PlantImages").
			Where("id NOT IN (?)", subQuery).
			Find(plants)
}

func (r *plantRepository) FindByCategoryID(categoryID int, plants *[]Plant) *gorm.DB {
	return r.db.Preload("PlantCategory").
			Preload("PlantCharacteristic").
			Preload("WateringSchedule").
			Preload("PlantInstructions").
			Preload("PlantInstructions.InstructionCategory").
			Preload("PlantFAQs").
			Preload("PlantImages").
			Where("plant_category_id = ?", categoryID).
			Find(plants)
}

func (r *plantRepository) CategoryExists(categoryID int) (bool, error) {
	var count int64
	err := r.db.Model(&PlantCategory{}).Where("id = ?", categoryID).Count(&count).Error
	if err != nil {
			return false, err
	}
	return count > 0, nil
}

func (r *plantRepository) FindAll(page, limit int) ([]Plant, error) {
	var plants []Plant
	var err error
	if page > 0 && limit > 0 {
			offset := (page - 1) * limit
			err = r.db.Preload("PlantCategory").Preload("PlantCharacteristic").Preload("WateringSchedule").
					Preload("PlantInstructions").Preload("PlantInstructions.InstructionCategory").Preload("PlantFAQs").Preload("PlantImages").
					Offset(offset).Limit(limit).Find(&plants).Error
	} else {
			err = r.db.Preload("PlantCategory").Preload("PlantCharacteristic").Preload("WateringSchedule").
					Preload("PlantInstructions").Preload("PlantInstructions.InstructionCategory").Preload("PlantFAQs").Preload("PlantImages").
					Find(&plants).Error
	}
	return plants, err
}

func (r *plantRepository) FindByID(id int) (Plant, error) {
	var plant Plant
	err := r.db.Preload("PlantCategory").Preload("PlantCharacteristic").Preload("WateringSchedule").
		Preload("PlantInstructions").Preload("PlantInstructions.InstructionCategory").Preload("PlantFAQs").Preload("PlantImages").First(&plant, id).Error
	return plant, err
}

func (r *plantRepository) Create(plant Plant) (Plant, error) {
	err := r.db.Create(&plant).Error
	return plant, err
}

func (r *plantRepository) Update(plant Plant) (Plant, error) {
	err := r.db.Save(&plant).Error
	return plant, err
}

func (r *plantRepository) FindByIDWithRelations(id int) (Plant, error) {
	var plant Plant
	err := r.db.Preload("PlantCategory").
			Preload("PlantCharacteristic").
			Preload("WateringSchedule").
			Preload("PlantInstructions").
			Preload("PlantInstructions.InstructionCategory").
			Preload("PlantFAQs").
			Preload("PlantImages").
			Where("id = ?", id).
			First(&plant).Error
	if err != nil {
			return Plant{}, err
	}
	return plant, nil
}

func (r *plantRepository) Delete(id int) error {
	if err := r.db.Delete(&Plant{}, id).Error; err != nil {
			return err
	}
	return nil
}

func (r *plantRepository) CountAll(count *int64) error {
	return r.db.Model(&Plant{}).Count(count).Error
}

func (r *plantRepository) SearchPlantsByName(name string, limit, offset int) ([]Plant, error) {
	var plants []Plant
	err := r.db.Preload("PlantCategory").
		Preload("PlantCharacteristic").
		Preload("WateringSchedule").
		Preload("PlantInstructions").
		Preload("PlantInstructions.InstructionCategory").
		Preload("PlantFAQs").
		Preload("PlantImages").
		Where("name LIKE ?", "%"+name+"%").
		Limit(limit).
		Offset(offset).
		Find(&plants).Error
	return plants, err
}

func (r *plantRepository) FindEarliestWateringTime(name string, limit, offset int) ([]Plant, error) {
	var plants []Plant
	err := r.db.Preload("PlantCategory").
		Preload("PlantCharacteristic").
		Preload("WateringSchedule").
		Preload("PlantInstructions").
		Preload("PlantInstructions.InstructionCategory").
		Preload("PlantFAQs").
		Preload("PlantImages").
		Where("name LIKE ?", "%"+name+"%").
		Limit(limit).
		Offset(offset).
		Find(&plants).Error
	return plants, err
}

func (r *plantRepository) CountPlantsByName(name string) (int64, error) {
	var count int64
	err := r.db.Model(&Plant{}).
		Where("name LIKE ?", "%"+name+"%").
		Count(&count).Error
	return count, err
}

func (r *plantRepository) ClearPlantInstructions(plantID int) error {
	return r.db.Where("plant_id = ?", plantID).Delete(&PlantInstruction{}).Error
}

func (r *plantRepository) ClearPlantFAQs(plantID int) error {
	return r.db.Where("plant_id = ?", plantID).Delete(&PlantFAQ{}).Error
}

func (r *plantRepository) ClearPlantImages(plantID int) error {
	return r.db.Where("plant_id = ?", plantID).Delete(&PlantImage{}).Error
}


