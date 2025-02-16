package html

import (
	"fmt"
	"time"

	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"

	"app/model"
)

// HomePage is the front page of the app.
func HomePage(props PageProps, posts []model.QuinePost, now time.Time) Node {
	props.Title = "QUINE Global Organization – Solving yesterday's problems for tomorrow – Global health, business, and software blog"
	props.Header = true

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
			return PostTeaser(t, now)
		}),
	}
}

func PostTeaser(post model.QuinePost, now time.Time) Node {
	return Div(
		Class("blog-post"),
		H1(
			A(
				Href(fmt.Sprintf("/post/%s", post.Id)),
				Text(post.Title),
			),
		),
		P(Text(post.Teaser), Text("…")),
		A(Href(fmt.Sprintf("/post/%s", post.Id)), Text("Read post")),
	)
}

func PostReader(post model.QuinePost, now time.Time) Node {
	footerSegue := post.FooterSegue
	if footerSegue == "" {
		footerSegue = "If you liked this post"
	}

	return Div(
		Div(
			Class("markdown"),
			H1(
				Text(post.Title),
			),
			Div(post.Content...),
			P(
				Style("margin-top: 3em"),
				Text(footerSegue),
				Text(", you can "),
				A(Href("http://blog.quineglobal.com/rss.xml"), Text("follow")),
				Text(" our thinkpieces via "),
				A(Href("http://blog.quineglobal.com/rss.xml"), Text("RSS")),
				Text(", or you can "),
				A(Href("https://www.linkedin.com/company/quine-global"), Text("find us on LinkedIn")),
				Text("!"),
			),
		),
	)
}
