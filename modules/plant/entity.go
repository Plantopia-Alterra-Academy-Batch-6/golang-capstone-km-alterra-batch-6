package plant

import (
	"time"

	"github.com/OctavianoRyan25/be-agriculture/modules/user"
)

type Plant struct {
	ID                    int                 `json:"id" gorm:"primaryKey"`
	Name                  string              `json:"name"`
	Description           string              `json:"description"`
	IsToxic               bool                `json:"is_toxic"`
	HarvestDuration       int                 `json:"harvest_duration"`
	Sunlight              string              `json:"sunlight"`
	PlantingTime          string              `json:"planting_time"`
	PlantCategoryID       int                 `json:"plant_category_id"`
	PlantCategory         PlantCategory       `json:"plant_category"`
	ClimateCondition      string              `json:"climate_condition"`
	PlantCharacteristicID int                 `json:"plant_characteristic_id"`
	PlantCharacteristic   PlantCharacteristic `json:"plant_characteristic" gorm:"foreignKey:PlantID;constraint:OnDelete:CASCADE"`
	WateringSchedule      PlantReminder       `json:"watering_schedule" gorm:"foreignKey:PlantID;constraint:OnDelete:CASCADE"`
	PlantInstructions     []PlantInstruction  `json:"plant_instructions" gorm:"foreignKey:PlantID;constraint:OnDelete:CASCADE"`
	AdditionalTips        string              `json:"additional_tips"`
	PlantFAQs             []PlantFAQ          `json:"plant_faqs" gorm:"foreignKey:PlantID;constraint:OnDelete:CASCADE"`
	PlantImages           []PlantImage        `json:"plant_images" gorm:"foreignKey:PlantID;constraint:OnDelete:CASCADE"`
	CreatedAt             time.Time           `json:"created_at"`
	UpdatedAt             time.Time           `json:"updated_at"`
}

type PlantProgress struct {
	ID        int       `json:"id" gorm:"primaryKey"`
	PlantID   int       `json:"plant_id"`
	UserID    int       `json:"user_id"`
	ImageURL  string    `json:"image_url"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type PlantCharacteristic struct {
	ID         int    `json:"id" gorm:"primaryKey"`
	PlantID    int    `json:"plant_id"`
	Height     int    `json:"height"`
	HeightUnit string `json:"height_unit"`
	Wide       int    `json:"wide"`
	WideUnit   string `json:"wide_unit"`
	LeafColor  string `json:"leaf_color"`
}

type PlantReminder struct {
	ID                   int       `json:"id" gorm:"primaryKey"`
	PlantID              int       `json:"plant_id"`
	WateringFrequency    int       `json:"watering_frequency"`
	Each                 string    `json:"each"`
	WateringAmount       int       `json:"watering_amount"`
	Unit                 string    `json:"unit"`
	WateringTime         string    `json:"watering_time"`
	WeatherCondition     string    `json:"weather_condition"`
	ConditionDescription string    `json:"condition_description"`
	CreatedAt            time.Time `json:"created_at"`
	UpdatedAt            time.Time `json:"updated_at"`
}

type PlantInstructionCategory struct {
	ID          int       `json:"id" gorm:"primaryKey"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	ImageURL    string    `json:"image_url"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type PlantInstruction struct {
	ID                    int                      `json:"id" gorm:"primaryKey"`
	PlantID               int                      `json:"plant_id"`
	InstructionCategoryID int                      `json:"instruction_category_id"`
	InstructionCategory   PlantInstructionCategory `json:"instruction_category" gorm:"foreignKey:InstructionCategoryID;references:ID"`
	StepNumber            int                      `json:"step_number"`
	StepTitle             string                   `json:"step_title"`
	StepDescription       string                   `json:"step_description"`
	StepImageURL          string                   `json:"step_image_url"`
	CreatedAt             time.Time                `json:"created_at"`
	UpdatedAt             time.Time                `json:"updated_at"`
}

type PlantFAQ struct {
	ID        int       `json:"id" gorm:"primaryKey"`
	PlantID   int       `json:"plant_id"`
	Question  string    `json:"question"`
	Answer    string    `json:"answer"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type PlantImage struct {
	ID        int       `json:"id" gorm:"primaryKey"`
	PlantID   int       `json:"plant_id"`
	FileName  string    `json:"file_name"`
	IsPrimary int       `json:"is_primary"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type PlantCategory struct {
	ID        int       `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name"`
	ImageURL  string    `json:"image_url"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserPlant struct {
	ID                   int       `json:"id" gorm:"primaryKey"`
	UserID               int       `json:"user_id"`
	PlantID              int       `json:"plant_id"`
	CustomizeName        string    `json:"customize_name"`
	InstructionCategory1 int       `json:"instruction_category_1" gorm:"default:1"`
	InstructionCategory2 int       `json:"instruction_category_2" gorm:"default:0"`
	InstructionCategory3 int       `json:"instruction_category_3" gorm:"default:0"`
	InstructionCategory4 int       `json:"instruction_category_4" gorm:"default:0"`
	CreatedAt            time.Time `json:"created_at"`
	UpdatedAt            time.Time `json:"updated_at"`

	Plant Plant     `json:"plant" gorm:"foreignKey:PlantID;references:ID"`
	User  user.User `json:"user" gorm:"foreignKey:UserID;references:ID"`
}

type UserPlantHistory struct {
	ID            int       `json:"id" gorm:"primaryKey"`
	UserID        int       `json:"user_id"`
	PlantID       int       `json:"plant_id"`
	PlantName     string    `json:"plant_name"`
	PlantCategory string    `json:"plant_category"`
	PlantImageURL string    `json:"plant_image_url"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`

	User user.User `json:"user" gorm:"foreignKey:UserID;references:ID"`
}

type PlantEarliestWatering struct {
	ID           int    `json:"id"`
	PlantID      int    `json:"plant_id"`
	WateringTime string `json:"watering_time"`
}
