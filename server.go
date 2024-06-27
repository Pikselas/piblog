package main

import (
	"encoding/json"
	"html/template"
	"math/rand"
	"net/http"
	"os"
	"strconv"
)

func main() {

	file_server := http.FileServer(http.Dir("./statics"))
	init_connection()

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
	http.HandleFunc("/create", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./statics/create.html")
	})
	http.HandleFunc("/get_blog_data/{id}", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		id, err := strconv.Atoi(r.PathValue("id"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		json.NewEncoder(w).Encode(FetchBlog(strconv.Itoa(id)))
	})
	http.HandleFunc("/create_blog", func(w http.ResponseWriter, r *http.Request) {

		type blog_data struct {
			Tags []string
			Desc string
			Data BlogPost
		}

		var newBlog blog_data

		err := json.NewDecoder(r.Body).Decode(&newBlog)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// convert base64 image to local files
		for itm := range newBlog.Data.Contents {
			if newBlog.Data.Contents[itm].Tag == "img" {
				id := rand.Int()
				img_file, _ := os.Create("./statics/" + strconv.Itoa(id))
				defer img_file.Close()
				img_file.WriteString(newBlog.Data.Contents[itm].Content)
				newBlog.Data.Contents[itm].Content = "/" + strconv.Itoa(id)
			}
		}

		newBlog.Data.Id = strconv.Itoa(rand.Int())
		InsertBlog(newBlog.Data, BlogDescription{Id: newBlog.Data.Id, Title: newBlog.Data.Title, Description: newBlog.Desc, Tags: newBlog.Tags})
	})
	http.HandleFunc("/search_tags", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(GetTags(""))
	})
	http.HandleFunc("/search_tags/{tag}", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(GetTags(r.PathValue("tag")))
	})
	http.HandleFunc("/search_by_tags", func(w http.ResponseWriter, r *http.Request) {
		var tags []string
		err := json.NewDecoder(r.Body).Decode(&tags)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(SearchBlogByTags(tags))
	})

	http.ListenAndServe(":8080", nil)
}
