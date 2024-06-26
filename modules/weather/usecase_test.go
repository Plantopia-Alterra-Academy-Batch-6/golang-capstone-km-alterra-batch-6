package weather

import (
	"testing"
	"time"

	"github.com/go-resty/resty/v2"
)

// MockRestyClient implements a mock of the Resty client for testing purposes.
type MockRestyClient struct{}

func (m *MockRestyClient) R() *resty.Request {
	return nil // Implement as needed for testing, or mock further.
}

func TestGetCurrentWeatherByCoordinates(t *testing.T) {
	service := &weatherService{}

	lat := 37.7749
	lon := -122.4194

	weather, err := service.GetCurrentWeatherByCoordinates(lat, lon)

	if err != nil {
		t.Errorf("Error fetching current weather: %v", err)
	}

	if weather == nil {
		t.Error("Expected weather object, got nil")
	}

	if weather.City == "" {
		t.Error("Expected non-empty city name")
	}

	if weather.CreatedAt.After(time.Now()) {
		t.Error("CreatedAt should not be in the future")
	}
}

func TestGetHourlyWeatherByCoordinates(t *testing.T) {
	service := &weatherService{}

	lat := 37.7749
	lon := -122.4194

	hourlyWeather, err := service.GetHourlyWeatherByCoordinates(lat, lon)

	if err != nil {
		t.Errorf("Error fetching hourly weather: %v", err)
	}

	if hourlyWeather == nil || len(hourlyWeather) == 0 {
		t.Error("Expected hourly weather data, got nil or empty slice")
	}

	for _, hw := range hourlyWeather {
		if hw.Timestamp.IsZero() {
			t.Error("Expected non-zero timestamp for hourly weather")
		}
	}
}

func TestGetDailyWeatherByCoordinates(t *testing.T) {
	service := &weatherService{}

	lat := 37.7749
	lon := -122.4194

	dailyWeather, err := service.GetDailyWeatherByCoordinates(lat, lon)

	if err != nil {
		t.Errorf("Error fetching daily weather: %v", err)
	}

	if dailyWeather == nil || len(dailyWeather) == 0 {
		t.Error("Expected daily weather data, got nil or empty slice")
	}

	for _, dw := range dailyWeather {
		if dw.Date.IsZero() {
			t.Error("Expected non-zero date for daily weather")
		}
	}
}


