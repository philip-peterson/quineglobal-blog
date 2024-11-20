package http

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	feeds "github.com/gorilla/feeds"
	. "maragu.dev/gomponents"
	hx "maragu.dev/gomponents-htmx/http"
	ghttp "maragu.dev/gomponents/http"

	"app/html"
	"app/model"
	"app/posts"
)

const body = `<?xml version='1.0' encoding='UTF-8'?>
	<rss version='2.0'>
		<channel>
		<title>QUINE Foundation software and business blog | RSS</title>
		<link>http://blog.quinefoundation.com/post/</link>
		<description>Software development and business insights</description>
		<language>en-us</language><item>
			<title>QUINE Core</title>
			<link>http://blog.quinefoundation.com/post/declarative-stateless</link>
			<description>The QUINE Foundation is strongly in favor of declarative and stateless systems. Recently, we have been battling entropy, specifically the entropy of hosted virtual machines and how their shelf lives are limited.
		In the pursuance of disaster-tolerance, sometimes we must tread new ground. The QUINE Foundation has now developed a stateless configuration for its entire host of services.</description>
			<pubDate>Sun, 25 Aug 2024 00:23:26 +0000</pubDate>
			</item>
			<item>
			<title>Look Where You're Headed</title>
			<link>http://blog.quinefoundation.com/post/look-where-youre-headed</link>
			<description>Years ago, I was standing in a cage in some guy's basement. The guy was someone I had hired to do personal training in San Francisco, so it wasn't quite as weird or terrifying as it sounds.
		As I was standing there, lifting weights improperly, the trainer said something that stuck with me years later and transcended the context of gym workouts, joining the world of software development: "Look straight ahead."</description>
			<pubDate>Sun, 25 Aug 2024 00:23:26 +0000</pubDate>
			</item><item>
			<title>Screws and Software</title>
			<link>http://blog.quinefoundation.com/post/screws-and-software</link>
			<description>What can screws teach us about coding?
		When screws are being made, they undergo rigorous bending and rolling. This is often known as working the metal. Once fully wrought, these metal pellets are fired in an oven, which might seem strange at first. After all, why take something after it’s been shaped and soften it up by heating?</description>
			<pubDate>Sun, 25 Aug 2024 00:23:26 +0000</pubDate>
			</item></channel>
	</rss>
`

type thingsGetter interface {
	GetThings(ctx context.Context) ([]model.Thing, error)
}

func rssRouter() chi.Router {
	r := chi.NewRouter()
	r.Use(render.SetContentType(render.ContentTypeJSON))
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("admin: index"))
	})
	return r
}

// Home handler for the home page, as well as HTMX partial for getting things.
func Home(r chi.Router, db thingsGetter) {
	r.Get("/", ghttp.Adapt(func(w http.ResponseWriter, r *http.Request) (Node, error) {
		things, err := db.GetThings(r.Context())
		if err != nil {
			return nil, err
		}

		if hx.IsRequest(r.Header) {
			return html.ThingsPartial(things, time.Now()), nil
		}

		return html.HomePage(html.PageProps{}, things, time.Now()), nil
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

		allPosts := []posts.QuinePost{posts.FirstPost}

		feed.Items = []*feeds.Item{}

		for _, p := range allPosts {
			feed.Items = append(feed.Items, &feeds.Item{
				Title:       p.Title,
				Link:        &feeds.Link{Href: fmt.Sprintf("https://blog.quinefoundation.com/post/%s", p.Id)},
				Description: "WIP WIP WIP",
				Author:      &author,
				Created:     p.Created,
				Updated:     p.Updated,
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
