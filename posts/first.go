package posts

import (
	"time"

	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

type QuinePost struct {
	Title   string
	Id      string // used as guid in rss, id in atom
	Updated time.Time
	Created time.Time
	Content []Node
}

var FirstPost = QuinePost{
	Title:   "Look Where You're Headed",
	Id:      "look-where-youre-headed",
	Created: time.Date(2024, 8, 25, 0, 23, 26, 0, time.UTC),
	Updated: time.Date(2024, 8, 25, 0, 23, 26, 0, time.UTC),
	Content: []Node{
		Div(Class("min-h-screen flex flex-col justify-between bg-white"),
			Div(Class("grow"),
				A(Href("/"), Class("inline-flex items-center text-xl font-semibold"),
					Img(Src("/images/logo.png"), Alt("Logo"), Class("h-12 w-auto bg-white rounded-full mr-4")),
					Text("Home"),
				),
			),
		),
	},
}
