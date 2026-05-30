package html

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"math"
	"math/rand/v2"
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
		Div(
			Class("post-teaser-content"),
			Time(
				Attr("datetime", post.Created.Format(time.RFC3339)),
				Class("post-date"),
				Text(post.Created.Format("January 2006")),
			),
			H2(Text(post.Title)),
			P(Class("teaser"), Text(post.Teaser), Text("…")),
			Span(Class("read-more"), Text("Read →")),
		),
		// Hover slash decoration (right half, sequential animation)
		Div(
			Class("post-slash-container"),
			Span(Class("slash")),
			Span(Class("slash")),
			Span(Class("slash")),
			Span(Class("slash")),
			Span(Class("slash")),
		),
	)
}

// cardPos holds the position for one scattered post card (in % and degrees).
type cardPos struct {
	Top  int
	Left int
	Rot  int
}

// generateLayout creates a random layout for n cards.
func generateLayout(n int) []cardPos {
	layout := make([]cardPos, n)
	for i := 0; i < n; i++ {
		layout[i] = cardPos{
			Top:  rand.IntN(68) + 6,   // 6%–73%
			Left: rand.IntN(78) + 4,   // 4%–81%
			Rot:  rand.IntN(15) - 7,   // -7° to +7°
		}
	}
	return layout
}

// scoreLayout returns a score for how well spaced the cards are.
// Higher is better (currently based on minimum pairwise distance).
func scoreLayout(layout []cardPos) float64 {
	if len(layout) <= 1 {
		return 999
	}
	minDist := 1000.0
	for i := 0; i < len(layout); i++ {
		for j := i + 1; j < len(layout); j++ {
			dx := float64(layout[i].Left - layout[j].Left)
			dy := float64(layout[i].Top - layout[j].Top)
			dist := math.Sqrt(dx*dx + dy*dy)
			if dist < minDist {
				minDist = dist
			}
		}
	}
	return minDist
}

// bestScatteredLayout tries many random layouts and returns the best one.
func bestScatteredLayout(n int) []cardPos {
	if n <= 0 {
		return nil
	}
	const attempts = 80
	best := generateLayout(n)
	bestScore := scoreLayout(best)

	for i := 0; i < attempts; i++ {
		candidate := generateLayout(n)
		s := scoreLayout(candidate)
		if s > bestScore {
			bestScore = s
			best = candidate
		}
	}
	return best
}

func PostReader(post model.QuinePost, otherPosts []model.QuinePost, now time.Time) Node {
	footerSegue := post.FooterSegue
	if footerSegue == "" {
		footerSegue = "If you liked this post"
	}

	// Limit to a few other posts for initial render
	otherPostsToShow := otherPosts
	if len(otherPostsToShow) > 4 {
		otherPostsToShow = otherPostsToShow[:4]
	}

	// Prepare full list for client-side re-shuffling (simple & light)
	type otherPostJSON struct {
		ID    string `json:"id"`
		Title string `json:"title"`
	}
	var allOtherPostsJSON []otherPostJSON
	for _, p := range otherPosts {
		allOtherPostsJSON = append(allOtherPostsJSON, otherPostJSON{ID: p.Id, Title: p.Title})
	}
	otherPostsJSON, _ := json.Marshal(allOtherPostsJSON)
	otherPostsDataAttr := base64.StdEncoding.EncodeToString(otherPostsJSON)

	return Div(
		Div(
			Class("markdown"),
			P(
				Class("back-to-home"),
				A(Href("/"), Class("back-to-posts"), Text("« Back to all posts")),
			),
			H1(
				Text(post.Title),
			),
			// Automatically show the date from the post's Created field
			P(
				Class("post-date"),
				Text("Written " + post.Created.Format("January 2, 2006")),
			),
			Div(post.Content...),
			P(
				Class("post-footer-note"),
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

		// Other posts section
		Div(
			Class("other-posts"),
			Attr("data-other-posts", otherPostsDataAttr),
			Div(
				Class("other-posts-header"),
				H3(Text("Other posts")),
				Button(
					ID("reshuffle-posts"),
					Class("reshuffle-btn"),
					Text("Re-shuffle"),
				),
			),
			Div(
				Class("scattered-posts"),
				func() Node {
					layout := bestScatteredLayout(len(otherPostsToShow))
					nodes := make([]Node, len(otherPostsToShow))
					for i, p := range otherPostsToShow {
						var style string
						if i < len(layout) {
							pos := layout[i]
							style = fmt.Sprintf("top:%d%%;left:%d%%;transform:rotate(%ddeg)", pos.Top, pos.Left, pos.Rot)
						}
						nodes[i] = A(
							Class("scattered-post"),
							Style(style),
							Href(fmt.Sprintf("/post/%s", p.Id)),
							Text(p.Title),
						)
					}
					return Group(nodes)
				}(),
			),
			Div(
				Class("other-posts-actions"),
				A(Href("/"), Class("view-all-posts"), Text("← View all posts")),
				Span(Class("other-posts-actions-spacer")),
				A(Href("https://quineglobal.com"), Class("global-projects"), Text("Check out all our cool Global Projects →")),
			),
		),
	)
}

// ComputingTicker renders a scrolling marquee of philosophical and computational claims.
func ComputingTicker() Node {
	facts := []string{
		`(cons '(4) '(3)) == '(4 3)`,
		`((lambda (x) x) 'quine) == 'quine`,
		`(car (cdr '(a b c))) => b`,
		`(append '() '(x y)) == '(x y)`,
		`quines output their own source code`,
		`((lambda (x) (x x)) (lambda (x) (x x))) never terminates`,
		`Nihilism is more provable but constructivism is the answer`,
		`Sometimes a local minimum can be escaped with minimal activation energy`,
		`Free will may be illusory, but the experience of choice is not`,
		`(omega (lambda (x) (x x))) never returns`,
	}

	sep := " • "
	nbsp := "\u00A0" // non-breaking space, prevents whitespace collapsing at marquee seam

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

	// Duplicate each row's content for seamless scrolling.
	// We append sep + nbsp so the seam between the two identical copies
	// always shows a full " • " (space-bullet-space) instead of a collapsed
	// "space bullet" at the loop point.
	row1 := joinFacts(row1Facts) + sep + nbsp
	row2 := joinFacts(row2Facts) + sep + nbsp

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
