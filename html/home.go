package html

import (
	"time"

	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"

	"app/model"
)

// HomePage is the front page of the app.
func HomePage(props PageProps, posts []model.QuinePost, now time.Time) Node {
	props.Title = "QUINE Foundation Blog"

	return page(props,
		Div(Class("prose prose-indigo prose-lg md:prose-xl"),
			Div(
				Posts(posts, now),
			),
		),
	)
}

func Posts(posts []model.QuinePost, now time.Time) Node {
	return Group{
		Map(posts, func(t model.QuinePost) Node {
			return Li(
				Class("blog-post"),
				H1(
					A(
						Href("/post/screws-and-software"),
						Text("Screws and Software"),
					),
				),
				P(Text("What can screws teach us about coding?")),
				A(Href("/post/screws-and-software"), Text("Read post")),
			)
		}),
	}
}
