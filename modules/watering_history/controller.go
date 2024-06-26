package wateringhistory

import (
	"fmt"
	"net/http"

	"github.com/OctavianoRyan25/be-agriculture/base"
	"github.com/OctavianoRyan25/be-agriculture/modules/user"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type WateringHistoryController struct {
	useCase WateringHistoryUseCase
}

func NeWateringHistoryController(useCase WateringHistoryUseCase) *WateringHistoryController {
	return &WateringHistoryController{
		useCase: useCase,
	}
}

func (c *WateringHistoryController) StoreWateringHistory(ctx echo.Context) error {
	userID := ctx.Get("user_id").(uint)
	req := new(WateringHistoryRequest)
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

	mapped := &WateringHistory{
		PlantID: req.PlantID,
		UserID:  int(userID),
	}

	wh, err := c.useCase.StoreWateringHistory(mapped)
	if err != nil {
		errRes := base.ErrorResponse{
			Status:  "error",
			Message: err.Error(),
			Code:    http.StatusInternalServerError,
		}
		return ctx.JSON(http.StatusInternalServerError, errRes)
	}

	mappedUser := user.MapUserToResponse(&wh.User)
	mappedPlant := MapPlantToPlantResponse(&wh.Plant)

	mappedres := &WateringHistoryResponse{
		Id:        wh.ID,
		Plant:     *mappedPlant,
		User:      *mappedUser,
		CreatedAt: wh.CreatedAt,
	}

	res := base.SuccessResponse{
		Status:  "success",
		Message: "Watering history created",
		Data:    mappedres,
	}

	return ctx.JSON(http.StatusCreated, res)
}

func (c *WateringHistoryController) GetAllWateringHistories(ctx echo.Context) error {
	userID := ctx.Get("user_id").(uint)

	wh, err := c.useCase.GetAllWateringHistories(userID)
	if err != nil {
		errRes := base.ErrorResponse{
			Status:  "error",
			Message: err.Error(),
			Code:    http.StatusInternalServerError,
		}
		return ctx.JSON(http.StatusInternalServerError, errRes)
	}

	var mappedRes []WateringHistoryResponse
	for _, v := range wh {
		mappedRes = append(mappedRes, WateringHistoryResponse{
			Id:        v.ID,
			Plant:     *MapPlantToPlantResponse(&v.Plant),
			User:      *user.MapUserToResponse(&v.User),
			CreatedAt: v.CreatedAt,
		})
	}

	res := base.SuccessResponse{
		Status:  "success",
		Message: "Watering histories fetched",
		Data:    mappedRes,
	}

	return ctx.JSON(http.StatusOK, res)
}

func (c *WateringHistoryController) GetLateWateringHistories(ctx echo.Context) error {
	userID := ctx.Get("user_id").(uint)
	if userID == 0 {
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

	notification, err := c.useCase.GetLateWateringHistories(userID)
	if err != nil {
		errRes := base.ErrorResponse{
			Status:  "error",
			Message: err.Error(),
			Code:    http.StatusInternalServerError,
		}
		return ctx.JSON(http.StatusInternalServerError, errRes)
	}

	mappedPlant := MapPlantToPlantResponse(&notification.Plant)

	res := base.SuccessResponse{
		Status:  "success",
		Message: "Late watering history fetched",
		Data: NotificationResponse{
			Id:        notification.Id,
			Title:     fmt.Sprintf("Hi there! You Forgot to Water your %s plant", notification.Plant.Name),
			Body:      "Remember, your plants need water to thrive. Let's water them",
			Plant:     *mappedPlant,
			UserID:    notification.UserId,
			CreatedAt: notification.CreatedAt,
		},
	}

	return ctx.JSON(http.StatusOK, res)
}
