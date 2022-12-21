package response

import "time"

type AddBlogResponse struct {
	ID       int       `json:"id"`
	CreateAt time.Time `json:"createAt"`
	UpdateAt time.Time `json:"updateAt"`
}
