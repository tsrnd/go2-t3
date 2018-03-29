package responsitory

import (
	"github.com/tsrnd/go2-t3/model"
)

// BlogResponse struct
type BlogResponse struct{}

// Edit blog
func (br BlogResponse) Edit(id int) map[string]interface{} {
	blog := model.Blog{}
	db := model.DBCon
	db.Find(&blog, id)
	data := map[string]interface{}{
		"Blog": blog,
	}
	return data
}

// Update Blog
func (br BlogResponse) Update(bl model.Blog, id int) {
	blog := model.Blog{}
	db := model.DBCon
	db.Find(&blog, id)
	db.Model(&blog).Update(model.Blog{Title: bl.Title, Description: bl.Description, Detail: bl.Detail})
}
