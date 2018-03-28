package handler

import (
	"html/template"
	"net/http"

	"github.com/tsrnd/go2-t3/reponsitory"
)

type (
	BlogHandler struct{}
)

func (bh BlogHandler) Index(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("views/blogs/index.html"))
	tmpl.ExecuteTemplate(w, "index", nil)
	// blogs := model.Blog{}
	// db := model.DBCon
	// db.First(&blogs, 2)
	// fmt.Print(blogs)
}

func (bh BlogHandler) Create(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("views/blogs/create.html"))
	tmpl.ExecuteTemplate(w, "create", nil)
}

func (bh BlogHandler) Store(w http.ResponseWriter, r *http.Request) {
	title := r.FormValue("title")
	description := r.FormValue("description")
	detail := r.FormValue("detail")
	reponsitory.CreateBlog(title, description, detail)
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}
