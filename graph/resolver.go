package graph

import (
	"posts_api/graph/model"
	"posts_api/internal"
	"sync"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	Store internal.DataStorage

	Comments    map[int][]chan *model.Comment
	Mu          sync.Mutex
	NewComments chan *model.Comment
}

func (r *Resolver) HandleNewComments() {
	for comment := range r.NewComments {
		r.Mu.Lock()
		postComments := r.Comments[comment.PostID]
		for _, c := range postComments {
			c <- comment
		}
		r.Mu.Unlock()
	}
}
