package repository

import (
	"strings"

	"github.com/tsrnd/go2-t3/model"
)

type BlogReponse struct {
}

type Message struct {
	Title       string
	Description string
	Detail      string
	Errors      map[string]string
}

func CreateBlog(title string, description string, detail string) {
	blog := model.Blog{Title: title, Description: description, Detail: detail}
	db := model.DBCon
	db.Create(&blog)
}

func (msg *Message) Validate() bool {
	msg.Errors = make(map[string]string)

	if strings.TrimSpace(msg.Title) == "" {
		msg.Errors["Title"] = "Please input title"
	}
	if strings.TrimSpace(msg.Description) == "" {
		msg.Errors["Description"] = "Please input description"
	}
	if strings.TrimSpace(msg.Detail) == "" {
		msg.Errors["Detail"] = "Please input detail"
	}
	return len(msg.Errors) == 0
}
