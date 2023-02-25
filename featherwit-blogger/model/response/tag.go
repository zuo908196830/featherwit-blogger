package response

import "featherwit-blogger/model"

type SearchTagResponse struct {
	Tag      *model.Tag           `json:"tag"`
	Children []*SearchTagResponse `json:"children"`
}
