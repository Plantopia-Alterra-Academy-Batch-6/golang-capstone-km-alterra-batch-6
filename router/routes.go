package router

import (
	"net/http"

	"github.com/OctavianoRyan25/be-agriculture/handler"
	"github.com/OctavianoRyan25/be-agriculture/middlewares"
	"github.com/OctavianoRyan25/be-agriculture/modules/admin"
	"github.com/OctavianoRyan25/be-agriculture/modules/article"
	bot "github.com/OctavianoRyan25/be-agriculture/modules/chatbot"
	"github.com/OctavianoRyan25/be-agriculture/modules/notification"
	"github.com/OctavianoRyan25/be-agriculture/modules/search"
	"github.com/OctavianoRyan25/be-agriculture/modules/user"
	wateringhistory "github.com/OctavianoRyan25/be-agriculture/modules/watering_history"
	"github.com/labstack/echo/v4"
)

func InitRoutes(e *echo.Echo, userController *user.UserController, adminController *admin.AdminController, plantCategoryHandler *handler.PlantCategoryHandler, plantHandler *handler.PlantHandler, plantUserHandler *handler.UserPlantHandler, weatherHandler *handler.WeatherHandler, plantInstructionCategoryHandler *handler.PlantInstructionCategoryHandler, plantProgressHandler *handler.PlantProgressHandler, search *search.SearchController, notification *notification.NotificationController, wateringhistory *wateringhistory.WateringHistoryController, fertilizer *handler.FertilizerHandler, aiFertilizer *handler.AIFertilizerRecommendationHandler, plantEarliestWateringHandler *handler.PlantEarliestWateringHandler, article *article.ArticleController) {
	group := e.Group("/api/v1")
	group.POST("/register", userController.RegisterUser)
	group.POST("/check-email", userController.CheckEmail)
	group.POST("/verify", userController.VerifyEmail)
	group.POST("/login", userController.Login)
	group.GET("/profile", userController.GetUserProfile, middlewares.Authentication())
	group.POST("/resendotp", userController.ResendOTP)
	group.POST("/forgot-password", userController.ResetPassword)

	groupAdmin := e.Group("/api/v1/admin")
	groupAdmin.POST("/register", adminController.RegisterUser)
	groupAdmin.POST("/login", adminController.Login)
	groupAdmin.GET("/profile", adminController.GetUserProfile, middlewares.Authentication())

	group.GET("/plants/categories", plantCategoryHandler.GetAll)
	group.GET("/plants/categories/:id", plantCategoryHandler.GetByID)
	groupAdmin.POST("/plants/categories", plantCategoryHandler.Create, middlewares.Authentication())
	groupAdmin.PUT("/plants/categories/:id", plantCategoryHandler.Update, middlewares.Authentication())
	groupAdmin.DELETE("/plants/categories/:id", plantCategoryHandler.Delete, middlewares.Authentication())

	group.GET("/plants/progress/:plant_id", plantProgressHandler.GetAllByUserIDAndPlantID, middlewares.Authentication())
	group.POST("/plants/progress", plantProgressHandler.UploadProgress, middlewares.Authentication())

	group.GET("/plants/instructions/categories", plantInstructionCategoryHandler.GetAll)
	group.GET("/plants/instructions/categories/:id", plantInstructionCategoryHandler.GetByID)
	groupAdmin.POST("/plants/instructions/categories", plantInstructionCategoryHandler.Create, middlewares.Authentication())
	groupAdmin.PUT("/plants/instructions/categories/:id", plantInstructionCategoryHandler.Update, middlewares.Authentication())
	groupAdmin.DELETE("/plants/instructions/categories/:id", plantInstructionCategoryHandler.Delete, middlewares.Authentication())

	group.GET("/plants", plantHandler.GetAll)
	group.GET("/plants/:id", plantHandler.GetByID)
	group.GET("/plants/search", plantHandler.SearchPlantsByName)
	group.GET("/plants/category/:category_id", plantHandler.GetPlantsByCategoryID)
	group.GET("/plants/recommendations", plantHandler.GetRecommendations, middlewares.Authentication())
	group.GET("/plants/instructions/:plant_id/:instruction_category_id", plantInstructionCategoryHandler.GetInstructionByCategoryID)
	groupAdmin.POST("/plants", plantHandler.Create, middlewares.Authentication())
	groupAdmin.PUT("/plants/:id", plantHandler.Update, middlewares.Authentication())
	groupAdmin.DELETE("/plants/:id", plantHandler.Delete, middlewares.Authentication())

	group.GET("/my/plants/:user_id", plantUserHandler.GetUserPlants, middlewares.Authentication())
	group.POST("/my/plants/add", plantUserHandler.AddUserPlant, middlewares.Authentication())
	group.PUT("/my/plants/:userPlantID/customize-name", plantUserHandler.UpdateCustomizeName, middlewares.Authentication())
	group.DELETE("/my/plants/:user_plant_id", plantUserHandler.DeleteUserPlantByID, middlewares.Authentication())
	group.POST("/my/plants/history", plantUserHandler.AddUserPlantHistory, middlewares.Authentication())
	group.GET("/my/plants/history", plantUserHandler.GetUserPlantHistoryByUserID, middlewares.Authentication())
	group.PUT("/my/plants/update-instructions", plantUserHandler.UpdateInstructionCategory, middlewares.Authentication())
	group.GET("/my/plant/details/:plant_id", plantUserHandler.GetUserPlantByUserIDAndPlantID, middlewares.Authentication())

	group.GET("/weather/current", weatherHandler.GetCurrentWeather, middlewares.Authentication())
	group.GET("/weather/hourly", weatherHandler.GetHourlyWeather, middlewares.Authentication())
	group.GET("/weather/daily", weatherHandler.GetDailyWeather, middlewares.Authentication())

	group.GET("/notifications/:id", notification.ReadNotification, middlewares.Authentication())
	group.GET("/notifications", notification.GetAllNotifications, middlewares.Authentication())
	group.DELETE("/notifications", notification.DeleteAllNotifications, middlewares.Authentication())

	groupFertilizer := e.Group("/api/v1")
	groupFertilizer.GET("/fertilizer", fertilizer.GetFertilizer)
	groupFertilizer.GET("/fertilizer/:Id", fertilizer.GetFertilizerById)
	groupFertilizer.POST("/fertilizer", fertilizer.CreateFertilizer, middlewares.Authentication())
	groupFertilizer.PUT("/fertilizer/:Id", fertilizer.UpdateFertilizer, middlewares.Authentication())
	groupFertilizer.DELETE("/fertilizer/:Id", fertilizer.DeleteFertilizer, middlewares.Authentication())

	group.POST("/create-customize-watering-reminder", notification.CreateCustomizeWateringReminder, middlewares.Authentication())
	group.POST("/watering-history", wateringhistory.StoreWateringHistory, middlewares.Authentication())
	group.GET("/watering-history", wateringhistory.GetAllWateringHistories, middlewares.Authentication())
	group.GET("/check-watering", wateringhistory.GetLateWateringHistories, middlewares.Authentication())
	group.GET("/watering-earliest", plantEarliestWateringHandler.GetEarliestWateringTime)

	group.POST("/chatbot", bot.ClassifyEnvironmentalIssue)

	group.GET("/recommend-fertilizer", aiFertilizer.GetFertilizerRecommendation)

	group.GET("/recommend-plants", aiFertilizer.GetPlantingRecommendation)

	group.GET("/login-google", user.LoginGoogle)

	group.GET("/auth/google/callback", user.CallbackGoogle)

	group.GET("/login-google-andro", user.LoginGoogleforAndro)

	group.GET("/auth-andro/google/callback", user.CallbackGoogleforAndro)

	group.GET("/auth", user.GetToken)

	group.POST("/article", article.StoreArticle, middlewares.Authentication())
	group.GET("/article", article.GetAllArticles)
	group.GET("/article/:id", article.GetArticle)
	group.PUT("/article/:id", article.UpdateArticle, middlewares.Authentication())
	group.DELETE("/article/:id", article.DeleteArticle, middlewares.Authentication())

	group.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	group.GET("/search", search.Search)
}
