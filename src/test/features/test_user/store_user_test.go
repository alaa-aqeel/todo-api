package testuser

import (
	"net/http"
	"testing"

	"github.com/alaa-aqeel/todo/src/test"
)

func TestCreateUserEndpoint(t *testing.T) {
	s := test.NewServer(t)
	defer s.Close()
	client := s.Client()

	t.Run("Create User", func(t *testing.T) {
		client.
			Request(http.MethodPost, "/api/v1/users").
			Expect().
			Status(http.StatusOK)

	})
}
