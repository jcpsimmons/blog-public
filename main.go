package main

import (
	"embed"
	"html/template"
	"io"
	"os"
	"strings"

	"net/http"

	"github.com/labstack/echo"
)

//go:embed static
var staticFiles embed.FS

type Template struct {
	templates *template.Template
}

type PageData struct {
	PageTitle string
	PageID    string
	Heading   string
	Content   template.HTML
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	if err := t.templates.ExecuteTemplate(w, name, data); err != nil {
		c.Logger().Error(err)
		return err
	}
	return nil
}

// func essayHandler(c echo.Context) error {
// 	essayName := c.Param("essay")
// 	fileName := essayName + ".md"
// 	post, err := GetMarkdown("essay/" + fileName)

// 	if err != nil {
// 		return err
// 	}

// 	title := FileNameToTitle(string(fileName))
// 	pageName := title + " - The Blog Dr. Josh C. Simmons"
// 	essayPageData := map[string]interface{}{
// 		"PageName":    pageName,
// 		"BodyID":      "Essay",
// 		"PageTitle":   title,
// 		"PageContent": string(post),
// 	}

// 	if err != nil {
// 		return err
// 	}
// 	return c.Render(http.StatusOK, "site-frame.html", essayPageData)
// }

func Home(c echo.Context) error {
	posts, err := GetAllEssaysAsListItems()

	if err != nil {
		return err
	}
	allPosts := "<ul>" + strings.Join(posts, "") + "</ul>"

	data := PageData{
		PageTitle: "The Blog of Dr. Josh C. Simmons",
		PageID:    "Home",
		Heading:   "Essays",
		Content:   template.HTML(allPosts),
	}

	return c.Render(http.StatusOK, "default", data)
}

// func aboutHandler(c echo.Context) error {
// 	fileName := "single-pages/about.md"
// 	post, err := GetMarkdown(fileName)
// 	if err != nil {
// 		return err
// 	}

// 	title := "About"
// 	pageName := "About - The Blog Dr. Josh C. Simmons"
// 	formattedPost, err := blogtemplate.SiteFrameWrap(title, string(post), pageName, "Essay")
// 	if err != nil {
// 		return err
// 	}
// 	return c.HTML(http.StatusOK, string(formattedPost))
// }

func Hello(c echo.Context) error {
	return c.Render(http.StatusOK, "default", map[string]interface{}{
		"name": "Josh",
		"test": string(c.Param("count")),
	})
}

func main() {
	e := echo.New()

	t := &Template{
		templates: template.Must(template.ParseGlob("./templates/*.html")),
	}
	e.Renderer = t

	e.GET("/", Home)
	// e.GET("/about", aboutHandler)
	// e.GET("/essay/:essay", essayHandler)
	e.Static("/static", "static")

	port := "4000"
	if os.Getenv("PORT") != "" {
		port = os.Getenv("PORT")
	}
	port = ":" + port
	e.Logger.Fatal(e.Start(port))
}
