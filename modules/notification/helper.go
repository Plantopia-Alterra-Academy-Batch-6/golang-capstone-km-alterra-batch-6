package notification

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/messaging"
	"github.com/OctavianoRyan25/be-agriculture/modules/plant"
	"github.com/OctavianoRyan25/be-agriculture/modules/user" // Updated import
	"github.com/robfig/cron/v3"
	"google.golang.org/api/option"
	"gorm.io/gorm"
)

// Initialize Firebase app
func InitFirebase() *firebase.App {
	// Production
	credential := os.Getenv("FIREBASE_CREDENTIAL")
	opt := option.WithCredentialsJSON([]byte(credential))
	//Development
	// opt := option.WithCredentialsFile("agriculture-af713-c7c8068614c1.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatalf("error initializing app: %v", err)
	}
	return app
}

// Send reminder notification and store it in the database
func SendReminder(user user.User, plant plant.Plant, useCase UseCase) error {
	// Simulating FCM messaging part is commented out
	app := InitFirebase()
	client, err := app.Messaging(context.Background())
	if err != nil {
		log.Fatalf("error getting FCM client: %v", err)
	}

	message := &messaging.Message{
		Token: user.FCMToken,
		Notification: &messaging.Notification{
			Title: "Watering Reminder",
			Body:  fmt.Sprintf("Hiii %s, It's time to water your plant: %s", user.Name, plant.Name),
		},
	}

	_, err = client.Send(context.Background(), message)
	if err != nil {
		log.Printf("Error sending FCM message: %v", err)
	} else {
		fmt.Printf("Reminder have pushed to %s for watering plant %s\n", user.Email, plant.Name)
	}

	// Store the notification in the database
	notification := &Notification{
		Title:     "Watering Reminder",
		Body:      fmt.Sprintf("Hiii %s, It's time to water your plant: %s", user.Name, plant.Name),
		UserId:    user.ID,
		PlantId:   plant.ID,
		IsRead:    false,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	_, err = useCase.StoreNotification(notification)
	if err != nil {
		log.Printf("Error storing notification: %v", err)
		return err
	}
	fmt.Printf("Notification stored for user %s\n", user.Email)
	return nil
}

func SendCustomReminder(reminder CustomizeWateringReminder, useCase UseCase) error {
	// Simulating FCM messaging part is commented out
	app := InitFirebase()
	client, err := app.Messaging(context.Background())
	if err != nil {
		log.Fatalf("error getting FCM client: %v", err)
	}

	user := reminder.User
	plant := reminder.Plant

	message := &messaging.Message{
		Token: user.FCMToken,
		Notification: &messaging.Notification{
			Title: "Watering Reminder",
			Body:  fmt.Sprintf("It's time to water your plant: %s", plant.Name),
		},
	}

	_, err = client.Send(context.Background(), message)
	if err != nil {
		log.Printf("Error sending FCM message: %v", err)
	} else {
		fmt.Printf("Reminder have pushed to %s for watering plant %s\n", user.Email, plant.Name)
	}

	// Store the notification in the database
	notification := &Notification{
		Title:     "Customize Watering Reminder",
		Body:      fmt.Sprintf("Hiii %s, It's time to water your plant: %s", user.Name, reminder.Plant.Name),
		UserId:    user.ID,
		PlantId:   reminder.Plant.ID,
		IsRead:    false,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	_, err = useCase.StoreNotification(notification)
	if err != nil {
		log.Printf("Error storing notification: %v", err)
		return err
	}
	fmt.Printf("Custom notification stored for user %s\n", user.Email)
	return nil
}

// Schedule watering reminders based on PlantReminder.WateringTime
func StartScheduler(db *gorm.DB, useCase UseCase) {
	// Define desired location for time zone (Asia/Jakarta)
	location, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		fmt.Printf("Error loading location: %v\n", err)
		return
	}

	c := cron.New()
	// c.AddFunc("@every 1m", func() {
	// 	handlerRegularReminder(db, useCase, location, "daily")
	// })
	c.AddFunc("0 * * * *", func() {
		handlerRegularReminder(db, useCase, location, "Day")
	})
	c.AddFunc("0 0 * * 0", func() {
		handlerRegularReminder(db, useCase, location, "Week")
	})
	c.AddFunc("0 0 1 * *", func() {
		handlerRegularReminder(db, useCase, location, "Month")
	})
	c.Start()
}

// func handlerRegularReminder(db *gorm.DB, useCase UseCase, location *time.Location, reminderType string) {
// 	fmt.Printf("Checking %s for plants to water...", reminderType)
// 	fmt.Printf("now is %v\n", time.Now().In(location))
// 	var plantsToWater []plant.Plant
// 	currentTime := time.Now().In(location)
// 	formattedTime := currentTime.Format("15:04")

// 	// Fetch all plants that need watering at the current time
// 	err := db.
// 		Preload("WateringSchedule").
// 		Joins("JOIN plant_reminders ON plant_reminders.plant_id = plants.id").
// 		Where("plant_reminders.watering_time = ? AND plant_reminders.each = ?", formattedTime, reminderType).
// 		Find(&plantsToWater).Error
// 	if err != nil {
// 		fmt.Printf("Failed to fetch plants to water: %v\n", err)
// 		return
// 	}

// 	// Check if any plants need watering
// 	if len(plantsToWater) == 0 {
// 		fmt.Println("No plants found for watering at this time.")
// 		return
// 	}

// 	// Iterate over each plant to water
// 	for _, plantToWater := range plantsToWater {
// 		var usersWithPlant []user.User

// 		// Find users who have this plant
// 		err := db.Model(&user.User{}).
// 			Joins("JOIN user_plants ON users.id = user_plants.user_id").
// 			Where("user_plants.plant_id = ?", plantToWater.ID).
// 			Find(&usersWithPlant).Error
// 		if err != nil {
// 			fmt.Printf("Failed to fetch users with plant %s: %v\n", plantToWater.Name, err)
// 			continue
// 		}

// 		// Notify each user
// 		for _, user := range usersWithPlant {
// 			err := SendReminder(user, plantToWater, useCase)
// 			if err != nil {
// 				fmt.Printf("Error sending reminder to user %s: %v\n", user.Email, err)
// 			}
// 		}
// 	}
// }

func handlerRegularReminder(db *gorm.DB, useCase UseCase, location *time.Location, reminderType string) {
	fmt.Printf("Checking %s for plants to water...\n", reminderType)

	currentTime := time.Now().In(location)
	formattedTime := currentTime.Format("15:04")

	// Fetch all plants that need watering at the current time
	var plantsToWater []plant.Plant
	err := db.
		Preload("WateringSchedule").
		Joins("JOIN plant_reminders ON plant_reminders.plant_id = plants.id").
		Where("plant_reminders.each = ?", reminderType).
		Find(&plantsToWater).Error
	if err != nil {
		fmt.Printf("Failed to fetch plants to water: %v\n", err)
		return
	}

	// Check if any plants need watering
	if len(plantsToWater) == 0 {
		fmt.Println("No plants found for watering at this time.")
		return
	}

	// Iterate over each plant to water
	for _, plantToWater := range plantsToWater {
		wateringTimes := strings.Split(plantToWater.WateringSchedule.WateringTime, ", ")
		for _, wateringTime := range wateringTimes {
			if wateringTime == formattedTime {
				var usersWithPlant []user.User

				// Find users who have this plant
				err := db.Model(&user.User{}).
					Joins("JOIN user_plants ON users.id = user_plants.user_id").
					Where("user_plants.plant_id = ?", plantToWater.ID).
					Find(&usersWithPlant).Error
				if err != nil {
					fmt.Printf("Failed to fetch users with plant %s: %v\n", plantToWater.Name, err)
					continue
				}

				// Notify each user
				for _, user := range usersWithPlant {
					err := SendReminder(user, plantToWater, useCase)
					if err != nil {
						fmt.Printf("Error sending reminder to user %s: %v\n", user.Email, err)
					}
				}
			}
		}
	}
}

func StartSchedulerForCustomizeWateringReminder(db *gorm.DB, useCase UseCase) {
	// Define desired location for time zone (Asia/Jakarta)
	location, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		fmt.Printf("Error loading location: %v\n", err)
		return
	}
	c := cron.New()
	// c.AddFunc("@every 1m", func() {
	// 	handleCustomizedReminders(db, useCase, location, "daily")
	// })
	c.AddFunc("0 * * * *", func() {
		handleCustomizedReminders(db, useCase, location, "Day")
	})

	c.AddFunc("0 0 * * 0", func() {
		handleCustomizedReminders(db, useCase, location, "Week")
	})

	c.AddFunc("0 0 1 * *", func() {
		handleCustomizedReminders(db, useCase, location, "Month")
	})

	c.Start()
}

func handleCustomizedReminders(db *gorm.DB, useCase UseCase, location *time.Location, reminderType string) {
	fmt.Printf("Checking for %s customized watering reminders...\n", reminderType)
	var reminders []CustomizeWateringReminder
	currentTime := time.Now().In(location)
	formattedTime := currentTime.Format("15:04")

	err := db.Preload("User").Preload("Plant").
		Where("type = ? AND time = ?", reminderType, formattedTime).
		Find(&reminders).Error
	if err != nil {
		fmt.Printf("Failed to fetch %s customized watering reminders: %v\n", reminderType, err)
		return
	}

	if len(reminders) == 0 {
		fmt.Printf("No %s customized watering reminders found for watering at this time.\n", reminderType)
		return
	}

	for _, reminder := range reminders {
		err := SendCustomReminder(reminder, useCase)
		if err != nil {
			fmt.Printf("Error sending %s customized watering reminder: %v\n", reminderType, err)
			continue
		}

		if !reminder.Recurring {
			if err := db.Delete(&reminder).Error; err != nil {
				fmt.Printf("Failed to delete one-time %s reminder: %v\n", reminderType, err)
			}
		}
	}
}

func MapPlantToPlantResponse(plant *plant.Plant) *PlantResponse {
	return &PlantResponse{
		ID:               plant.ID,
		Name:             plant.Name,
		Description:      plant.Description,
		IsToxic:          plant.IsToxic,
		HarvestDuration:  plant.HarvestDuration,
		Sunlight:         plant.Sunlight,
		PlantingTime:     plant.PlantingTime,
		ClimateCondition: plant.ClimateCondition,
		PlantImage:       MapPlantImagesToPlantImageResponses(plant.PlantImages),
		CreatedAt:        plant.CreatedAt,
	}
}

func MapPlantImagesToPlantImageResponses(images []plant.PlantImage) []PlantImageResponse {
	var plantImageResponses []PlantImageResponse

	for _, image := range images {
		if image.IsPrimary == 1 {
			plantImageResponses = append(plantImageResponses, PlantImageResponse{
				ID:       image.ID,
				FileName: image.FileName,
			})
		}
	}

	return plantImageResponses
}
