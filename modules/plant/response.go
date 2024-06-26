package plant

import "time"

type PlantResponse struct {
	ID                  int                          `json:"id"`
	Name                string                       `json:"name"`
	Description         string                       `json:"description"`
	IsToxic             bool                         `json:"is_toxic"`
	HarvestDuration     int                          `json:"harvest_duration"`
	PlantCategory       PlantCategoryClimateResponse `json:"plant_category"`
	ClimateCondition    string                       `json:"climate_condition"`
	PlantingTime        string                       `json:"planting_time"`
	Sunlight            string                       `json:"sunlight"`
	PlantCharacteristic PlantCharacteristicResponse  `json:"plant_characteristic"`
	WateringSchedule    PlantReminderResponse        `json:"watering_schedule"`
	PlantInstruction    []PlantInstructionResponse   `json:"plant_instructions"`
	AdditionalTips      string                       `json:"additional_tips"`
	PlantFAQ            []PlantFAQResponse           `json:"plant_faqs"`
	PlantImages         []PlantImageResponse         `json:"plant_images"`
	CreatedAt           time.Time                    `json:"created_at"`
}

type PlantCategoryClimateResponse struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	ImageURL string `json:"image_url"`
}

type PlantEarliestWateringResponse struct {
	ID           int    `json:"id"`
	PlantID      int    `json:"plant_id"`
	WateringTime string `json:"watering_time"`
}

type PlantCharacteristicResponse struct {
	ID         int    `json:"id"`
	Height     int    `json:"height"`
	HeightUnit string `json:"height_unit"`
	Wide       int    `json:"wide"`
	WideUnit   string `json:"wide_unit"`
	LeafColor  string `json:"leaf_color"`
}

type PlantImageResponse struct {
	ID        int    `json:"id"`
	PlantID   int    `json:"plant_id"`
	FileName  string `json:"file_name"`
	IsPrimary int    `json:"is_primary"`
}

type PlantReminderResponse struct {
	ID                   int    `json:"id"`
	PlantID              int    `json:"plant_id"`
	WateringFrequency    int    `json:"watering_frequency"`
	Each                 string `json:"each"`
	WateringAmount       int    `json:"watering_amount"`
	Unit                 string `json:"unit"`
	WateringTime         string `json:"watering_time"`
	WeatherCondition     string `json:"weather_condition"`
	ConditionDescription string `json:"condition_description"`
}

type PlantInstructionResponse struct {
	ID                  int                              `json:"id"`
	PlantID             int                              `json:"plant_id"`
	InstructionCategory PlantInstructionCategoryResponse `json:"instruction_category"`
	StepNumber          int                              `json:"step_number"`
	StepTitle           string                           `json:"step_title"`
	StepDescription     string                           `json:"step_description"`
	StepImageURL        string                           `json:"step_image_url"`
}

type PlantInstructionCategoryResponse struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	ImageURL    string `json:"image_url"`
}

type PlantFAQResponse struct {
	ID        int       `json:"id"`
	PlantID   int       `json:"plant_id"`
	Question  string    `json:"question"`
	Answer    string    `json:"answer"`
	CreatedAt time.Time `json:"created_at"`
}

type PlantProgressResponse struct {
	ID        int       `json:"id"`
	PlantID   int       `json:"plant_id"`
	UserID    int       `json:"user_id"`
	ImageURL  string    `json:"image_url"`
	CreatedAt time.Time `json:"created_at"`
}

func NewPlantResponse(plant Plant) PlantResponse {
	return PlantResponse{
		ID:                  plant.ID,
		Name:                plant.Name,
		Description:         plant.Description,
		IsToxic:             plant.IsToxic,
		HarvestDuration:     plant.HarvestDuration,
		ClimateCondition:    plant.ClimateCondition,
		PlantingTime:        plant.PlantingTime,
		Sunlight:            plant.Sunlight,
		PlantCategory:       NewPlantCategoryResponse(plant.PlantCategory),
		PlantCharacteristic: NewPlantCharacteristicResponse(plant.PlantCharacteristic),
		WateringSchedule:    NewPlantReminderResponse(plant.WateringSchedule),
		PlantInstruction:    NewPlantInstructionResponses(plant.PlantInstructions),
		AdditionalTips:      plant.AdditionalTips,
		PlantFAQ:            NewPlantFAQResponses(plant.PlantFAQs),
		PlantImages:         NewPlantImageResponses(plant.PlantImages),
		CreatedAt:           plant.CreatedAt,
	}
}

func NewPlantCategoryResponse(category PlantCategory) PlantCategoryClimateResponse {
	return PlantCategoryClimateResponse{
		ID:       category.ID,
		Name:     category.Name,
		ImageURL: category.ImageURL,
	}
}

func NewPlantEarliestWateringResponse(category PlantEarliestWatering) PlantEarliestWateringResponse {
	return PlantEarliestWateringResponse{
		ID:           category.ID,
		PlantID:      category.PlantID,
		WateringTime: category.WateringTime,
	}
}

func NewPlantProgressResponse(progress PlantProgress) PlantProgressResponse {
	return PlantProgressResponse{
		ID:        progress.ID,
		PlantID:   progress.PlantID,
		UserID:    progress.UserID,
		ImageURL:  progress.ImageURL,
		CreatedAt: progress.CreatedAt,
	}
}

func NewPlantInstructionCategoryResponse(category PlantInstructionCategory) PlantInstructionCategoryResponse {
	return PlantInstructionCategoryResponse{
		ID:          category.ID,
		Name:        category.Name,
		Description: category.Description,
		ImageURL:    category.ImageURL,
	}
}

func NewPlantCharacteristicResponse(characteristic PlantCharacteristic) PlantCharacteristicResponse {
	return PlantCharacteristicResponse{
		ID:         characteristic.ID,
		Height:     characteristic.Height,
		HeightUnit: characteristic.HeightUnit,
		Wide:       characteristic.Wide,
		WideUnit:   characteristic.WideUnit,
		LeafColor:  characteristic.LeafColor,
	}
}

func NewPlantImageResponses(images []PlantImage) []PlantImageResponse {
	var responses []PlantImageResponse

	for _, img := range images {
		responses = append(responses, NewPlantImageResponse(img))
	}

	return responses
}

func NewPlantImageResponse(image PlantImage) PlantImageResponse {
	return PlantImageResponse{
		ID:        image.ID,
		PlantID:   image.PlantID,
		FileName:  image.FileName,
		IsPrimary: image.IsPrimary,
	}
}

func NewPlantReminderResponse(reminder PlantReminder) PlantReminderResponse {
	return PlantReminderResponse{
		ID:                   reminder.ID,
		PlantID:              reminder.PlantID,
		WateringFrequency:    reminder.WateringFrequency,
		Each:                 reminder.Each,
		WateringAmount:       reminder.WateringAmount,
		Unit:                 reminder.Unit,
		WateringTime:         reminder.WateringTime,
		WeatherCondition:     reminder.WeatherCondition,
		ConditionDescription: reminder.ConditionDescription,
	}
}

func NewPlantInstructionResponses(instructions []PlantInstruction) []PlantInstructionResponse {
	var responses []PlantInstructionResponse

	for _, instruction := range instructions {
		responses = append(responses, NewPlantInstructionResponse(instruction))
	}

	return responses
}

func NewPlantInstructionResponse(instruction PlantInstruction) PlantInstructionResponse {
	return PlantInstructionResponse{
		ID:                  instruction.ID,
		PlantID:             instruction.PlantID,
		InstructionCategory: NewPlantInstructionCategoryResponse(instruction.InstructionCategory),
		StepNumber:          instruction.StepNumber,
		StepTitle:           instruction.StepTitle,
		StepDescription:     instruction.StepDescription,
		StepImageURL:        instruction.StepImageURL,
	}
}

func NewPlantFAQResponses(faqs []PlantFAQ) []PlantFAQResponse {
	var responses []PlantFAQResponse

	for _, faq := range faqs {
		responses = append(responses, NewPlantFAQResponse(faq))
	}

	return responses
}

func NewPlantFAQResponse(faq PlantFAQ) PlantFAQResponse {
	return PlantFAQResponse{
		ID:        faq.ID,
		PlantID:   faq.PlantID,
		Question:  faq.Question,
		Answer:    faq.Answer,
		CreatedAt: faq.CreatedAt,
	}
}

type UserPlantResponse struct {
	ID                   int           `json:"id"`
	Plant                PlantResponse `json:"plant"`
	CustomizeName        string        `json:"customize_name"`
	InstructionCategory1 int           `json:"instruction_category_1"`
	InstructionCategory2 int           `json:"instruction_category_2"`
	InstructionCategory3 int           `json:"instruction_category_3"`
	InstructionCategory4 int           `json:"instruction_category_4"`
	CreatedAt            time.Time     `json:"created_at"`
}

func NewUserPlantResponse(userPlant UserPlant) UserPlantResponse {
	return UserPlantResponse{
		ID:                   userPlant.ID,
		Plant:                NewPlantResponse(userPlant.Plant),
		CustomizeName:        userPlant.CustomizeName,
		InstructionCategory1: userPlant.InstructionCategory1,
		InstructionCategory2: userPlant.InstructionCategory2,
		InstructionCategory3: userPlant.InstructionCategory3,
		InstructionCategory4: userPlant.InstructionCategory4,
		CreatedAt:            userPlant.CreatedAt,
	}
}

func NewUserPlantResponses(userPlants []UserPlant) map[int][]UserPlantResponse {
	responses := make(map[int][]UserPlantResponse)

	for _, userPlant := range userPlants {
		userID := userPlant.UserID
		if _, ok := responses[userID]; !ok {
			responses[userID] = make([]UserPlantResponse, 0)
		}
		responses[userID] = append(responses[userID], NewUserPlantResponse(userPlant))
	}

	return responses
}

type UserPlantHistoryResponse struct {
	ID            int       `json:"id"`
	UserID        int       `json:"user_id"`
	PlantName     string    `json:"plant_name"`
	PlantCategory string    `json:"plant_category"`
	ImageURL      string    `json:"image_url"`
	CreatedAt     time.Time `json:"created_at"`
}

func NewUserPlantHistoryResponse(userPlantHistory UserPlantHistory) UserPlantHistoryResponse {
	return UserPlantHistoryResponse{
		ID:            userPlantHistory.ID,
		UserID:        userPlantHistory.UserID,
		PlantName:     userPlantHistory.PlantName,
		PlantCategory: userPlantHistory.PlantCategory,
		ImageURL:      userPlantHistory.PlantImageURL,
		CreatedAt:     userPlantHistory.CreatedAt,
	}
}

type PlantInstructionStepResponse struct {
	ID              int    `json:"id"`
	StepNumber      int    `json:"step_number"`
	StepTitle       string `json:"step_title"`
	StepDescription string `json:"step_description"`
	StepImageURL    string `json:"step_image_url"`
}

type PlantInstructionsGroupedResponse struct {
	PlantID             int                              `json:"plant_id"`
	InstructionCategory PlantInstructionCategoryResponse `json:"instruction_category"`
	Steps               []PlantInstructionStepResponse   `json:"steps"`
}

func NewPlantInstructionStepResponses(instructions []PlantInstruction) PlantInstructionsGroupedResponse {
	if len(instructions) == 0 {
		return PlantInstructionsGroupedResponse{}
	}

	firstInstruction := instructions[0]

	groupedResponse := PlantInstructionsGroupedResponse{
		PlantID: firstInstruction.PlantID,
		InstructionCategory: PlantInstructionCategoryResponse{
			ID:          firstInstruction.InstructionCategory.ID,
			Name:        firstInstruction.InstructionCategory.Name,
			Description: firstInstruction.InstructionCategory.Description,
			ImageURL:    firstInstruction.InstructionCategory.ImageURL,
		},
	}

	for _, instruction := range instructions {
		groupedResponse.Steps = append(groupedResponse.Steps, PlantInstructionStepResponse{
			ID:              instruction.ID,
			StepNumber:      instruction.StepNumber,
			StepTitle:       instruction.StepTitle,
			StepDescription: instruction.StepDescription,
			StepImageURL:    instruction.StepImageURL,
		})
	}

	return groupedResponse
}
