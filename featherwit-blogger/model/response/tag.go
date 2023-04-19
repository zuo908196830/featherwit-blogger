package response

import "featherwit-blogger/model"

type SearchTagResponse struct {
	Tags []*model.Tag `json:"tag"`
}
