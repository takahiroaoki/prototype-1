package handler

import (
	"html/template"
	"net/http"
)

type MainHandler struct {
	tmpl *template.Template
}

func (h *MainHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	data := MainDto{
		Greeting: "Hello World2!",
		Sample: SampleDto{
			data: "test",
		},
	}
	h.tmpl.Execute(w, data)
}

func NewMainHandler() http.Handler {
	return &MainHandler{
		tmpl: template.Must(template.ParseFiles("resources/templates/main/index.html")),
	}
}

type MainDto struct {
	Greeting string
	Sample   SampleDto
}

type SampleDto struct {
	data string
}
