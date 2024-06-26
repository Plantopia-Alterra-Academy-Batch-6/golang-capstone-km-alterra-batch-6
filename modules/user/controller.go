package user

import (
	"net/http"

	"github.com/OctavianoRyan25/be-agriculture/base"
	"github.com/OctavianoRyan25/be-agriculture/utils/helper"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type UserController struct {
	userUseCase UserUseCase
}

func NewUserController(userUseCase UserUseCase) *UserController {
	return &UserController{
		userUseCase: userUseCase,
	}
}

func (c *UserController) RegisterUser(ctx echo.Context) error {
	req := new(UserRequest)
	//Mapping UserRequest to User
	err := ctx.Bind(&req)
	if err != nil {
		errRes := base.ErrorResponse{
			Status:  "error",
			Message: err.Error(),
			Code:    http.StatusUnprocessableEntity,
		}
		return ctx.JSON(http.StatusUnprocessableEntity, errRes)
	}

	validate := validator.New()

	err = validate.Struct(req)
	if err != nil {
		errRes := base.ErrorResponse{
			Status:  "error",
			Message: err.Error(),
			Code:    http.StatusBadRequest,
		}
		return ctx.JSON(http.StatusBadRequest, errRes)
	}

	mapped := MapUserRequestToUser(req)
	mapped.Password = HashPass(req.Password)

	registeredUser, code, err := c.userUseCase.RegisterUser(mapped)
	if err != nil {
		errRes := base.ErrorResponse{
			Status:  "error",
			Message: err.Error(),
			Code:    code,
		}
		return ctx.JSON(code, errRes)
	}

	code, err = c.userUseCase.SendEmailVerification(registeredUser)
	if err != nil {
		errRes := base.ErrorResponse{
			Status:  "error",
			Message: err.Error(),
			Code:    code,
		}
		return ctx.JSON(code, errRes)
	}

	mappedRes := MapUserToResponse(registeredUser)

	res := base.SuccessResponse{
		Status:  "success",
		Message: "User registered",
		Data:    mappedRes,
	}

	return ctx.JSON(code, res)
}

func (c *UserController) CheckEmail(ctx echo.Context) error {
	req := new(CheckEmailRequest)
	err := ctx.Bind(&req)
	if err != nil {
		errRes := base.ErrorResponse{
			Status:  "error",
			Message: err.Error(),
			Code:    http.StatusUnprocessableEntity,
		}
		return ctx.JSON(http.StatusUnprocessableEntity, errRes)
	}

	code, err := c.userUseCase.CheckEmail(req.Email)
	if err != nil {
		errRes := base.ErrorResponse{
			Status:  "error",
			Message: err.Error(),
			Code:    code,
		}
		return ctx.JSON(code, errRes)
	}

	res := base.SuccessResponse{
		Status:  "success",
		Message: "Email is available",
	}

	return ctx.JSON(code, res)
}

func (c *UserController) VerifyEmail(ctx echo.Context) error {
	req := new(OTPRequest)
	err := ctx.Bind(&req)
	if err != nil {
		errRes := base.ErrorResponse{
			Status:  "error",
			Message: err.Error(),
			Code:    http.StatusUnprocessableEntity,
		}
		return ctx.JSON(http.StatusUnprocessableEntity, errRes)
	}

	code, err := c.userUseCase.VerifyEmail(req.Email, req.OTP)
	if err != nil {
		errRes := base.ErrorResponse{
			Status:  "error",
			Message: err.Error(),
			Code:    code,
		}
		return ctx.JSON(code, errRes)
	}

	res := base.SuccessResponse{
		Status:  "success",
		Message: "Email verified",
	}

	return ctx.JSON(code, res)
}

func (c *UserController) Login(ctx echo.Context) error {
	req := new(LoginRequest)
	err := ctx.Bind(&req)
	if err != nil {
		errRes := base.ErrorResponse{
			Status:  "error",
			Message: err.Error(),
			Code:    http.StatusUnprocessableEntity,
		}
		return ctx.JSON(http.StatusUnprocessableEntity, errRes)
	}

	validate := validator.New()

	err = validate.Struct(req)
	if err != nil {
		errRes := base.ErrorResponse{
			Status:  "error",
			Message: err.Error(),
			Code:    http.StatusBadRequest,
		}
		return ctx.JSON(http.StatusBadRequest, errRes)
	}

	mapped := MapLoginRequestToUser(req)
	user, code, err := c.userUseCase.Login(mapped)
	if err != nil {
		errRes := base.ErrorResponse{
			Status:  "error",
			Message: err.Error(),
			Code:    code,
		}
		return ctx.JSON(code, errRes)
	}

	comparePass := ComparePass([]byte(user.Password), []byte(req.Password))
	if !comparePass {
		errRes := base.ErrorResponse{
			Status:  "error",
			Message: "Invalid email or password",
			Code:    400,
		}
		return ctx.JSON(http.StatusBadRequest, errRes)
	}

	token, err := helper.GenerateToken(uint(user.ID), user.Email, "user")
	if err != nil {
		errRes := base.ErrorResponse{
			Status:  "error",
			Message: "Failed to generate token",
			Code:    500,
		}
		return ctx.JSON(http.StatusInternalServerError, errRes)
	}
	resToken := LoginResponse{
		Token: token,
	}
	res := base.SuccessResponse{
		Status:  "success",
		Message: "Login success",
		Data:    resToken,
	}

	return ctx.JSON(code, res)
}

func (c *UserController) GetUserProfile(ctx echo.Context) error {
	userId := ctx.Get("user_id").(uint)
	if userId == 0 {
		errRes := base.ErrorResponse{
			Status:  "error",
			Message: "Bad request",
			Code:    http.StatusBadRequest,
		}
		return ctx.JSON(http.StatusBadRequest, errRes)
	}
	role := ctx.Get("role").(string)
	if role != "user" {
		errRes := base.ErrorResponse{
			Status:  "error",
			Message: "Forbidden access",
			Code:    http.StatusForbidden,
		}
		return ctx.JSON(http.StatusForbidden, errRes)

	}

	user, code, err := c.userUseCase.GetUserProfile(userId)
	if err != nil {
		errRes := base.ErrorResponse{
			Status:  "error",
			Message: err.Error(),
			Code:    code,
		}
		return ctx.JSON(code, errRes)
	}

	mappedRes := MapUserToResponse(user)

	res := base.SuccessResponse{
		Status:  "success",
		Message: "User profile",
		Data:    mappedRes,
	}

	return ctx.JSON(code, res)
}

func (c *UserController) ResendOTP(ctx echo.Context) error {
	req := new(CheckEmailRequest)
	err := ctx.Bind(&req)
	if err != nil {
		errRes := base.ErrorResponse{
			Status:  "error",
			Message: err.Error(),
			Code:    http.StatusUnprocessableEntity,
		}
		return ctx.JSON(http.StatusUnprocessableEntity, errRes)
	}
	user, code, err := c.userUseCase.GetUser(req.Email)
	if err != nil {
		errRes := base.ErrorResponse{
			Status:  "error",
			Message: err.Error(),
			Code:    code,
		}
		return ctx.JSON(code, errRes)
	}
	res, err := c.userUseCase.SendEmailVerification(user)
	if err != nil {
		errRes := base.ErrorResponse{
			Status:  "error",
			Message: err.Error(),
			Code:    code,
		}
		return ctx.JSON(code, errRes)
	}
	resSuccess := base.SuccessResponse{
		Status:  "success",
		Message: "OTP sent",
		Data:    res,
	}
	return ctx.JSON(code, resSuccess)
}

func (c *UserController) ResetPassword(ctx echo.Context) error {
	req := new(ResetPasswordRequest)
	err := ctx.Bind(&req)
	if err != nil {
		errRes := base.ErrorResponse{
			Status:  "error",
			Message: err.Error(),
			Code:    http.StatusUnprocessableEntity,
		}
		return ctx.JSON(http.StatusUnprocessableEntity, errRes)
	}
	validate := validator.New()

	err = validate.Struct(req)
	if err != nil {
		errRes := base.ErrorResponse{
			Status:  "error",
			Message: err.Error(),
			Code:    http.StatusBadRequest,
		}
		return ctx.JSON(http.StatusBadRequest, errRes)
	}

	user, code, err := c.userUseCase.GetUser(req.Email)
	if err != nil {
		errRes := base.ErrorResponse{
			Status:  "error",
			Message: err.Error(),
			Code:    code,
		}
		return ctx.JSON(code, errRes)
	}
	user.Password = HashPass(req.NewPassword)
	code, err = c.userUseCase.ResetPassword(user.Email, user.Password)
	if err != nil {
		errRes := base.ErrorResponse{
			Status:  "error",
			Message: err.Error(),
			Code:    code,
		}
		return ctx.JSON(code, errRes)
	}
	res := base.SuccessResponse{
		Status:  "success",
		Message: "Password change successfully",
	}
	return ctx.JSON(code, res)
}
