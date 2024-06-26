package notification

import (
	"time"

	"github.com/OctavianoRyan25/be-agriculture/modules/plant"
	"github.com/OctavianoRyan25/be-agriculture/modules/user"
)

type Notification struct {
	Id        int `gorm:"primaryKey"`
	Title     string
	Body      string
	UserId    int `gorm:"foreignKey:UserID;references:Id"`
	PlantId   int `gorm:"foreignKey:PlantId;references:Id"`
	IsRead    bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

type CustomizeWateringReminder struct {
	Id        int `gorm:"primaryKey"`
	UserId    int `gorm:"foreignKey:UserID;references:ID"`
	User      user.User
	PlantId   int `gorm:"foreignKey:PlantId;references:ID"`
	Plant     plant.Plant
	Time      string
	Recurring bool
	Type      string // "daily", "weekly", "monthly", "yearly
	CreatedAt time.Time
	UpdatedAt time.Time
}
