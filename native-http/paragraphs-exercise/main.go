package main

import (
	"html/template"
	"log"
	"net/http"
)

type Paragraphs struct {
	P1 string
	P2 string
}

func getParagraphs(w http.ResponseWriter, r *http.Request) {

	paragraphs := Paragraphs{
		P1: "paragraph 1",
		P2: "paragraph 2",
	}
	tmpl := template.Must(template.ParseFiles("public/paragraphs.html"))
	tmpl.Execute(w, paragraphs)
}

func changeParagraphs(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/html")
	html := `
	<p id="para1" hx-swap-oob="true"> Paragraph 3 oob </p>
	<p id="para2" hx-swap-oob="true"> Paragraph 4 oob </p>
	`
	w.Write([]byte(html))
}
func main() {
	srvr := http.NewServeMux()

	srvr.HandleFunc("GET /paragraphs", getParagraphs)
	srvr.HandleFunc("GET /change", changeParagraphs)

	fs := http.FileServer(http.Dir("public"))
	srvr.Handle("/", http.StripPrefix("", fs))

	log.Fatal(http.ListenAndServe(":8080", srvr))
}
