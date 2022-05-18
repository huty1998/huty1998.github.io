package main

import (
	"context"
	"fmt"
	"myserver/middleware"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	//middleware
	r.Use(middleware.M1())

	r.GET("/test", func(ctx *gin.Context) {
		req, _ := ctx.Get("key1") //same ctx as which in middleware??
		fmt.Println("main")
		ctx.JSON(http.StatusOK, map[string]interface{}{"req": req})
	})

	r.NoRoute(func(ctx *gin.Context) { //404
		ctx.String(http.StatusNotFound, "404 NOT FOUND, NO ROUTE")
	})

	// r.Run(":728") //gin
	ctx, cancel := context.WithCancel(context.Background())

	server := &http.Server{Addr: ":728", Handler: r} //combine this router with http.Server
	server.ListenAndServe()                          // listen and server http request
	<-ctx.Done()                                     //release port
	if err := server.Shutdown(ctx); err != nil {
		fmt.Printf("Server forced to shutdown: " + err.Error())
	}

	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGKILL, os.Interrupt)
	if <-c != nil {
		cancel()
	}
	<-ctx.Done()
	// var wg sync.WaitGroup
	// wg.Add(1)
	// wg.Wait() https://blog.csdn.net/textdemo123/article/details/103416781
}
