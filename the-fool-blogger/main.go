package main

import (
	_ "github.com/go-sql-driver/mysql"
	"os"
	"os/signal"
	"syscall"
	"the-fool-blogger/global"
	"the-fool-blogger/router"
	"the-fool-blogger/server"
)

func main() {
	global.NewDbEngine()
	server := server.NewHttpServer(router.Init())
	server.Start()
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
}
