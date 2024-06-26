package handler

import (
	"net/http"
	"strconv"

	"github.com/OctavianoRyan25/be-agriculture/modules/plant"
	"github.com/OctavianoRyan25/be-agriculture/utils/helper"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type UserPlantHandler struct {
	service plant.UserPlantService
}

func NewUserPlantHandler(service plant.UserPlantService) *UserPlantHandler {
	return &UserPlantHandler{service}
}

func (h *UserPlantHandler) GetUserPlantByUserIDAndPlantID(c echo.Context) error {
	plantID, err := strconv.Atoi(c.Param("plant_id"))
	if err != nil {
			response := helper.APIResponse("Invalid plant ID", http.StatusBadRequest, "error", nil)
			return c.JSON(http.StatusBadRequest, response)
	}

	userID, ok := c.Get("user_id").(uint)
	if !ok {
			response := helper.APIResponse("Invalid user ID from token", http.StatusInternalServerError, "error", nil)
			return c.JSON(http.StatusInternalServerError, response)
	}

	userPlant, err := h.service.GetUserPlantByUserIDAndPlantID(int(userID), plantID)
	if err != nil {
			response := helper.APIResponse("Failed to fetch user plant", http.StatusInternalServerError, "error", nil)
			return c.JSON(http.StatusInternalServerError, response)
	}

	response := helper.APIResponse("User plant fetched successfully", http.StatusOK, "success", userPlant)
	return c.JSON(http.StatusOK, response)
}


func (h *UserPlantHandler) UpdateInstructionCategory(c echo.Context) error {
	var request plant.UpdateInstructionCategoryInput
	if err := c.Bind(&request); err != nil {
		response := helper.APIResponse("Invalid request", http.StatusBadRequest, "error", nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	validate := validator.New()
	if err := validate.Struct(request); err != nil {
		errors := helper.FormatValidationError(err)
		response := helper.APIResponse(errors[0], http.StatusBadRequest, "error", nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	err := h.service.UpdateInstructionCategory(plant.UpdateInstructionCategoryInput{
		UserPlantID:         request.UserPlantID,
		InstructionCategory: request.InstructionCategory,
	})
	if err != nil {
		response := helper.APIResponse(err.Error(), http.StatusBadRequest, "error", nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	response := helper.APIResponse("Instruction category updated successfully", http.StatusOK, "success", nil)
	return c.JSON(http.StatusOK, response)
}

func (h *UserPlantHandler) AddUserPlant(c echo.Context) error {
	var request plant.AddUserPlantInput
	if err := c.Bind(&request); err != nil {
		response := helper.APIResponse("Invalid request", http.StatusBadRequest, "error", nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	validate := validator.New()
	if err := validate.Struct(request); err != nil {
		errors := helper.FormatValidationError(err)
		response := helper.APIResponse(errors[0], http.StatusBadRequest, "error", nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	userID, ok := c.Get("user_id").(uint)
	if !ok {
		response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
		return c.JSON(http.StatusUnauthorized, response)
	}

	request.UserID = int(userID)

	// Check if the plant already exists for the user
	exists, err := h.service.CheckUserPlantExistsForAdd(request.UserID, request.PlantID)
	if err != nil {
		response := helper.APIResponse("Failed to add plant to user", http.StatusInternalServerError, "error", nil)
		return c.JSON(http.StatusInternalServerError, response)
	}
	if exists {
		response := helper.APIResponse("Plant already exists for the user", http.StatusBadRequest, "error", nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	userPlant, err := h.service.AddUserPlant(request)
	if err != nil {
		response := helper.APIResponse("Failed to add plant to user", http.StatusInternalServerError, "error", nil)
		return c.JSON(http.StatusInternalServerError, response)
	}

	response := helper.APIResponse("Plant added to user successfully", http.StatusCreated, "success", userPlant)
	return c.JSON(http.StatusCreated, response)
}

func (h *UserPlantHandler) GetUserPlants(c echo.Context) error {
	userID, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		response := helper.APIResponse("Invalid user ID", http.StatusBadRequest, "error", nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	limitStr := c.QueryParam("limit")
	pageStr := c.QueryParam("page")

	var limit, page int

	if limitStr == "" {
		limit = -1
	} else {
		limit, err = strconv.Atoi(limitStr)
		if err != nil || limit <= 0 {
			response := helper.APIResponse("Invalid limit parameter", http.StatusBadRequest, "error", nil)
			return c.JSON(http.StatusBadRequest, response)
		}
	}

	if pageStr == "" {
		page = 1
	} else {
		page, err = strconv.Atoi(pageStr)
		if err != nil || page <= 0 {
			response := helper.APIResponse("Invalid page parameter", http.StatusBadRequest, "error", nil)
			return c.JSON(http.StatusBadRequest, response)
		}
	}

	totalCount, err := h.service.CountByUserID(userID)
	if err != nil {
		response := helper.APIResponse("Failed to count user plants", http.StatusInternalServerError, "error", nil)
		return c.JSON(http.StatusInternalServerError, response)
	}

	if limit > 0 && int64((page-1)*limit) >= totalCount {
		response := helper.APIResponse("Page exceeds available data", http.StatusBadRequest, "error", nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	userPlantResponses, err := h.service.GetUserPlantsByUserID(userID, limit, page)
	if err != nil {
		response := helper.APIResponse("Failed to fetch user plants", http.StatusInternalServerError, "error", nil)
		return c.JSON(http.StatusInternalServerError, response)
	}

	var plants []plant.UserPlantResponse
	for _, userPlants := range userPlantResponses {
		for _, userPlant := range userPlants {
			plants = append(plants, userPlant)
		}
	}

	responseData := struct {
		UserID     int                       `json:"user_id"`
		Plants     []plant.UserPlantResponse `json:"plants"`
		Limit      int                       `json:"limit"`
		Page       int                       `json:"page"`
		TotalCount int64                     `json:"total_count"`
		TotalPages int                       `json:"total_pages"`
	}{
		UserID:     userID,
		Plants:     plants,
		Limit:      limit,
		Page:       page,
		TotalCount: totalCount,
		TotalPages: int((totalCount + int64(limit) - 1) / int64(limit)),
	}

	response := helper.APIResponse("User plants fetched successfully", http.StatusOK, "success", responseData)
	return c.JSON(http.StatusOK, response)
}

func (h *UserPlantHandler) DeleteUserPlantByID(c echo.Context) error {
	userPlantID, err := strconv.Atoi(c.Param("user_plant_id"))
	if err != nil {
		response := helper.APIResponse("Invalid user plant ID", http.StatusBadRequest, "error", nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	userID := c.Get("user_id").(uint)

	userPlant, err := h.service.GetUserPlantByID(userPlantID)
	if err != nil {
		response := helper.APIResponse("User plant not found", http.StatusNotFound, "error", nil)
		return c.JSON(http.StatusNotFound, response)
	}

	if userPlant.UserID != int(userID) {
		response := helper.APIResponse("You do not have permission to delete this plant", http.StatusForbidden, "error", nil)
		return c.JSON(http.StatusForbidden, response)
	}

	deletedUserPlant, err := h.service.DeleteUserPlantByID(userPlantID)
	if err != nil {
		response := helper.APIResponse("Failed to delete user plant", http.StatusInternalServerError, "error", nil)
		return c.JSON(http.StatusInternalServerError, response)
	}

	response := helper.APIResponse("User plant deleted successfully", http.StatusOK, "success", deletedUserPlant)
	return c.JSON(http.StatusOK, response)
}

func (h *UserPlantHandler) AddUserPlantHistory(c echo.Context) error {
	var input plant.UserPlantHistoryInput
	if err := c.Bind(&input); err != nil {
		response := helper.APIResponse("Invalid request", http.StatusBadRequest, "error", nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	userID, ok := c.Get("user_id").(uint)
	if !ok {
		response := helper.APIResponse("Invalid user ID from token", http.StatusInternalServerError, "error", nil)
		return c.JSON(http.StatusInternalServerError, response)
	}
	input.UserID = int(userID)

	if input.PlantID == 0 {
		response := helper.APIResponse("Plant ID is required", http.StatusBadRequest, "error", nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	plantExists, err := h.service.CheckPlantExists(input.PlantID)
	if err != nil {
		response := helper.APIResponse("Failed to verify plant ID", http.StatusInternalServerError, "error", nil)
		return c.JSON(http.StatusInternalServerError, response)
	}
	if !plantExists {
		response := helper.APIResponse("Plant ID does not exist", http.StatusBadRequest, "error", nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	history, err := h.service.AddUserPlantHistory(input)
	if err != nil {
		response := helper.APIResponse("Failed to create user plant history", http.StatusInternalServerError, "error", nil)
		return c.JSON(http.StatusInternalServerError, response)
	}

	response := helper.APIResponse("User plant history created successfully", http.StatusCreated, "success", history)
	return c.JSON(http.StatusCreated, response)
}

func (h *UserPlantHandler) GetUserPlantHistoryByUserID(c echo.Context) error {
	userID, ok := c.Get("user_id").(uint)
	if !ok {
		response := helper.APIResponse("Invalid user ID from token", http.StatusInternalServerError, "error", nil)
		return c.JSON(http.StatusInternalServerError, response)
	}

	history, err := h.service.GetUserPlantHistoryByUserID(int(userID))
	if err != nil {
		response := helper.APIResponse("Failed to get user plant history", http.StatusInternalServerError, "error", nil)
		return c.JSON(http.StatusInternalServerError, response)
	}

	response := helper.APIResponse("User plant history fetched successfully", http.StatusOK, "success", history)
	return c.JSON(http.StatusOK, response)
}

func (h *UserPlantHandler) UpdateCustomizeName(c echo.Context) error {
	var input struct {
		CustomizeName string `json:"customize_name" form:"customize_name" validate:"required"`
	}

	if err := c.Bind(&input); err != nil {
		response := helper.APIResponse("Invalid request", http.StatusBadRequest, "error", nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	userPlantID, err := strconv.Atoi(c.Param("userPlantID"))
	if err != nil {
		response := helper.APIResponse("Invalid user plant ID", http.StatusBadRequest, "error", nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	exists, err := h.service.CheckUserPlantExists(userPlantID)
	if err != nil {
		response := helper.APIResponse("Error checking user plant existence", http.StatusInternalServerError, "error", nil)
		return c.JSON(http.StatusInternalServerError, response)
	}

	if !exists {
		response := helper.APIResponse("User plant ID not found", http.StatusNotFound, "error", nil)
		return c.JSON(http.StatusNotFound, response)
	}

	updatedPlant, err := h.service.UpdateCustomizeName(userPlantID, input.CustomizeName)
	if err != nil {
		response := helper.APIResponse("Failed to update customize name", http.StatusInternalServerError, "error", nil)
		return c.JSON(http.StatusInternalServerError, response)
	}

	response := helper.APIResponse("Customize name updated successfully", http.StatusOK, "success", updatedPlant)
	return c.JSON(http.StatusOK, response)
}
