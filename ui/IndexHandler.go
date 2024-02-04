package ui

import (
	"log"
	"matthewhope/forum/repo"
	"matthewhope/forum/transport"
	"net/http"
)

type IndexViewModel struct {
	Title string
	Posts []transport.PostDTO
}

type IndexHandler struct {
	templateName string
	Repo         repo.IRepository
}

func NewIndexHandler(r repo.IRepository) *IndexHandler {
	return &IndexHandler{templateName: "Index", Repo: r}
}

func (h *IndexHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.get(w, r)
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}

func (h *IndexHandler) get(w http.ResponseWriter, r *http.Request) {
	posts, err := h.Repo.GetAllPosts()
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
	data := make([]transport.PostDTO, len(posts))
	for i, p := range posts {
		data[i] = transport.PostDTO(p)
	}
	vm := IndexViewModel{Title: "Home Page", Posts: data}
	err = tmpl.ExecuteTemplate(w, h.templateName, vm)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
}
