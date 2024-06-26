package article

import "gorm.io/gorm"

type Repository interface {
	StoreArticle(*Article) (*Article, error)
	GetArticle(int) (*Article, error)
	GetAllArticles() ([]Article, error)
	UpdateArticle(a *Article, id int) (*Article, error)
	DeleteArticle(int) error
}

type articleRepository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *articleRepository {
	return &articleRepository{
		db: db,
	}
}

func (r *articleRepository) StoreArticle(a *Article) (*Article, error) {
	err := r.db.Create(a).Error
	if err != nil {
		return nil, err
	}
	return a, nil
}

func (r *articleRepository) GetArticle(id int) (*Article, error) {
	var a Article
	err := r.db.Where("id = ?", id).First(&a).Error
	if err != nil {
		return nil, err
	}
	return &a, nil
}

func (r *articleRepository) GetAllArticles() ([]Article, error) {
	var articles []Article
	err := r.db.Find(&articles).Error
	if err != nil {
		return nil, err
	}
	return articles, nil
}

func (r *articleRepository) UpdateArticle(a *Article, id int) (*Article, error) {
	err := r.db.Where("id = ?", id).Updates(a).Error
	if err != nil {
		return nil, err
	}
	return a, nil
}

func (r *articleRepository) DeleteArticle(id int) error {
	err := r.db.Where("id = ?", id).Delete(&Article{}).Error
	if err != nil {
		return err
	}
	return nil
}
