package controller

import (
	"html/template"
	"net/http"
)

type (
	BlogController struct{}
)

func (bc BlogController) Index(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("views/blogs/index.html"))
	tmpl.ExecuteTemplate(w, "index", nil)
}
