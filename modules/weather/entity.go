package weather

import (
	"time"
)

type Weather struct {
	ID          uint      `gorm:"primary_key" json:"id"`
	City        string    `json:"city"`
	Temperature float64   `json:"temperature"`
	RealFeel	 float64   `json:"real_feel"`
	Pressure    int       `json:"pressure"`
	Humidity    int       `json:"humidity"`
	WindSpeed   float64   `json:"wind_speed"`
	Main        string    `json:"main"`
	Description string    `json:"description"`
	Icon        string    `json:"icon"`
	Sunrise     int64 `json:"sunrise"`
	CreatedAt   time.Time `json:"created_at"`
}

type HourlyWeather struct {
	ID          uint      `gorm:"primary_key" json:"id"`
	City        string    `json:"city"`
	Temperature float64   `json:"temperature"`
	RealFeel	 float64   `json:"real_feel"`
	Pressure    int       `json:"pressure"`
	Humidity    int       `json:"humidity"`
	WindSpeed   float64   `json:"wind_speed"`
	Main        string    `json:"main"`
	Description string    `json:"description"`
	Icon        string    `json:"icon"`
	Timestamp   time.Time `json:"timestamp"`
}

type DailyWeather struct {
	ID          uint      `gorm:"primary_key" json:"id"`
	City        string    `json:"city"`
	Temperature float64   `json:"temperature"`
	RealFeel	 float64   `json:"real_feel"`
	Pressure    int       `json:"pressure"`
	Humidity    int       `json:"humidity"`
	WindSpeed   float64   `json:"wind_speed"`
	Main        string    `json:"main"`
	Description string    `json:"description"`
	Sunrise     int64 `json:"sunrise"`
	Icon        string    `json:"icon"`
	Date        time.Time `json:"date"`
}
