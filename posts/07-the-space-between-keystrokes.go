package posts

import (
	"app/model"
	"time"

	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

var spaceBetweenKeystrokes = model.QuinePost{
	Title:       "The Space Between Keystrokes",
	Id:          "space-between-keystrokes",
	Teaser:      "Introducing Quine Hyper",
	Created:     time.Date(2026, 5, 29, 0, 0, 0, 0, time.UTC),
	Updated:     time.Date(2026, 5, 29, 0, 0, 0, 0, time.UTC),
	FooterSegue: "To follow the development of Quine Hyper",
	Content: []Node{
		P(Text("Today we are releasing Quine Hyper, a desktop terminal built for focus.")),

		P(Text("Hyper is built on top of the excellent open-source Hyper project. Our goal was simple: give people a terminal that gets out of the way. Clean defaults, calm typography, nothing demanding your attention that isn't your own work.")),

		P(
			Text("This release includes updated dependencies, a collection of bugfixes, and new features. Downloads and release notes are available at "),
			A(Href("https://hyper.quineglobal.com"), Text("hyper.quineglobal.com")),
			Text("."),
		),

		P(Text("Quine Hyper is available today for:")),

		Ul(
			Li(Text("macOS (Apple Silicon and Intel)")),
			Li(Text("Windows (x64)")),
			Li(Text("Linux (x64 and ARM64, AppImage and RPM)")),
		),
	},
}
