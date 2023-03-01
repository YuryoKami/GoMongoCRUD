package routes

import (
	"GoMongoCRUD/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoute(router *gin.Engine) {
	router.GET("/user/:userId", controllers.GetAUser())
	router.GET("/users/list", controllers.GetAllUsers())
	router.POST("/user", controllers.CreateUser())
	//router.POST("/users/list", controllers.CreateUsers())
	router.PUT("/user/:userId", controllers.EditAUser())
	router.DELETE("delete/user/:userId", controllers.DeleteAUser())
}
