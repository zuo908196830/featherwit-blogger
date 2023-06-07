package request

type BlogRequest struct {
	Title   string `json:"title" form:"title" binding:"required"`
	Content string `json:"content" form:"content" binding:"required"`
	Profile string `json:"profile" form:"profile"`
	Cover   string `json:"cover" form:"cover"`
}

type Page struct {
	Limit  int `json:"limit" uri:"limit" form:"limit"`
	Offset int `json:"offset" uri:"offset" form:"offset"`
}

type BlogIdRequest struct {
	ID int64 `json:"id" uri:"id" form:"title"`
}

type UpdateBlogRequest struct {
	ID      int64  `json:"id" form:"id" bind:"required"`
	Title   string `json:"title" form:"title"`
	Content string `json:"content" form:"content"`
	Profile string `json:"profile" form:"profile"`
}

type SearchBlogRequest struct {
	Page
	TagId int64  `json:"tagId" form:"tagId"`
	Name  string `json:"name" form:"name"`
	Order int    `json:"order" form:"order" bind:"required"`
}
