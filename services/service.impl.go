package services

import (
	"context"
	"errors"
	"time"

	"github.com/YuryoKami/GoMongoCRUD/models"
	"github.com/YuryoKami/GoMongoCRUD/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserServiceImpl struct {
	userCollection *mongo.Collection
	ctx            context.Context
}

func NewUserService(userCollection *mongo.Collection, ctx context.Context) UserService {
	return &UserServiceImpl{userCollection, ctx}
}

// CreateUser implements UserService
func (u *UserServiceImpl) CreateUser(user *models.CreateUserRequest) (*models.DBUser, error) {
	user.Registered = time.Now()

	res, err := u.userCollection.InsertOne(u.ctx, user)
	if err != nil {
		if err2, ok := err.(mongo.WriteException); ok && err2.WriteErrors[0].Code == 11000 {
			return nil, errors.New("this user already exists")
		}
		return nil, err
	}

	opt := options.Index()
	opt.SetUnique(true)
	index := mongo.IndexModel{Keys: bson.M{"email": 1}, Options: opt}
	if _, err := u.userCollection.Indexes().CreateOne(u.ctx, index); err != nil {
		return nil, errors.New("could not create an index for 'email'")
	}

	var newUser *models.DBUser
	query := bson.M{"_id": res.InsertedID}
	if err = u.userCollection.FindOne(u.ctx, query).Decode(&newUser); err != nil {
		return nil, err
	}

	return newUser, nil
}

// DeleteUser implements UserService
func (u *UserServiceImpl) DeleteUser(id string) error {
	obId, _ := primitive.ObjectIDFromHex(id)
	query := bson.M{"_id": obId}

	res, err := u.userCollection.DeleteOne(u.ctx, query)
	if err != nil {
		return err
	}

	if res.DeletedCount == 0 {
		return errors.New("cannot find any user with this ID: " + id)
	}

	return nil
}

// GetUserByID implements UserService
func (u *UserServiceImpl) GetUserByID(id string) (*models.DBUser, error) {
	obId, _ := primitive.ObjectIDFromHex(id)
	query := bson.M{"_id": obId}
	var post *models.DBUser

	if err := u.userCollection.FindOne(u.ctx, query).Decode(&post); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("cannot find any user with this ID: " + id)
		}
		return nil, err
	}

	return post, nil
}

// GetUsersList implements UserService
func (u *UserServiceImpl) GetUsersList() ([]*models.DBUser, error) {
	query := bson.M{}
	cursor, err := u.userCollection.Find(u.ctx, query)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(u.ctx)

	var users []*models.DBUser
	for cursor.Next(u.ctx) {
		user := &models.DBUser{}
		err := cursor.Decode(user)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}
	if len(users) == 0 {
		return []*models.DBUser{}, nil
	}

	return users, nil
}

// UpdateUser implements UserService
func (u *UserServiceImpl) UpdateUser(id string, user *models.UpdateUserRequest) (*models.DBUser, error) {
	doc, err := utils.ToDoc(user)
	if err != nil {
		return nil, err
	}

	obID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	query := bson.D{{Key: "_id", Value: obID}}
	update := bson.D{{Key: "$set", Value: doc}}
	res := u.userCollection.FindOneAndUpdate(u.ctx, query, update, options.FindOneAndUpdate().SetReturnDocument(1))

	var updatedPost *models.DBUser
	if err := res.Decode(&updatedPost); err != nil {
		return nil, errors.New("cannot find any user with this ID: " + id)
	}

	return updatedPost, nil
}
