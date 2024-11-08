package handler

import (
	"html/template"
	"net/http"
)

type MainHandler struct {
	tmpl *template.Template
}

func (h *MainHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	pageDto := MainPageDto{
		LinkList: []Link{
			{
				URL: "/",
				ID:  "choiceA",
			},
			{
				URL: "/",
				ID:  "choiceB",
			},
		},
	}
	h.tmpl.Execute(w, pageDto)
}

func NewMainHandler() http.Handler {
	return &MainHandler{
		tmpl: template.Must(template.ParseFiles("resources/templates/main/index.html")),
	}
}

type MainPageDto struct {
	LinkList []Link
}

type Link struct {
	URL string
	ID  string
}
