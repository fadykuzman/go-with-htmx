package main

import (
	"log"
	"net/http"
)

func main() {
	srvr := http.NewServeMux()

	fs := http.FileServer(http.Dir("public"))
	srvr.Handle("/", http.StripPrefix("", fs))

	log.Fatal(http.ListenAndServe(":8080", srvr))
}
