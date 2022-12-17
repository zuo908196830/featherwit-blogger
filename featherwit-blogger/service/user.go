package service

import "fmt"

type UserService struct{}

func (u *UserService) Login() {
	fmt.Println("in service login")
}
