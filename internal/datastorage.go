package internal

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"os"
	"posts_api/graph/model"
	"posts_api/internal/postgres"
	"posts_api/internal/poststore"
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

const defaultDbPort = "5432"
const defaultDbHost = "localhost"
const defaultDbUser = "postgres"
const defaultDbPassword = ""
const defaultDbName = "postgres"

func NewDataStorage(storageType string) (DataStorage, *sql.DB) {
	switch storageType {
	case "in-memory":
		store := poststore.New()
		return store, nil
	case "postgres":
		port, exists := os.LookupEnv("POSTGRES_PORT")
		if !exists {
			port = defaultDbPort
		}
		host, exists := os.LookupEnv("POSTGRES_HOST")
		if !exists {
			port = defaultDbHost
		}
		user, exists := os.LookupEnv("POSTGRES_USER")
		if !exists {
			port = defaultDbUser
		}
		password, exists := os.LookupEnv("POSTGRES_PASSWORD")
		if !exists {
			port = defaultDbPassword
		}
		dbname, exists := os.LookupEnv("POSTGRES_DB")
		if !exists {
			port = defaultDbName
		}
		connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

		// connStr := "host=postgres port=5432 user=api_tester password=testing dbname=postApi sslmode=disable"
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
