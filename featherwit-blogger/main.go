package main

import (
	"featherwit-blogger/global"
	"featherwit-blogger/router"
	"featherwit-blogger/server"
	_ "github.com/go-sql-driver/mysql"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	global.NewDbEngine()
	global.InitDbEngine()
	server := server.NewHttpServer(router.Init())
	server.Start()
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
}
