package controller

import "github.com/poltro90/tech-panel/model"

// UserControllerMock mocks the IUserControllerInterface
type UserControllerMock struct {
	users []model.User
	err   error
}

// NewUserControllerMock returns a new mock implementing IUserController interface
func NewUserControllerMock(users []model.User, err error) UserControllerMock {
	return UserControllerMock{
		users,
		err,
	}
}

// GetCompanyUsers mocks the controller function for retrieving a list of users given the company ID
func (ucm UserControllerMock) GetCompanyUsers(companyID string) ([]model.User, error) {
	return ucm.users, ucm.err
}
