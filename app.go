package main

import (
	"html/template"
	"net/http"
)

func init() {
	tpl = template.Must(template.ParseGlob("frontend.html"))
}

var (
	tpl *template.Template
)

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "frontend.html", nil)
}
