package routes

import (
	"GoMongoCRUD/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoute(router *gin.Engine) {
	router.GET("/user/:userId", controllers.GetAUser())
	router.GET("/users", controllers.GetAllUsers())
	router.POST("/user", controllers.CreateUser())
	router.PUT("/user/:userId", controllers.EditAUser())
	router.DELETE("/user/:userId", controllers.DeleteAUser())
}
