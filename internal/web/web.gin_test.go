package web

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"thiagofo92/study-api-gin/internal/web/routers"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var st Setup

type Setup struct {
	server *httptest.Server
}

func (s *Setup) BeforeAll() {
	g := gin.Default()
	rt := routers.NewGinRouter(g)

	rt.Build()
	s.server = httptest.NewServer(g)
}

func (s *Setup) AfterAll() {
	s.server.Close()
}

func TestMain(m *testing.M) {
	st = Setup{}
	st.BeforeAll()
	defer st.AfterAll()
	m.Run()

}

func TestRouters(t *testing.T) {
	t.Run("User router", func(t *testing.T) {
		url := st.server.URL + "/v1/user/1234"
		req, err := http.NewRequest("GET", url, nil)

		assert.Nil(t, err)

		http.DefaultClient.Do(req)

		defer req.Body.Close()

		var user map[string]interface{}
		err = json.NewDecoder(req.Body).Decode(&user)

		assert.Nil(t, err)

		assert.Equal(t, nil, user)
	})

	t.Run("Book router", func(t *testing.T) {

	})
}

func fecth(method string, url, payload interface{}) (interface{}, error) {
	return nil, nil
}
