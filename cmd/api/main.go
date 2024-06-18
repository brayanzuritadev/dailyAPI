package main

import (
	//"log"

	"fmt"

	"github.com/brayanzuritadev/dailyapi/cmd/api/routes"
	"github.com/gin-gonic/gin"
	//"github.com/joho/godotenv"
)

func main() {
	/*err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}*/

	fmt.Println("nuevament epor aqui")
	setup()
}

func setup() {
	ginEngine := gin.Default()
	routes.SetRoutes(ginEngine)
	ginEngine.Run()
}
