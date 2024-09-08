package postgres

import (
	"context"
	"database/sql"

	"post-comment-service/internal/domain"
	"post-comment-service/pkg/errors"
)

type CommentRepository struct {
	db *sql.DB
}

func NewCommentRepository(db *sql.DB) *CommentRepository {
	return &CommentRepository{db: db}
}

func (r *CommentRepository) CreateComment(ctx context.Context, comment *domain.Comment) error {
	query := `INSERT INTO comments (post_id, user_id, content, created_at) VALUES ($1, $2, $3, $4) RETURNING id`
	err := r.db.QueryRowContext(ctx, query, comment.PostID, comment.UserID, comment.Content, comment.CreatedAt).Scan(&comment.ID)
	if err != nil {
		return errors.Wrap(err, "failed to create comment")
	}
	return nil
}

func (r *CommentRepository) GetComments(ctx context.Context, postID int64) ([]*domain.Comment, error) {
	query := `SELECT id, post_id, user_id, content, created_at FROM comments WHERE post_id = $1 ORDER BY created_at DESC`
	rows, err := r.db.QueryContext(ctx, query, postID)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get comments")
	}
	defer rows.Close()

	var comments []*domain.Comment
	for rows.Next() {
		comment := &domain.Comment{}
		err := rows.Scan(&comment.ID, &comment.PostID, &comment.UserID, &comment.Content, &comment.CreatedAt)
		if err != nil {
			return nil, errors.Wrap(err, "failed to scan comment")
		}
		comments = append(comments, comment)
	}
	return comments, nil
}
