package html

import (
	"time"

	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"

	"app/model"
)

func PostPage(props PageProps, post model.QuinePost, now time.Time) Node {
	props.Title = "QUINE Foundation Blog"

	return page(props,
		Div(Class("prose prose-indigo prose-lg md:prose-xl"),
			Div(
				PostReader(post, now),
			),
		),
	)
}
