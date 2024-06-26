// handler.go
package handler

import (
	"net/http"

	"github.com/OctavianoRyan25/be-agriculture/modules/plant"
	"github.com/OctavianoRyan25/be-agriculture/utils/helper"
	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/labstack/echo/v4"
)

type PlantEarliestWateringHandler struct {
	service    plant.PlantEarliestWateringService
	cloudinary *cloudinary.Cloudinary
}

func NewPlantEarliestWateringHandler(service plant.PlantEarliestWateringService, cloudinary *cloudinary.Cloudinary) *PlantEarliestWateringHandler {
	return &PlantEarliestWateringHandler{service, cloudinary}
}

func (h *PlantEarliestWateringHandler) GetEarliestWateringTime(c echo.Context) error {

    schedule, err := h.service.FindEarliestWateringTime()
	if err != nil {
		response := helper.APIResponse("Schedule not found", http.StatusNotFound, "error", nil)
		return c.JSON(http.StatusNotFound, response)
	}
	response := helper.APIResponse("Schedule fetched successfully", http.StatusOK, "success", schedule)
	return c.JSON(http.StatusOK, response)
}
