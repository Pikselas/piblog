package main

import (
	"encoding/json"
	"html/template"
	"net/http"
)

type BlogData struct {
	Tag     string
	Content string
}

type BlogPost struct {
	Title    string
	Contents []BlogData
}

func main() {

	blg := BlogPost{
		Title: "My Blog",
		Contents: []BlogData{
			{Tag: "h1", Content: "This is the first blog post"},
			{Tag: "p", Content: "This is the second blog post"},
			{Tag: "pre", Content: "#include <stdio.h> \n int main() { \n printf(\"Hello, World!\"); \n return 0; \n}"},
		},
	}

	file_server := http.FileServer(http.Dir("./statics"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/" {
			http.ServeFile(w, r, "./statics/home.html")
		} else {
			file_server.ServeHTTP(w, r)
		}
	})

	http.HandleFunc("/blog/{id}", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("./statics/temp.html"))
		tmpl.Execute(w, struct{ BlogID string }{r.PathValue("id")})
	})

	http.HandleFunc("/get_blog_data/{id}", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(blg)
	})

	http.ListenAndServe(":8080", nil)
}
