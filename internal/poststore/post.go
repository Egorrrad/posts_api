package poststore

import (
	"GraphQL_api/graph/model"
	"time"
)

// CreatePost creates a new post in the store.
func (ts *InMemoryStorage) CreatePost(user *model.User, text string, date time.Time) int {
	ts.Lock()
	defer ts.Unlock()

	post := &model.Post{
		ID:           ts.nextPostId,
		User:         user,
		Text:         text,
		Date:         date,
		AllowComment: true,
	}

	ts.posts[ts.nextPostId] = post
	ts.nextPostId++
	return post.ID
}

// AllowCommentPost allows users to comment on this post.
// If post doesn't belong to this user or no such post exists, error is returned.
func (ts *InMemoryStorage) AllowCommentPost(id int, user *model.User, allow bool) (bool, error) {
	ts.Lock()
	defer ts.Unlock()

	t, err := ts.PostExists(id)
	if err != nil {
		return false, err
	}
	_, err = ts.isPostBelongsToUser(t, user)
	if err != nil {
		return false, err
	}

	ts.posts[id].AllowComment = allow
	return true, nil
}

// GetPost retrieves a post from the store, by id. If no such id exists, an
// error is returned.
func (ts *InMemoryStorage) GetPost(id int) (*model.Post, error) {
	ts.Lock()
	defer ts.Unlock()

	t, err := ts.PostExists(id)
	return t, err
}

// DeletePost deletes the post with the given id. If no such id exists, an error
// is returned.
func (ts *InMemoryStorage) DeletePost(id int) (bool, error) {
	ts.Lock()
	defer ts.Unlock()

	_, err := ts.PostExists(id)
	if err != nil {
		return false, err
	}

	delete(ts.posts, id)
	return true, nil
}

// GetAllPosts returns all the posts in the store.
func (ts *InMemoryStorage) GetAllPosts() ([]*model.Post, error) {
	ts.Lock()
	defer ts.Unlock()

	allPosts := make([]*model.Post, 0, len(ts.posts))
	for _, post := range ts.posts {
		allPosts = append(allPosts, post)
	}
	return allPosts, nil
}
