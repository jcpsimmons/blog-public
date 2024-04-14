package main

import (
	"embed"
	"fmt"
	"strings"

	"net/http"

	"github.com/jcpsimmons/glog/essay"
	"github.com/jcpsimmons/glog/template"
)

//go:embed static
var staticFiles embed.FS

func essayHandler(w http.ResponseWriter, r *http.Request) {
	fileName := strings.TrimPrefix(r.URL.Path, "/essay/") + ".md"
	post, err := essay.GetEssay(fileName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	title := essay.FileNameToTitle(string(fileName))
	pageName := title + " - The Blog Dr. Josh C. Simmons"
	formattedPost, err := template.SiteFrameWrap(title, string(post), pageName, "Essay")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(formattedPost)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("homeHandler", r.URL.Path)

	posts, err := essay.GetAllEssaysAsListItems()

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
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

func main() {
	http.Handle("/static/", http.FileServer(http.FS(staticFiles)))
	http.HandleFunc("/essay/", essayHandler)
	http.HandleFunc("/", homeHandler)

	err := http.ListenAndServe(":4000", nil)
	if err != nil {
		fmt.Println("ListenAndServe:", err)
	}
}
