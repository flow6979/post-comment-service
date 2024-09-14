package router

import (
	"post-comment-service/internal/adapters/repositories/http"
	"post-comment-service/internal/config"
	"post-comment-service/pkg/middleware"

	"github.com/gorilla/mux"
)

func NewRouter(cfg *config.Config) *mux.Router {
	r := mux.NewRouter()

	// Initialize handlers
	postHandler := http.NewPostHandler(cfg.PostService)
	userHandler := http.NewUserHandler(cfg.UserService)

	// Public routes
	r.HandleFunc("/register", userHandler.Register).Methods("POST")
	r.HandleFunc("/login", userHandler.Login).Methods("POST")

	// Public route for listing posts
	r.HandleFunc("/posts", postHandler.ListPosts).Methods("GET")

	// Protected routes
	protected := r.PathPrefix("").Subrouter()
	protected.Use(middleware.AuthMiddleware)

	protected.HandleFunc("/posts", postHandler.CreatePost).Methods("POST")
	protected.HandleFunc("/posts/{id}", postHandler.GetPost).Methods("GET")
	protected.HandleFunc("/posts/{postID}/comments", postHandler.CreateComment).Methods("POST")
	protected.HandleFunc("/posts/{postID}/comments", postHandler.GetComments).Methods("GET")

	return r
}
