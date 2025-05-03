package html

import (
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func CreditsPage(props PageProps) Node {
	props.Title = "QUINE – Website Credits"
	props.Header = false

	return page(props,
		Div(Class("prose prose-indigo prose-lg md:prose-xl"),
			P(
				Text("Roboto font is provided by Google and authored by Christian Robertson. "),
				A(Href("https://blog.quineglobal.com/static/fonts/LICENSE.txt"), Text("License here.")),
			),
		),
	)
}
