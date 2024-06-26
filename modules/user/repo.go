package user

import (
	"errors"
	"fmt"

	"github.com/OctavianoRyan25/be-agriculture/constants"
	"gorm.io/gorm"
)

type Repository interface {
	RegisterUser(*User) (*User, error)
	IsDuplicateEmail(string) (bool, error)
	VerifyEmail(string, string) error
	IsValidated(string) (bool, error)
	Login(*User) (*User, error)
	GetUserProfile(uint) (*User, error)
	GetUser(string) (*User, error)
	ResetPassword(string, string) error
	UpdateFCMToken(uint, string) error
}

type userRepository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *userRepository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) RegisterUser(user *User) (*User, error) {
	err := r.db.Create(user).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *userRepository) IsDuplicateEmail(email string) (bool, error) {
	var user User
	err := r.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, fmt.Errorf("error querying database: %v", err)
	}
	return true, nil
}

func (r *userRepository) VerifyEmail(email, otp string) error {
	var user User
	err := r.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return err
	}

	if user.OTP != otp {
		return errors.New(constants.ErrInvalidOTP)
	}

	user.Is_Active = true
	err = r.db.Save(&user).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *userRepository) IsValidated(email string) (bool, error) {
	var user User
	err := r.db.Where("email = ? AND is_active = ?", email, true).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, fmt.Errorf("error querying database: %v", err)
	}
	return true, nil
}

func (r *userRepository) Login(user *User) (*User, error) {
	// var u User
	err := r.db.Where("email = ?", user.Email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *userRepository) GetUserProfile(id uint) (*User, error) {
	var user User
	err := r.db.Where("id = ?", id).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) GetUser(email string) (*User, error) {
	var user User
	err := r.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) ResetPassword(email, password string) error {
	var user User
	err := r.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return err
	}

	user.Password = password
	err = r.db.Save(&user).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *userRepository) UpdateFCMToken(id uint, fcmToken string) error {
	var user User
	err := r.db.Where("id = ?", id).First(&user).Error
	if err != nil {
		return err
	}

	user.FCMToken = fcmToken
	err = r.db.Save(&user).Error
	if err != nil {
		return err
	}

	return nil
}
