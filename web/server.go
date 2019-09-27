package main

import (
	"github.com/zyp461476492/docker-app/web/router"
	"log"
	"net/http"
	"time"
)

func main() {
	router.ConfigRouter()
	s := &http.Server{
		Addr:           "0.0.0.0:8080",
		Handler:        nil,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Println("服务已启动")
	log.Fatal(s.ListenAndServe())
}
