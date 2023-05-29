package request

type AddComment struct {
	BlogId   int64  `json:"blogId"`
	ParentId int64  `json:"parentId"` //如果是二级评论的话，所属一级评论id，一级评论时为负数
	ReplyId  int64  `json:"replyId"`
	Content  string `json:"Content"`
}

type GetBlogComment struct {
	BlogId int64 `json:"blogId" uri:"blogId"`
	Limit  int   `json:"limit" uri:"limit"`
	Offset int   `json:"offset" uri:"offset"`
}
