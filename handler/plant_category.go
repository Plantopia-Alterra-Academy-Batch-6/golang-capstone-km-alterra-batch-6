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

type PlantCategoryHandler struct {
	service plant.PlantCategoryService
	cloudinary  *cloudinary.Cloudinary
}

func NewPlantCategoryHandler(service plant.PlantCategoryService, cloudinary  *cloudinary.Cloudinary) *PlantCategoryHandler {
	return &PlantCategoryHandler{service , cloudinary}
}

func (h *PlantCategoryHandler) GetAll(c echo.Context) error {
	categories, err := h.service.FindAll()
	
	if err != nil {
		response := helper.APIResponse("Failed to get plant categories", http.StatusInternalServerError, "error", nil)
		return c.JSON(http.StatusInternalServerError, response)
	}

	response := helper.APIResponse("Plant categories fetched successfully", http.StatusOK, "success", categories)
	return c.JSON(http.StatusOK, response)
}

func (h *PlantCategoryHandler) GetByID(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		response := helper.APIResponse("Invalid ID", http.StatusBadRequest, "error", nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	category, err := h.service.FindByID(id)
	if err != nil {
		response := helper.APIResponse("Plant category not found", http.StatusNotFound, "error", nil)
		return c.JSON(http.StatusNotFound, response)
	}

	response := helper.APIResponse("Plant category fetched successfully", http.StatusOK, "success", category)
	return c.JSON(http.StatusOK, response)
}

func (h *PlantCategoryHandler) Create(c echo.Context) error {
	var input plant.PlantCategoryClimateInput

	if err := c.Bind(&input); err != nil {
		response := helper.APIResponse("Failed to bind input", http.StatusBadRequest, "error", nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	role := c.Get("role").(string)
	if role != "admin" {
		response := helper.APIResponse("Only admin can create", http.StatusUnauthorized, "error", nil)
		return c.JSON(http.StatusUnauthorized, response)
	}

	validate := validator.New()
	if err := validate.Struct(input); err != nil {
		errors := helper.FormatValidationError(err)
		response := helper.APIResponse("Validation error", http.StatusBadRequest, "error", errors)
		return c.JSON(http.StatusBadRequest, response)
	}

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
		response := helper.APIResponse("Failed to create plant category", http.StatusInternalServerError, "error", nil)
		return c.JSON(http.StatusInternalServerError, response)
	}

	response := helper.APIResponse("Plant category created successfully", http.StatusCreated, "success", nil)
	return c.JSON(http.StatusCreated, response)
}

func (h *PlantCategoryHandler) Update(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		response := helper.APIResponse("Invalid ID", http.StatusBadRequest, "error", nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	role := c.Get("role").(string)
	if role != "admin" {
		response := helper.APIResponse("Only admin can update", http.StatusUnauthorized, "error", nil)
		return c.JSON(http.StatusUnauthorized, response)
	}

	var input plant.PlantCategoryClimateInput

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

	// Handle file upload if a new file is provided
	file, err := c.FormFile("image_url")
	if err != nil && err != http.ErrMissingFile {
		response := helper.APIResponse("Failed to get uploaded image", http.StatusBadRequest, "error", nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	const maxFileSize = 2 * 1024 * 1024
	if file.Size > maxFileSize {
		response := helper.APIResponse("File size exceeds 2MB", http.StatusBadRequest, "error", nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	var imageURL string
	if file != nil {
		fileReader, err := file.Open()
		if err != nil {
			response := helper.APIResponse("Failed to open uploaded file", http.StatusInternalServerError, "error", nil)
			return c.JSON(http.StatusInternalServerError, response)
		}
		defer fileReader.Close()

		params := uploader.UploadParams{
			Folder: "be-agriculture",
		}

		uploadResult, err := h.cloudinary.Upload.Upload(context.Background(), fileReader, params)
		if err != nil {
			response := helper.APIResponse("Failed to upload image to Cloudinary", http.StatusInternalServerError, "error", nil)
			return c.JSON(http.StatusInternalServerError, response)
		}

		imageURL = uploadResult.SecureURL
	}

	// Pass the image URL only if a new image was uploaded
	category, err := h.service.Update(id, input, imageURL)
	if err != nil {
		response := helper.APIResponse("Failed to update plant category", http.StatusInternalServerError, "error", nil)
		return c.JSON(http.StatusInternalServerError, response)
	}

	response := helper.APIResponse("Plant category updated successfully", http.StatusOK, "success", category)
	return c.JSON(http.StatusOK, response)
}


func (h *PlantCategoryHandler) Delete(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		response := helper.APIResponse("Invalid ID", http.StatusBadRequest, "error", nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	role := c.Get("role").(string)
	if role != "admin" {
		response := helper.APIResponse("Only admin can delete", http.StatusUnauthorized, "error", nil)
		return c.JSON(http.StatusUnauthorized, response)
	}

	category, err := h.service.FindByID(id)
	if err != nil {
		response := helper.APIResponse("Plant category not found", http.StatusNotFound, "error", nil)
		return c.JSON(http.StatusNotFound, response)
	}

	err = h.service.Delete(id)
	if err != nil {
		response := helper.APIResponse("Failed to delete plant category", http.StatusInternalServerError, "error", nil)
		return c.JSON(http.StatusInternalServerError, response)
	}

	response := helper.APIResponse("Plant category deleted successfully", http.StatusOK, "success", category)
	return c.JSON(http.StatusOK, response)
}

