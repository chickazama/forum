package api

import (
	"encoding/json"
	"matthewhope/forum/repo"
	"net/http"
)

type PostsHandler struct {
	Repo repo.IRepository
}

func NewPostsHandler(r repo.IRepository) *PostsHandler {
	return &PostsHandler{Repo: r}
}

func (h *PostsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.get(w, r)
		return
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}

func (h *PostsHandler) get(w http.ResponseWriter, r *http.Request) {
	ret, err := h.Repo.GetAllPosts()
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(w).Encode(&ret)
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
}
