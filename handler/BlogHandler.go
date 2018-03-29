package handler

import (
	"html/template"
	"net/http"

	"github.com/tsrnd/go2-t3/repository"
)

type BlogHandler struct {
	// br BlogResponse
}

func (bh BlogHandler) Index(w http.ResponseWriter, r *http.Request) {
	var br repository.BlogResponse
	blogs := br.GetAll()
	data := map[string]interface{}{
		"Blogs": blogs,
	}

	tmpl, err := template.ParseFiles("views/blogs/index.html")
	if err != nil {
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}
	tmpl.Execute(w, data)
}
