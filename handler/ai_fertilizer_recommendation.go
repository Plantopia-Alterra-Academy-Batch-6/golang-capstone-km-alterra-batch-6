package handler

import (
	"net/http"

	"github.com/OctavianoRyan25/be-agriculture/modules/ai"
	"github.com/OctavianoRyan25/be-agriculture/utils/helper"
	"github.com/labstack/echo/v4"
)

type AIFertilizerRecommendationHandler struct {
	Service ai.AIFertilizerRecommendationService
}

func NewAIFertilizerRecommendationHandler(service ai.AIFertilizerRecommendationService) *AIFertilizerRecommendationHandler {
	return &AIFertilizerRecommendationHandler{Service: service}
}

func (h *AIFertilizerRecommendationHandler) GetFertilizerRecommendation(c echo.Context) error {
	plantName := c.QueryParam("plant_name")
	if plantName == "" {
		return c.JSON(http.StatusBadRequest, helper.APIResponse("plant_name is required", http.StatusBadRequest, "error", nil))
	}

	recommendations, err := h.Service.GetFertilizerRecommendation(plantName)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.APIResponse("Failed to get recommendation", http.StatusInternalServerError, "error", err.Error()))
	}

	if len(recommendations) == 0 {
		return c.JSON(http.StatusInternalServerError, helper.APIResponse("No valid recommendations found", http.StatusInternalServerError, "error", nil))
	}

	var responseData []map[string]interface{}
	for _, rec := range recommendations {
		responseData = append(responseData, map[string]interface{}{
			"plant_name":     rec.PlantName,
			"recommendation": rec.Recommendation,
		})
	}

	return c.JSON(http.StatusOK, helper.APIResponse("Recommendation retrieved successfully", http.StatusOK, "success", responseData))
}

func (h *AIFertilizerRecommendationHandler) GetPlantingRecommendation(c echo.Context) error {
	plantName := c.QueryParam("plant_name")
	if plantName == "" {
		return c.JSON(http.StatusBadRequest, helper.APIResponse("plant_name is required", http.StatusBadRequest, "error", nil))
	}

	recommendations, err := h.Service.GetPlantingRecommendation(plantName)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.APIResponse("Failed to get recommendation", http.StatusInternalServerError, "error", err.Error()))
	}

	if len(recommendations) == 0 {
		return c.JSON(http.StatusInternalServerError, helper.APIResponse("No valid recommendations found", http.StatusInternalServerError, "error", nil))
	}

	var responseData []map[string]interface{}
	for _, rec := range recommendations {
		responseData = append(responseData, map[string]interface{}{
			"plant_name":     rec.PlantName,
			"recommendation": rec.Recommendation,
		})
	}

	return c.JSON(http.StatusOK, helper.APIResponse("Recommendation retrieved successfully", http.StatusOK, "success", responseData))
}
