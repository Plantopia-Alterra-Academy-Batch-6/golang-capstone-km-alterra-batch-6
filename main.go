package main

import (
	"fmt"
	"os"

	"github.com/OctavianoRyan25/be-agriculture/configs"
	"github.com/OctavianoRyan25/be-agriculture/handler"
	"github.com/OctavianoRyan25/be-agriculture/modules/admin"
	"github.com/OctavianoRyan25/be-agriculture/modules/ai"
	"github.com/OctavianoRyan25/be-agriculture/modules/article"
	"github.com/OctavianoRyan25/be-agriculture/modules/fertilizer"
	"github.com/OctavianoRyan25/be-agriculture/modules/notification"
	"github.com/OctavianoRyan25/be-agriculture/modules/plant"
	"github.com/OctavianoRyan25/be-agriculture/modules/search"
	"github.com/OctavianoRyan25/be-agriculture/modules/user"
	wateringhistory "github.com/OctavianoRyan25/be-agriculture/modules/watering_history"
	"github.com/OctavianoRyan25/be-agriculture/modules/weather"
	"github.com/OctavianoRyan25/be-agriculture/router"
	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	fmt.Println("Ini Branch Development!")

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	db, err := configs.InitDB()
	if err != nil {
		panic("Failed to connect database")
	}

	err = configs.AutoMigrate(db)
	if err != nil {
		panic("Failed to migrate database")
	}

	cloudinary, err := initCloudinary()
	if err != nil {
		fmt.Println("Failed to initialize Cloudinary:", err)
		return
	}

	repo := user.NewRepository(db)
	useCase := user.NewUseCase(repo)
	controller := user.NewUserController(useCase)

	repoAdmin := admin.NewRepository(db)
	useCaseAdmin := admin.NewUseCase(repoAdmin)
	controllerAdmin := admin.NewUserController(*useCaseAdmin)

	plantCategoryRepository := plant.NewPlantCategoryRepository(db)
	plantCategoryService := plant.NewPlantCategoryService(plantCategoryRepository)
	plantCategoryHandler := handler.NewPlantCategoryHandler(plantCategoryService, cloudinary)

	plantProgressRepository := plant.NewPlantProgressRepository(db)
	plantProgressService := plant.NewPlantProgressService(plantProgressRepository)
	plantProgressHandler := handler.NewPlantProgressHandler(plantProgressService, cloudinary)

	plantInstructionCategoryRepository := plant.NewPlantInstructionCategoryRepository(db)
	plantInstructionCategoryService := plant.NewPlantInstructionCategoryService(plantInstructionCategoryRepository)
	plantInstructionCategoryHandler := handler.NewPlantInstructionCategoryHandler(plantInstructionCategoryService, cloudinary)

	plantRepository := plant.NewPlantRepository(db)
	plantService := plant.NewPlantService(plantRepository, plantCategoryRepository)
	plantHandler := handler.NewPlantHandler(plantService, cloudinary)

	plantUserRepository := plant.NewUserPlantRepository(db)
	plantUserService := plant.NewUserPlantService(plantUserRepository)
	plantUserHandler := handler.NewUserPlantHandler(plantUserService)

	plantEarliestWateringRepository := plant.NewPlantEarliestWateringRepository(db)
	plantEarliestWateringService := plant.NewPlantEarliestWateringService(plantEarliestWateringRepository)
	plantEarliestWateringHandler := handler.NewPlantEarliestWateringHandler(plantEarliestWateringService, cloudinary)

	weatherService := weather.NewWeatherService()
	weatherHandler := handler.NewWeatherHandler(weatherService)

	searchRepository := search.NewRepository(db)
	searchUsecase := search.NewUsecase(searchRepository)
	searchController := search.NewSearchController(searchUsecase)

	// Initialize the notification repository and use case
	notificationRepo := notification.NewRepository(db)
	notificationUseCase := notification.NewUseCase(notificationRepo)
	notificationController := notification.NewNotificationController(notificationUseCase)

	// Initialize Firebase
	//firebaseApp := notification.InitFirebase()

	// Schedule watering reminders
	notification.StartScheduler(db, notificationUseCase)
	notification.StartSchedulerForCustomizeWateringReminder(db, notificationUseCase)

	// Initialize the watering history repository and use case
	wateringHistoryRepo := wateringhistory.NewRepository(db)
	wateringHistoryUseCase := wateringhistory.NewUseCase(wateringHistoryRepo)
	wateringHistoryController := wateringhistory.NeWateringHistoryController(wateringHistoryUseCase)

	fertilizerRepo := fertilizer.NewFertilizerRepository(db)
	fertilizerUseCase := fertilizer.NewFertilizerService(fertilizerRepo)
	fertilizerHandler := handler.NewFertilizerHandler(fertilizerUseCase, cloudinary)

	articleRepo := article.NewRepository(db)
	articleUseCase := article.NewUseCase(articleRepo)
	articleController := article.NewArticleController(articleUseCase)

	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		panic("OPENAI_API_KEY environment variable is not set")
	}

	aiFertilizerRecommendationService := ai.NewPlantService(apiKey)
	aiFertilizerRecommendationHandler := handler.NewAIFertilizerRecommendationHandler(aiFertilizerRecommendationService)

	router.InitRoutes(e, controller, controllerAdmin, plantCategoryHandler, plantHandler, plantUserHandler, weatherHandler, plantInstructionCategoryHandler, plantProgressHandler, searchController, notificationController, wateringHistoryController, fertilizerHandler, aiFertilizerRecommendationHandler, plantEarliestWateringHandler, articleController)

	e.Logger.Fatal(e.Start(":8080"))
}

func initCloudinary() (*cloudinary.Cloudinary, error) {
	//Production
	cloudinaryURL := os.Getenv("CLOUDINARY_URL")
	// cloudinaryURL := "cloudinary://key:secret@cloud_name"
	cloudinary, err := cloudinary.NewFromURL(cloudinaryURL)
	if err != nil {
		return nil, err
	}
	return cloudinary, nil
}
