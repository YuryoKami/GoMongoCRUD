package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TODELETE_User struct {
	Id       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name     string             `json:"name,omitempty" validate:"required"`
	Location string             `json:"location,omitempty" validate:"required"`
	Title    string             `json:"title,omitempty" validate:"required"`
}

type User struct {
	ID         primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Password   string             `json:"password" bson:"password" binding:"required"`
	IsActive   bool               `json:"isActive,omitempty" bson:"isActive,omitempty"`
	Balance    string             `json:"balance,omitempty" bson:"balance,omitempty"`
	Age        string             `json:"age,omitempty" bson:"age,omitempty"`
	Name       string             `json:"name" bson:"name" binding:"required"`
	Gender     string             `json:"gender,omitempty" bson:"gender,omitempty"`
	Company    string             `json:"company,omitempty" bson:"company,omitempty"`
	Email      string             `json:"email,omitempty" bson:"email,omitempty"`
	Phone      string             `json:"phone,omitempty" bson:"phone,omitempty"`
	Address    string             `json:"address,omitempty" bson:"address,omitempty"`
	About      string             `json:"about,omitempty" bson:"about,omitempty"`
	Registered string             `json:"registered,omitempty" bson:"registered,omitempty"`
	Latitude   float64            `json:"latitude,omitempty" bson:"latitude,omitempty"`
	Longitude  float32            `json:"longitude,omitempty" bson:"longitude,omitempty"`
	Tags       []string           `json:"tags,omitempty" bson:"tags,omitempty"`
	Friends    []Friends          `json:"friends,omitempty" bson:"friends,omitempty"`
	Data       string             `json:"data,omitempty" bson:"data,omitempty"`
}

type Friends struct {
	ID   int    `json:"id" bson:"id" binding:"required"`
	Name string `json:"name" bson:"name" binding:"required"`
}
