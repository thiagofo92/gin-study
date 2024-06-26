package usecase

import (
	"testing"
	repmongo "thiagofo92/study-api-gin/infra/repository/rep_mongo"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
)

func TestUsecaseBook(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))

	mt.Run("success to rent a book", func(mt *mtest.T) {
		ruser := repmongo.NewUserRepository(mt.DB)
		rbook := repmongo.NewBooksRepository(mt.DB)
		usecase := NewBookUseCase(ruser, rbook)

		mt.AddMockResponses(mtest.CreateCursorResponse(0, "library.book", mtest.FirstBatch, bson.D{}))
		mt.AddMockResponses(mtest.CreateCursorResponse(0, "library.book", mtest.FirstBatch, bson.D{}))

		id := primitive.NewObjectID().Hex()
		idBook := primitive.NewObjectID().Hex()
		err := usecase.Rent(id, idBook)

		assert.Nil(t, err)
	})
}
