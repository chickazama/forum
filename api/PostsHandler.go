package api

import (
	"encoding/json"
	"log"
	"matthewhope/forum/models"
	"matthewhope/forum/repo"
	"net/http"
	"time"
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
	case http.MethodPost:
		h.post(w, r)
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

func (h *PostsHandler) post(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
	p := new(models.Post)
	ctime := time.Now().UTC().UnixMilli()
	p.Category = r.PostFormValue("category")
	p.Content = r.PostFormValue("content")
	p.CreatedAt = ctime
	p.UpdatedAt = ctime
	p.UserID = 1
	err = p.Validate()
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}
	ret, err := h.Repo.CreatePost(*p)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(&ret)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
}
