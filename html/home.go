package html

import (
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

	// Prepare full list of other posts for client-side re-shuffling
	type otherPostJSON struct {
		ID    string `json:"id"`
		Title string `json:"title"`
	}
	var allOtherPostsJSON []otherPostJSON
	for _, p := range otherPosts {
		allOtherPostsJSON = append(allOtherPostsJSON, otherPostJSON{
			ID:    p.Id,
			Title: p.Title,
		})
	}
	otherPostsJSON, _ := json.Marshal(allOtherPostsJSON)

	return Div(
		Div(
			Class("markdown"),
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
			Div(
				Class("other-posts-header"),
				H3(Text("Other posts")),
				Button(
					ID("reshuffle-posts"),
					Class("reshuffle-btn"),
					Text("Re-shuffle"),
				),
			),
			Script(Type("application/json"), ID("other-posts-data"), Text(string(otherPostsJSON))),
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
				A(Href("https://quineglobal.com"), Class("global-projects"), Text("Check out all our cool Global Projects >")),
			),
		),

		// Client-side re-shuffle script for Other posts
		Script(Raw(`
			(function() {
				const container = document.querySelector('.scattered-posts');
				const dataEl = document.getElementById('other-posts-data');
				const btn = document.getElementById('reshuffle-posts');

				if (!container || !dataEl || !btn) return;

				let allPosts = [];
				try {
					allPosts = JSON.parse(dataEl.textContent);
				} catch (e) {
					console.error('Failed to parse other posts data');
					return;
				}

				function generateLayout(n) {
					const layout = [];
					for (let i = 0; i < n; i++) {
						layout.push({
							top: Math.floor(Math.random() * 68) + 6,
							left: Math.floor(Math.random() * 78) + 4,
							rot: Math.floor(Math.random() * 15) - 7
						});
					}
					return layout;
				}

				function scoreLayout(layout) {
					if (layout.length <= 1) return 999;
					let minDist = 1000;
					for (let i = 0; i < layout.length; i++) {
						for (let j = i + 1; j < layout.length; j++) {
							const dx = layout[i].left - layout[j].left;
							const dy = layout[i].top - layout[j].top;
							const dist = Math.sqrt(dx * dx + dy * dy);
							if (dist < minDist) minDist = dist;
						}
					}
					return minDist;
				}

				function bestScatteredLayout(n, attempts = 60) {
					if (n <= 0) return [];
					let best = generateLayout(n);
					let bestScore = scoreLayout(best);
					for (let i = 0; i < attempts; i++) {
						const candidate = generateLayout(n);
						const s = scoreLayout(candidate);
						if (s > bestScore) {
							bestScore = s;
							best = candidate;
						}
					}
					return best;
				}

				function shuffleArray(array) {
					const arr = array.slice();
					for (let i = arr.length - 1; i > 0; i--) {
						const j = Math.floor(Math.random() * (i + 1));
						[arr[i], arr[j]] = [arr[j], arr[i]];
					}
					return arr;
				}

				function renderScatteredPosts(posts) {
					if (!posts || posts.length === 0) return;
					const layout = bestScatteredLayout(posts.length);
					container.innerHTML = '';

					posts.forEach((post, i) => {
						const pos = layout[i] || { top: 20, left: 10, rot: 0 };
						const a = document.createElement('a');
						a.className = 'scattered-post';
						a.href = '/post/' + post.id;
						a.style.top = pos.top + '%';
						a.style.left = pos.left + '%';
						a.style.transform = 'rotate(' + pos.rot + 'deg)';
						a.textContent = post.title;
						container.appendChild(a);
					});
				}

				btn.addEventListener('click', () => {
					if (allPosts.length === 0) return;
					// Pick up to 4 random posts
					const shuffled = shuffleArray(allPosts);
					const selected = shuffled.slice(0, Math.min(4, shuffled.length));
					renderScatteredPosts(selected);
				});
			})();
		`)),
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
