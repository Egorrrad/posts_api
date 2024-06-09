package postgres

import (
	"GraphQL_api/graph/model"
	"database/sql"
	"time"
)

type CommentModel struct {
	DB *sql.DB
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
