package repmongo

import (
	"testing"
	inputapp "thiagofo92/study-api-gin/app/input_app"
	"thiagofo92/study-api-gin/app/outputapp"
	"thiagofo92/study-api-gin/share/convert"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
)

func TestUser(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	mt.Run("Success to create a user", func(mt *mtest.T) {
		rep := NewUserRepository(mt.DB)

		var mockErr mtest.WriteError
		mockErr.Message = "error to create a new user"

		input := inputapp.UserInput{
			Name:  "Dev Silva",
			Email: "test@test.com",
		}
		mockInput := convert.BsonArray(&input)

		mt.AddMockResponses(mtest.CreateCursorResponse(0, "library.client", mtest.FirstBatch, mockInput))

		expected := outputapp.UserOutPut{
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
}
