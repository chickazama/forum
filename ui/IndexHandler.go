package ui

import (
	"log"
	"net/http"
)

type IndexHandler struct {
	templateName string
}

func NewIndexHandler() *IndexHandler {
	return &IndexHandler{templateName: "Index"}
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
	err := tmpl.ExecuteTemplate(w, h.templateName, nil)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
}
