package main

import ("fmt"
		"net/http"
		"html/template")

// https://gowebexamples.com/templates/

type Example struct {
	Title string
	Content string
}

func notemplate_handler(w http.ResponseWriter, r *http.Request) { // localhost:8000/notemplate
    fmt.Fprintf(w, "<h1>Inline tags</h1>")
}

func template_handler(w http.ResponseWriter, r *http.Request) { // localhost:8000/template
    eg := Example{Title: "Some title", Content: "Some content"}
	tmpl, _ := template.ParseFiles("htmltemplate1.html")
	tmpl.Execute(w, eg)
}

func main() {
	http.HandleFunc("/notemplate", notemplate_handler)
	http.HandleFunc("/template", template_handler)
	http.ListenAndServe(":8000", nil)
}