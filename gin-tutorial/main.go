package main

import (
	"context"
	"fmt"
	"gin-tutorial/pkg/settings"
	"gin-tutorial/routers"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	//router := gin.Default()
	//router.GET("/", func(c *gin.Context) {
	//	c.JSON(200, gin.H{
	//		"message": "pong",
	//	})
	//})

	//router := routers.InitRouter()
	//
	//s := http.Server{
	//	Addr:           fmt.Sprintf(":%d", settings.HTTPPort),
	//	Handler:        router,
	//	ReadTimeout:    settings.ReadTimeout,
	//	WriteTimeout:   settings.WriteTimeout,
	//	MaxHeaderBytes: 1 << 20,
	//}
	//err := s.ListenAndServe()
	//if err != nil {
	//	return
	//}

	// Zero downtime restarts for golang HTTP servers.
	//endless.DefaultReadTimeOut = settings.ReadTimeout
	//endless.DefaultWriteTimeOut = settings.WriteTimeout
	//endless.DefaultMaxHeaderBytes = 1 << 20
	//endPoint := fmt.Sprintf(":%d", settings.HTTPPort)
	//
	//server := endless.NewServer(endPoint, routers.InitRouter())
	//server.BeforeBegin = func(add string) {
	//	log.Printf("Actual pid is %d", syscall.Getpid())
	//}
	//
	//err := server.ListenAndServe()
	//if err != nil {
	//	log.Printf("Server err: %v", err)
	//}

	// http.Shutdown
	router := routers.InitRouter()
	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", settings.HTTPPort),
		Handler:        router,
		ReadTimeout:    settings.ReadTimeout,
		WriteTimeout:   settings.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		if err := s.ListenAndServe(); err != nil {
			log.Printf("Listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit

	log.Println("Shutdown server ....")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := s.Shutdown(ctx); err != nil {
		log.Fatal("Server shutdown: ", err)
	}
	log.Println("Server exiting....")
}
