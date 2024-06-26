package search

import "github.com/OctavianoRyan25/be-agriculture/modules/plant"

type Usecase interface {
	Search(params PlantSearchParams) ([]plant.Plant, error)
}

type searchUsecase struct {
	repo Repository
}

func NewUsecase(repo Repository) *searchUsecase {
	return &searchUsecase{
		repo: repo,
	}
}

func (uc *searchUsecase) Search(params PlantSearchParams) ([]plant.Plant, error) {
	plants, err := uc.repo.Search(params)
	if err != nil {
		return nil, err
	}
	if len(plants) == 0 {
		return nil, nil
	}
	return plants, nil
}
