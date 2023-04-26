package request

type StarRequest struct {
	BlogId []int64 `json:"blogId" form:"blogId" uri:"blogId"`
}
