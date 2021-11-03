package handler

import (
	"errors"
	"net/http/httptest"
	"testing"

	"github.com/poltro90/tech-panel/controller"
	"github.com/poltro90/tech-panel/model"
)

func Test_GetUsers(t *testing.T) {
	t.Run("Retrieve users without error", func(t *testing.T) {
		// Case with no error
		users := []model.User{
			{
				FirstName: "Mario",
				LastName:  "Rossi",
			},
			{
				FirstName: "John",
				LastName:  "Doe",
			},
		}
		controllerMock := controller.NewUserControllerMock(users, nil)

		handler := NewUserHandler(controllerMock)
		w := httptest.NewRecorder()
		r := httptest.NewRequest("POST", "/users", nil)
		handler.GetUsers(w, r)

		// here you can test the users
	})

	t.Run("Retrieve users with error", func(t *testing.T) {
		err := errors.New("expected error")
		controllerMock := controller.NewUserControllerMock(nil, err)

		handler := NewUserHandler(controllerMock)
		w := httptest.NewRecorder()
		r := httptest.NewRequest("POST", "/users", nil)
		handler.GetUsers(w, r)

		// here you can test the error
	})
}
