package reponsitory

import "github.com/tsrnd/go2-t3/model"

type BlogReponse struct {
}

func CreateBlog(title string, description string, detail string) {
	blog := model.Blog{Title: title, Description: description, Detail: detail}
	db := model.DBCon
	db.Create(&blog)
}
