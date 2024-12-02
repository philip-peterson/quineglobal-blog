package http

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	feeds "github.com/gorilla/feeds"
	strip "github.com/grokify/html-strip-tags-go"
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

// Home handler for the home page, as well as HTMX partial for getting things.
func Home(r chi.Router, db thingsGetter) {
	r.Get("/", ghttp.Adapt(func(w http.ResponseWriter, r *http.Request) (Node, error) {
		allPosts := posts.AllPosts

		if hx.IsRequest(r.Header) {
			return html.Posts(allPosts, time.Now()), nil
		}

		return html.HomePage(html.PageProps{}, allPosts, time.Now()), nil
	}))

	r.Get("/rss.xml", func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Set("Content-Type", "text/xml")

		now := time.Now()

		author := feeds.Author{Name: "Philip Peterson", Email: "peterson@sent.com"}

		feed := &feeds.Feed{
			Title:       "QUINE Foundation software, global health, and business blog",
			Link:        &feeds.Link{Href: "https://blog.quinefoundation.com/"},
			Description: "Software development, global health, and business insights",
			Author:      &author,
			Created:     now,
		}

		allPosts := posts.AllPosts

		feed.Items = []*feeds.Item{}

		for _, p := range allPosts {

			var bodyBuffer bytes.Buffer
			for _, node := range p.Content {
				node.Render(&bodyBuffer)
			}

			body := bodyBuffer.String()
			bodyStripped := strip.StripTags(body)

			feed.Items = append(feed.Items, &feeds.Item{
				Title:       p.Title,
				Link:        &feeds.Link{Href: fmt.Sprintf("https://blog.quinefoundation.com/post/%s", p.Id)},
				Description: bodyStripped,
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
