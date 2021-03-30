package web

import (
	"log"
	"net/http"
)

var fname string

func webHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, fname)
}

func ServeFile(filename string) error {
	fname = filename
	http.HandleFunc("/", webHandler)
	log.Fatal(http.ListenAndServe(":8082", nil))
	return nil
}
