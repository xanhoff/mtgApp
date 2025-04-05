package main

import (
	"example/mtgApp/controllers"
	"example/mtgApp/database"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

	database.InitDB(os.Getenv("DB"))

	router := gin.Default()

	router.GET("/player", controllers.GetAllPlayers)
	router.GET("/player/:playerid", controllers.GetPlayer)

	router.GET("/game", controllers.GetAllGames)
	router.GET("/game/:gameid", controllers.GetGame)

	router.Run("localhost:8080")
}
