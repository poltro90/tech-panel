package mysql

import "github.com/poltro90/tech-panel/model"

// GetUsersBycompanyID retrieves from a MySQL database the list of users for a given company
func (s MySQLStore) GetUsersBycompanyID(companyID string) ([]model.User, error) {
	query := getUserByCompanyIDQuery()
	rows, err := s.db.Query(query, companyID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users = make([]model.User, 0)
	for rows.Next() {
		var user = model.User{}
		err := rows.Scan(&user.FirstName, &user.LastName, &user.DOB)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

/*
 * This is an unexported function and doesn't need to be commented.
 * It cannot be used outside the package mysql, where it has been defined.
 */
func getUserByCompanyIDQuery() string {
	return "SELECT first_name, last_name, dob FROM users WHERE company_id = ?"
}
