package main

import (
	"log"
	"net/http"
)

func main() {
	http.Handle("/tmpfiles", http.StripPrefix("/tmpfiles/", http.FileServer(http.Dir("/tmp/holmes"))))
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		log.Fatal(err)
	}
}
