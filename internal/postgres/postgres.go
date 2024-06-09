package postgres

import (
	"GraphQL_api/graph/model"
	"time"
)

type PostgresStorage struct {
	// реализация для PostgreSQL
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

func (s *PostgresStorage) CreateComment(userId int, postId int, text string, date time.Time) (int, error) {
	//TODO implement me
	panic("implement me")
}

func (s *PostgresStorage) CreateCommentToComment(userId int, commentId int, text string, date time.Time) (int, error) {
	//TODO implement me
	panic("implement me")
}

func (s *PostgresStorage) GetComment(id int) (*model.Comment, error) {
	//TODO implement me
	panic("implement me")
}

func (s *PostgresStorage) DeleteComment(id int) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func (s *PostgresStorage) GetAllComments() ([]*model.Comment, error) {
	//TODO implement me
	panic("implement me")
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
