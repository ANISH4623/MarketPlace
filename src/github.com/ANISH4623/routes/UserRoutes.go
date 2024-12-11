package routes

import (
	"github.com/ANISH4623/controllers"
	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine) {
	router.POST("/users/", controllers.CreateUser())
}
