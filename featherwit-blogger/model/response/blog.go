package response

import (
	"featherwit-blogger/model"
	"time"
)

type AddBlogResponse struct {
	ID       int64     `json:"id"`
	CreateAt time.Time `json:"createAt"`
	UpdateAt time.Time `json:"updateAt"`
}

type BlogCountResponse struct {
	Total int64 `json:"total"`
}

type SearchBlogs struct {
	Total int64         `json:"total"`
	Blogs []*model.Blog `json:"blogs"`
}
