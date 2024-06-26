package admin

import (
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type Repository interface {
	RegisterUser(*Admin) (*Admin, error)
	IsDuplicateEmail(string) (bool, error)
	Login(*Admin) (*Admin, error)
	GetUserProfile(uint) (*Admin, error)
}

type adminRespository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *adminRespository {
	return &adminRespository{
		db: db,
	}
}

func (r *adminRespository) RegisterUser(user *Admin) (*Admin, error) {
	err := r.db.Create(user).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *adminRespository) IsDuplicateEmail(email string) (bool, error) {
	var admin Admin
	err := r.db.Where("email = ?", email).First(&admin).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, fmt.Errorf("error querying database: %v", err)
	}
	return true, nil
}

func (r *adminRespository) Login(admin *Admin) (*Admin, error) {
	// var u User
	err := r.db.Where("email = ?", admin.Email).First(&admin).Error
	if err != nil {
		return nil, err
	}
	return admin, nil
}

func (r *adminRespository) GetUserProfile(id uint) (*Admin, error) {
	var admin Admin
	err := r.db.Where("id = ?", id).First(&admin).Error
	if err != nil {
		return nil, err
	}
	return &admin, nil
}
