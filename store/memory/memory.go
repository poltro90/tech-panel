package memory

import "github.com/poltro90/tech-panel/model"

// MemoryStore is the in-memory store implementation for IUserStore interface
type MemoryStore struct {
	companyUsers map[string][]model.User
}

// AddCompanyUser appends a new user to the in-memory store for company users
func (ms MemoryStore) AddCompanyUser(companyID string, user model.User) {
	if ms.companyUsers[companyID] == nil {
		ms.companyUsers[companyID] = make([]model.User, 0)
	}
	ms.companyUsers[companyID] = append(ms.companyUsers[companyID], user)
}

// GetUsersBycompanyID returns the list of users
func (ms MemoryStore) GetUsersBycompanyID(companyID string) ([]model.User, error) {
	return ms.companyUsers[companyID], nil
}
