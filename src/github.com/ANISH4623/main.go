package main

import (
	"fmt"
	"os"

	"github.com/ANISH4623/database"
	"github.com/ANISH4623/helpers"
	"github.com/ANISH4623/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	// http package has methods for dealing with requests
	helpers.LoadConfig(".env")
	database.Connect()
	router := gin.New()
	port := os.Getenv("PORT")
	routes.UserRoutes(router)
	productGroup := router.Group("/users/:userId")
	routes.ProductRoutes(productGroup)
	err := router.Run(":" + port)
	if err != nil {
		fmt.Println(err)
	}

}
