package handler

import (
	"net/http"

	"strconv"

	"github.com/OctavianoRyan25/be-agriculture/modules/fertilizer"
	"github.com/OctavianoRyan25/be-agriculture/utils/helper"
	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type FertilizerHandler struct {
	service    fertilizer.FertilizerService
	cloudinary *cloudinary.Cloudinary
}

func NewFertilizerHandler(service fertilizer.FertilizerService, cloudinary *cloudinary.Cloudinary) *FertilizerHandler {
	return &FertilizerHandler{service, cloudinary}
}

func (h *FertilizerHandler) GetFertilizer(c echo.Context) error {
	categories, err := h.service.GetFertilizer()

	if err != nil {
		response := helper.APIResponse("Failed to get fertilizer", http.StatusInternalServerError, "error", nil)
		return c.JSON(http.StatusInternalServerError, response)
	}

	response := helper.APIResponse("Fertilizer fetched successfully", http.StatusOK, "success", categories)
	return c.JSON(http.StatusOK, response)
}

func (h *FertilizerHandler) GetFertilizerById(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("Id"))

	if err != nil {
		response := helper.APIResponse("Invalid ID", http.StatusBadRequest, "error", nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	category, err := h.service.GetFertilizerByID(id)
	if err != nil {
		response := helper.APIResponse("Fertilizer not found", http.StatusNotFound, "error", nil)
		return c.JSON(http.StatusNotFound, response)
	}

	response := helper.APIResponse("Fertilizer fetched successfully", http.StatusOK, "success", category)
	return c.JSON(http.StatusOK, response)
}

func (h *FertilizerHandler) CreateFertilizer(c echo.Context) error {

	var input fertilizer.FertilizerInput

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

	//  fertilizer, err = h.service.CreateFertilizer(input)
	// if err != nil {
	// 	response := helper.APIResponse("Failed to create fertilizer", http.StatusInternalServerError, "error", nil)
	// 	return c.JSON(http.StatusInternalServerError, response)
	// }

	response := helper.APIResponse("Fertilizer created successfully", http.StatusCreated, "success", nil)
	return c.JSON(http.StatusCreated, response)
}

// Update a Fertilizer by Id
func (h *FertilizerHandler) UpdateFertilizer(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("Id"))

	if err != nil {
		response := helper.APIResponse("Invalid ID", http.StatusBadRequest, "error", nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	role := c.Get("role").(string)
	if role != "admin" {
		response := helper.APIResponse("Only admin can update", http.StatusUnauthorized, "error", nil)
		return c.JSON(http.StatusUnauthorized, response)
	}

	var input fertilizer.FertilizerInput

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

	category, err := h.service.UpdateFertilizer(id, input)
	if err != nil {
		response := helper.APIResponse("Failed to update fertilizer", http.StatusInternalServerError, "error", nil)
		return c.JSON(http.StatusInternalServerError, response)
	}

	response := helper.APIResponse("Fertilizer updated successfully", http.StatusOK, "success", category)
	return c.JSON(http.StatusOK, response)
}

// Delete a Fertilizer by Id
func (h *FertilizerHandler) DeleteFertilizer(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("Id"))

	if err != nil {
		response := helper.APIResponse("Invalid ID", http.StatusBadRequest, "error", nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	role := c.Get("role").(string)
	if role != "admin" {
		response := helper.APIResponse("Only admin can delete", http.StatusUnauthorized, "error", nil)
		return c.JSON(http.StatusUnauthorized, response)
	}

	category, err := h.service.GetFertilizerByID(id)
	if err != nil {
		response := helper.APIResponse("Fertilizer not found", http.StatusNotFound, "error", nil)
		return c.JSON(http.StatusNotFound, response)
	}

	err = h.service.DeleteFertilizer(id)
	if err != nil {
		response := helper.APIResponse("Failed to delete fertilizer", http.StatusInternalServerError, "error", nil)
		return c.JSON(http.StatusInternalServerError, response)
	}

	response := helper.APIResponse("Fertilizer deleted successfully", http.StatusOK, "success", category)
	return c.JSON(http.StatusOK, response)
}
