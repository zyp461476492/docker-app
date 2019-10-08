package main

import (
	"flag"
	"github.com/zyp461476492/docker-app/web/router"
	"log"
	"net/http"
	"time"
)

var port string

func main() {
	flag.StringVar(&port, "port", "8080", "port")
	flag.Parse()
	log.Printf("listen port [%s]", port)

	router.ConfigRouter()
	addr := "0.0.0.0:" + port
	s := &http.Server{
		Addr:           addr,
		Handler:        nil,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Println("server is up")
	log.Panicln(s.ListenAndServe())

}
