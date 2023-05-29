package response

import "featherwit-blogger/model"

type CommentsTree struct {
	Comment         *model.Comment     `json:"comment"`
	ChildrenCount   int                `json:"childrenCount"`
	User            *UserShow          `json:"user"`
	ChildrenComment []*ChildrenComment `json:"childrenComment"`
}

type ChildrenComment struct {
	User          *UserShow      `json:"user"`
	Comment       *model.Comment `json:"comment"`
	ReplyNickname string         `json:"replyNickname"`
}

type Comments struct {
	Count        int             `json:"count"`
	CommentsTree []*CommentsTree `json:"comments"`
}
