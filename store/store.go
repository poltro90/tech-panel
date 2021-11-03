package store

import "github.com/poltro90/tech-panel/model"

type IUserStore interface {
	GetUsersBycompanyID(companyID string) ([]model.User, error)
}
