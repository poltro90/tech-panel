package store

import "github.com/poltro90/tech-panel/model"

type UserStoreMock struct {
	users []model.User
	err   error
}

// NewUserStoreMock returns a new mock implementing IUserStore interface
func NewUserStoreMock(users []model.User, err error) UserStoreMock {
	return UserStoreMock{
		users,
		err,
	}
}

// GetUsersBycompanyID mocks the implementation of the store function for retrieving users given a company ID
func (usm UserStoreMock) GetUsersBycompanyID(companyID string) ([]model.User, error) {
	return usm.users, usm.err
}
