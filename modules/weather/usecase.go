package weather

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/go-resty/resty/v2"
)

type WeatherService interface {
	GetCurrentWeatherByCoordinates(lat, lon float64) (*Weather, error)
	GetHourlyWeatherByCoordinates(lat, lon float64) ([]HourlyWeather, error)
	GetDailyWeatherByCoordinates(lat, lon float64) ([]DailyWeather, error)
}

type weatherService struct{}

func NewWeatherService() WeatherService {
	return &weatherService{}
}

func (s *weatherService) GetCurrentWeatherByCoordinates(lat, lon float64) (*Weather, error) {
	client := resty.New()
	apiKey := os.Getenv("OPENWEATHER_API_KEY")
	resp, err := client.R().
			SetQueryParams(map[string]string{
					"lat":   fmt.Sprintf("%f", lat),
					"lon":   fmt.Sprintf("%f", lon),
					"appid": apiKey,
					"units": "metric",
			}).
			Get("http://api.openweathermap.org/data/2.5/weather")

	if err != nil {
			return nil, err
	}

	var apiResponse map[string]interface{}
	if err := json.Unmarshal(resp.Body(), &apiResponse); err != nil {
			return nil, err
	}

	main := apiResponse["main"].(map[string]interface{})
	weatherDesc := apiResponse["weather"].([]interface{})[0].(map[string]interface{})
	wind := apiResponse["wind"].(map[string]interface{})
	sys := apiResponse["sys"].(map[string]interface{})
	timezoneOffset := int64(apiResponse["timezone"].(float64)) // Offset in seconds

	sunriseUnix := int64(sys["sunrise"].(float64))
	localSunrise := time.Unix(sunriseUnix, 0).UTC().Add(time.Duration(timezoneOffset) * time.Second)

	weather := &Weather{
			ID:          1,
			City:        apiResponse["name"].(string),
			Temperature: main["temp"].(float64),
			RealFeel:    main["feels_like"].(float64),
			Pressure:    int(main["pressure"].(float64)),
			Humidity:    int(main["humidity"].(float64)),
			WindSpeed:   wind["speed"].(float64),
			Main:        weatherDesc["main"].(string),
			Description: weatherDesc["description"].(string),
			Icon:        weatherDesc["icon"].(string),
			Sunrise:     localSunrise.Unix(),
			CreatedAt:   time.Now(),
	}

	return weather, nil
}


func (s *weatherService) GetHourlyWeatherByCoordinates(lat, lon float64) ([]HourlyWeather, error) {
	client := resty.New()
	apiKey := os.Getenv("OPENWEATHER_API_KEY")

	resp, err := client.R().
			SetQueryParams(map[string]string{
					"lat":   fmt.Sprintf("%f", lat),
					"lon":   fmt.Sprintf("%f", lon),
					"appid": apiKey,
					"units": "metric",
			}).
			Get("https://pro.openweathermap.org/data/2.5/forecast/hourly")

	if err != nil {
			return nil, err
	}

	var apiResponse map[string]interface{}
	if err := json.Unmarshal(resp.Body(), &apiResponse); err != nil {
			return nil, err
	}

	if cod, ok := apiResponse["cod"].(string); ok && cod != "200" {
			return nil, fmt.Errorf("API returned non-200 status code: %s", cod)
	}

	cityInfo := apiResponse["city"].(map[string]interface{})
	timezoneOffset := cityInfo["timezone"].(float64) 

	list := apiResponse["list"].([]interface{})
	var hourlyWeathers []HourlyWeather

	for i, item := range list {
			forecastItem := item.(map[string]interface{})
			main := forecastItem["main"].(map[string]interface{})
			weatherDesc := forecastItem["weather"].([]interface{})[0].(map[string]interface{})
			wind := forecastItem["wind"].(map[string]interface{})

			timestamp := int64(forecastItem["dt"].(float64))
			localTime := time.Unix(timestamp, 0).UTC().Add(time.Duration(timezoneOffset) * time.Second)

			hourlyWeather := HourlyWeather{
					ID:          uint(i + 1),
					City:        cityInfo["name"].(string),
					Timestamp:   localTime,
					Temperature: main["temp"].(float64),
					RealFeel:    main["feels_like"].(float64),
					Pressure:    int(main["pressure"].(float64)),
					Humidity:    int(main["humidity"].(float64)),
					WindSpeed:   wind["speed"].(float64),
					Main:        weatherDesc["main"].(string),
					Description: weatherDesc["description"].(string),
					Icon:        weatherDesc["icon"].(string),
			}

			hourlyWeathers = append(hourlyWeathers, hourlyWeather)
	}

	return hourlyWeathers, nil
}

func (s *weatherService) GetDailyWeatherByCoordinates(lat, lon float64) ([]DailyWeather, error) {
	client := resty.New()
	apiKey := os.Getenv("OPENWEATHER_API_KEY")

	resp, err := client.R().
			SetQueryParams(map[string]string{
					"lat":   fmt.Sprintf("%f", lat),
					"lon":   fmt.Sprintf("%f", lon),
					"appid": apiKey,
					"units": "metric",
					"cnt":   "7",
			}).
			Get(fmt.Sprintf("http://api.openweathermap.org/data/2.5/forecast/daily?lat=%f&lon=%f&cnt=7&appid=%s", lat, lon, apiKey))

	if err != nil {
			return nil, err
	}

	var apiResponse map[string]interface{}
	if err := json.Unmarshal(resp.Body(), &apiResponse); err != nil {
			return nil, err
	}

	if cod, ok := apiResponse["cod"].(string); ok && cod != "200" {
			return nil, fmt.Errorf("API returned non-200 status code: %s", cod)
	}

	cityInfo := apiResponse["city"].(map[string]interface{})
	timezoneOffset := int64(cityInfo["timezone"].(float64)) // Offset in seconds

	list := apiResponse["list"].([]interface{})
	var dailyWeathers []DailyWeather

	for i, item := range list {
			forecastItem := item.(map[string]interface{})
			temp := forecastItem["temp"].(map[string]interface{})
			feelsLike := forecastItem["feels_like"].(map[string]interface{})
			weatherDesc := forecastItem["weather"].([]interface{})[0].(map[string]interface{})

			dt := int64(forecastItem["dt"].(float64))
			localDate := time.Unix(dt, 0).UTC().Add(time.Duration(timezoneOffset) * time.Second)

			dailyWeather := DailyWeather{
					ID:          uint(i + 1),
					City:        cityInfo["name"].(string),
					Date:        localDate,
					Temperature: temp["day"].(float64),
					RealFeel:    feelsLike["day"].(float64),
					Pressure:    int(forecastItem["pressure"].(float64)),
					Humidity:    int(forecastItem["humidity"].(float64)),
					WindSpeed:   forecastItem["speed"].(float64),
					Main:        weatherDesc["main"].(string),
					Description: weatherDesc["description"].(string),
					Icon:        weatherDesc["icon"].(string),
			}

			dailyWeathers = append(dailyWeathers, dailyWeather)
	}

	return dailyWeathers, nil
}
