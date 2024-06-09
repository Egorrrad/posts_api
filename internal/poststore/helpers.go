package poststore

import (
	"GraphQL_api/graph/model"
	"fmt"
)

// UserExists checks user in the store, by id.
// If no such user, error is returned
func (ts *InMemoryStorage) UserExists(id int) (*model.User, error) {
	t, ok := ts.users[id]
	if !ok {
		return nil, fmt.Errorf("user with id=%d not found", id)
	}
	return t, nil
}

// PostExists checks user in the store, by id.
// If no such post, error is returned
func (ts *InMemoryStorage) PostExists(id int) (*model.Post, error) {
	t, ok := ts.posts[id]
	if !ok {
		return nil, fmt.Errorf("post with id=%d not found", id)
	}
	return t, nil
}

// CommentExists checks user in the store, by id.
// If no such comment, error is returned
func (ts *InMemoryStorage) CommentExists(id int) (*model.Comment, error) {
	t, ok := ts.comments[id]
	if !ok {
		return nil, fmt.Errorf("comment with id=%d not found", id)
	}
	return t, nil
}

// isPostBelongsToUser checks post to belong to user.
// If user don't belong to post, error is returned
func (ts *InMemoryStorage) isPostBelongsToUser(post *model.Post, user *model.User) (bool, error) {
	if post.User.ID == user.ID {
		return true, nil
	}
	return false, fmt.Errorf("you can't edit post with id=%d", post.ID)
}
