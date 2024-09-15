package main

import (
	"database/sql"
	"net/http"

	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"

	"post-comment-service/internal/adapters/repositories/repos"
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

	logger.Info.Printf("Successfully connected to the database")

	// Run migrations
	if err := runMigrations(db); err != nil {
		logger.Error.Fatalf("Failed to run migrations: %v", err)
	}

	postRepo := repos.NewPostRepository(db)
	commentRepo := repos.NewCommentRepository(db)
	userRepo := repos.NewUserRepository(db)

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
