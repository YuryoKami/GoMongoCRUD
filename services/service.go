package services

import (
	"github.com/YuryoKami/GoMongoCRUD/models"
)

type UserService interface {
	GetUserByID(id string) (*models.DBUser, error)
	GetUsersList() ([]*models.DBUser, error)
	CreateUser(user *models.CreateUserRequest) (*models.DBUser, error)
	UpdateUser(id string, user *models.UpdateUserRequest) (*models.DBUser, error)
	DeleteUser(id string) error
}
