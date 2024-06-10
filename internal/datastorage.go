package internal

import (
	"GraphQL_api/graph/model"
	"GraphQL_api/internal/postgres"
	"GraphQL_api/internal/poststore"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"time"
)

type DataStorage interface {
	CreatePost(user *model.User, text string, date time.Time) (int, error)
	AllowCommentPost(id int, user *model.User, allow bool) (bool, error)
	GetPost(id int) (*model.Post, error)
	DeletePost(id int) (bool, error)
	GetAllPosts() ([]*model.Post, error)
	CreateComment(userId int, postId int, text string, date time.Time) (int, error)
	CreateCommentToComment(userId int, commentId int, text string, date time.Time) (int, error)
	GetComment(id int) (*model.Comment, error)
	DeleteComment(id int) (bool, error)
	GetAllComments() ([]*model.Comment, error)
	CreateUser(firstName string, lastName string) (int, error)
	GetUser(id int) (*model.User, error)
	DeleteUser(id int) (bool, error)
	GetAllUsers() ([]*model.User, error)
}

func NewDataStorage(storageType string) (DataStorage, *sql.DB) {
	switch storageType {
	case "in-memory":
		store := poststore.New()
		return store, nil
	case "postgres":
		connStr := "host=localhost port=5432 user=api_tester password=testing dbname=postApi sslmode=disable"
		store, err := postgres.OpenDB(connStr)
		if err != nil {
			fmt.Println(err)
		}
		// defer postgres.CloseDB(store)
		return store, store.DB
	default:
		panic("Unsupported storage type")
	}
}
