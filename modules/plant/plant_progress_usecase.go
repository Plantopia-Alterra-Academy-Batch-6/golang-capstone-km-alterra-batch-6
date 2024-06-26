package plant

type PlantProgressService interface {
	FindProgressByUserIDAndPlantID(userID int, plantID int) ([]PlantProgressResponse, error)
	FindByID(id int) (PlantProgressResponse, error)
	Create(input PlantProgressInput, fileLocation string) (PlantProgressResponse, error)
	Delete(id int) error
}

type plantProgressService struct {
	repository PlantProgressRepository
}

func NewPlantProgressService(repository PlantProgressRepository) PlantProgressService {
	return &plantProgressService{repository}
}

func (s *plantProgressService) FindProgressByUserIDAndPlantID(userID int, plantID int) ([]PlantProgressResponse, error) {
	plantProgresses, err := s.repository.FindByUserIDAndPlantID(userID, plantID)
	if err != nil {
		return nil, err
	}

	var responses []PlantProgressResponse
	for _, progress := range plantProgresses {
		responses = append(responses, NewPlantProgressResponse(progress))
	}

	return responses, nil
}

func (s *plantProgressService) FindByID(id int) (PlantProgressResponse, error) {
	itemProgress, err := s.repository.FindByID(id)
	if err != nil {
		return PlantProgressResponse{}, err
	}

	return NewPlantProgressResponse(itemProgress), nil
}

func (s *plantProgressService) Create(input PlantProgressInput, imageURL string) (PlantProgressResponse, error) {
	itemProgress := PlantProgress{
		UserID:   input.UserID,
		PlantID:  input.PlantID,
		ImageURL: imageURL,
	}

	newItemProgress, err := s.repository.Create(itemProgress)
	if err != nil {
		return PlantProgressResponse{}, err
	}

	return NewPlantProgressResponse(newItemProgress), nil
}

func (s *plantProgressService) Delete(id int) error {
	itemProgress, err := s.repository.FindByID(id)
	if err != nil {
		return err
	}

	return s.repository.Delete(itemProgress)
}
