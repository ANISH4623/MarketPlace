package routes

import (
	"github.com/ANISH4623/controllers"
	"github.com/gin-gonic/gin"
)

func ProductRoutes(router *gin.RouterGroup) {
	router.GET("/", controllers.GetProducts)
	router.GET("/:id", controllers.GetProduct)
	router.POST("/", controllers.CreateProduct)
	router.PUT("/:id", controllers.UpdateProduct)
	router.DELETE("/:id", controllers.DeleteProduct)
}
