package request

type ConcernRequest struct {
	ConcernUsers string `json:"concernUser" form:"concernUser" uri:"concernUser"`
}
