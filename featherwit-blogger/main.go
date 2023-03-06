package main

import (
	"featherwit-blogger/global"
	"featherwit-blogger/router"
	"featherwit-blogger/server"
	"os"
	"os/signal"
	"syscall"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	global.InitGlobal()
	server := server.NewHttpServer(router.Init())
	server.Start()
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
}
