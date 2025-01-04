package http

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	netHtml "golang.org/x/net/html"

	"github.com/go-chi/chi/v5"
	feeds "github.com/gorilla/feeds"
	lo "github.com/samber/lo"
	. "maragu.dev/gomponents"
	hx "maragu.dev/gomponents-htmx/http"
	ghttp "maragu.dev/gomponents/http"

	"app/html"
	"app/model"
	"app/posts"
)

type thingsGetter interface {
	GetThings(ctx context.Context) ([]model.Thing, error)
}

// func rssRouter() chi.Router {
// 	r := chi.NewRouter()
// 	r.Use(render.SetContentType(render.ContentTypeJSON))
// 	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
// 		w.Write([]byte("admin: index"))
// 	})
// 	return r
// }

type HttpErrorResponse struct {
	ErrorMessage string
}

type NotFound HttpErrorResponse

func (a NotFound) StatusCode() int {
	return 404
}

func (a NotFound) Error() string {
	return a.ErrorMessage
}

type errorWithStatusCode interface {
	StatusCode() int
}

var _ errorWithStatusCode = NotFound{}

// Home handler for the home page, as well as HTMX partial for getting things.
func Home(r chi.Router, db thingsGetter) {
	r.Get("/", ghttp.Adapt(func(w http.ResponseWriter, r *http.Request) (Node, error) {
		allPosts := posts.AllPosts

		if hx.IsRequest(r.Header) {
			return html.Posts(allPosts, time.Now()), nil
		}

		return html.HomePage(html.PageProps{}, allPosts, time.Now()), nil
	}))

	r.Get("/credits", ghttp.Adapt(func(w http.ResponseWriter, r *http.Request) (Node, error) {
		return html.CreditsPage(html.PageProps{}), nil
	}))

	r.Get("/about", ghttp.Adapt(func(w http.ResponseWriter, r *http.Request) (Node, error) {
		return html.AboutPage(html.PageProps{}), nil
	}))

	r.Route("/post", func(r chi.Router) {
		r.Get("/", func(writer http.ResponseWriter, req *http.Request) {
			// redirect em from /post to /
			http.Redirect(writer, req, "/", http.StatusTemporaryRedirect)
		})

		r.With(PostFetcherErrorer).
			Get("/{postID}", ghttp.Adapt(func(rw http.ResponseWriter, r *http.Request) (Node, error) {
				ctx := r.Context()
				post := ctx.Value(postCtxKey{}).(model.QuinePost)

				return html.PostPage(html.PageProps{}, post, time.Now()), nil
			}))
	})

	r.Get("/rss.xml", func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Set("Content-Type", "text/xml")

		var latest time.Time

		allPosts := posts.AllPosts

		for _, p := range allPosts {
			if p.Created.After(latest) {
				latest = p.Created
			}
			if p.Updated.After(latest) {
				latest = p.Updated
			}
		}

		author := feeds.Author{Name: "Philip", Email: "philip@quinefoundation.com"}

		feed := &feeds.Feed{
			Title:       "QUINE Global Organization – Solving yesterday's problems for tomorrow – Global health, business, and software blog",
			Link:        &feeds.Link{Href: "https://blog.quineglobal.com/"},
			Description: "Software development, global health, and business insights",
			Author:      &author,
			Created:     latest,
		}

		feed.Items = []*feeds.Item{}

		for _, p := range allPosts {

			var bodyBuffer bytes.Buffer
			for _, node := range p.Content {
				node.Render(&bodyBuffer)
			}

			body := bodyBuffer.String()
			description, err := htmlToText(body)
			if err != nil {
				fmt.Fprintf(os.Stderr, "failed to render description")
			}

			feed.Items = append(feed.Items, &feeds.Item{
				Title:       p.Title,
				Link:        &feeds.Link{Href: fmt.Sprintf("https://blog.quineglobal.com/post/%s", p.Id)},
				Description: description,
				Author:      &author,
				Created:     p.Created,
				Updated:     p.Updated,
				Content:     body,
			})
		}

		rss, err := feed.ToRss()
		if err != nil {
			http.Error(rw, "Could not render RSS feed", http.StatusInternalServerError)
			return
		}

		rw.Write([]byte(rss))
	})
}

func PostFetcherErrorer(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		postID := chi.URLParam(r, "postID")

		post, found := lo.Find(posts.AllPosts, func(q model.QuinePost) bool {
			return q.Id == postID
		})

		if !found {
			http.Error(w, "404, post not found", 404)
			return
		}

		ctx := context.WithValue(r.Context(), postCtxKey{}, post)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

type postCtxKey struct{}

func htmlToText(body string) (string, error) {
	r := strings.NewReader(body)
	doc, err := netHtml.Parse(r)
	if err != nil {
		return "", err
	}

	var out bytes.Buffer

	walk(doc, &out)

	return out.String(), nil
}

func walk(node *netHtml.Node, out *bytes.Buffer) {
	if node.Type == netHtml.ElementNode && (node.Data == "p" || node.Data == "div" || node.Data == "li") {
		out.WriteString("\n")
	}
	if node.Type == netHtml.TextNode {
		out.WriteString(node.Data)
	} else {
		children := node.ChildNodes()
		for c := range children {
			walk(c, out)
		}
	}
	if node.Type == netHtml.ElementNode && (node.Data == "p" || node.Data == "div" || node.Data == "li") {
		out.WriteString("\n")
	}
}
