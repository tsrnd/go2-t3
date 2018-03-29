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
}

func (bh BlogHandler) Create(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("views/blogs/create.html"))
	tmpl.ExecuteTemplate(w, "create", nil)
}

func (bh BlogHandler) Store(w http.ResponseWriter, r *http.Request) {
	title := r.FormValue("title")
	description := r.FormValue("description")
	detail := r.FormValue("detail")
	msg := &reponsitory.Message{
		Title:       title,
		Description: description,
		Detail:      detail,
	}
	if msg.Validate() == false {
		tmpl := template.Must(template.ParseFiles("views/blogs/create.html"))
		tmpl.ExecuteTemplate(w, "create", msg)
		return
	}
	reponsitory.CreateBlog(title, description, detail)
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}
