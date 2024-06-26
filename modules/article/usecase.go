package article

import "time"

type UseCase interface {
	StoreArticle(*Article) (*Article, error)
	GetArticle(int) (*Article, error)
	GetAllArticles() ([]Article, error)
	UpdateArticle(a *Article, id int) (*Article, error)
	DeleteArticle(int) error
}

type useCase struct {
	repo Repository
}

func NewUseCase(repo Repository) *useCase {
	return &useCase{
		repo: repo,
	}
}

func (uc *useCase) StoreArticle(a *Article) (*Article, error) {
	location, _ := time.LoadLocation("Asia/Jakarta")
	a.CreatedAt = time.Now().In(location)
	a.UpdatedAt = time.Now().In(location)
	return uc.repo.StoreArticle(a)
}

func (uc *useCase) GetArticle(id int) (*Article, error) {
	return uc.repo.GetArticle(id)
}

func (uc *useCase) GetAllArticles() ([]Article, error) {
	return uc.repo.GetAllArticles()
}

func (uc *useCase) UpdateArticle(a *Article, id int) (*Article, error) {
	location, _ := time.LoadLocation("Asia/Jakarta")
	a.UpdatedAt = time.Now().In(location)
	return uc.repo.UpdateArticle(a, id)
}

func (uc *useCase) DeleteArticle(id int) error {
	return uc.repo.DeleteArticle(id)
}
