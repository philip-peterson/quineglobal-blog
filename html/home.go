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
		ComputingTicker(),
		Posts(posts, now),
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
	url := fmt.Sprintf("/post/%s", post.Id)
	return A(
		Class("post-teaser"),
		Href(url),
		Time(
			Attr("datetime", post.Created.Format(time.RFC3339)),
			Class("post-date"),
			Text(post.Created.Format("January 2006")),
		),
		H2(Text(post.Title)),
		P(Class("teaser"), Text(post.Teaser), Text("…")),
		Span(Class("read-more"), Text("Read →")),
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

// ComputingTicker renders a scrolling marquee of delightfully nonsensical
// computing / Lisp-flavored facts on the home page.
func ComputingTicker() Node {
	facts := []string{
		`(cons '(4) '(3)) == '(4 3)`,
		`((lambda (x) x) 'quine) == 'quine`,
		`(car (cdr '(a b c))) => b`,
		`(append '() '(x y)) == '(x y)`,
		`(map (lambda (x) (* x x)) '(1 2 3)) == '(1 4 9)`,
		`'(this is not evaluated)`,
		`(eq? (list) '()) => #f`,
		`(+ 40 2) == 42`,
		"`(1 ,@'(2 3) 4) => (1 2 3 4)`",
		`quines output their own source code`,
		`(eval (read "(+ 1 1)")) => 2`,
		`(omega (lambda (x) (x x))) never returns`,
	}

	sep := " • "

	// Split facts roughly in half for two scrolling lines
	mid := (len(facts) + 1) / 2
	row1Facts := facts[:mid]
	row2Facts := facts[mid:]

	joinFacts := func(fs []string) string {
		s := ""
		for i, f := range fs {
			if i > 0 {
				s += sep
			}
			s += f
		}
		return s
	}

	// Duplicate each row's content for seamless scrolling
	row1 := joinFacts(row1Facts) + sep
	row2 := joinFacts(row2Facts) + sep

	return Div(
		Class("computing-ticker"),
		Div(
			Class("computing-ticker-row"),
			Div(Class("computing-ticker-inner"), Span(Text(row1)), Span(Text(row1))),
		),
		Div(
			Class("computing-ticker-row"),
			Div(Class("computing-ticker-inner"), Span(Text(row2)), Span(Text(row2))),
		),
	)
}
