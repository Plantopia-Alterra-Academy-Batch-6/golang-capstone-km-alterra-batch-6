package admin

import (
	"net/http"

	"github.com/OctavianoRyan25/be-agriculture/base"
	"github.com/OctavianoRyan25/be-agriculture/utils/helper"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type AdminController struct {
	adminUseCase adminUseCase
}

func NewUserController(adminUseCase adminUseCase) *AdminController {
	return &AdminController{
		adminUseCase: adminUseCase,
	}
}

func (c *AdminController) RegisterUser(ctx echo.Context) error {
	req := new(AdminRequest)
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

	registeredUser, code, err := c.adminUseCase.RegisterUser(mapped)
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

func (c *AdminController) CheckEmail(ctx echo.Context) error {
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

	code, err := c.adminUseCase.CheckEmail(req.Email)
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

func (c *AdminController) Login(ctx echo.Context) error {
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
	user, code, err := c.adminUseCase.Login(mapped)
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

	token, err := helper.GenerateToken(uint(user.ID), user.Email, "admin")
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

func (c *AdminController) GetUserProfile(ctx echo.Context) error {
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
	if role != "admin" {
		errRes := base.ErrorResponse{
			Status:  "error",
			Message: "Forbidden access",
			Code:    http.StatusForbidden,
		}
		return ctx.JSON(http.StatusForbidden, errRes)

	}

	user, code, err := c.adminUseCase.GetUserProfile(userId)
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
