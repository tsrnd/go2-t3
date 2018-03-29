package repository

import (
	"log"

	. "github.com/tsrnd/go2-t3/model"
)

type BlogResponse struct{}

func (br BlogResponse) GetAll() []Blog {
	blogs := []Blog{}
	err := DBCon.Find(&blogs)
	if err.Error != nil {
		log.Fatal(err)
	}
	return blogs
}

// Edit blog
func (br BlogResponse) Edit(id int) map[string]interface{} {
	blog := Blog{}
	db := DBCon
	db.Find(&blog, id)
	data := map[string]interface{}{
		"Blog": blog,
	}
	return data
}

// Update Blog
func (br BlogResponse) Update(bl Blog, id int) {
	blog := Blog{}
	db := DBCon
	db.Find(&blog, id)
	db.Model(&blog).Update(Blog{Title: bl.Title, Description: bl.Description, Detail: bl.Detail})
}
