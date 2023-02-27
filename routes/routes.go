package routes

import (
	"github.com/YuryoKami/GoMongoCRUD/controllers"
	"github.com/gin-gonic/gin"
)

type UserRouteController struct {
	userController controllers.UserController
}

func NewUserControllerRoute(userController controllers.UserController) UserRouteController {
	return UserRouteController{userController}
}

func (r *UserRouteController) PostRoute(rg *gin.RouterGroup) {
	router := rg.Group("/users")

	router.GET("/users/list", r.userController.GetUsersList)
	router.GET("/user/:id", r.userController.GetUserByID)
	router.POST("/add/user", r.userController.CreateUser)
	//TODO router.POST("/add/users", r.userController.CreateUsers)
	router.PATCH("/user/:id", r.userController.UpdateUser)
	router.DELETE("delete/user/:id", r.userController.DeleteUser)
}
