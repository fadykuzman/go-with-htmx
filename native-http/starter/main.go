package main

import (
	"log"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	msg := "Hello"
	w.Write([]byte(msg))
}

func about(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "public/index.html")
}

func htmx(w http.ResponseWriter, r *http.Request) {
	msg := "Hi, HTMX"
	w.Write([]byte(msg))
}

func main() {
	http.HandleFunc("/htmx", htmx)
	http.HandleFunc("/about", about)
	http.HandleFunc("/", hello)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
