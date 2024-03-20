package repmongo

import (
	"testing"
	inputapp "thiagofo92/study-api-gin/app/input_app"
	"thiagofo92/study-api-gin/app/outputapp"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
)

func TestBook(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))

	mt.Run("Success to add book", func(mt *mtest.T) {
		rp := NewBooksRepository(mt.DB)
		input := inputapp.BookInput{}

		mt.AddMockResponses(mtest.CreateCursorResponse(1, "library.books", mtest.FirstBatch, bson.D{}))
		resul, err := rp.Add(input)

		assert.Nil(t, err)
		assert.NotEmpty(t, resul.Id)
	})

	mt.Run("Error to add book", func(mt *mtest.T) {
		rp := NewBooksRepository(mt.DB)
		input := inputapp.BookInput{}
		var mockError mtest.WriteError

		mockError.Message = "test error to add book"

		mt.AddMockResponses(mtest.CreateWriteErrorsResponse(mockError))
		result, err := rp.Add(input)

		assert.NotNil(t, err)
		assert.Zero(t, result)
	})

	mt.Run("Success to update the book by id", func(mt *mtest.T) {
		rp := NewBooksRepository(mt.DB)
		input := inputapp.BookInput{}

		mt.AddMockResponses(mtest.CreateCursorResponse(1, "library.books", mtest.FirstBatch, bson.D{{"_id", ""}}))
		id := primitive.NewObjectID().Hex()
		err := rp.Update(id, input)

		assert.Nil(t, err)
	})

	mt.Run("Error to convert ID to ObjectID", func(mt *mtest.T) {
		rp := NewBooksRepository(mt.DB)
		input := inputapp.BookInput{}

		mt.AddMockResponses(mtest.CreateCursorResponse(1, "library.books", mtest.FirstBatch, bson.D{{"_id", ""}}))
		err := rp.Update("", input)

		assert.NotNil(t, err)
		assert.Contains(t, err.Error(), "error to convert string to objectId")
	})

	mt.Run("Error to update DB", func(mt *mtest.T) {
		rp := NewBooksRepository(mt.DB)
		input := inputapp.BookInput{}

		var mockErr mtest.WriteError
		mockErr.Message = "Mock error"

		mt.AddMockResponses(mtest.CreateWriteErrorsResponse(mockErr))
		id := primitive.NewObjectID().Hex()
		err := rp.Update(id, input)

		assert.NotNil(t, err)
		assert.Contains(t, err.Error(), "error to add book")
	})

	mt.Run("Find book by ID", func(mt *mtest.T) {
		rp := NewBooksRepository(mt.DB)

		input := inputapp.BookInput{
			Name: "Test",
		}

		buff, _ := bson.Marshal(input)

		var mock bson.D
		_ = bson.Unmarshal(buff, &mock)

		mt.AddMockResponses(mtest.CreateCursorResponse(1, "library.book", mtest.FirstBatch, mock))
		id := primitive.NewObjectID().Hex()

		output, err := rp.FindById(id)
		expected := outputapp.BookOutput{
			Name: "Test",
		}

		assert.Nil(t, err)
		assert.Equal(t, expected, output)
	})

	mt.Run("Error to find book by ID - Invalid ObjectID", func(mt *mtest.T) {
		rp := NewBooksRepository(mt.DB)

		var mockErr mtest.WriteError
		mockErr.Message = "Mock error"

		mt.AddMockResponses(mtest.CreateWriteErrorsResponse(mockErr))

		_, err := rp.FindById("")

		assert.NotNil(t, err)
		assert.Contains(t, err.Error(), "error to convert string to objectID")
	})

	mt.Run("Error to find book by ID - DB error", func(mt *mtest.T) {
		rp := NewBooksRepository(mt.DB)

		var mockErr mtest.WriteError
		mockErr.Message = "Mock error"

		mt.AddMockResponses(mtest.CreateWriteErrorsResponse(mockErr))
		id := primitive.NewObjectID().Hex()

		_, err := rp.FindById(id)

		assert.NotNil(t, err)
		assert.Contains(t, err.Error(), "error to find book in DB")
	})

	mt.Run("Delete book by ID", func(mt *mtest.T) {
		rp := NewBooksRepository(mt.DB)

		mt.AddMockResponses(mtest.CreateCursorResponse(1, "library.book", mtest.FirstBatch, bson.D{{}}))
		id := primitive.NewObjectID().Hex()

		_, err := rp.Delete(id)

		assert.Nil(t, err)
	})

	mt.Run("Error to delete book - convert string to object ID", func(mt *mtest.T) {
		rp := NewBooksRepository(mt.DB)

		mt.AddMockResponses(mtest.CreateCursorResponse(1, "library.book", mtest.FirstBatch, bson.D{{}}))

		_, err := rp.Delete("")

		assert.NotNil(t, err)
		assert.Contains(t, err.Error(), "error to convert string to objectID")
	})

	mt.Run("Error to delete book - convert string to object ID", func(mt *mtest.T) {
		rp := NewBooksRepository(mt.DB)

		mt.AddMockResponses(mtest.CreateCursorResponse(1, "library.book", mtest.FirstBatch, bson.D{{}}))
		id := primitive.NewObjectID().Hex()

		_, err := rp.Delete(id)

		assert.NotNil(t, err)
		assert.Contains(t, err.Error(), "error to find book in DB")
	})
}
