package wateringhistory

import (
	"time"

	"github.com/OctavianoRyan25/be-agriculture/modules/plant"
	"github.com/OctavianoRyan25/be-agriculture/modules/user"
)

type WateringHistory struct {
	ID        int `gorm:"primaryKey"`
	PlantID   int
	Plant     plant.Plant `gorm:"foreignKey:PlantID;references:ID"`
	UserID    int
	User      user.User `gorm:"foreignKey:UserID;references:ID"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Notification struct {
	Id        int `gorm:"primaryKey"`
	Title     string
	Body      string
	UserId    int         `gorm:"foreignKey:UserID;references:ID"`
	PlantId   int         `gorm:"foreignKey:PlantId;references:ID"`
	Plant     plant.Plant `gorm:"foreignKey:PlantId;references:ID"`
	IsRead    bool
	CreatedAt time.Time
	UpdatedAt time.Time
}
