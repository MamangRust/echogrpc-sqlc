// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: comment.sql

package db

import (
	"context"
)

const createComment = `-- name: CreateComment :one
INSERT INTO comments (id_post_comment,user_name_comment,comment) VALUES ($1,$2,$3) RETURNING id, id_post_comment, user_name_comment, comment, created_at, updated_at
`

type CreateCommentParams struct {
	IDPostComment   int32  `json:"id_post_comment"`
	UserNameComment string `json:"user_name_comment"`
	Comment         string `json:"comment"`
}

func (q *Queries) CreateComment(ctx context.Context, arg CreateCommentParams) (*Comment, error) {
	row := q.db.QueryRowContext(ctx, createComment, arg.IDPostComment, arg.UserNameComment, arg.Comment)
	var i Comment
	err := row.Scan(
		&i.ID,
		&i.IDPostComment,
		&i.UserNameComment,
		&i.Comment,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return &i, err
}

const deleteComment = `-- name: DeleteComment :exec
DELETE FROM comments WHERE id = $1
`

func (q *Queries) DeleteComment(ctx context.Context, id int32) error {
	_, err := q.db.ExecContext(ctx, deleteComment, id)
	return err
}

const getComment = `-- name: GetComment :one
SELECT id, id_post_comment,user_name_comment,comment, created_at, updated_at FROM comments WHERE id = $1 LIMIT 1
`

func (q *Queries) GetComment(ctx context.Context, id int32) (*Comment, error) {
	row := q.db.QueryRowContext(ctx, getComment, id)
	var i Comment
	err := row.Scan(
		&i.ID,
		&i.IDPostComment,
		&i.UserNameComment,
		&i.Comment,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return &i, err
}

const getComments = `-- name: GetComments :many
SELECT id, id_post_comment, user_name_comment, comment, created_at, updated_at FROM comments ORDER BY id DESC
`

func (q *Queries) GetComments(ctx context.Context) ([]*Comment, error) {
	rows, err := q.db.QueryContext(ctx, getComments)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []*Comment
	for rows.Next() {
		var i Comment
		if err := rows.Scan(
			&i.ID,
			&i.IDPostComment,
			&i.UserNameComment,
			&i.Comment,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, &i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateComment = `-- name: UpdateComment :one
UPDATE comments SET id_post_comment=$2,user_name_comment=$3,comment=$4 WHERE id=$1 RETURNING id, id_post_comment, user_name_comment, comment, created_at, updated_at
`

type UpdateCommentParams struct {
	ID              int32  `json:"id"`
	IDPostComment   int32  `json:"id_post_comment"`
	UserNameComment string `json:"user_name_comment"`
	Comment         string `json:"comment"`
}

func (q *Queries) UpdateComment(ctx context.Context, arg UpdateCommentParams) (*Comment, error) {
	row := q.db.QueryRowContext(ctx, updateComment,
		arg.ID,
		arg.IDPostComment,
		arg.UserNameComment,
		arg.Comment,
	)
	var i Comment
	err := row.Scan(
		&i.ID,
		&i.IDPostComment,
		&i.UserNameComment,
		&i.Comment,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return &i, err
}