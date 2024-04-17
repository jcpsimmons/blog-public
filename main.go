package main

import (
	"context"
	"embed"
	"html/template"
	"io"
	"os"
	"os/signal"
	"strings"
	"time"

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

func Essay(c echo.Context) error {
	essayName := c.Param("essay")
	c.Logger().Info("Reading essay: ", essayName)
	fileName := essayName + ".md"
	post, err := GetMarkdown("essay/" + fileName)
	c.Logger().Info("Reading file: ", fileName)
	if err != nil {
		return err
	}

	title := FileNameToTitle(string(fileName))
	pageTitle := title + " - The Blog Dr. Josh C. Simmons"

	data := PageData{
		PageTitle: pageTitle,
		PageID:    "Essay",
		Heading:   title,
		Content:   template.HTML(post),
	}
	return c.Render(http.StatusOK, "default", data)
}

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

func About(c echo.Context) error {
	fileName := "single-pages/about.md"
	aboutText, err := GetMarkdown(fileName)
	if err != nil {
		return err
	}

	heading := "About"
	pageTitle := "About - The Blog Dr. Josh C. Simmons"

	data := PageData{
		PageTitle: pageTitle,
		PageID:    "Home",
		Heading:   heading,
		Content:   template.HTML(aboutText),
	}
	return c.Render(http.StatusOK, "default", data)

}

func main() {
	e := echo.New()

	t := &Template{
		templates: template.Must(template.ParseGlob("./templates/*.html")),
	}
	e.Renderer = t

	e.GET("/", Home)
	e.GET("/about", About)
	e.GET("/essay/:essay", Essay)

	e.Static("/static", "static")

	port := "4000"
	if os.Getenv("PORT") != "" {
		port = os.Getenv("PORT")
	}
	port = ":" + port

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()
	// Start server
	go func() {
		if err := e.Start(port); err != nil && err != http.ErrServerClosed {
			e.Logger.Fatal("shutting down the server")
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with a timeout of 10 seconds.
	<-ctx.Done()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}
