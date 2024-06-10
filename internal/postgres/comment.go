package postgres

import (
	"GraphQL_api/graph/model"
	"database/sql"
	_ "encoding/json"
	"errors"
	"time"
)

func (s *PostgresStorage) CreateComment(userId int, postId int, text string, date time.Time) (int, error) {
	stmt := `INSERT INTO comments (user_id, post_id, text) VALUES($1,$2,$3) RETURNING id`

	lastInsertId := 0
	err := s.DB.QueryRow(stmt, userId, postId, text).Scan(&lastInsertId)
	id := lastInsertId
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (s *PostgresStorage) CreateCommentToComment(userId int, commentId int, text string, date time.Time) (int, error) {
	stmt := `SELECT post_id FROM comments WHERE id = $1`
	row := s.DB.QueryRow(stmt, commentId)

	var postId int
	err := row.Scan(&postId)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return 0, ErrNoRecord
		} else {
			return 0, err
		}
	}

	stmt = `INSERT INTO comments (user_id, post_id, parent_comment, text) VALUES($1,$2,$3,$4) RETURNING id`

	lastInsertId := 0
	err = s.DB.QueryRow(stmt, userId, postId, commentId, text).Scan(&lastInsertId)

	id := lastInsertId
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (s *PostgresStorage) GetComment(id int) (*model.Comment, error) {
	stmt := `SELECT id, user_id, text, created_at FROM comments WHERE id = $1`

	row := s.DB.QueryRow(stmt, id)

	comment := &model.Comment{}

	var userId int
	err := row.Scan(&comment.ID, &userId, &comment.Text, &comment.Date)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNoRecord
		} else {
			return nil, err
		}
	}

	comment.User, err = s.GetUser(userId)
	if err != nil {
		return nil, err
	}

	comment.Comments, err = s.getCommentComments(comment.ID)
	if err != nil {
		return nil, err
	}

	return comment, nil
}

func (s *PostgresStorage) DeleteComment(id int) (bool, error) {
	stmt := `DELETE FROM comments WHERE id=$1`

	err := s.DB.QueryRow(stmt, id).Err()

	if err != nil {
		return false, err
	}
	return true, nil
}

func (s *PostgresStorage) GetAllComments() ([]*model.Comment, error) {
	stmt := `SELECT id, user_id, text, created_at FROM comments`

	rows, err := s.DB.Query(stmt)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var comments []*model.Comment

	for rows.Next() {
		comment := &model.Comment{}
		var userId int
		err = rows.Scan(&comment.ID, &userId, &comment.Text, &comment.Date)

		if err != nil {
			return nil, err
		}

		comment.User, err = s.GetUser(userId)
		if err != nil {
			return nil, err
		}

		comment.Comments, err = s.getCommentComments(comment.ID)
		if err != nil {
			return nil, err
		}

		comments = append(comments, comment)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return comments, nil
}

func (s *PostgresStorage) getCommentComments(commentId int) ([]*model.Comment, error) {
	var comments []*model.Comment
	commentMap := make(map[int]*model.Comment)

	stmt := `
        WITH RECURSIVE comment_tree AS (
            SELECT id, user_id, parent_comment, text, created_at
            FROM comments
            WHERE id = $1
            UNION ALL
            SELECT c.id, c.user_id, c.parent_comment, c.text, c.created_at
            FROM comments c
            JOIN comment_tree ct ON c.parent_comment = ct.id
        )
        SELECT id, user_id, parent_comment, text, created_at FROM comment_tree;
    `
	rows, err := s.DB.Query(stmt, commentId)

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

		commentMap[comment.ID] = comment

		if parrentCommentID.Valid {
			if int(parrentCommentID.Int64) == commentId {
				comments = append(comments, comment)
			}
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
