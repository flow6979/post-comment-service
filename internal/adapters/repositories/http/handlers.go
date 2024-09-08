package http

import (
	"encoding/json"
	"net/http"
	"strconv"

	"post-comment-service/internal/domain"
	"post-comment-service/internal/ports"
	"post-comment-service/pkg/errors"

	"github.com/gorilla/mux"
)

type PostHandler struct {
	postService ports.PostService
}

func NewPostHandler(postService ports.PostService) *PostHandler {
	return &PostHandler{postService: postService}
}

func (h *PostHandler) CreatePost(w http.ResponseWriter, r *http.Request) {
	var post domain.Post
	if err := json.NewDecoder(r.Body).Decode(&post); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.postService.CreatePost(r.Context(), &post); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(post)
}

func (h *PostHandler) GetPost(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		http.Error(w, "Invalid post ID", http.StatusBadRequest)
		return
	}

	post, err := h.postService.GetPost(r.Context(), id)
	if err != nil {
		if errors.Is(err, errors.ErrNotFound) {
			http.Error(w, "Post not found", http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	json.NewEncoder(w).Encode(post)
}

func (h *PostHandler) ListPosts(w http.ResponseWriter, r *http.Request) {
	posts, err := h.postService.ListPosts(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(posts)
}

func (h *PostHandler) CreateComment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	postID, err := strconv.ParseInt(vars["postID"], 10, 64)
	if err != nil {
		http.Error(w, "Invalid post ID", http.StatusBadRequest)
		return
	}

	var comment domain.Comment
	if err := json.NewDecoder(r.Body).Decode(&comment); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	comment.PostID = postID

	if err := h.postService.CreateComment(r.Context(), &comment); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(comment)
}

func (h *PostHandler) GetComments(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	postID, err := strconv.ParseInt(vars["postID"], 10, 64)
	if err != nil {
		http.Error(w, "Invalid post ID", http.StatusBadRequest)
		return
	}

	comments, err := h.postService.GetComments(r.Context(), postID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(comments)
}
