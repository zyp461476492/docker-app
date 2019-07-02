package web

import (
	"log"
	"net/http"
	"time"
)

func testCon(w http.ResponseWriter, r *http.Request) {

}

func test() {
	http.Handle("/test", nil)

	s := &http.Server{
		Addr:           ":8080",
		Handler:        nil,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Fatal(s.ListenAndServe())
}
