package notification

import (
	"net/http"
	"strconv"

	"github.com/OctavianoRyan25/be-agriculture/base"
	"github.com/labstack/echo/v4"
)

type NotificationController struct {
	UseCase UseCase
}

func NewNotificationController(useCase UseCase) *NotificationController {
	return &NotificationController{
		UseCase: useCase,
	}
}

func (c *NotificationController) ReadNotification(ctx echo.Context) error {
	userId := ctx.Get("user_id").(uint)
	params := ctx.Param("id")
	convId, _ := strconv.Atoi(params)
	// if userId == 0 {
	// 	errRes := base.ErrorResponse{
	// 		Status:  "error",
	// 		Message: "Id cannot be empty",
	// 		Code:    http.StatusBadRequest,
	// 	}
	// 	return ctx.JSON(http.StatusBadRequest, errRes)
	// }
	notification, err := c.UseCase.ReadNotification(convId)
	if err != nil {
		errRes := base.ErrorResponse{
			Status:  "error",
			Message: err.Error(),
			Code:    http.StatusInternalServerError,
		}
		return ctx.JSON(http.StatusInternalServerError, errRes)
	}

	if userId != uint(notification.UserId) {
		errRes := base.ErrorResponse{
			Status:  "error",
			Message: "You are not authorized to read this notification",
			Code:    http.StatusUnauthorized,
		}
		return ctx.JSON(http.StatusUnauthorized, errRes)
	}

	mapped := &NotificationResponse{
		Id:        notification.Id,
		Title:     notification.Title,
		Body:      notification.Body,
		UserId:    notification.UserId,
		IsRead:    notification.IsRead,
		PlantId:   notification.PlantId,
		CreatedAt: notification.CreatedAt,
	}

	res := base.SuccessResponse{
		Status:  "success",
		Message: "Notification stored",
		Data:    mapped,
	}
	return ctx.JSON(http.StatusOK, res)
}

func (c *NotificationController) GetAllNotifications(ctx echo.Context) error {
	userId := ctx.Get("user_id").(uint)
	if userId == 0 {
		errRes := base.ErrorResponse{
			Status:  "error",
			Message: "Id cannot be empty",
			Code:    http.StatusBadRequest,
		}
		return ctx.JSON(http.StatusBadRequest, errRes)
	}
	notifications, err := c.UseCase.GetAllNotifications(userId)
	if err != nil {
		errRes := base.ErrorResponse{
			Status:  "error",
			Message: err.Error(),
			Code:    http.StatusInternalServerError,
		}
		return ctx.JSON(http.StatusInternalServerError, errRes)
	}
	var mapped []NotificationResponse
	for _, v := range notifications {
		mapped = append(mapped, NotificationResponse{
			Id:        v.Id,
			Title:     v.Title,
			Body:      v.Body,
			UserId:    v.UserId,
			IsRead:    v.IsRead,
			PlantId:   v.PlantId,
			CreatedAt: v.CreatedAt,
		})
	}
	res := base.SuccessResponse{
		Status:  "success",
		Message: "Notifications fetched",
		Data:    mapped,
	}
	return ctx.JSON(http.StatusOK, res)
}

func (c *NotificationController) DeleteAllNotifications(ctx echo.Context) error {
	userId := ctx.Get("user_id").(uint)
	if userId == 0 {
		errRes := base.ErrorResponse{
			Status:  "error",
			Message: "Id cannot be empty",
			Code:    http.StatusBadRequest,
		}
		return ctx.JSON(http.StatusBadRequest, errRes)
	}
	err := c.UseCase.DeleteAllNotifications(userId)
	if err != nil {
		errRes := base.ErrorResponse{
			Status:  "error",
			Message: err.Error(),
			Code:    http.StatusInternalServerError,
		}
		return ctx.JSON(http.StatusInternalServerError, errRes)
	}
	res := base.SuccessResponse{
		Status:  "success",
		Message: "Notifications deleted",
	}
	return ctx.JSON(http.StatusOK, res)
}

func (c *NotificationController) CreateCustomizeWateringReminder(ctx echo.Context) error {
	UserID := ctx.Get("user_id").(uint)
	reminder := new(CustomizeWateringReminderRequest)
	if err := ctx.Bind(reminder); err != nil {
		errRes := base.ErrorResponse{
			Status:  "error",
			Message: err.Error(),
			Code:    http.StatusBadRequest,
		}
		return ctx.JSON(http.StatusBadRequest, errRes)
	}
	reminderModel := &CustomizeWateringReminder{
		PlantId:   reminder.PlantID,
		UserId:    int(UserID),
		Time:      reminder.Time,
		Recurring: reminder.Recurring,
		Type:      reminder.Type,
	}
	_, err := c.UseCase.CreateCustomizeWateringReminder(reminderModel)
	if err != nil {
		errRes := base.ErrorResponse{
			Status:  "error",
			Message: err.Error(),
			Code:    http.StatusInternalServerError,
		}
		return ctx.JSON(http.StatusInternalServerError, errRes)
	}
	mapped := &CustomizeWateringReminderResponse{
		Id:        reminderModel.Id,
		PlantID:   reminderModel.PlantId,
		Plant:     *MapPlantToPlantResponse(&reminderModel.Plant),
		UserID:    reminderModel.UserId,
		User:      reminderModel.User,
		Time:      reminderModel.Time,
		Recurring: reminderModel.Recurring,
		Type:      reminderModel.Type,
		CreatedAt: reminderModel.CreatedAt,
	}
	res := base.SuccessResponse{
		Status:  "success",
		Message: "Customize watering reminder created",
		Data:    mapped,
	}
	return ctx.JSON(http.StatusCreated, res)
}
