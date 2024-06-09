package model

import (
	"encoding/base64"
	"fmt"
)

type UsersConnection struct {
	Users []*User
	From  int
	To    int
}

func (u *UsersConnection) TotalCount() int {
	return len(u.Users)
}

func (u *UsersConnection) PageInfo() PageInfo {
	return PageInfo{
		StartCursor: EncodeCursor(u.From),
		EndCursor:   EncodeCursor(u.To - 1),
		HasNextPage: u.To < len(u.Users),
	}
}

type PostsConnection struct {
	Posts []*Post
	From  int
	To    int
}

func (p *PostsConnection) TotalCount() int {
	return len(p.Posts)
}

func (p *PostsConnection) PageInfo() PageInfo {
	return PageInfo{
		StartCursor: EncodeCursor(p.From),
		EndCursor:   EncodeCursor(p.To - 1),
		HasNextPage: p.To < len(p.Posts),
	}
}

type CommentsConnection struct {
	Comments []*Comment
	From     int
	To       int
}

func (c *CommentsConnection) TotalCount() int {
	return len(c.Comments)
}

func (c *CommentsConnection) PageInfo() PageInfo {
	return PageInfo{
		StartCursor: EncodeCursor(c.From),
		EndCursor:   EncodeCursor(c.To - 1),
		HasNextPage: c.To < len(c.Comments),
	}
}

func EncodeCursor(i int) string {
	return base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("cursor%d", i+1)))
}
