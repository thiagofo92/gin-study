package repmongo

import (
	"context"
	"fmt"
	inputapp "thiagofo92/study-api-gin/internal/app/input_app"
	"thiagofo92/study-api-gin/internal/app/outputapp"
	"thiagofo92/study-api-gin/internal/infra/repository/rep_mongo/schema"

	"go.mongodb.org/mongo-driver/bson"
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
	data := schema.UserSchema{
		Name:         input.Name,
		Password:     input.Password,
		Email:        input.Email,
		RentedBooks:  []string{},
		BooksHistory: []string{},
	}
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

func (u *UserRepository) Update(idStr string, input inputapp.UserInput) error {

	id, err := primitive.ObjectIDFromHex(idStr)

	if err != nil {
		return fmt.Errorf("error to convert string to object ID %w", err)
	}

	data := schema.UserSchema{
		Name:        input.Name,
		Password:    input.Password,
		Email:       input.Email,
		RentedBooks: input.RentedBooks,
	}

	v := bson.D{{
		Key:   "$set",
		Value: data,
	}}

	_, err = u.coll.UpdateByID(context.TODO(), id, v)

	if err != nil {
		return fmt.Errorf("error to create user %w", err)
	}

	return nil
}

func (u *UserRepository) FindById(idStr string) (outputapp.UserOutPut, error) {

	id, err := primitive.ObjectIDFromHex(idStr)

	if err != nil {
		return outputapp.UserOutPut{}, err
	}

	filter := bson.D{{Key: "_id", Value: id}}
	res := u.coll.FindOne(context.TODO(), filter)
	var output outputapp.UserOutPut

	err = res.Decode(&output)

	if err != nil {
		return outputapp.UserOutPut{}, fmt.Errorf("error to create user %w", err)
	}

	return output, nil
}

func (u *UserRepository) RentBook(idStr string, bookId string) error {

	id, err := primitive.ObjectIDFromHex(idStr)

	if err != nil {
		return fmt.Errorf("error to convert string to ObjectID %w", err)
	}

	v := bson.D{{
		Key: "$push",
		Value: bson.D{{
			Key:   "rentedBooks",
			Value: bookId,
		},
			{
				Key:   "booksHistory",
				Value: bookId,
			},
		},
	}}
	_, err = u.coll.UpdateByID(context.TODO(), id, v)

	if err != nil {
		return fmt.Errorf("error to update dabase %w", err)
	}

	return nil
}

func (u *UserRepository) ReturnBook(idStr string, bookId string) error {
	id, err := primitive.ObjectIDFromHex(idStr)

	if err != nil {
		return fmt.Errorf("return a book - error to convert string to ObjectId %w", err)
	}

	v := bson.D{{
		Key: "$pull",
		Value: bson.D{{
			Key: "rentedBooks",
			Value: bson.D{{
				Key: "$elemMatch",
				Value: bson.D{{
					Key:   "$eq",
					Value: id,
				}},
			}},
		}},
	}}
	_, err = u.coll.UpdateByID(context.TODO(), id, v)

	if err != nil {
		return fmt.Errorf("return a book - error to update database %w", err)
	}

	return nil
}
