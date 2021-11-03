package model

import "time"

// User is representing a user
type User struct {
	FirstName string
	LastName  string
	DOB       time.Time
}
