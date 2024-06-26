package plant

import (
	"errors"
	"fmt"
)

type UserPlantService interface {
	AddUserPlant(input AddUserPlantInput) (UserPlantResponse, error)
	GetUserPlantsByUserID(userID int, limit int, offset int) (map[int][]UserPlantResponse, error)
	DeleteUserPlantByID(userPlantID int) (UserPlantResponse, error)
	GetUserPlantByID(userPlantID int) (UserPlant, error)
	CountByUserID(userID int) (int64, error)

	AddUserPlantHistory(input UserPlantHistoryInput) (UserPlantHistoryResponse, error)
	GetUserPlantHistoryByUserID(userID int) ([]UserPlantHistoryResponse, error)
	CheckPlantExists(plantID int) (bool, error)
	CheckUserPlantExists(userPlantID int) (bool, error)
	UpdateCustomizeName(userPlantID int, customizeName string) (UserPlantResponse, error)
	CheckUserPlantExistsForAdd(userID, plantID int) (bool, error)
	UpdateInstructionCategory(input UpdateInstructionCategoryInput) error
	GetUserPlantByUserIDAndPlantID(userID int, plantID int) (UserPlantResponse, error)
}

type userPlantService struct {
	repository UserPlantRepository
}

func NewUserPlantService(repository UserPlantRepository) UserPlantService {
	return &userPlantService{repository}
}

func (s *userPlantService) GetUserPlantByUserIDAndPlantID(userID int, plantID int) (UserPlantResponse, error) {
	userPlant, err := s.repository.GetUserPlantByUserIDAndPlantID(userID, plantID)
	if err != nil {
			return UserPlantResponse{}, err
	}
	return NewUserPlantResponse(userPlant), nil
}

func (s *userPlantService) UpdateInstructionCategory(input UpdateInstructionCategoryInput) error {
	userPlant, err := s.repository.GetUserPlantByID(input.UserPlantID)
	if err != nil || userPlant.ID == 0 {
		return errors.New("user plant not found")
	}

	var fieldToUpdate string
	switch input.InstructionCategory {
	case 1:
		fieldToUpdate = "instruction_category1"
	case 2:
		fieldToUpdate = "instruction_category2"
	case 3:
		fieldToUpdate = "instruction_category3"
	case 4:
		fieldToUpdate = "instruction_category4"
	default:
		return errors.New("invalid instruction category")
	}

	err = s.repository.UpdateInstructionCategory(input.UserPlantID, fieldToUpdate)
	return err
}

func (s *userPlantService) CheckPlantExists(plantID int) (bool, error) {
	return s.repository.CheckPlantExists(plantID)
}

func (s *userPlantService) GetUserPlantHistoryByUserID(userID int) ([]UserPlantHistoryResponse, error) {
	histories, err := s.repository.GetUserPlantHistoryByUserID(userID)
	if err != nil {
		return nil, err
	}

	var response []UserPlantHistoryResponse
	for _, history := range histories {
		response = append(response, NewUserPlantHistoryResponse(history))
	}

	return response, nil
}

func (s *userPlantService) AddUserPlantHistory(input UserPlantHistoryInput) (UserPlantHistoryResponse, error) {
	plant, err := s.repository.GetPlantByID(input.PlantID)
	if err != nil {
		return UserPlantHistoryResponse{}, err
	}

	category, err := s.repository.GetPlantCategoryByID(plant.PlantCategoryID)
	if err != nil {
		return UserPlantHistoryResponse{}, err
	}

	image, err := s.repository.GetPrimaryPlantImageByPlantID(input.PlantID)
	if err != nil {
		return UserPlantHistoryResponse{}, err
	}

	userPlantHistory := UserPlantHistory{
		UserID:        input.UserID,
		PlantID:       input.PlantID,
		PlantName:     plant.Name,
		PlantCategory: category.Name,
		PlantImageURL: image.FileName,
	}

	createdUserPlantHistory, err := s.repository.Create(userPlantHistory)
	if err != nil {
		return UserPlantHistoryResponse{}, err
	}

	return NewUserPlantHistoryResponse(createdUserPlantHistory), nil
}

func (s *userPlantService) AddUserPlant(input AddUserPlantInput) (UserPlantResponse, error) {
	// Check if the plant already exists in the user's list
	userPlants, err := s.repository.GetUserPlantsByUserID(input.UserID, 0, 0)
	if err != nil {
		return UserPlantResponse{}, err
	}

	for _, userPlant := range userPlants {
		if userPlant.PlantID == input.PlantID {
			return UserPlantResponse{}, fmt.Errorf("plant with ID %d is already added by the user", input.PlantID)
		}
	}

	exists, err := s.CheckPlantExists(input.PlantID)
	if err != nil {
		return UserPlantResponse{}, err
	}
	if !exists {
		return UserPlantResponse{}, fmt.Errorf("plant with ID %d does not exist", input.PlantID)
	}

	// Default CustomizeName to Plant Name if not provided
	if input.CustomizeName == "" {
		plantData, err := s.repository.GetPlantByID(input.PlantID)
		if err != nil {
			return UserPlantResponse{}, err
		}
		input.CustomizeName = plantData.Name
	}

	userPlant := UserPlant{
		UserID:        input.UserID,
		PlantID:       input.PlantID,
		CustomizeName: input.CustomizeName,
	}

	createdUserPlant, err := s.repository.AddUserPlant(userPlant)
	if err != nil {
		return UserPlantResponse{}, err
	}

	// Fetch complete user plant data
	completeUserPlant, err := s.repository.GetUserPlantByID(createdUserPlant.ID)
	if err != nil {
		return UserPlantResponse{}, err
	}

	return NewUserPlantResponse(completeUserPlant), nil
}

func (s *userPlantService) UpdateCustomizeName(userPlantID int, customizeName string) (UserPlantResponse, error) {
	exists, err := s.CheckUserPlantExists(userPlantID)
	if err != nil {
		return UserPlantResponse{}, err
	}
	if !exists {
		return UserPlantResponse{}, errors.New("user plant ID not found")
	}

	err = s.repository.UpdateCustomizeName(userPlantID, customizeName)
	if err != nil {
		return UserPlantResponse{}, err
	}

	// Retrieve updated user plant
	userPlant, err := s.repository.GetUserPlantByID(userPlantID)
	if err != nil {
		return UserPlantResponse{}, err
	}

	return NewUserPlantResponse(userPlant), nil
}

func (s *userPlantService) CheckUserPlantExists(userPlantID int) (bool, error) {
	return s.repository.CheckUserPlantExists(userPlantID)
}

func (s *userPlantService) GetUserPlantsByUserID(userID int, limit int, page int) (map[int][]UserPlantResponse, error) {
	offset := (page - 1) * limit

	if limit <= 0 {
		limit = -1
		offset = -1
	}

	userPlants, err := s.repository.GetUserPlantsByUserID(userID, limit, offset)
	if err != nil {
		return nil, err
	}

	return NewUserPlantResponses(userPlants), nil
}

func (s *userPlantService) DeleteUserPlantByID(userPlantID int) (UserPlantResponse, error) {
	deletedUserPlant, err := s.repository.GetUserPlantByID(userPlantID)
	if err != nil {
		return UserPlantResponse{}, err
	}

	err = s.repository.DeleteUserPlantByID(userPlantID)
	if err != nil {
		return UserPlantResponse{}, err
	}

	deletedUserPlantResponse := NewUserPlantResponse(deletedUserPlant)

	return deletedUserPlantResponse, nil
}

func (s *userPlantService) GetUserPlantByID(userPlantID int) (UserPlant, error) {
	return s.repository.GetUserPlantByID(userPlantID)
}

func (s *userPlantService) CountByUserID(userID int) (int64, error) {
	var count int64
	err := s.repository.CountByUserID(userID, &count)
	return count, err
}

func (s *userPlantService) CheckUserPlantExistsForAdd(userID, plantID int) (bool, error) {
	return s.repository.CheckUserPlantExistsForAdd(userID, plantID)
}
