package mysql

import "database/sql"

// MySQLStore implements the IUserStore interface
type MySQLStore struct {
	db sql.DB
}

// NewMySQLStore builds a new implementation of the MySQL store.
// Given a database connection as input, returns a new instance
// of the store, to be leveraged within the mysql package
func NewMySQLStore(db sql.DB) MySQLStore {
	return MySQLStore{
		db: db,
	}
}
