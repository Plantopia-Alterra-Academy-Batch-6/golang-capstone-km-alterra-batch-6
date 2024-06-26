package handler

import (
	"net/http"
	"strconv"

	"github.com/OctavianoRyan25/be-agriculture/modules/weather"
	"github.com/OctavianoRyan25/be-agriculture/utils/helper"
	"github.com/labstack/echo/v4"
)

type WeatherHandler struct {
	Service weather.WeatherService
}

func NewWeatherHandler(service weather.WeatherService) *WeatherHandler {
	return &WeatherHandler{Service: service}
}

func (h *WeatherHandler) GetCurrentWeather(c echo.Context) error {
	latStr := c.QueryParam("lat")
	lonStr := c.QueryParam("lon")

	lat, err := strconv.ParseFloat(latStr, 64)
	if err != nil {
		response := helper.APIResponse("Invalid latitude", http.StatusBadRequest, "error", nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	lon, err := strconv.ParseFloat(lonStr, 64)
	if err != nil {
		response := helper.APIResponse("Invalid longitude", http.StatusBadRequest, "error", nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	weather, err := h.Service.GetCurrentWeatherByCoordinates(lat, lon)
	if err != nil {
		response := helper.APIResponse("Failed to get current weather", http.StatusInternalServerError, "error", nil)
		return c.JSON(http.StatusInternalServerError, response)
	}

	response := helper.APIResponse("Current weather data", http.StatusOK, "success", weather)
	return c.JSON(http.StatusOK, response)
}

func (h *WeatherHandler) GetHourlyWeather(c echo.Context) error {
	latStr := c.QueryParam("lat")
	lonStr := c.QueryParam("lon")

	lat, err := strconv.ParseFloat(latStr, 64)
	if err != nil {
			response := helper.APIResponse("Invalid latitude", http.StatusBadRequest, "error", nil)
			return c.JSON(http.StatusBadRequest, response)
	}

	lon, err := strconv.ParseFloat(lonStr, 64)
	if err != nil {
			response := helper.APIResponse("Invalid longitude", http.StatusBadRequest, "error", nil)
			return c.JSON(http.StatusBadRequest, response)
	}

	hourlyWeather, err := h.Service.GetHourlyWeatherByCoordinates(lat, lon)
	if err != nil {
			response := helper.APIResponse("Failed to get hourly weather", http.StatusInternalServerError, "error", nil)
			return c.JSON(http.StatusInternalServerError, response)
	}

	response := helper.APIResponse("Hourly weather data", http.StatusOK, "success", hourlyWeather)
	return c.JSON(http.StatusOK, response)
}

func (h *WeatherHandler) GetDailyWeather(c echo.Context) error {
	latStr := c.QueryParam("lat")
	lonStr := c.QueryParam("lon")

	lat, err := strconv.ParseFloat(latStr, 64)
	if err != nil {
			response := helper.APIResponse("Invalid latitude", http.StatusBadRequest, "error", nil)
			return c.JSON(http.StatusBadRequest, response)
	}

	lon, err := strconv.ParseFloat(lonStr, 64)
	if err != nil {
			response := helper.APIResponse("Invalid longitude", http.StatusBadRequest, "error", nil)
			return c.JSON(http.StatusBadRequest, response)
	}

	dailyWeather, err := h.Service.GetDailyWeatherByCoordinates(lat, lon)
	if err != nil {
			response := helper.APIResponse("Failed to get daily weather", http.StatusInternalServerError, "error", nil)
			return c.JSON(http.StatusInternalServerError, response)
	}

	response := helper.APIResponse("Daily weather data", http.StatusOK, "success", dailyWeather)
	return c.JSON(http.StatusOK, response)
}