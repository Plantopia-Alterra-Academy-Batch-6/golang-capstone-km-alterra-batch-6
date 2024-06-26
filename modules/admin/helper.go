package admin

import (
	"math/rand"
)

func MapUserRequestToUser(userRequest *AdminRequest) *Admin {
	return &Admin{
		Name:      userRequest.Name,
		Email:     userRequest.Email,
		Password:  userRequest.Password,
		Url_Image: userRequest.Url_Image,
	}
}

func MapUserToResponse(user *Admin) *AdminResponse {
	return &AdminResponse{
		ID:         user.ID,
		Name:       user.Name,
		Email:      user.Email,
		Url_Image:  user.Url_Image,
		Created_at: user.Created_at,
	}
}

func MapLoginRequestToUser(loginRequest *LoginRequest) *Admin {
	return &Admin{
		Email:    loginRequest.Email,
		Password: loginRequest.Password,
	}
}

func RandomOTP() string {
	var code [6]byte
	_, err := rand.Read(code[:])
	if err != nil {
		panic(err) // atau tangani error sesuai kebutuhan
	}
	for i := range code {
		code[i] = (code[i] % 10) + '0' // Ubah ke angka 0-9
	}
	return string(code[:])
}
