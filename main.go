package main

import (
	"context"
	"crypto/tls"
	"embed"
	"html/template"
	"io"
	"os"
	"os/signal"
	"strings"
	"time"

	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"golang.org/x/crypto/acme"
	"golang.org/x/crypto/acme/autocert"
)

//go:embed static
var staticFiles embed.FS

//go:embed templates
var templateFiles embed.FS

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
	e.Logger.SetLevel(log.INFO)
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())

	t := &Template{
		templates: template.Must(template.ParseGlob("./templates/*.html")),
	}
	e.Renderer = t

	e.GET("/", Home)
	e.GET("/about", About)
	e.GET("/essay/:essay", Essay)

	e.Static("/static", "static")

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	go func() {
		// check for --dev flag
		if len(os.Args) > 1 && os.Args[1] == "--dev" {
			e.Logger.Fatal(e.Start(":4000"))

		} else {

			// redirect http to https
			e.Pre(middleware.HTTPSRedirect())
			// redirect www to non-www
			e.Pre(middleware.HTTPSNonWWWRedirect())
			autoTLSManager := autocert.Manager{
				Prompt: autocert.AcceptTOS,
				// Cache certificates to avoid issues with rate limits (https://letsencrypt.org/docs/rate-limits)
				Cache: autocert.DirCache("/var/www/.cache"),
				//HostPolicy: autocert.HostWhitelist("<DOMAIN>"),
			}
			s := http.Server{
				Addr:    ":443",
				Handler: e, // set Echo as handler
				TLSConfig: &tls.Config{
					//Certificates: nil, // <-- s.ListenAndServeTLS will populate this field
					GetCertificate: autoTLSManager.GetCertificate,
					NextProtos:     []string{acme.ALPNProto},
				},
				//ReadTimeout: 30 * time.Second, // use custom timeouts
			}
			if err := s.ListenAndServeTLS("", ""); err != http.ErrServerClosed {
				e.Logger.Fatal(err)
			}
		}
	}()

	<-ctx.Done()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}
