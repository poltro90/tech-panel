package controller

import "github.com/poltro90/tech-panel/model"

// IUserController is the interface for the user controller
type IUserController interface {
	// GetCompanyUsers returns a list of users given the company ID
	GetCompanyUsers(companyID string) ([]model.User, error)
}
