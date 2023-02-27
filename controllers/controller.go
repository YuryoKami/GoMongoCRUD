package controllers

import (
	"net/http"
	"strings"

	"github.com/YuryoKami/GoMongoCRUD/models"
	"github.com/YuryoKami/GoMongoCRUD/services"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService services.UserService
}

func NewUserController(userService services.UserService) UserController {
	return UserController{userService}
}

func (uc *UserController) CreateUser(ctx *gin.Context) {
	var user *models.CreateUserRequest
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	newUser, err := uc.userService.CreateUser(user)
	if err != nil {
		if strings.Contains(err.Error(), "this user already exists") {
			ctx.JSON(http.StatusConflict, gin.H{"status": "fail", "message": err.Error()})
			return
		}
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"status": "success", "user": newUser})
}

func (uc *UserController) UpdateUser(ctx *gin.Context) {
	userId := ctx.Param("userId")
	var user *models.UpdateUserRequest
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	updatedUser, err := uc.userService.UpdateUser(userId, user)
	if err != nil {
		if strings.Contains(err.Error(), "cannot find any user with this ID") {
			ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": err.Error()})
			return
		}
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": updatedUser})
}

func (uc *UserController) DeleteUser(ctx *gin.Context) {
	userId := ctx.Param("userId")
	err := uc.userService.DeleteUser(userId)
	if err != nil {
		if strings.Contains(err.Error(), "cannot find any user with this ID") {
			ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": err.Error()})
			return
		}
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
		return
	}
	ctx.JSON(http.StatusNoContent, nil)
}

func (uc *UserController) GetUserByID(ctx *gin.Context) {
	userId := ctx.Param("userId")
	user, err := uc.userService.GetUserByID(userId)
	if err != nil {
		if strings.Contains(err.Error(), "cannot find any user with this ID") {
			ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": err.Error()})
			return
		}
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"status": "success", "user": user})
}

func (uc *UserController) GetUsersList(ctx *gin.Context) {
	users, err := uc.userService.GetUsersList()
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"status": "success", "results": len(users), "users": users})
}
