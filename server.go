package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"main/ToOcto"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func createSlug(title string) string {
	title = strings.ToLower(title)
	title = strings.ReplaceAll(title, " ", "-")
	re := regexp.MustCompile(`[^\w-]`)
	title = re.ReplaceAllString(title, "")
	title = url.QueryEscape(title)
	return title
}

func create_blog_util(create_blog_func func(BlogPost, BlogDescription)) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
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
		ID := createSlug(newBlog.Data.Title)
		// convert base64 image to local files
		img_count := 0
		raw_src_path := "https://raw.githubusercontent.com/Pikselas/pikselasblogcontent/main/images/%s/%d"

		// user, octo_err := ToOcto.NewOctoUser(ENV["EMAIL"], ENV["GH_TOKEN"])
		user, octo_err := ToOcto.NewOctoUser(os.Getenv("EMAIL"), os.Getenv("GH_TOKEN"))
		if octo_err != nil {
			http.Error(w, octo_err.Error(), http.StatusInternalServerError)
			return
		}
		for itm := range newBlog.Data.Contents {
			if newBlog.Data.Contents[itm].Tag == "img" {
				path := "images/" + ID + "/" + strconv.Itoa(img_count)
				octo_err = user.Transfer("pikselasblogcontent", path, bytes.NewBufferString(newBlog.Data.Contents[itm].Content))
				if octo_err != nil {
					oct_err2 := user.Update("pikselasblogcontent", path, bytes.NewBufferString(newBlog.Data.Contents[itm].Content))
					if oct_err2 != nil {
						http.Error(w, fmt.Sprintf("Error saving image: %s <br/> Error updating image: %s", octo_err.Error(), oct_err2.Error()), http.StatusInternalServerError)
						return
					}
				}
				newBlog.Data.Contents[itm].Content = fmt.Sprintf(raw_src_path, ID, img_count)
				img_count++
			}
		}
		newBlog.Data.Id = ID
		newBlog.Data.lastUpdated = time.Now()
		create_blog_func(newBlog.Data, BlogDescription{Id: ID, Title: newBlog.Data.Title, Description: newBlog.Desc, Tags: newBlog.Tags})
	}
}

func convert_to_json(data interface{}) template.JS {
	json_data, err := json.Marshal(data)
	if err != nil {
		return "{}"
	}
	return template.JS(json_data)
}

func main() {

	// parseEnv()
	// init_connection(ENV["DB_URL"])
	init_connection(os.Getenv("DB_URL"))
	file_server := http.FileServer(http.Dir("./statics"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/" {
			http.ServeFile(w, r, "./statics/front.html")
		} else {
			file_server.ServeHTTP(w, r)
		}
	})

	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./statics/home.html")
	})

	http.HandleFunc("/blog/{id}", func(w http.ResponseWriter, r *http.Request) {
		blog_template := template.Must(template.ParseFiles("./templates/blog_template.html"))
		blog_template.Execute(w, struct {
			BlogD BlogPost
		}{FetchBlog(r.PathValue("id"))})
	})
	http.HandleFunc("/create", func(w http.ResponseWriter, r *http.Request) {
		blog_template := template.Must(template.New("create_template.html").Funcs(template.FuncMap{
			"toJSON": convert_to_json,
		}).ParseFiles("./templates/create_template.html"))
		err := blog_template.Execute(w, struct {
			SubmitUrl string
			Blog      BlogPost
			Desc      string
		}{
			SubmitUrl: "/create_blog",
			Blog:      BlogPost{Title: "Create New Blog"},
			Desc:      "Blog description is here!",
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})
	http.HandleFunc("/edit/{id}", func(w http.ResponseWriter, r *http.Request) {
		blog_template := template.Must(template.New("create_template.html").Funcs(template.FuncMap{
			"toJSON": convert_to_json,
		}).ParseFiles("./templates/create_template.html"))
		err := blog_template.Execute(w, struct {
			SubmitUrl string
			Blog      BlogPost
			Desc      string
		}{
			SubmitUrl: "/update_blog",
			Blog:      FetchBlog(r.PathValue("id")),
			Desc:      SearchBlogById(r.PathValue("id"))[0].Description,
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})
	http.HandleFunc("/get_blog_data/{id}", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(FetchBlog(r.PathValue("id")))
	})
	http.HandleFunc("/create_blog", create_blog_util(InsertBlog))
	http.HandleFunc("/update_blog", create_blog_util(UpdateBlog))
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
	http.HandleFunc("/search_by_title/{title}", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(SearchBlogByTitle(r.PathValue("title")))
	})

	http.ListenAndServe(":8080", nil)
}
