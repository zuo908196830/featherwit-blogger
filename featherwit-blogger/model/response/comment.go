package response

import "featherwit-blogger/model"

type CommentsTree struct {
	Comment         *model.Comment   `json:"comment"`
	ChildrenComment []*model.Comment `json:"childrenComment"`
}
