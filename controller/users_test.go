package controller

import (
	"errors"
	"testing"

	"github.com/poltro90/tech-panel/model"

	"github.com/poltro90/tech-panel/store"
	"github.com/stretchr/testify/assert"
)

func Test_GetCompanyUsers(t *testing.T) {
	t.Run("Correctly retrieve users", func(t *testing.T) {
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
		storeMock := store.NewUserStoreMock(users, nil)

		c := NewUserController(storeMock)
		actualUsers, actualErr := c.GetCompanyUsers("fake-company-id")
		assert.Equal(t, users, actualUsers)
		assert.NoError(t, actualErr)
	})

	t.Run("Retrieve users with error", func(t *testing.T) {
		err := errors.New("test-error")
		storeMock := store.NewUserStoreMock(nil, err)

		c := NewUserController(storeMock)
		actualUsers, actualErr := c.GetCompanyUsers("fake-company-id")
		assert.Equal(t, err, actualErr)
		assert.Nil(t, actualUsers)
	})
}
