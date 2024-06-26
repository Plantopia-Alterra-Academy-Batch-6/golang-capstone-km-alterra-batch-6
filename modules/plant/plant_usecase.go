package plant

import (
	"time"
)

type PlantService interface {
	FindAll(page, limit int) ([]PlantResponse, error)
	FindByID(id int) (PlantResponse, error)
	CreatePlant(input CreatePlantInput) (PlantResponse, error)
	UpdatePlant(id int, input UpdatePlantInput) (PlantResponse, error)
	DeletePlant(id int) (PlantResponse, error)
	CountAll() (int64, error)
	SearchPlantsByName(name string, page, limit int) ([]Plant, int64, error)
	GetRecommendations(userID int) ([]PlantResponse, error)
	FindByCategoryID(categoryID int) ([]PlantResponse, error)
	CategoryExists(categoryID int) (bool, error)
}

type plantService struct {
	repository             PlantRepository
	plantCategoryRepository PlantCategoryRepository
}

func NewPlantService(repository PlantRepository, plantCategoryRepository PlantCategoryRepository) PlantService {
	return &plantService{repository, plantCategoryRepository}
}

func (s *plantService) GetRecommendations(userID int) ([]PlantResponse, error) {
	var plants []Plant
	err := s.repository.FindRecommendations(userID, &plants).Error
	if err != nil {
			return nil, err
	}

	var plantResponses []PlantResponse
	for _, plant := range plants {
			plantResponses = append(plantResponses, NewPlantResponse(plant))
	}

	return plantResponses, nil
}

func (s *plantService) FindByCategoryID(categoryID int) ([]PlantResponse, error) {
	var plants []Plant
	err := s.repository.FindByCategoryID(categoryID, &plants).Error
	if err != nil {
			return nil, err
	}

	var plantResponses []PlantResponse
	for _, plant := range plants {
			plantResponses = append(plantResponses, NewPlantResponse(plant))
	}

	return plantResponses, nil
}

func (s *plantService) CategoryExists(categoryID int) (bool, error) {
	return s.repository.CategoryExists(categoryID)
}

func (s *plantService) FindAll(page, limit int) ([]PlantResponse, error) {
	plants, err := s.repository.FindAll(page, limit)
	if err != nil {
			return nil, err
	}

	var responses []PlantResponse
	for _, plant := range plants {
			responses = append(responses, NewPlantResponse(plant))
	}

	return responses, nil
}

func (s *plantService) FindByID(id int) (PlantResponse, error) {
	plant, err := s.repository.FindByID(id)
	if err != nil {
		return PlantResponse{}, err
	}

	return NewPlantResponse(plant), nil
}

func (s *plantService) CreatePlant(input CreatePlantInput) (PlantResponse, error) {
	category, err := s.plantCategoryRepository.FindByID(input.PlantCategoryID)
	if err != nil {
		return PlantResponse{}, err
	}

	plant := Plant{
		Name:               input.Name,
		Description:        input.Description,
		IsToxic:            input.IsToxic,
		HarvestDuration:    input.HarvestDuration,
		Sunlight: input.Sunlight,
		PlantingTime: input.PlantingTime,
		PlantCategoryID:    input.PlantCategoryID,
		PlantCategory:      category, 
		ClimateCondition:   input.ClimateCondition,
		CreatedAt:          time.Now(),
		UpdatedAt:          time.Now(),
		PlantInstructions:  make([]PlantInstruction, len(input.PlantInstructions)),
		AdditionalTips:     input.AdditionalTips,
		PlantFAQs:          make([]PlantFAQ, len(input.PlantFAQs)),
		PlantImages:        make([]PlantImage, len(input.PlantImages)),
		PlantCharacteristic: PlantCharacteristic{
			Height:     input.PlantCharacteristic.Height,
			HeightUnit: input.PlantCharacteristic.HeightUnit,
			Wide:       input.PlantCharacteristic.Wide,
			WideUnit:   input.PlantCharacteristic.WideUnit,
			LeafColor:  input.PlantCharacteristic.LeafColor,
		},
		WateringSchedule: PlantReminder{
			WateringFrequency:   input.WateringSchedule.WateringFrequency,
			Each:                input.WateringSchedule.Each,
			WateringAmount:      input.WateringSchedule.WateringAmount,
			Unit:                input.WateringSchedule.Unit,
			WateringTime:        input.WateringSchedule.WateringTime,
			WeatherCondition:    input.WateringSchedule.WeatherCondition,
			ConditionDescription: input.WateringSchedule.ConditionDescription,
		},
	}

	for i, instruction := range input.PlantInstructions {
		plant.PlantInstructions[i] = PlantInstruction{
			InstructionCategoryID: instruction.InstructionCategoryID,
			StepNumber:      instruction.StepNumber,
			StepTitle:       instruction.StepTitle,
			StepDescription: instruction.StepDescription,
			StepImageURL:    instruction.StepImageURL,
			CreatedAt:       time.Now(),
			UpdatedAt:       time.Now(),
		}
	}

	for i, faq := range input.PlantFAQs {
		plant.PlantFAQs[i] = PlantFAQ{
			Question:  faq.Question,
			Answer:    faq.Answer,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}
	}

	for i, image := range input.PlantImages {
		plant.PlantImages[i] = PlantImage{
			FileName:  image.FileName,
			IsPrimary: image.IsPrimary,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}
	}

	newPlant, err := s.repository.Create(plant)
	if err != nil {
		return PlantResponse{}, err
	}

	return NewPlantResponse(newPlant), nil
}

func (s *plantService) UpdatePlant(id int, input UpdatePlantInput) (PlantResponse, error) {
	plant, err := s.repository.FindByIDWithRelations(id)
	if err != nil {
		return PlantResponse{}, err
	}

	category, err := s.plantCategoryRepository.FindByID(input.PlantCategoryID)
	if err != nil {
		return PlantResponse{}, err
	}

	// Update basic fields
	plant.Name = input.Name
	plant.Description = input.Description
	plant.IsToxic = input.IsToxic
	plant.HarvestDuration = input.HarvestDuration
	plant.Sunlight = input.Sunlight
	plant.PlantingTime = input.PlantingTime
	plant.PlantCategoryID = input.PlantCategoryID
	plant.PlantCategory = category
	plant.ClimateCondition = input.ClimateCondition
	plant.AdditionalTips = input.AdditionalTips
	plant.UpdatedAt = time.Now()

	// Update PlantCharacteristic
	plant.PlantCharacteristic = PlantCharacteristic{
		Height:     input.PlantCharacteristic.Height,
		HeightUnit: input.PlantCharacteristic.HeightUnit,
		Wide:       input.PlantCharacteristic.Wide,
		WideUnit:   input.PlantCharacteristic.WideUnit,
		LeafColor:  input.PlantCharacteristic.LeafColor,
	}

	// Update WateringSchedule
	plant.WateringSchedule = PlantReminder{
		WateringFrequency:    input.WateringSchedule.WateringFrequency,
		Each:                 input.WateringSchedule.Each,
		WateringAmount:       input.WateringSchedule.WateringAmount,
		Unit:                 input.WateringSchedule.Unit,
		WateringTime:         input.WateringSchedule.WateringTime,
		WeatherCondition:     input.WateringSchedule.WeatherCondition,
		ConditionDescription: input.WateringSchedule.ConditionDescription,
	}

	// Clear existing instructions, FAQs, and images
	if err := s.repository.ClearPlantInstructions(id); err != nil {
		return PlantResponse{}, err
	}
	if err := s.repository.ClearPlantFAQs(id); err != nil {
		return PlantResponse{}, err
	}
	if err := s.repository.ClearPlantImages(id); err != nil {
		return PlantResponse{}, err
	}

	// Clear the slices to ensure no old data is retained
	plant.PlantInstructions = []PlantInstruction{}
	plant.PlantFAQs = []PlantFAQ{}
	plant.PlantImages = []PlantImage{}

	// Add new instructions
	for _, instruction := range input.PlantInstructions {
		newInstruction := PlantInstruction{
			PlantID:              id,
			InstructionCategoryID: instruction.InstructionCategoryID,
			StepNumber:      instruction.StepNumber,
			StepTitle:       instruction.StepTitle,
			StepDescription: instruction.StepDescription,
			StepImageURL:    instruction.StepImageURL,
			CreatedAt:       time.Now(),
			UpdatedAt:       time.Now(),
		}
		plant.PlantInstructions = append(plant.PlantInstructions, newInstruction)
	}

	// Add new FAQs
	for _, faq := range input.PlantFAQs {
		newFAQ := PlantFAQ{
			PlantID:    id,
			Question:   faq.Question,
			Answer:     faq.Answer,
			CreatedAt:  time.Now(),
			UpdatedAt:  time.Now(),
		}
		plant.PlantFAQs = append(plant.PlantFAQs, newFAQ)
	}

	// Add new images
	for _, image := range input.PlantImages {
		newImage := PlantImage{
			PlantID:    id,
			FileName:   image.FileName,
			IsPrimary:  image.IsPrimary,
			CreatedAt:  time.Now(),
			UpdatedAt:  time.Now(),
		}
		plant.PlantImages = append(plant.PlantImages, newImage)
	}

	updatedPlant, err := s.repository.Update(plant)
	if err != nil {
		return PlantResponse{}, err
	}

	return NewPlantResponse(updatedPlant), nil
}

func (s *plantService) DeletePlant(id int) (PlantResponse, error) {
	plant, err := s.repository.FindByIDWithRelations(id)
	if err != nil {
			return PlantResponse{}, err
	}

	if err := s.repository.Delete(id); err != nil {
			return PlantResponse{}, err
	}

	return NewPlantResponse(plant), nil
}

func (s *plantService) CountAll() (int64, error) {
	var count int64
	err := s.repository.CountAll(&count)
	return count, err
}

func (s *plantService) SearchPlantsByName(name string, page, limit int) ([]Plant, int64, error) {
	offset := (page - 1) * limit
	plants, err := s.repository.SearchPlantsByName(name, limit, offset)
	if err != nil {
		return nil, 0, err
	}

	totalCount, err := s.repository.CountPlantsByName(name)
	if err != nil {
		return nil, 0, err
	}

	return plants, totalCount, nil
}


