package postgres

import (
	"GraphQL_api/graph/model"
	"database/sql"
	"errors"
)

func (s *PostgresStorage) CreateUser(firstName string, lastName string) (int, error) {
	stmt := `INSERT INTO users ("firstName", "lastName") VALUES($1,$2) RETURNING id`

	lastInsertId := 0
	err := s.DB.QueryRow(stmt, firstName, lastName).Scan(&lastInsertId)

	id := lastInsertId
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (s *PostgresStorage) GetUser(id int) (*model.User, error) {
	stmt := `SELECT id, "firstName", "lastName" FROM users WHERE id = $1`

	row := s.DB.QueryRow(stmt, id)

	usr := &model.User{}

	err := row.Scan(&usr.ID, &usr.FirstName, &usr.LastName)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNoRecord
		} else {
			return nil, err
		}
	}

	return usr, nil
}

func (s *PostgresStorage) DeleteUser(id int) (bool, error) {
	stmt := `DELETE FROM users WHERE id=$1`

	err := s.DB.QueryRow(stmt, id).Err()

	if err != nil {
		return false, err
	}

	return true, nil
}

func (s *PostgresStorage) GetAllUsers() ([]*model.User, error) {
	stmt := `SELECT id, "firstName", "lastName" FROM users`

	rows, err := s.DB.Query(stmt)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var users []*model.User

	for rows.Next() {
		user := &model.User{}
		err = rows.Scan(&user.ID, &user.FirstName, &user.LastName)

		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}
