package handler

import (
	"html/template"
	"net/http"
)

type MainHandler struct {
	tmpl *template.Template
}

func (h *MainHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.tmpl.Execute(w, nil)
}

func NewMainHandler() http.Handler {
	return &MainHandler{
		tmpl: template.Must(template.ParseFiles("resources/templates/main/index.html")),
	}
}
