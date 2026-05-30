// Package HTML holds all the common HTML components and utilities.
package html

import (
	"crypto/sha256"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"

	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/components"
	. "maragu.dev/gomponents/html"
)

var hashOnce sync.Once
var appCSSPath, appJSPath, htmxJSPath string

// PageProps are properties for the [page] component.
type PageProps struct {
	Title       string
	Description string
	Header      bool
}

// page layout with header, footer, and container to restrict width and set base padding.
func page(props PageProps, children ...Node) Node {
	// Hash the paths for easy cache busting on changes
	hashOnce.Do(func() {
		appCSSPath = getHashedPath("public/styles/app.css")
		htmxJSPath = getHashedPath("public/scripts/htmx.js")
		appJSPath = getHashedPath("public/scripts/app.js")
	})

	maybeHeader := Group{}
	if props.Header {
		maybeHeader = append(maybeHeader, header())
	}

	return HTML5(HTML5Props{
		Title:       props.Title,
		Description: props.Description,
		Language:    "en",
		Head: []Node{
			Link(Rel("stylesheet"), Href(appCSSPath)),
			Script(Src(htmxJSPath), Defer()),
			Script(Src("https://code.jquery.com/jquery-3.7.1.slim.min.js"), Defer()),
			Script(Src(appJSPath), Defer()),
			Meta(Name("viewport"), Content("width=device-width, initial-scale=1")),
		},
		Body: Group{
			Div(Class("blog"),
				maybeHeader,
				Div(
					Group(children),
				),
			),
			actualFooter(),

			// Dark mode toggle (fixed in corner)
			Button(
				ID("theme-toggle"),
				Class("theme-toggle"),
				Attr("aria-label", "Toggle dark mode"),
				Attr("title", "Toggle theme"),
				Span(Class("icon")),
			),
		},
	})
}

func header() Node {
	return Header(
		Div(
			Class("site-hierarchy"),
			A(Href("https://quineglobal.com"), Class("hierarchy-parent"), Text("QUINE Global")),
			Div(Class("hierarchy-connector")),
			Span(Class("hierarchy-current"), Text("Blog")),
		),
		Img(Src("/images/quine_global_logo.png"), Height("72")),
		Div(
			Text("All posts by "),
			Strong(Text("QUINE Global")),
			Text("."),
		),
	)
}

func actualFooter() Node {
	return Div(
		Class("footer"),
		A(Href("/"), Text("Home")),
		A(Href("/about"), Text("About")),
		A(Href("/rss.xml"), Text("RSS")),
		A(Href("/credits"), Text("Credits")),
	)
}

func backToHome() Node {
	return P(
		Class("back-to-home"),
		A(Href("/"), Text("← Home")),
	)
}

func getHashedPath(path string) string {
	externalPath := strings.TrimPrefix(path, "public")
	ext := filepath.Ext(path)
	if ext == "" {
		panic("no extension found")
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return fmt.Sprintf("%v.x%v", strings.TrimSuffix(externalPath, ext), ext)
	}

	return fmt.Sprintf("%v.%x%v", strings.TrimSuffix(externalPath, ext), sha256.Sum256(data), ext)
}
