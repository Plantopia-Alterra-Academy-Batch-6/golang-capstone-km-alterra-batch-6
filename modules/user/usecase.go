package user

import (
	"bytes"
	"errors"
	"fmt"
	"html/template"
	"path/filepath"
	"time"

	"github.com/OctavianoRyan25/be-agriculture/constants"
	"gopkg.in/gomail.v2"
)

type UserUseCase interface {
	RegisterUser(*User) (*User, int, error)
	CheckEmail(string) (int, error)
	SendEmailVerification(*User) (int, error)
	VerifyEmail(string, string) (int, error)
	Login(*User) (*User, int, error)
	GetUserProfile(uint) (*User, int, error)
	GetUser(string) (*User, int, error)
	ResetPassword(string, string) (int, error)
}

type userUseCase struct {
	repo Repository
}

func NewUseCase(repo Repository) *userUseCase {
	return &userUseCase{
		repo: repo,
	}
}

func (uc *userUseCase) RegisterUser(user *User) (*User, int, error) {
	duplicate, err := uc.repo.IsDuplicateEmail(user.Email)
	if err != nil {
		return nil, constants.ErrCodeBadRequest, err
	}
	if duplicate {
		return nil, constants.ErrCodeEmailAlreadyExist, errors.New(constants.ErrEmailAlreadyExist)
	}
	user.Is_Active = false
	user.OTP = RandomOTP()
	user.Created_at = time.Now()
	user.Updated_at = time.Now()

	registeredUser, err := uc.repo.RegisterUser(user)
	return registeredUser, constants.CodeSuccess, err
}

func (uc *userUseCase) CheckEmail(email string) (int, error) {
	duplicate, err := uc.repo.IsDuplicateEmail(email)
	if err != nil {
		return constants.ErrCodeBadRequest, err
	}
	if duplicate {
		return constants.ErrCodeEmailAlreadyExist, errors.New(constants.ErrEmailAlreadyExist)
	}
	return constants.CodeSuccess, nil
}

func (uc *userUseCase) SendEmailVerification(user *User) (int, error) {
	m := gomail.NewMessage()
	m.SetHeader("From", EMAIL_FROM)
	m.SetHeader("To", user.Email)
	m.SetHeader("Subject", "Email Verification")

	OTP := user.OTP
	username := user.Name
	//Production
	path := filepath.Join("/app", "modules", "user", "template", "base.html")
	//Development
	// path := filepath.Join("modules", "user", "template", "base.html")
	template, err := template.ParseFiles(path)
	if err != nil {
		return constants.ErrCodeBadRequest, err
	}

	var body bytes.Buffer
	data := struct {
		Username string
		OTP      string
	}{
		Username: username,
		OTP:      OTP,
	}

	err = template.Execute(&body, data)
	if err != nil {
		return constants.ErrCodeBadRequest, err
	}
	m.SetBody("text/html", body.String())
	d := gomail.NewDialer(SMTP_HOST, 587, SMTP_USER, SMTP_PASSWORD)

	if err := d.DialAndSend(m); err != nil {
		return constants.ErrCodeBadRequest, err
	}

	return constants.CodeSuccess, nil
}

func (uc *userUseCase) VerifyEmail(email, otp string) (int, error) {
	err := uc.repo.VerifyEmail(email, otp)
	if err != nil {
		return constants.ErrCodeBadRequest, err
	}
	return constants.CodeSuccess, nil
}

func (uc *userUseCase) Login(user *User) (*User, int, error) {
	request_fcm := user.FCMToken
	user, err := uc.repo.Login(user)
	if err != nil {
		return nil, constants.ErrCodeBadRequest, err
	}
	if !user.Is_Active {
		return nil, constants.ErrCodeEmailNotValidatedYet, errors.New(constants.ErrEmailNotValidatedYet)
	}
	response_fcm := user.FCMToken
	fmt.Println("FCM token changed")
	if request_fcm != response_fcm {
		user.FCMToken = request_fcm
		err = uc.repo.UpdateFCMToken(uint(user.ID), user.FCMToken)
		if err != nil {
			return nil, constants.ErrCodeBadRequest, err
		}
	}
	fmt.Println("FCM token updated")
	return user, constants.CodeSuccess, nil
}

func (uc *userUseCase) GetUserProfile(id uint) (*User, int, error) {
	user, err := uc.repo.GetUserProfile(id)
	if err != nil {
		return nil, constants.ErrCodeBadRequest, err
	}
	return user, constants.CodeSuccess, nil
}

func (uc *userUseCase) GetUser(email string) (*User, int, error) {
	user, err := uc.repo.GetUser(email)
	if err != nil {
		return nil, constants.ErrCodeBadRequest, err
	}
	return user, constants.CodeSuccess, nil
}

func (uc *userUseCase) ResetPassword(email, password string) (int, error) {
	err := uc.repo.ResetPassword(email, password)
	if err != nil {
		return constants.ErrCodeBadRequest, err
	}
	return constants.CodeSuccess, nil
}
