package poststore

import (
	"posts_api/graph/model"
	"time"
)

// CreateComment creates a new comment in the store.
// If user or post don't exist, an error is returned.
func (ts *InMemoryStorage) CreateComment(userId int, postId int, text string, date time.Time) (int, error) {
	ts.Lock()
	defer ts.Unlock()

	user, err := ts.UserExists(userId)
	if err != nil {
		return 0, err
	}
	post, err := ts.PostExists(postId)
	if err != nil {
		return 0, err
	}

	comment := &model.Comment{
		ID:   ts.nextCommentId,
		User: user,
		Text: text,
		Date: date,
	}

	ts.comments[ts.nextCommentId] = comment
	ts.nextCommentId++
	ts.posts[post.ID].Comments = append(ts.posts[post.ID].Comments, comment)
	return comment.ID, nil
}

// CreateCommentToComment creates new comment to parent comment in the store.
// If user or parent comment don't exist, an error is returned
func (ts *InMemoryStorage) CreateCommentToComment(userId int, commentId int, text string, date time.Time) (int, error) {
	ts.Lock()
	defer ts.Unlock()

	user, err := ts.UserExists(userId)
	if err != nil {
		return 0, err
	}

	parentComment, err := ts.CommentExists(commentId)
	if err != nil {
		return 0, err
	}

	comment := &model.Comment{
		ID:   ts.nextCommentId,
		User: user,
		Text: text,
		Date: date,
	}

	ts.comments[ts.nextCommentId] = comment
	ts.nextCommentId++
	ts.comments[parentComment.ID].Comments = append(ts.comments[parentComment.ID].Comments, comment)
	return comment.ID, nil
}

// GetComment retrieves a comment from the store, by id. If no such id exists, an
// error is returned.
func (ts *InMemoryStorage) GetComment(id int) (*model.Comment, error) {
	ts.Lock()
	defer ts.Unlock()

	comment, err := ts.CommentExists(id)
	return comment, err
}

// DeleteComment deletes the comment with the given id. If no such id exists, an error
// is returned.
func (ts *InMemoryStorage) DeleteComment(id int) (bool, error) {
	ts.Lock()
	defer ts.Unlock()

	_, err := ts.CommentExists(id)
	if err != nil {
		return false, err
	}

	delete(ts.comments, id)
	return true, nil
}

// GetAllComments returns all comments in the store.
func (ts *InMemoryStorage) GetAllComments() ([]*model.Comment, error) {
	ts.Lock()
	defer ts.Unlock()

	allComments := make([]*model.Comment, 0, len(ts.comments))
	for _, comment := range ts.comments {
		allComments = append(allComments, comment)
	}
	return allComments, nil
}
