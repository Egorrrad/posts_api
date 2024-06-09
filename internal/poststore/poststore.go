package poststore

import (
	"GraphQL_api/graph/model"
	"sync"
)

type InMemoryStorage struct {
	sync.Mutex

	posts      map[int]*model.Post
	nextPostId int

	comments      map[int]*model.Comment
	nextCommentId int

	users      map[int]*model.User
	nextUserId int
}

func New() *InMemoryStorage {
	ts := &InMemoryStorage{}
	ts.posts = make(map[int]*model.Post)
	ts.nextPostId = 0

	ts.comments = make(map[int]*model.Comment)
	ts.nextCommentId = 0

	ts.users = make(map[int]*model.User)
	ts.nextUserId = 0

	return ts
}
