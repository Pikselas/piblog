package main

import (
	"encoding/json"
	"html/template"
	"math/rand"
	"net/http"
	"os"
	"strconv"
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

	var blg []BlogPost = make([]BlogPost, 0)

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
		json.NewEncoder(w).Encode(blg[id])
	})

	http.HandleFunc("/create_blog", func(w http.ResponseWriter, r *http.Request) {

		var newBlog BlogPost
		err := json.NewDecoder(r.Body).Decode(&newBlog)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// convert base64 image to local files
		for itm := range newBlog.Contents {
			if newBlog.Contents[itm].Tag == "img" {
				id := rand.Int()
				img_file, _ := os.Create("./statics/" + strconv.Itoa(id))
				defer img_file.Close()
				img_file.WriteString(newBlog.Contents[itm].Content)
				newBlog.Contents[itm].Content = "/" + strconv.Itoa(id)
			}
		}
		blg = append(blg, newBlog)
	})

	http.ListenAndServe(":8080", nil)
}
