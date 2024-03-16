package repmongo

import (
	"testing"
	inputapp "thiagofo92/study-api-gin/app/input_app"
	"thiagofo92/study-api-gin/share/convert"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
)

func TestBook(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))

	mt.Run("Success to add book", func(mt *mtest.T) {
		rp := NewBooksRepository(mt.DB)
		input := inputapp.BookInput{}

		bs := convert.BsonArray(&input)
		mt.AddMockResponses(mtest.CreateCursorResponse(1, "library.books", mtest.FirstBatch, bs))
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
}
