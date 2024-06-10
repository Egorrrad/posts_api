package poststore

import (
	"github.com/stretchr/testify/assert"
	_ "os"
	"testing"
	"time"
)

var store *InMemoryStorage

func TestPostStore(t *testing.T) {
	// Создаем in-memory хранилище перед выполнением всех тестов
	store = New()
}

func TestPostStore_CreateUser(t *testing.T) {
	// Create a store and a single user.
	testCases := []struct {
		name       string
		testScript func()
	}{
		{"create user", func() {
			id, err := store.CreateUser("tester", "admin")
			assert.NoError(t, err)
			user, ok := store.users[id]

			assert.Equal(t, true, ok)
			assert.Equal(t, id, user.ID)
		}},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.testScript()
		})
	}
}

func TestPostStore_GetUser(t *testing.T) {
	testCases := []struct {
		name       string
		testScript func()
	}{
		{"get user in store", func() {
			id := 0
			user, err := store.GetUser(id)

			assert.NoError(t, err)
			assert.Equal(t, id, user.ID)
			assert.Equal(t, "tester", user.FirstName)
			assert.Equal(t, "admin", user.LastName)
		}},
		{"get user not in store", func() {
			id := 20
			user, err := store.GetUser(id)

			assert.Nil(t, user)
			assert.Error(t, err)
		}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.testScript()
		})
	}
}

func TestPostStore_DeleteUser(t *testing.T) {
	testCases := []struct {
		name       string
		testScript func()
	}{
		{"delete user in store", func() {
			id := 0
			ok, err := store.DeleteUser(id)

			assert.NoError(t, err)
			assert.Equal(t, true, ok)
			user, ok := store.users[id]
			assert.Equal(t, false, ok)
			assert.Nil(t, user)
		}},
		{"delete user not in store", func() {
			id := 20
			ok, err := store.DeleteUser(id)

			assert.Equal(t, false, ok)
			assert.Error(t, err)
		}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.testScript()
		})
	}
}

func TestPostStore_GetAllUsers(t *testing.T) {
	testCases := []struct {
		name       string
		testScript func()
	}{
		{"get users in empty store", func() {
			users, err := store.GetAllUsers()

			assert.NoError(t, err)
			assert.Empty(t, users)
		}},
		{"get users in store", func() {
			id, err := store.CreateUser("tester001", "admin2")
			assert.NoError(t, err)
			users, err := store.GetAllUsers()

			assert.NoError(t, err)
			oneUser := users[0]
			assert.Equal(t, id, oneUser.ID)
			assert.Equal(t, "tester001", oneUser.FirstName)
			assert.Equal(t, "admin2", oneUser.LastName)
		}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.testScript()
		})
	}
}

func TestPostStore_CreatePost(t *testing.T) {
	testCases := []struct {
		name       string
		testScript func()
	}{
		{"create post", func() {
			user, _ := store.GetUser(1)
			text := "first post"
			currentTime := time.Now()
			id, err := store.CreatePost(user, text, currentTime)
			assert.NoError(t, err)
			post, ok := store.posts[id]

			assert.Equal(t, true, ok)
			assert.Equal(t, id, post.ID)
			assert.WithinDuration(t, time.Now(), post.Date, time.Second)
			assert.Equal(t, text, post.Text)
			assert.Equal(t, true, post.AllowComment)
		}},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.testScript()
		})
	}
}

func TestPostStore_GetPost(t *testing.T) {
	testCases := []struct {
		name       string
		testScript func()
	}{
		{"get post in store", func() {
			id := 0
			text := "first post"
			post, err := store.GetPost(id)

			assert.NoError(t, err)
			assert.Equal(t, id, post.ID)
			assert.Equal(t, text, post.Text)
			assert.Equal(t, true, post.AllowComment)
		}},
		{"get post not in store", func() {
			id := 20
			post, err := store.GetPost(id)

			assert.Nil(t, post)
			assert.Error(t, err)
		}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.testScript()
		})
	}
}

func TestPostStore_DeletePost(t *testing.T) {
	testCases := []struct {
		name       string
		testScript func()
	}{
		{"delete post in store", func() {
			id := 0
			ok, err := store.DeletePost(id)

			assert.NoError(t, err)
			assert.Equal(t, true, ok)
			post, ok := store.posts[id]
			assert.Equal(t, false, ok)
			assert.Nil(t, post)
		}},
		{"delete post not in store", func() {
			id := 20
			ok, err := store.DeletePost(id)

			assert.Equal(t, false, ok)
			assert.Error(t, err)
		}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.testScript()
		})
	}
}

func TestPostStore_GetAllPosts(t *testing.T) {
	testCases := []struct {
		name       string
		testScript func()
	}{
		{"get posts in empty store", func() {
			posts, err := store.GetAllPosts()

			assert.NoError(t, err)
			assert.Empty(t, posts)
		}},
		{"get posts in store", func() {
			user, _ := store.GetUser(1)
			text := "first post"
			currentTime := time.Now()
			id, err := store.CreatePost(user, text, currentTime)
			assert.NoError(t, err)
			posts, err := store.GetAllPosts()

			assert.NoError(t, err)
			onePost := posts[0]
			assert.Equal(t, id, onePost.ID)
			assert.Equal(t, text, onePost.Text)
		}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.testScript()
		})
	}
}

func TestPostStore_AllowCommentPost(t *testing.T) {
	testCases := []struct {
		name       string
		testScript func()
	}{
		{"allow comment not user post", func() {
			id, err := store.CreateUser("other user", "none")
			assert.NoError(t, err)
			user, err := store.GetUser(id)
			postId := 1
			ok, err := store.AllowCommentPost(postId, user, false)

			assert.Error(t, err)
			assert.Equal(t, false, ok)
		}},
		{"allow comment user post", func() {
			id := 1
			user, err := store.GetUser(id)
			postId := 1
			ok, err := store.AllowCommentPost(postId, user, false)

			assert.NoError(t, err)
			assert.Equal(t, true, ok)
			post, err := store.GetPost(postId)
			assert.Equal(t, false, post.AllowComment)
		}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.testScript()
		})
	}
}

func TestPostStore_CreateComment(t *testing.T) {
	testCases := []struct {
		name       string
		testScript func()
	}{
		{"create comment", func() {
			userId := 1
			postId := 1
			text := "first comment"
			currentTime := time.Now()
			id, err := store.CreateComment(userId, postId, text, currentTime)

			assert.NoError(t, err)
			comment, ok := store.comments[id]

			assert.Equal(t, true, ok)
			assert.Equal(t, id, comment.ID)
			assert.WithinDuration(t, time.Now(), comment.Date, time.Second)
			assert.Equal(t, text, comment.Text)
			assert.Equal(t, postId, comment.PostID)

			post, _ := store.GetPost(postId)
			assert.Equal(t, comment.ID, post.Comments[0].ID)
		}},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.testScript()
		})
	}
}

func TestPostStore_CreateCommentToComment(t *testing.T) {
	testCases := []struct {
		name       string
		testScript func()
	}{
		{"create comment", func() {
			userId := 1
			commentId := 0
			postId := 1
			text := "second comment (to comment)"
			currentTime := time.Now()
			id, err := store.CreateCommentToComment(userId, commentId, text, currentTime)

			assert.NoError(t, err)
			comment, ok := store.comments[id]

			assert.Equal(t, true, ok)
			assert.Equal(t, id, comment.ID)
			assert.WithinDuration(t, time.Now(), comment.Date, time.Second)
			assert.Equal(t, text, comment.Text)
			assert.Equal(t, postId, comment.PostID)

			parentComment, _ := store.GetComment(commentId)
			assert.Equal(t, comment.ID, parentComment.Comments[0].ID) // ????
		}},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.testScript()
		})
	}
}

func TestPostStore_GetComment(t *testing.T) {
	testCases := []struct {
		name       string
		testScript func()
	}{
		{"get comment in store", func() {
			id := 0
			userId := 1
			postId := 1
			text := "first comment"
			comment, err := store.GetComment(id)

			assert.NoError(t, err)
			assert.Equal(t, id, comment.ID)
			assert.Equal(t, text, comment.Text)
			assert.Equal(t, userId, comment.User.ID)
			assert.Equal(t, postId, comment.PostID)

			post, _ := store.GetPost(postId)
			assert.Equal(t, comment.ID, post.Comments[0].ID)
		}},
		{"get comment not in store", func() {
			id := 20
			comment, err := store.GetComment(id)

			assert.Nil(t, comment)
			assert.Error(t, err)
		}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.testScript()
		})
	}
}

func TestPostStore_DeleteComment(t *testing.T) {
	testCases := []struct {
		name       string
		testScript func()
	}{
		{"delete comment in store", func() {
			id := 0
			ok, err := store.DeleteComment(id)

			assert.NoError(t, err)
			assert.Equal(t, true, ok)
			comment, ok := store.comments[id]
			assert.Equal(t, false, ok)
			assert.Nil(t, comment)
		}},
		{"delete comment not in store", func() {
			id := 20
			ok, err := store.DeleteComment(id)

			assert.Equal(t, false, ok)
			assert.Error(t, err)
		}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.testScript()
		})
	}
}

func TestPostStore_GetAllComments(t *testing.T) {
	testCases := []struct {
		name       string
		testScript func()
	}{
		{"get comments in store", func() {
			id := 1
			userId := 1
			postId := 1
			text := "second comment (to comment)"
			comments, err := store.GetAllComments()

			assert.NoError(t, err)
			oneComment := comments[0]

			assert.Equal(t, id, oneComment.ID)
			assert.Equal(t, text, oneComment.Text)
			assert.Equal(t, userId, oneComment.User.ID)
			assert.Equal(t, postId, oneComment.PostID)

		}},
		{"get comments in empty store", func() {
			id := 1
			_, err := store.DeleteComment(id)
			comments, err := store.GetAllComments()

			assert.NoError(t, err)
			assert.Empty(t, comments)
		}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.testScript()
		})
	}
}
