package repmongo

import (
	"context"
	"fmt"
	inputapp "thiagofo92/study-api-gin/app/input_app"
	"thiagofo92/study-api-gin/app/outputapp"
	"thiagofo92/study-api-gin/infra/repository/rep_mongo/schema"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type BooksRepository struct {
	coll *mongo.Collection
}

func NewBooksRepository(rep *mongo.Database) *BooksRepository {
	coll := rep.Collection("books")
	return &BooksRepository{
		coll: coll,
	}
}

func (b *BooksRepository) Add(input inputapp.BookInput) (outputapp.BookOutput, error) {
	data := schema.BookSchema(input)

	resul, err := b.coll.InsertOne(context.TODO(), data)

	if err != nil {
		return outputapp.BookOutput{}, fmt.Errorf("error to add book %w", err)
	}

	return outputapp.BookOutput{
		Id:         resul.InsertedID.(primitive.ObjectID).Hex(),
		Name:       input.Name,
		Author:     input.Author,
		Categories: input.Categories,
		Available:  input.Available,
		Rented:     input.Rented,
	}, nil
}

func (b *BooksRepository) Update(id string, input inputapp.BookInput) error {
	data := schema.BookSchema(input)

	u := bson.D{{
		Key:   "$set",
		Value: data,
	}}

	idbs, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return fmt.Errorf("error to convert string to objectId %w", err)
	}

	_, err = b.coll.UpdateByID(context.TODO(), idbs, u)

	if err != nil {
		return fmt.Errorf("error to add book %w", err)
	}

	return nil
}

func (b *BooksRepository) FindById(idStr string) (outputapp.BookOutput, error) {
	id, err := primitive.ObjectIDFromHex(idStr)

	if err != nil {
		return outputapp.BookOutput{}, fmt.Errorf("error to convert string to objectID %w", err)
	}

	filter := bson.D{{Key: "_id", Value: id}}
	res := b.coll.FindOne(context.TODO(), filter)

	var output outputapp.BookOutput

	err = res.Decode(&output)

	if err != nil {
		return outputapp.BookOutput{}, fmt.Errorf("error to find book in DB %w", err)
	}

	return output, nil
}

func (b *BooksRepository) Delete(idStr string) (int64, error) {
	id, err := primitive.ObjectIDFromHex(idStr)

	if err != nil {
		return 0, fmt.Errorf("error to convert string to objectID %w", err)
	}

	filter := bson.D{{Key: "_id", Value: id}}
	res, err := b.coll.DeleteOne(context.TODO(), filter)

	if err != nil {
		return 0, fmt.Errorf("error to find book in DB %w", err)
	}

	return res.DeletedCount, nil
}

func (b *BooksRepository) UpdateRent(idStr string, count int) error {
	id, err := primitive.ObjectIDFromHex(idStr)

	if err != nil {
		return fmt.Errorf("error to convert string to objectID %w", err)
	}

	filter := bson.D{{Key: "_id", Value: id}}
	v := bson.D{{Key: "$inc", Value: bson.D{{Key: "rented", Value: count}}}}

	res, err := b.coll.UpdateOne(context.TODO(), filter, v)

	if err != nil {
		return fmt.Errorf("error to find book in DB %w", err)
	}

	fmt.Println("Count rented books", res.ModifiedCount)
	return nil
}
