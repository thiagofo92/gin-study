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
