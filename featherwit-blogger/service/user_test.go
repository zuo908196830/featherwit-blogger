package service

import (
	"featherwit-blogger/global"
	"log"
	"testing"
)

func TestSearchAttentionUser(t *testing.T) {
	global.InitGlobal()
	s, err := UserServiceApp.SearchAttentionUser("xxx", 10, 0, nil)
	if err != nil {
		log.Fatalf("error:%v", err)
	}
	for _, username := range *s {
		log.Println(username)
	}
}
