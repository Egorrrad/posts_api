package postgres

import (
	"GraphQL_api/graph/model"
	"database/sql"
	"time"
)

type PostModel struct {
	DB *sql.DB
}

func (s *PostgresStorage) CreatePost(user *model.User, text string, date time.Time) int {
	//TODO implement me
	panic("implement me")
}

func (s *PostgresStorage) AllowCommentPost(id int, user *model.User, allow bool) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func (s *PostgresStorage) GetPost(id int) (*model.Post, error) {
	//TODO implement me
	panic("implement me")
}

func (s *PostgresStorage) DeletePost(id int) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func (s *PostgresStorage) GetAllPosts() ([]*model.Post, error) {
	//TODO implement me
	panic("implement me")
}
