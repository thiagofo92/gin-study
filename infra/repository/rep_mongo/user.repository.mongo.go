package repmongo

import (
	"context"
	"fmt"
	inputapp "thiagofo92/study-api-gin/app/input_app"
	"thiagofo92/study-api-gin/app/outputapp"
	"thiagofo92/study-api-gin/infra/repository/rep_mongo/schema"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository struct {
	coll *mongo.Collection
}

func NewUserRepository(db *mongo.Database) *UserRepository {
	return &UserRepository{
		coll: db.Collection("users"),
	}
}

func (u *UserRepository) Create(input inputapp.UserInput) (outputapp.UserOutPut, error) {
	data := schema.UserSchema(input)
	resul, err := u.coll.InsertOne(context.TODO(), data)

	if err != nil {
		return outputapp.UserOutPut{}, fmt.Errorf("error to create user %w", err)
	}

	return outputapp.UserOutPut{
		Id:    resul.InsertedID.(primitive.ObjectID).Hex(),
		Email: input.Email,
		Name:  input.Name,
	}, nil
}
