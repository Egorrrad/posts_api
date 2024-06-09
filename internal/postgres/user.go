package postgres

import (
	"GraphQL_api/graph/model"
	"database/sql"
)

type UserModel struct {
	DB *sql.DB
}

func (s *PostgresStorage) CreateUser(firstName string, lastName string) int {
	//TODO implement me
	panic("implement me")
}

func (s *PostgresStorage) GetUser(id int) (*model.User, error) {
	//TODO implement me
	panic("implement me")
}

func (s *PostgresStorage) DeleteUser(id int) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func (s *PostgresStorage) GetAllUsers() ([]*model.User, error) {
	//TODO implement me
	panic("implement me")
}
