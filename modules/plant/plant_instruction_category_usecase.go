package plant

type PlantInstructionCategoryService interface {
	FindAll() ([]PlantInstructionCategoryResponse, error)
	FindByID(id int) (PlantInstructionCategoryResponse, error)
	FindInstructionByCategoryID(plantID int, instructionCategoryID int) (PlantInstructionsGroupedResponse, error)
	Create(input PlantInstructionCategoryInput, fileLocation string) (PlantInstructionCategoryResponse, error)
	Update(id int, input PlantInstructionCategoryInput, imageURL string) (PlantInstructionCategoryResponse, error)
	Delete(id int) error
}

type plantInstructionCategoryService struct {
	repository PlantInstructionCategoryRepository
}

func NewPlantInstructionCategoryService(repository PlantInstructionCategoryRepository) PlantInstructionCategoryService {
	return &plantInstructionCategoryService{repository}
}

func (s *plantInstructionCategoryService) FindAll() ([]PlantInstructionCategoryResponse, error) {
	categories, err := s.repository.FindAll()
	if err != nil {
		return nil, err
	}

	var responses []PlantInstructionCategoryResponse
	for _, category := range categories {
		responses = append(responses, NewPlantInstructionCategoryResponse(category))
	}

	return responses, nil
}

func (s *plantInstructionCategoryService) FindByID(id int) (PlantInstructionCategoryResponse, error) {
	category, err := s.repository.FindByID(id)
	if err != nil {
		return PlantInstructionCategoryResponse{}, err
	}

	return NewPlantInstructionCategoryResponse(category), nil
}

func (s *plantInstructionCategoryService) FindInstructionByCategoryID(plantID int, instructionCategoryID int) (PlantInstructionsGroupedResponse, error) {
	instructions, err := s.repository.FindInstructionByCategoryID(plantID, instructionCategoryID)
	if err != nil {
		return PlantInstructionsGroupedResponse{}, err
	}

	response := NewPlantInstructionStepResponses(instructions)

	return response, nil
}

func (s *plantInstructionCategoryService) Create(input PlantInstructionCategoryInput, imageURL string) (PlantInstructionCategoryResponse, error) {
	category := PlantInstructionCategory{
		Name:        input.Name,
		Description: input.Description,
		ImageURL:    imageURL,
	}

	newCategory, err := s.repository.Create(category)
	if err != nil {
		return PlantInstructionCategoryResponse{}, err
	}

	return NewPlantInstructionCategoryResponse(newCategory), nil
}

func (s *plantInstructionCategoryService) Update(id int, input PlantInstructionCategoryInput, imageURL string) (PlantInstructionCategoryResponse, error) {
	category, err := s.repository.FindByID(id)
	if err != nil {
		return PlantInstructionCategoryResponse{}, err
	}

	category.Name = input.Name
	category.Description = input.Description
	if imageURL != "" {
		category.ImageURL = imageURL
	}

	updatedCategory, err := s.repository.Update(category)
	if err != nil {
		return PlantInstructionCategoryResponse{}, err
	}

	return NewPlantInstructionCategoryResponse(updatedCategory), nil
}

func (s *plantInstructionCategoryService) Delete(id int) error {
	category, err := s.repository.FindByID(id)
	if err != nil {
		return err
	}

	return s.repository.Delete(category)
}
