package handler

import (
	"html/template"
	"net/http"

	"github.com/tsrnd/go2-t3/repository"
)

type (
	BlogHandler struct{}
)

func (bh BlogHandler) Index(w http.ResponseWriter, r *http.Request) {
	var br repository.BlogResponse
	data := br.GetAll()
	tmpl := template.Must(template.ParseFiles("views/blogs/index.html"))
	tmpl.ExecuteTemplate(w, "index", data)
}
