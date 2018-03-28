package repository

import (
	"github.com/tsrnd/go2-t3/model"
)

type BlogResponse struct{}

func (br BlogResponse) GetAll() map[string]interface{} {
	blogs := []model.Blog{}
	db := model.DBCon
	defer db.Close()
	db.Find(&blogs)
	data := map[string]interface{}{
		"Blogs": blogs,
	}
	return data
}
