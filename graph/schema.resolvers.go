package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.47-dev

import (
	"context"
	"encoding/base64"
	"fmt"
	"posts_api/graph/model"
	"strconv"
	"strings"
	"time"
)

// Edges is the resolver for the edges field.
func (r *commentsConnectionResolver) Edges(ctx context.Context, obj *model.CommentsConnection) ([]*model.CommentsEdge, error) {
	edges := make([]*model.CommentsEdge, obj.To-obj.From)

	for i := range edges {
		edges[i] = &model.CommentsEdge{
			Node:   obj.Comments[obj.From+i],
			Cursor: model.EncodeCursor(obj.From + i),
		}
	}

	return edges, nil
}

// CreatePost is the resolver for the createPost field.
func (r *mutationResolver) CreatePost(ctx context.Context, input model.NewPost) (int, error) {
	user, err := r.Store.GetUser(input.UserID)
	if err != nil {
		return 0, err
	}
	currentTime := time.Now()
	return r.Store.CreatePost(user, input.Text, currentTime)
}

// DeletePost is the resolver for the deletePost field.
func (r *mutationResolver) DeletePost(ctx context.Context, id int) (bool, error) {
	return r.Store.DeletePost(id)
}

// CreateComment is the resolver for the createComment field.
func (r *mutationResolver) CreateComment(ctx context.Context, input model.NewComment) (int, error) {
	currentTime := time.Now()
	postId := input.PostID
	post, err := r.Store.GetPost(postId)
	if err != nil {
		return 0, err
	}
	if !post.AllowComment {
		return 0, fmt.Errorf("can't comment this post")
	}

	commentId, err := r.Store.CreateComment(input.UserID, input.PostID, input.Text, currentTime)
	if err != nil {
		return 0, err
	}

	newComment, err := r.Store.GetComment(commentId)
	if err != nil {
		return 0, err
	}

	r.NewComments <- newComment

	return commentId, nil
}

// CreateCommentToComment is the resolver for the createCommentToComment field.
func (r *mutationResolver) CreateCommentToComment(ctx context.Context, input model.NewCommentToComment) (int, error) {
	currentTime := time.Now()
	return r.Store.CreateCommentToComment(input.UserID, input.CommentID, input.Text, currentTime)
}

// DeleteComment is the resolver for the deleteComment field.
func (r *mutationResolver) DeleteComment(ctx context.Context, id int) (bool, error) {
	return r.Store.DeleteComment(id)
}

// AllowComments is the resolver for the allowComments field.
func (r *mutationResolver) AllowComments(ctx context.Context, input model.AllowComment) (bool, error) {
	user, err := r.Store.GetUser(input.UserID)
	if err != nil {
		return false, err
	}
	return r.Store.AllowCommentPost(input.PostID, user, input.AllowComment)
}

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, firstName string, lastName string) (int, error) {
	return r.Store.CreateUser(firstName, lastName)
}

// Edges is the resolver for the edges field.
func (r *postsConnectionResolver) Edges(ctx context.Context, obj *model.PostsConnection) ([]*model.PostsEdge, error) {
	edges := make([]*model.PostsEdge, obj.To-obj.From)

	for i := range edges {
		edges[i] = &model.PostsEdge{
			Node:   obj.Posts[obj.From+i],
			Cursor: model.EncodeCursor(obj.From + i),
		}
	}

	return edges, nil
}

// GetPost is the resolver for the getPost field.
func (r *queryResolver) GetPost(ctx context.Context, id int) (*model.Post, error) {
	return r.Store.GetPost(id)
}

// GetUser is the resolver for the getUser field.
func (r *queryResolver) GetUser(ctx context.Context, id int) (*model.User, error) {
	return r.Store.GetUser(id)
}

// GetComment is the resolver for the getComment field.
func (r *queryResolver) GetComment(ctx context.Context, id int) (*model.Comment, error) {
	return r.Store.GetComment(id)
}

// UsersConnection is the resolver for the usersConnection field.
func (r *queryResolver) UsersConnection(ctx context.Context, first *int, after *string) (*model.UsersConnection, error) {
	allUsers, _ := r.Store.GetAllUsers()

	from := 0

	if after != nil {
		b, err := base64.StdEncoding.DecodeString(*after)

		if err != nil {
			return nil, err
		}

		i, err := strconv.Atoi(strings.TrimPrefix(string(b), "cursor"))

		if err != nil {
			return nil, err
		}

		from = i
	}

	to := len(allUsers)

	if to == 0 && after != nil {
		return nil, fmt.Errorf("cursor %s not exsists", *after)
	}

	if first != nil {
		to = from + *first

		if to > len(allUsers) {
			to = len(allUsers)
		}
	}

	return &model.UsersConnection{
		Users: allUsers,
		From:  from,
		To:    to,
	}, nil
}

// PostsConnection is the resolver for the postsConnection field.
func (r *queryResolver) PostsConnection(ctx context.Context, first *int, after *string) (*model.PostsConnection, error) {
	allPosts, _ := r.Store.GetAllPosts()

	from := 0
	if after != nil {
		b, err := base64.StdEncoding.DecodeString(*after)

		if err != nil {
			return nil, err
		}

		i, err := strconv.Atoi(strings.TrimPrefix(string(b), "cursor"))

		if err != nil {
			return nil, err
		}

		from = i
	}

	to := len(allPosts)

	if to == 0 && after != nil {
		return nil, fmt.Errorf("cursor %s not exsists", *after)
	}

	if first != nil {
		to = from + *first

		if to > len(allPosts) {
			to = len(allPosts)
		}
	}

	return &model.PostsConnection{
		Posts: allPosts,
		From:  from,
		To:    to,
	}, nil
}

// CommentsConnection is the resolver for the commentsConnection field.
func (r *queryResolver) CommentsConnection(ctx context.Context, first *int, after *string) (*model.CommentsConnection, error) {
	allComments, _ := r.Store.GetAllComments()

	from := 0
	if after != nil {
		b, err := base64.StdEncoding.DecodeString(*after)

		if err != nil {
			return nil, err
		}

		i, err := strconv.Atoi(strings.TrimPrefix(string(b), "cursor"))

		if err != nil {
			return nil, err
		}

		from = i
	}

	to := len(allComments)

	if to == 0 && after != nil {
		return nil, fmt.Errorf("cursor %s not exsists", *after)
	}

	if first != nil {
		to = from + *first

		if to > len(allComments) {
			to = len(allComments)
		}
	}

	return &model.CommentsConnection{
		Comments: allComments,
		From:     from,
		To:       to,
	}, nil
}

// CommentPublished is the resolver for the commentPublished field.
func (r *subscriptionResolver) CommentPublished(ctx context.Context, postID int) (<-chan *model.Comment, error) {
	r.Mu.Lock()
	defer r.Mu.Unlock()

	channel := make(chan *model.Comment, 10)
	_, ok := r.Comments[postID]
	if !ok {
		r.Comments[postID] = make([]chan *model.Comment, 0)
	}
	r.Comments[postID] = append(r.Comments[postID], channel)

	go func() {
		<-ctx.Done()
		r.Mu.Lock()
		defer r.Mu.Unlock()
		for i, c := range r.Comments[postID] {
			if c == channel {
				r.Comments[postID] = append(r.Comments[postID][:i], r.Comments[postID][i+1:]...)
				close(c)
				break
			}
		}
	}()

	return channel, nil
}

// Edges is the resolver for the edges field.
func (r *usersConnectionResolver) Edges(ctx context.Context, obj *model.UsersConnection) ([]*model.UsersEdge, error) {
	edges := make([]*model.UsersEdge, obj.To-obj.From)

	for i := range edges {
		edges[i] = &model.UsersEdge{
			Node:   obj.Users[obj.From+i],
			Cursor: model.EncodeCursor(obj.From + i),
		}
	}

	return edges, nil
}

// CommentsConnection returns CommentsConnectionResolver implementation.
func (r *Resolver) CommentsConnection() CommentsConnectionResolver {
	return &commentsConnectionResolver{r}
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// PostsConnection returns PostsConnectionResolver implementation.
func (r *Resolver) PostsConnection() PostsConnectionResolver { return &postsConnectionResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

// Subscription returns SubscriptionResolver implementation.
func (r *Resolver) Subscription() SubscriptionResolver { return &subscriptionResolver{r} }

// UsersConnection returns UsersConnectionResolver implementation.
func (r *Resolver) UsersConnection() UsersConnectionResolver { return &usersConnectionResolver{r} }

type commentsConnectionResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type postsConnectionResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type subscriptionResolver struct{ *Resolver }
type usersConnectionResolver struct{ *Resolver }
