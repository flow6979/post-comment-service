package repos

import (
	"context"
	"database/sql"

	"post-comment-service/internal/domain"
	"post-comment-service/pkg/errors"
)

type PostRepository struct {
	db *sql.DB
}

func NewPostRepository(db *sql.DB) *PostRepository {
	return &PostRepository{db: db}
}

func (r *PostRepository) CreatePost(ctx context.Context, post *domain.Post) error {
	query := `INSERT INTO posts (user_id, title, content, created_at) VALUES ($1, $2, $3, $4) RETURNING id`
	err := r.db.QueryRowContext(ctx, query, post.UserID, post.Title, post.Content, post.CreatedAt).Scan(&post.ID)
	if err != nil {
		return errors.Wrap(err, "failed to create post")
	}
	return nil
}

func (r *PostRepository) GetPost(ctx context.Context, id int64) (*domain.Post, error) {
	query := `SELECT id, user_id, title, content, created_at FROM posts WHERE id = $1`
	post := &domain.Post{}
	err := r.db.QueryRowContext(ctx, query, id).Scan(&post.ID, &post.UserID, &post.Title, &post.Content, &post.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.NotFound("post not found")
		}
		return nil, errors.Wrap(err, "failed to get post")
	}
	return post, nil
}

func (r *PostRepository) ListPosts(ctx context.Context) ([]*domain.Post, error) {
	query := `SELECT id, user_id, title, content, created_at FROM posts ORDER BY created_at DESC`
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, errors.Wrap(err, "failed to list posts")
	}
	defer rows.Close()

	var posts []*domain.Post
	for rows.Next() {
		post := &domain.Post{}
		err := rows.Scan(&post.ID, &post.UserID, &post.Title, &post.Content, &post.CreatedAt)
		if err != nil {
			return nil, errors.Wrap(err, "failed to scan post")
		}
		posts = append(posts, post)
	}
	return posts, nil
}
