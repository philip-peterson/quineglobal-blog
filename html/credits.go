package html

import (
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func CreditsPage(props PageProps) Node {
	props.Title = "QUINE – Website Credits"
	props.Header = false

	return page(props,
		Div(
			Class("markdown"),
			backToHome(),

			P(
				Text("This site uses system fonts for speed and reliability."),
			),
		),
	)
}
