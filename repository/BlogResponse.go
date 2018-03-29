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
