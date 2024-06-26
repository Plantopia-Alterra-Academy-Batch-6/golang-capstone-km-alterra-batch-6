package plant

type PlantCategoryService interface {
	FindAll() ([]PlantCategoryClimateResponse, error)
	FindByID(id int) (PlantCategoryClimateResponse, error)
	Create(input PlantCategoryClimateInput, fileLocation string) (PlantCategoryClimateResponse, error)
	Update(id int, input PlantCategoryClimateInput, imageURL string) (PlantCategoryClimateResponse, error)
	Delete(id int) error
}

type plantCategoryService struct {
	repository PlantCategoryRepository
}

func NewPlantCategoryService(repository PlantCategoryRepository) PlantCategoryService {
	return &plantCategoryService{repository}
}

func (s *plantCategoryService) FindAll() ([]PlantCategoryClimateResponse, error) {
	categories, err := s.repository.FindAll()
	if err != nil {
		return nil, err
	}

	var responses []PlantCategoryClimateResponse
	for _, category := range categories {
		responses = append(responses, NewPlantCategoryResponse(category))
	}

	return responses, nil
}

func (s *plantCategoryService) FindByID(id int) (PlantCategoryClimateResponse, error) {
	category, err := s.repository.FindByID(id)
	if err != nil {
		return PlantCategoryClimateResponse{}, err
	}

	return NewPlantCategoryResponse(category), nil
}

func (s *plantCategoryService) Create(input PlantCategoryClimateInput, imageURL string) (PlantCategoryClimateResponse, error) {
	category := PlantCategory{
		Name:     input.Name,
		ImageURL: imageURL,
	}

	newCategory, err := s.repository.Create(category)
	if err != nil {
		return PlantCategoryClimateResponse{}, err
	}

	return NewPlantCategoryResponse(newCategory), nil
}

func (s *plantCategoryService) Update(id int, input PlantCategoryClimateInput, imageURL string) (PlantCategoryClimateResponse, error) {
	category, err := s.repository.FindByID(id)
	if err != nil {
		return PlantCategoryClimateResponse{}, err
	}

	category.Name = input.Name
	if imageURL != "" {
		category.ImageURL = imageURL
	}

	updatedCategory, err := s.repository.Update(category)
	if err != nil {
		return PlantCategoryClimateResponse{}, err
	}

	return NewPlantCategoryResponse(updatedCategory), nil
}

func (s *plantCategoryService) Delete(id int) error {
	category, err := s.repository.FindByID(id)
	if err != nil {
		return err
	}

	return s.repository.Delete(category)
}
