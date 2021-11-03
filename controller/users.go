package controller

import (
	"github.com/poltro90/tech-panel/model"
	"github.com/poltro90/tech-panel/store"
)

// UserController is the implementation for the IUserController interface
type UserController struct {
	userStore store.IUserStore
}

// NewUserController returns a new UserController given a UserStore
func NewUserController(userStore store.IUserStore) UserController {
	return UserController{
		userStore: userStore,
	}
}

// GetCompanyUsers returns the list of users for a company given it's ID
func (uc UserController) GetCompanyUsers(companyID string) ([]model.User, error) {
	return uc.userStore.GetUsersBycompanyID(companyID)
}
