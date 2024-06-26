package fertilizer

type FertilizerService interface {
	CreateFertilizer(input FertilizerInput) (FertilizerResponse, error)
	GetFertilizer() ([]FertilizerResponse, error)
	GetFertilizerByID(id int) (FertilizerResponse, error)
	DeleteFertilizer(id int) error
	UpdateFertilizer(id int, input FertilizerInput) (FertilizerResponse, error)
}

type fertilizerService struct {
	repository FertilizerRepository
}

func NewFertilizerService(repository FertilizerRepository) FertilizerService {
	return &fertilizerService{repository}
}

func (s *fertilizerService) GetFertilizer() ([]FertilizerResponse, error) {
	categories, err := s.repository.GetFertilizer()
	if err != nil {
		return nil, err
	}

	var responses []FertilizerResponse
	for _, category := range categories {
		responses = append(responses, NewFertilizerResponse(category))
	}

	return responses, nil
}

func (s *fertilizerService) GetFertilizerByID(id int) (FertilizerResponse, error) {
	category, err := s.repository.GetFertilizerByID(id)
	if err != nil {
		return FertilizerResponse{}, err
	}

	return NewFertilizerResponse(category), nil
}

// UpdateFertilizer implements FertilizerService.
func (s *fertilizerService) UpdateFertilizer(id int, input FertilizerInput) (FertilizerResponse, error) {
	category, err := s.repository.GetFertilizerByID(id)
	if err != nil {
		return FertilizerResponse{}, err
	}

	category.Name = input.Name

	updatedCategory, err := s.repository.UpdateFertilizer(category)
	if err != nil {
		return FertilizerResponse{}, err
	}

	return NewFertilizerResponse(updatedCategory), nil
}

func (s *fertilizerService) CreateFertilizer(input FertilizerInput) (FertilizerResponse, error) {
	category := Fertilizer{
		Id: input.Id,
		Name:     input.Name,
		Compostition: input.Compostition,
		CreateAt: input.CreateAt,
	}
	newCategory, err := s.repository.CreateFertilizer(category)
	if err != nil {
		return FertilizerResponse{}, err
	}

	return NewFertilizerResponse(newCategory), nil
}

func (s *fertilizerService) DeleteFertilizer(id int) error {
	category, err := s.repository.GetFertilizerByID(id)
	if err != nil {
		return err
	}

	return s.repository.DeleteFertilizer(category)
}
