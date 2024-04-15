package main

import (
	"embed"
	"log"
	"os"
	"strings"

	"net/http"

	"github.com/jcpsimmons/blog-public/template"
)

//go:embed static
var staticFiles embed.FS

func essayHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("essayHandler on path:", r.URL.Path)
	fileName := strings.TrimPrefix(r.URL.Path, "/essay/") + ".md"
	post, err := GetMarkdown("essay/" + fileName)

	if err != nil || string(post) == "" {
		log.Println("Error getting markdown:", err)
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}

	title := FileNameToTitle(string(fileName))
	pageName := title + " - The Blog Dr. Josh C. Simmons"
	formattedPost, err := template.SiteFrameWrap(title, string(post), pageName, "Essay")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(formattedPost)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("homeHandler on path:", r.URL.Path)
	posts, err := GetAllEssaysAsListItems()

	if err != nil || r.URL.Path != "/" {
		http.Error(w, "not found", http.StatusNotFound)
		return
	}
	allPosts := "<ul>" + strings.Join(posts, "") + "</ul>"

	formattedPost, err := template.SiteFrameWrap("Essays", string(allPosts), "The Blog of Dr. Josh C. Simmons", "Home")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(formattedPost)
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("aboutHandler on path:", r.URL.Path)

	fileName := "single-pages/about.md"
	post, err := GetMarkdown(fileName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	title := "About"
	pageName := "About - The Blog Dr. Josh C. Simmons"
	formattedPost, err := template.SiteFrameWrap(title, string(post), pageName, "Essay")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(formattedPost)
}

func main() {
	log.Println("Starting server...")

	log.Println("Getting port...")
	port := "4000"
	if os.Getenv("PORT") != "" {
		port = os.Getenv("PORT")
	}
	port = ":" + port

	log.Println("Listening on port", port)

	http.Handle("/static/", http.FileServer(http.FS(staticFiles)))
	http.HandleFunc("/essay/", essayHandler)
	http.HandleFunc("/about", aboutHandler)
	http.HandleFunc("/", homeHandler)

	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Println("ListenAndServe:", err)
	}
}
