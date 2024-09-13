package main

import (
	"database/sql"
	"net/http"

	"post-comment-service/internal/adapters/repositories/postgres"
	"post-comment-service/internal/adapters/repositories/router"
	"post-comment-service/internal/application"
	"post-comment-service/internal/config"
	"post-comment-service/pkg/logger"
)

func main() {
	logger.Info.Printf("Starting things up for you boi !!")

	cfg, err := config.Load()
	if err != nil {
		logger.Error.Fatalf("Failed to load configuration: %v", err)
	}

	db, err := sql.Open("postgres", cfg.DatabaseURL)
	if err != nil {
		logger.Error.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		logger.Error.Fatalf("Failed to ping database: %v", err)
	}

	postRepo := postgres.NewPostRepository(db)
	commentRepo := postgres.NewCommentRepository(db)
	userRepo := postgres.NewUserRepository(db)

	postService := application.NewPostService(postRepo, commentRepo)
	userService := application.NewUserService(userRepo)

	cfg.PostService = postService
	cfg.UserService = userService

	r := router.NewRouter(cfg)

	logger.Info.Printf("Server starting on %s", cfg.ServerAddress)
	if err := http.ListenAndServe(cfg.ServerAddress, r); err != nil {
		logger.Error.Fatalf("Server failed to start: %v", err)
	}
}
