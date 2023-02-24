package service

import (
	"featherwit-blogger/global"
	"featherwit-blogger/model"
	"featherwit-blogger/model/request"
	"log"
	"testing"
)

func TestMA(t *testing.T) {
	global.InitGlobal()
	param := request.AddTagRequest{
		Tags: []*request.TagTree{
			&request.TagTree{
				Name: "1",
				Children: []*request.TagTree{
					&request.TagTree{
						Name: "1.1",
					},
				},
			},
			&request.TagTree{
				Name: "2",
			},
		},
	}
	tags := make([]*model.Tag, 0)
	TagServiceApp.MutilAddTag(param.Tags, "", "", 1, &tags)
	log.Println(tags)
	TagServiceApp.AddTag(&tags, nil)
}
