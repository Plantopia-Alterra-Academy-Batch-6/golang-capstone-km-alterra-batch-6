package plant

type PlantCategoryClimateInput struct {
	Name     string `form:"name" validate:"required"`
	ImageURL string `form:"image_url"`
}

type PlantInstructionCategoryInput struct {
	Name        string `form:"name" validate:"required"`
	Description string `form:"description" validate:"required"`
	ImageURL    string `form:"image_url"`
}

type CreatePlantInput struct {
	Name                string                         `form:"name" validate:"required"`
	Description         string                         `form:"description" validate:"required"`
	IsToxic             bool                           `form:"is_toxic"`
	HarvestDuration     int                            `form:"harvest_duration" validate:"required"`
	Sunlight            string                         `form:"sunlight" validate:"required"`
	PlantingTime        string                         `form:"planting_time" validate:"required"`
	PlantCategoryID     int                            `form:"plant_category_id" validate:"required"`
	ClimateCondition    string                         `form:"climate_condition" validate:"required"`
	PlantCharacteristic CreatePlantCharacteristicInput `form:"plant_characteristic" validate:"required"`
	WateringSchedule    CreateWateringScheduleInput    `form:"watering_schedule"`
	PlantInstructions   []CreatePlantInstructionInput  `form:"plant_instructions"`
	AdditionalTips      string                         `form:"additional_tips"`
	PlantFAQs           []CreatePlantFAQInput          `form:"plant_faqs"`
	PlantImages         []CreatePlantImageInput        `form:"plant_images" validate:"required,dive"`
}

type UpdatePlantInput struct {
	Name                string                         `form:"name" validate:"required"`
	Description         string                         `form:"description" validate:"required"`
	IsToxic             bool                           `form:"is_toxic"`
	HarvestDuration     int                            `form:"harvest_duration" validate:"required"`
	Sunlight            string                         `form:"sunlight" validate:"required"`
	PlantingTime        string                         `form:"planting_time" validate:"required"`
	PlantCategoryID     int                            `form:"plant_category_id" validate:"required"`
	ClimateCondition    string                         `form:"climate_condition" validate:"required"`
	WateringSchedule    CreateWateringScheduleInput    `form:"watering_schedule"`
	PlantCharacteristic CreatePlantCharacteristicInput `form:"plant_characteristic" validate:"required"`
	PlantInstructions   []CreatePlantInstructionInput  `form:"plant_instructions"`
	AdditionalTips      string                         `form:"additional_tips"`
	PlantFAQs           []CreatePlantFAQInput          `form:"plant_faqs"`
	PlantImages         []CreatePlantImageInput        `form:"plant_images" validate:"required,dive"`
}

type UpdateInstructionCategoryInput struct {
	UserPlantID         int `json:"user_plant_id" validate:"required"`
	InstructionCategory int `json:"instruction_category" validate:"required,min=1,max=4"`
}

type CreateWateringScheduleInput struct {
	WateringFrequency    int    `form:"watering_frequency" validate:"required"`
	Each                 string `form:"each" validate:"required"`
	WateringAmount       int    `form:"watering_amount" validate:"required"`
	Unit                 string `form:"unit" validate:"required"`
	WateringTime         string `form:"watering_time" validate:"required"`
	WeatherCondition     string `form:"weather_condition"`
	ConditionDescription string `form:"condition_description"`
}

type CreatePlantCharacteristicInput struct {
	Height     int    `form:"height" validate:"required"`
	HeightUnit string `form:"height_unit" validate:"required"`
	Wide       int    `form:"wide" validate:"required"`
	WideUnit   string `form:"wide_unit" validate:"required"`
	LeafColor  string `form:"leaf_color" validate:"required"`
}

type CreatePlantInstructionInput struct {
	InstructionCategoryID int    `form:"instruction_category_id" validate:"required"`
	StepNumber            int    `form:"step_number" validate:"required"`
	StepTitle             string `form:"step_title" validate:"required"`
	StepDescription       string `form:"step_description" validate:"required"`
	StepImageURL          string `form:"step_image_url"`
}

type CreatePlantFAQInput struct {
	Question string `form:"question" validate:"required"`
	Answer   string `form:"answer" validate:"required"`
}

type CreatePlantImageInput struct {
	FileName  string `form:"file_name" validate:"required"`
	IsPrimary int    `form:"is_primary"`
}

type AddUserPlantInput struct {
	UserID        int    `json:"user_id" form:"user_id"`
	PlantID       int    `json:"plant_id" form:"plant_id" validate:"required"`
	CustomizeName string `json:"customize_name" form:"customize_name"`
}

type PlantProgressInput struct {
	PlantID  int    `form:"plant_id" validate:"required"`
	UserID   int    `json:"user_id" form:"user_id"`
	ImageURL string `form:"image_url"`
}

type UserPlantHistoryInput struct {
	UserID        int    `json:"user_id" form:"user_id"`
	PlantID       int    `json:"plant_id" form:"plant_id" validate:"required"`
	PlantName     string `json:"plant_name" form:"plant_name"`
	PlantCategory string `json:"plant_category" form:"plant_category"`
	ImageURL      string `json:"image_url" form:"image_url"`
}