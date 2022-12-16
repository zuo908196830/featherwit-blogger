package main

import (
	_ "github.com/go-sql-driver/mysql"
	"os"
	"os/signal"
	"syscall"
	"the-fool-blogger/global"
	"the-fool-blogger/router"
)

func main() {
	global.NewDbEngine()
	//server := server.NewHttpServer(router.Init())
	//server.Start()
	router.Init().Run(":8080")

	quit := make(chan os.Signal)
	signal.Notify(quit,syscall.SIGINT, syscall.SIGTERM)
	<- quit
}
