package html

import (
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func AboutPage(props PageProps) Node {
	props.Title = "QUINE – About"
	props.Header = false

	return page(props,
		Div(Class("prose prose-indigo prose-lg md:prose-xl"),
			P(
				Text("Can tech be good? Quine Global believes it can. We are fighting to improve Earth for humanity by solving technological problems that block global health. We also want to be good stewards to nature generally."),
			),

			Figure(
				Img(Src("/images/forest.png"), Alt("Raking the forest leaves")),
				FigCaption(Text("The forest doesn't rake itself.")),
			),

			P(
				Text("Why aren't server rooms placed in the arctic, where cooling is done by nature already? What is next for ensuring quality of the air and sea? That's what we want to find out."),
			),

			P(
				Text("For more about why we exist, there is also "),

				A(Href("/post/what-is-quine"), Text("this blog post")),

				Text("."),
			),

			P(
				Text("Would you like to join us? You can "),
				A(Href("https://www.linkedin.com/company/quine-global"), Text("follow us on LinkedIn")),
				Text(", or "),
				A(Href("mailto:philip@quineglobal.com"), Text("email the author.")),
			),
		),
	)
}
