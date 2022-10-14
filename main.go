package main

import (
	"chillroom/cache"
	"chillroom/configs"
	"chillroom/controllers"
	"chillroom/repositories"
	"chillroom/services"

	"github.com/gin-gonic/gin"
)

func init() {
	configs.LoadEnvVariables()
	configs.ConnectDB()
	configs.SyncDB()
}

func main() {
	server := gin.Default()
	apiCache := cache.NewRedisCache("localhost:6379", 0, 10)

	authRepository := repositories.NewAuthRepository()
	authService := services.NewAuthService(&authRepository)
	AuthController := controllers.NewAuthController(&authService)
	server.POST("/register", AuthController.Register)
	server.POST("/login", AuthController.Login)

	movieService := services.NewMovieService()
	MovieController := controllers.NewMovieController(&movieService, &apiCache)
	server.GET("/movies/trending", MovieController.GetTrending)
	server.GET("/movies/:id", MovieController.FindByID)
	server.GET("/movies/search", MovieController.SearchMovie)

	gameService := services.NewGameService()
	GameController := controllers.NewGameController(&gameService, &apiCache)
	server.GET("/games/trending", GameController.GetTrending)
	server.GET("/games/:id", GameController.FindByID)
	server.GET("/games/search", GameController.SearchGames)

	server.Run()
}