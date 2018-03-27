package handler

import (
	"html/template"
	"net/http"

	"github.com/tsrnd/go2-t3/model"
)

type (
	BlogHandler struct{}
)

func (bh BlogHandler) Index(w http.ResponseWriter, r *http.Request) {
	blogs := []model.Blog{}
	db := model.DBCon
	db.Find(&blogs)
	data := map[string]interface{}{
		"Blogs": blogs,
	}
	tmpl := template.Must(template.ParseFiles("views/blogs/index.html"))
	tmpl.ExecuteTemplate(w, "index", data)
}
