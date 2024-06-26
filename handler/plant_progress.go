package handler

import (
	"context"
	"net/http"
	"strconv"

	"github.com/OctavianoRyan25/be-agriculture/modules/plant"
	"github.com/OctavianoRyan25/be-agriculture/utils/helper"
	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type PlantProgressHandler struct {
	service plant.PlantProgressService
	cloudinary  *cloudinary.Cloudinary
}

func NewPlantProgressHandler(service plant.PlantProgressService, cloudinary  *cloudinary.Cloudinary) *PlantProgressHandler {
	return &PlantProgressHandler{service, cloudinary}
}

func (h *PlantProgressHandler) GetAllByUserIDAndPlantID(c echo.Context) error {
	userID, ok := c.Get("user_id").(uint)
    if !ok {
        response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
        return c.JSON(http.StatusUnauthorized, response)
    }

	plantID, err := strconv.Atoi(c.Param("plant_id"))
	if err != nil {
		response := helper.APIResponse("Invalid plant ID", http.StatusBadRequest, "error", nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	plantProgress, err := h.service.FindProgressByUserIDAndPlantID(int(userID), plantID)
	if err != nil {
		response := helper.APIResponse("Failed to fetch plant progress", http.StatusInternalServerError, "error", nil)
		return c.JSON(http.StatusInternalServerError, response)
	}

	responseData := struct {
		PlantProgress []plant.PlantProgressResponse `json:"plant_progress"`
	}{
		PlantProgress: plantProgress,
	}

	response := helper.APIResponse("Plant progress fetched successfully", http.StatusOK, "success", responseData)
	return c.JSON(http.StatusOK, response)
}

func (h *PlantProgressHandler) UploadProgress(c echo.Context) error {
	var input plant.PlantProgressInput

	if err := c.Bind(&input); err != nil {
		response := helper.APIResponse("Failed to bind input", http.StatusBadRequest, "error", nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	validate := validator.New()
	if err := validate.Struct(input); err != nil {
		errors := helper.FormatValidationError(err)
		response := helper.APIResponse("Validation error", http.StatusBadRequest, "error", errors)
		return c.JSON(http.StatusBadRequest, response)
	}

	userID, ok := c.Get("user_id").(uint)
    if !ok {
        response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
        return c.JSON(http.StatusUnauthorized, response)
    }

	input.UserID = int(userID)

	file, err := c.FormFile("image_url")
	if err != nil {
		response := helper.APIResponse("Failed to get uploaded image", http.StatusBadRequest, "error", nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	const maxFileSize = 2 * 1024 * 1024
	if file.Size > maxFileSize {
		response := helper.APIResponse("File size exceeds 2MB", http.StatusBadRequest, "error", nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	fileReader, err := file.Open()
	if err != nil {
		response := helper.APIResponse("Failed to open uploaded file", http.StatusInternalServerError, "error", nil)
		return c.JSON(http.StatusInternalServerError, response)
	}
	defer fileReader.Close()

	params := uploader.UploadParams{
		Folder:    "be-agriculture",
		// Overwrite: true,
	}

	uploadResult, err := h.cloudinary.Upload.Upload(context.Background(), fileReader, params)
	if err != nil {
		response := helper.APIResponse("Failed to upload image to Cloudinary", http.StatusInternalServerError, "error", nil)
		return c.JSON(http.StatusInternalServerError, response)
	}

	imageURL := uploadResult.SecureURL

	_, err = h.service.Create(input, imageURL)
	if err != nil {
		response := helper.APIResponse("Failed to add plant progress", http.StatusInternalServerError, "error", nil)
		return c.JSON(http.StatusInternalServerError, response)
	}

	response := helper.APIResponse("Plant progress add successfully", http.StatusCreated, "success", nil)
	return c.JSON(http.StatusCreated, response)
}