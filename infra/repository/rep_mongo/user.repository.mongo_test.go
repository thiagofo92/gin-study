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

func TestUser(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	mt.Run("Success to create a user", func(mt *mtest.T) {
		rep := NewUserRepository(mt.DB)

		mockInput := bson.D{{Key: "Name", Value: "Dev Silva"}, {Key: "Email", Value: "test@test.com"}, {Key: "Password", Value: "1234"}}
		mt.AddMockResponses(mtest.CreateCursorResponse(0, "library.client", mtest.FirstBatch, mockInput))

		expected := outputapp.UserOutPut{
			Name:     "Dev Silva",
			Email:    "test@test.com",
			Password: "1234",
		}

		input := inputapp.UserInput{
			Name:  "Dev Silva",
			Email: "test@test.com",
		}
		resul, _ := rep.Create(input)

		assert.NotEmpty(t, resul.Id)
		assert.Equal(t, expected.Name, resul.Name)
		assert.Equal(t, expected.Email, resul.Email)
		assert.Empty(t, resul.Password)
	})

	mt.Run("Error to create a new user", func(mt *mtest.T) {
		rep := NewUserRepository(mt.DB)

		var mockErr mtest.WriteError
		mockErr.Message = "error to create a new user"

		mt.AddMockResponses(mtest.CreateWriteErrorsResponse(mockErr))

		_, err := rep.Create(inputapp.UserInput{})

		assert.NotNil(t, err)
	})

	mt.Run("Success to update user by id", func(mt *mtest.T) {
		rep := NewUserRepository(mt.DB)

		mt.AddMockResponses(mtest.CreateCursorResponse(0, "library.client", mtest.FirstBatch, bson.D{{}}))

		input := inputapp.UserInput{
			Name:  "Dev Silva",
			Email: "test@test.com",
		}

		id := primitive.NewObjectID().Hex()
		err := rep.Update(id, input)

		assert.Nil(t, err)
	})
	mt.Run("Error to update user - convert string to object ID", func(mt *mtest.T) {
		rep := NewUserRepository(mt.DB)

		mt.AddMockResponses(mtest.CreateCursorResponse(0, "library.client", mtest.FirstBatch, bson.D{{}}))

		input := inputapp.UserInput{
			Name:  "Dev Silva",
			Email: "test@test.com",
		}

		err := rep.Update("", input)

		assert.NotNil(t, err)
	})
	mt.Run("Error to update user - update DB", func(mt *mtest.T) {
		rep := NewUserRepository(mt.DB)

		var mockError mtest.WriteError
		mockError.Message = "Error mock"
		mt.AddMockResponses(mtest.CreateWriteErrorsResponse(mockError))

		input := inputapp.UserInput{
			Name:  "Dev Silva",
			Email: "test@test.com",
		}

		id := primitive.NewObjectID().Hex()
		err := rep.Update(id, input)

		assert.NotNil(t, err)
	})

	mt.Run("Find user by id", func(mt *mtest.T) {
		rep := NewUserRepository(mt.DB)

		input := inputapp.UserInput{
			Password:     "1234",
			RentedBooks:  []string{},
			BooksHistory: []string{},
			Name:         "Dev Silva",
			Email:        "test@test.com",
		}

		buff, _ := bson.Marshal(input)

		var mock bson.D

		bson.Unmarshal(buff, &mock)

		mt.AddMockResponses(mtest.CreateCursorResponse(1, "library.user", mtest.FirstBatch, mock))

		id := primitive.NewObjectID().Hex()
		output, _ := rep.FindById(id)

		expected := outputapp.UserOutPut{
			Password: "1234",
			Name:     "Dev Silva",
			Email:    "test@test.com",
		}

		assert.Equal(t, expected, output)
	})

	mt.Run("Error to find user by id - convert string to object ID", func(mt *mtest.T) {
		rep := NewUserRepository(mt.DB)

		mt.AddMockResponses(mtest.CreateCursorResponse(1, "library.user", mtest.FirstBatch, bson.D{}))

		_, err := rep.FindById("")

		assert.NotNil(t, err)
	})

	mt.Run("Error to find user by id - DB", func(mt *mtest.T) {
		rep := NewUserRepository(mt.DB)

		var mock mtest.WriteError
		mock.Message = "Mock error"

		mt.AddMockResponses(mtest.CreateWriteErrorsResponse(mock))

		id := primitive.NewObjectID().Hex()
		_, err := rep.FindById(id)

		assert.NotNil(t, err)
	})

	mt.Run("success to rent a book", func(mt *mtest.T) {
		rep := NewUserRepository(mt.DB)

		mt.AddMockResponses(mtest.CreateCursorResponse(0, "library.user", mtest.FirstBatch, bson.D{}))

		idUser := primitive.NewObjectID().Hex()

		err := rep.RentBook(idUser, "")

		assert.Nil(t, err, nil)
	})

	mt.Run("error to rent a book", func(mt *mtest.T) {
		rep := NewUserRepository(mt.DB)

		var wer mtest.WriteConcernError
		wer.Message = "test error"

		mt.AddMockResponses(mtest.CreateWriteConcernErrorResponse(wer))

		idUser := primitive.NewObjectID().Hex()

		err := rep.RentBook(idUser, "")

		assert.NotNil(t, err)
	})

	mt.Run("error to rent a book - convert string to objectId", func(mt *mtest.T) {
		rep := NewUserRepository(mt.DB)

		var wer mtest.WriteConcernError
		wer.Message = "test error"

		mt.AddMockResponses(mtest.CreateWriteConcernErrorResponse(wer))

		err := rep.RentBook("", "")

		assert.Contains(t, err.Error(), "convert string to ObjectID")
	})

	mt.Run("success to give a book back", func(mt *mtest.T) {
		rep := NewUserRepository(mt.DB)

		mt.AddMockResponses(mtest.CreateCursorResponse(0, "library.user", mtest.FirstBatch, bson.D{}))

		idUser := primitive.NewObjectID().Hex()

		err := rep.ReturnBook(idUser, "")

		assert.Nil(t, err, nil)
	})

	mt.Run("error to give a book back - database", func(mt *mtest.T) {
		rep := NewUserRepository(mt.DB)

		var wer mtest.WriteConcernError
		wer.Message = "test error"

		mt.AddMockResponses(mtest.CreateWriteConcernErrorResponse(wer))

		idUser := primitive.NewObjectID().Hex()

		err := rep.ReturnBook(idUser, "")

		assert.NotNil(t, err)
	})

	mt.Run("error to rent a book - convert string to objectId", func(mt *mtest.T) {
		rep := NewUserRepository(mt.DB)

		var wer mtest.WriteConcernError
		wer.Message = "test error"

		mt.AddMockResponses(mtest.CreateWriteConcernErrorResponse(wer))

		err := rep.ReturnBook("", "")

		assert.Contains(t, err.Error(), "convert string to ObjectId")
	})
}
