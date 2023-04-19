package request

type BlogRequest struct {
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
	Profile string `json:"profile"`
}

type Page struct {
	Limit  int `json:"limit" uri:"limit"`
	Offset int `json:"offset" uri:"offset"`
}

type BlogIdRequest struct {
	ID int64 `json:"id" uri:"id"`
}

type UpdateBlogRequest struct {
	ID      int64  `json:"id" bind:"required"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

type SearchBlogRequest struct {
	TagId int64  `json:"tagId"`
	Name  string `json:"name"`
}
