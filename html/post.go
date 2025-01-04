package html

import (
	"time"

	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"

	"app/model"
)

func PostPage(props PageProps, post model.QuinePost, now time.Time) Node {
	props.Title = post.Title + " | QUINE"

	return page(props,
		Div(
			Div(
				PostReader(post, now),
			),
		),
	)
}
