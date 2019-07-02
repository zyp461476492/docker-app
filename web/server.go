package main

import (
	"github.com/zyp461476492/sync-app/web/router"
	"log"
	"net/http"
	"time"
)

func main() {
	router.ConfigRouter()
	s := &http.Server{
		Addr:           "127.0.0.1:8080",
		Handler:        nil,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	for {
		log.Fatal(s.ListenAndServe())
	}

}
