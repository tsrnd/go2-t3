package handler

import (
	"html/template"
	"net/http"
)

type (
	BlogHandler struct{}
)

func (bc BlogHandler) Index(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("views/blogs/index.html"))
	tmpl.ExecuteTemplate(w, "index", nil)
}
