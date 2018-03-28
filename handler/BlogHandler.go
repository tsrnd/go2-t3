package handler

import (
	"html/template"
	"net/http"
	"strconv"

	"github.com/tsrnd/go2-t3/model"
	"github.com/tsrnd/go2-t3/responsitory"

	"github.com/gorilla/mux"
)

// BlogHandler struct
type BlogHandler struct{}

var br responsitory.BlogResponse

// Edit blog
func (bh BlogHandler) Edit(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	data := br.Edit(id)
	tmpl := template.Must(template.ParseFiles("views/blogs/edit.html"))
	tmpl.ExecuteTemplate(w, "edit", data)
}

// Update Blog
func (bh BlogHandler) Update(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	Title := r.FormValue("title")
	Description := r.FormValue("description")
	Detail := r.FormValue("detail")
	blog := model.Blog{
		Title:       Title,
		Description: Description,
		Detail:      Detail,
	}
	br.Update(blog, id)
}
