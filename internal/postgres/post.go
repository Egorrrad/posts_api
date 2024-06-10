package postgres

import (
	"GraphQL_api/graph/model"
	"database/sql"
	"errors"
	"time"
)

func (s *PostgresStorage) CreatePost(user *model.User, text string, date time.Time) (int, error) {
	stmt := `INSERT INTO posts (user_id, text) VALUES($1,$2) RETURNING id`

	lastInsertId := 0
	err := s.DB.QueryRow(stmt, user.ID, text).Scan(&lastInsertId)

	id := lastInsertId
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (s *PostgresStorage) AllowCommentPost(id int, user *model.User, allow bool) (bool, error) {
	stmt := `UPDATE posts SET "allowComment" = $1
                WHERE id = $2;`

	// проверка существования коммента
	// проверка того, что пост принадлежит пользователю
	err := s.DB.QueryRow(stmt, allow, id).Err()

	if err != nil {
		return false, err
	}
	return true, nil
}

func (s *PostgresStorage) GetPost(id int) (*model.Post, error) {
	stmt := `SELECT id, user_id, text, "allowComment", created_at FROM posts WHERE id = $1`

	row := s.DB.QueryRow(stmt, id)

	post := &model.Post{}

	var userId int
	err := row.Scan(&post.ID, &userId, &post.Text, &post.AllowComment, &post.Date)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNoRecord
		} else {
			return nil, err
		}
	}

	post.User, err = s.GetUser(userId)
	if err != nil {
		return nil, err
	}

	// получение комментариев должно быть иерархическим
	post.Comments, err = s.getPostComments(post.ID)
	if err != nil {
		return nil, err
	}

	return post, nil
}

func (s *PostgresStorage) DeletePost(id int) (bool, error) {
	stmt := `DELETE FROM posts WHERE id=$1`

	err := s.DB.QueryRow(stmt, id).Err()

	if err != nil {
		return false, err
	}
	return true, nil
}

func (s *PostgresStorage) GetAllPosts() ([]*model.Post, error) {
	stmt := `SELECT id, user_id, text, "allowComment", created_at FROM posts`

	rows, err := s.DB.Query(stmt)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var posts []*model.Post

	for rows.Next() {
		post := &model.Post{}
		var userId int
		err = rows.Scan(&post.ID, &userId, &post.Text, &post.AllowComment, &post.Date)

		if err != nil {
			return nil, err
		}

		post.User, err = s.GetUser(userId)
		if err != nil {
			return nil, err
		}

		// тоже комментарии надо
		post.Comments, err = s.getPostComments(post.ID)
		if err != nil {
			return nil, err
		}

		posts = append(posts, post)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return posts, nil
}

func (s *PostgresStorage) getPostComments(postId int) ([]*model.Comment, error) {
	var comments []*model.Comment
	commentMap := make(map[int]*model.Comment)

	stmt := `
        WITH RECURSIVE comment_tree AS (
            SELECT id, user_id, parent_comment, text, created_at
            FROM comments
            WHERE parent_comment IS NULL and post_id = $1
            UNION ALL
            SELECT c.id, c.user_id, c.parent_comment, c.text, c.created_at
            FROM comments c
            JOIN comment_tree ct ON c.parent_comment = ct.id
        )
        SELECT id, user_id, parent_comment, text, created_at FROM comment_tree;
    `
	rows, err := s.DB.Query(stmt, postId)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		comment := &model.Comment{}
		var commentUserId int
		var parrentCommentID sql.NullInt64
		if err = rows.Scan(&comment.ID, &commentUserId, &parrentCommentID, &comment.Text, &comment.Date); err != nil {
			return nil, err
		}

		comment.User, err = s.GetUser(commentUserId)
		if err != nil {
			return nil, err
		}

		if !parrentCommentID.Valid {
			comments = append(comments, comment)
		}
		commentMap[comment.ID] = comment

		if parrentCommentID.Valid {
			if parent, ok := commentMap[int(parrentCommentID.Int64)]; ok {
				parent.Comments = append(parent.Comments, comment)
			}
		}

	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return comments, err
}
