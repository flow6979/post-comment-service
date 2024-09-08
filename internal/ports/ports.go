package ports

import (
	"context"

	"post-comment-service/internal/domain"
)

type PostRepository interface {
	CreatePost(ctx context.Context, post *domain.Post) error
	GetPost(ctx context.Context, id int64) (*domain.Post, error)
	ListPosts(ctx context.Context) ([]*domain.Post, error)
}

type CommentRepository interface {
	CreateComment(ctx context.Context, comment *domain.Comment) error
	GetComments(ctx context.Context, postID int64) ([]*domain.Comment, error)
}

type UserRepository interface {
	CreateUser(ctx context.Context, user *domain.User) error
	GetUser(ctx context.Context, username string) (*domain.User, error)
}

type PostService interface {
	CreatePost(ctx context.Context, post *domain.Post) error
	GetPost(ctx context.Context, id int64) (*domain.Post, error)
	ListPosts(ctx context.Context) ([]*domain.Post, error)
	CreateComment(ctx context.Context, comment *domain.Comment) error
	GetComments(ctx context.Context, postID int64) ([]*domain.Comment, error)
}

type UserService interface {
	CreateUser(ctx context.Context, user *domain.User) error
	AuthenticateUser(ctx context.Context, username, password string) (*domain.User, error)
}
