package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DBUser struct {
	ID         primitive.ObjectID `json:"id" bson:"_id" binding:"required"`
	Password   string             `json:"password" bson:"password" binding:"required"`
	IsActive   bool               `json:"isActive" bson:"isActive" binding:"required"`
	Balance    string             `json:"balance" bson:"balance" binding:"required"`
	Age        int                `json:"age" bson:"age" binding:"required"`
	Name       string             `json:"name" bson:"name" binding:"required"`
	Gender     string             `json:"gender" bson:"gender" binding:"required"`
	Company    string             `json:"company,omitempty" bson:"company,omitempty"`
	Email      string             `json:"email" bson:"email" binding:"required"`
	Phone      string             `json:"phone,omitempty" bson:"phone,omitempty"`
	Address    string             `json:"address" bson:"address" binding:"required"`
	About      string             `json:"about,omitempty" bson:"about,omitempty"`
	Registered time.Time          `json:"registered" bson:"registered" binding:"required"`
	Latitude   string             `json:"latitude" bson:"latitude" binding:"required"`
	Longitude  string             `json:"longitude" bson:"longitude" binding:"required"`
	Tags       []string           `json:"tags,omitempty" bson:"tags,omitempty"`
	Friends    []Friends          `json:"friends,omitempty" bson:"friends,omitempty"`
	Data       string             `json:"data" bson:"data" binding:"required"`
}

type CreateUserRequest struct {
	Password   string    `json:"password" bson:"password" binding:"required"`
	IsActive   bool      `json:"isActive" bson:"isActive" binding:"required"`
	Balance    string    `json:"balance" bson:"balance" binding:"required"`
	Age        int       `json:"age" bson:"age" binding:"required"`
	Name       string    `json:"name" bson:"name" binding:"required"`
	Gender     string    `json:"gender" bson:"gender" binding:"required"`
	Company    string    `json:"company,omitempty" bson:"company,omitempty"`
	Email      string    `json:"email" bson:"email" binding:"required"`
	Phone      string    `json:"phone,omitempty" bson:"phone,omitempty"`
	Address    string    `json:"address" bson:"address" binding:"required"`
	About      string    `json:"about,omitempty" bson:"about,omitempty"`
	Registered time.Time `json:"registered" bson:"registered" binding:"required"`
	Latitude   string    `json:"latitude" bson:"latitude" binding:"required"`
	Longitude  string    `json:"longitude" bson:"longitude" binding:"required"`
	Tags       []string  `json:"tags,omitempty" bson:"tags,omitempty"`
	Friends    []Friends `json:"friends,omitempty" bson:"friends,omitempty"`
	Data       string    `json:"data" bson:"data" binding:"required"`
}

type UpdateUserRequest struct {
	ID         primitive.ObjectID `json:"id" bson:"_id" binding:"required"`
	Password   string             `json:"password" bson:"password" binding:"required"`
	IsActive   bool               `json:"isActive" bson:"isActive" binding:"required"`
	Balance    string             `json:"balance" bson:"balance" binding:"required"`
	Age        int                `json:"age" bson:"age" binding:"required"`
	Name       string             `json:"name" bson:"name" binding:"required"`
	Gender     string             `json:"gender" bson:"gender" binding:"required"`
	Company    string             `json:"company,omitempty" bson:"company,omitempty"`
	Email      string             `json:"email" bson:"email" binding:"required"`
	Phone      string             `json:"phone,omitempty" bson:"phone,omitempty"`
	Address    string             `json:"address" bson:"address" binding:"required"`
	About      string             `json:"about,omitempty" bson:"about,omitempty"`
	Registered time.Time          `json:"registered" bson:"registered" binding:"required"`
	Latitude   string             `json:"latitude" bson:"latitude" binding:"required"`
	Longitude  string             `json:"longitude" bson:"longitude" binding:"required"`
	Tags       []string           `json:"tags,omitempty" bson:"tags,omitempty"`
	Friends    []Friends          `json:"friends,omitempty" bson:"friends,omitempty"`
	Data       string             `json:"data" bson:"data" binding:"required"`
}

type Friends struct {
	ID   primitive.ObjectID `json:"id" bson:"_id" binding:"required"`
	Name string             `json:"name" bson:"name" binding:"required"`
}
