package internal

import (
	"GraphQL_api/graph/model"
	"GraphQL_api/internal/postgres"
	"GraphQL_api/internal/poststore"
	"time"
)

type DataStorage interface {
	CreatePost(user *model.User, text string, date time.Time) int
	AllowCommentPost(id int, user *model.User, allow bool) (bool, error)
	GetPost(id int) (*model.Post, error)
	DeletePost(id int) (bool, error)
	GetAllPosts() ([]*model.Post, error)
	CreateComment(userId int, postId int, text string, date time.Time) (int, error)
	CreateCommentToComment(userId int, commentId int, text string, date time.Time) (int, error)
	GetComment(id int) (*model.Comment, error)
	DeleteComment(id int) (bool, error)
	GetAllComments() ([]*model.Comment, error)
	CreateUser(firstName string, lastName string) int
	GetUser(id int) (*model.User, error)
	DeleteUser(id int) (bool, error)
	GetAllUsers() ([]*model.User, error)
}

func NewDataStorage(storageType string) DataStorage {
	switch storageType {
	case "in-memory":
		store := poststore.New()
		return store
	case "postgres":
		return &postgres.PostgresStorage{}
	default:
		panic("Unsupported storage type")
	}
}
