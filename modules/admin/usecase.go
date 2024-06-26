package admin

import (
	"errors"
	"time"

	"github.com/OctavianoRyan25/be-agriculture/constants"
)

type AdminUseCase interface {
	RegisterUser(*Admin) (*Admin, int, error)
	CheckEmail(string) (int, error)
	Login(*Admin) (*Admin, int, error)
	GetUserProfile(uint) (*Admin, int, error)
}

type adminUseCase struct {
	repo Repository
}

func NewUseCase(repo Repository) *adminUseCase {
	return &adminUseCase{
		repo: repo,
	}
}

func (uc *adminUseCase) RegisterUser(user *Admin) (*Admin, int, error) {
	duplicate, err := uc.repo.IsDuplicateEmail(user.Email)
	if err != nil {
		return nil, constants.ErrCodeBadRequest, err
	}
	if duplicate {
		return nil, constants.ErrCodeEmailAlreadyExist, errors.New(constants.ErrEmailAlreadyExist)
	}
	user.Created_at = time.Now()
	user.Updated_at = time.Now()

	registeredUser, err := uc.repo.RegisterUser(user)
	return registeredUser, constants.CodeSuccess, err
}

func (uc *adminUseCase) CheckEmail(email string) (int, error) {
	duplicate, err := uc.repo.IsDuplicateEmail(email)
	if err != nil {
		return constants.ErrCodeBadRequest, err
	}
	if duplicate {
		return constants.ErrCodeEmailAlreadyExist, errors.New(constants.ErrEmailAlreadyExist)
	}
	return constants.CodeSuccess, nil
}

func (uc *adminUseCase) Login(user *Admin) (*Admin, int, error) {
	user, err := uc.repo.Login(user)
	if err != nil {
		return nil, constants.ErrCodeBadRequest, err
	}
	return user, constants.CodeSuccess, nil
}

func (uc *adminUseCase) GetUserProfile(id uint) (*Admin, int, error) {
	user, err := uc.repo.GetUserProfile(id)
	if err != nil {
		return nil, constants.ErrCodeBadRequest, err
	}
	return user, constants.CodeSuccess, nil
}
