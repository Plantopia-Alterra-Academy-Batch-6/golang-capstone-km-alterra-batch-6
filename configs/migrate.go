package configs

import (
	"github.com/OctavianoRyan25/be-agriculture/modules/admin"
	"github.com/OctavianoRyan25/be-agriculture/modules/article"
	"github.com/OctavianoRyan25/be-agriculture/modules/fertilizer"
	"github.com/OctavianoRyan25/be-agriculture/modules/notification"
	"github.com/OctavianoRyan25/be-agriculture/modules/plant"
	"github.com/OctavianoRyan25/be-agriculture/modules/user"
	wateringhistory "github.com/OctavianoRyan25/be-agriculture/modules/watering_history"
	"gorm.io/gorm"
)

func AutoMigrate(db *gorm.DB) error {
	if err := db.AutoMigrate(&user.User{}, &admin.Admin{}, &plant.PlantCategory{}, &plant.Plant{}, &plant.PlantImage{}, &plant.PlantInstruction{}, &plant.PlantFAQ{}, &plant.PlantReminder{}, &plant.PlantCharacteristic{}, &plant.UserPlant{}, &plant.PlantInstructionCategory{}, &plant.PlantProgress{}, &notification.Notification{}, &notification.CustomizeWateringReminder{}, &wateringhistory.WateringHistory{}, &plant.UserPlantHistory{}, &fertilizer.Fertilizer{}, &plant.PlantEarliestWatering{}, &article.Article{}); err != nil {
		return err
	}
	return nil
}
