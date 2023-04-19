package request

type AddTagRequest struct {
	TagNames []string `json:"tagNames"`
}

type AddTagBlogRequest struct {
	BlogId int64   `json:"blogId" binding:"required"`
	TagIds []int64 `json:"tags" binding:"required"`
}
