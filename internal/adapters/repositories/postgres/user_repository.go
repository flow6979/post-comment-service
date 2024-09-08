package postgres

import (
	"context"
	"database/sql"

	"post-comment-service/internal/domain"
	"post-comment-service/pkg/errors"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) CreateUser(ctx context.Context, user *domain.User) error {
	query := `INSERT INTO users (username, password) VALUES ($1, $2) RETURNING id`
	err := r.db.QueryRowContext(ctx, query, user.Username, user.Password).Scan(&user.ID)
	if err != nil {
		return errors.Wrap(err, "failed to create user")
	}
	return nil
}

func (r *UserRepository) GetUser(ctx context.Context, username string) (*domain.User, error) {
	query := `SELECT id, username, password FROM users WHERE username = $1`
	user := &domain.User{}
	err := r.db.QueryRowContext(ctx, query, username).Scan(&user.ID, &user.Username, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.NotFound("user not found")
		}
		return nil, errors.Wrap(err, "failed to get user")
	}
	return user, nil
}
