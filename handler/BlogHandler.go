package handler

import (
	"html/template"
	"net/http"
	"strconv"

	"github.com/tsrnd/go2-t3/model"
	"github.com/tsrnd/go2-t3/repository"

	"github.com/gorilla/mux"
)

// BlogHandler struct
type BlogHandler struct{}

// Edit blog
func (bh BlogHandler) Edit(w http.ResponseWriter, r *http.Request) {
	var br repository.BlogResponse

	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	data := br.Edit(id)
	tmpl := template.Must(template.ParseFiles("views/blogs/edit.html"))
	tmpl.ExecuteTemplate(w, "edit", data)
}

// Update Blog
func (bh BlogHandler) Update(w http.ResponseWriter, r *http.Request) {
	var br repository.BlogResponse

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
	http.Redirect(w, r, "/blogs", http.StatusMovedPermanently)
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

func (bh BlogHandler) Create(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("views/blogs/create.html"))
	tmpl.ExecuteTemplate(w, "create", nil)
}

func (bh BlogHandler) Store(w http.ResponseWriter, r *http.Request) {
	title := r.FormValue("title")
	description := r.FormValue("description")
	detail := r.FormValue("detail")
	msg := &repository.Message{
		Title:       title,
		Description: description,
		Detail:      detail,
	}
	if msg.Validate() == false {
		tmpl := template.Must(template.ParseFiles("views/blogs/create.html"))
		tmpl.ExecuteTemplate(w, "create", msg)
		return
	}
	repository.CreateBlog(title, description, detail)
	http.Redirect(w, r, "/blogs", http.StatusMovedPermanently)
}
