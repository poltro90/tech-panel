package handler

import (
	"encoding/json"
	"net/http"
)

// GetUsers serves the request for retrieving the list of users for a given company
func (h UserHandler) GetUsers(w http.ResponseWriter, req *http.Request) {
	// Read params from the request
	companyID := "abcdef0123456789"

	// Controller
	users, err := h.userController.GetCompanyUsers(companyID)
	if err != nil {
		panic("orror error")
	}

	// Prepare the response
	response, err := json.Marshal(users)
	if err != nil {
		panic("orror error")
	}

	// HTTP response
	w.Write(response)
}
