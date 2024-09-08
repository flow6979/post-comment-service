package application

import (
	"context"
	"time"

	"post-comment-service/internal/domain"
	"post-comment-service/internal/ports"
	"post-comment-service/pkg/errors"

	"golang.org/x/crypto/bcrypt"
)

type PostService struct {
	postRepo    ports.PostRepository
	commentRepo ports.CommentRepository
}

func NewPostService(postRepo ports.PostRepository, commentRepo ports.CommentRepository) *PostService {
	return &PostService{
		postRepo:    postRepo,
		commentRepo: commentRepo,
	}
}

func (s *PostService) CreatePost(ctx context.Context, post *domain.Post) error {
	post.CreatedAt = time.Now()
	return s.postRepo.CreatePost(ctx, post)
}

func (s *PostService) GetPost(ctx context.Context, id int64) (*domain.Post, error) {
	return s.postRepo.GetPost(ctx, id)
}

func (s *PostService) ListPosts(ctx context.Context) ([]*domain.Post, error) {
	return s.postRepo.ListPosts(ctx)
}

func (s *PostService) CreateComment(ctx context.Context, comment *domain.Comment) error {
	comment.CreatedAt = time.Now()
	return s.commentRepo.CreateComment(ctx, comment)
}

func (s *PostService) GetComments(ctx context.Context, postID int64) ([]*domain.Comment, error) {
	return s.commentRepo.GetComments(ctx, postID)
}

type UserService struct {
	userRepo ports.UserRepository
}

func NewUserService(userRepo ports.UserRepository) *UserService {
	return &UserService{
		userRepo: userRepo,
	}
}

func (s *UserService) CreateUser(ctx context.Context, user *domain.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return errors.Wrap(err, "failed to hash password")
	}
	user.Password = string(hashedPassword)
	return s.userRepo.CreateUser(ctx, user)
}

func (s *UserService) AuthenticateUser(ctx context.Context, username, password string) (*domain.User, error) {
	user, err := s.userRepo.GetUser(ctx, username)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get user")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, errors.Wrap(err, "invalid password")
	}

	return user, nil
}
