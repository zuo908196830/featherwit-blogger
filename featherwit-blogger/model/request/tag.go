package request

type AddTagRequest struct {
	Tags []*TagTree `json:"tags"`
}

type TagTree struct {
	Name     string     `json:"name"`
	Children []*TagTree `json:"children"`
}

type TagRequest struct {
	ID1    string `json:"id1"`
	ID2    string `json:"id2"`
	ID3    string `json:"id3"`
	Height int    `json:"height"`
}

type AddTagBlogRequest struct {
	BlogId int64         `json:"blogId" binding:"required"`
	Tags   []*TagRequest `json:"tags"`
}
