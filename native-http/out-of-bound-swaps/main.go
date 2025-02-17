package main

import (
	"html/template"
	"log"
	"net/http"
)

func demo(w http.ResponseWriter, r *http.Request) {
	html := `
	<div>new 1</div>
	<div id="target2" hx-swap-oob="true">
	new 2
	</div>
	<div id="target2" hx-swap-oob="afterend">
	<div>after 2</div>
	</div>
	<div hx-swap-oob="innerHTML:#target3"> new 3 </div>
`
	tmpl, _ := template.New("demo").Parse(html)
	tmpl.Execute(w, tmpl)
}

func main() {
	http.HandleFunc("GET /demo/", demo)
	dir := http.Dir("public")
	fs := http.FileServer(dir)
	http.Handle("/", http.StripPrefix("", fs))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
