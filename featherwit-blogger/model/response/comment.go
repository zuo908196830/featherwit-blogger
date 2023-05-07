package response

import "featherwit-blogger/model"

type CommentsTree struct {
	Comment         *model.Comment   `json:"comment"`
	ChildrenCount   int              `json:"childrenCount"`
	ChildrenComment []*model.Comment `json:"childrenComment"`
}
