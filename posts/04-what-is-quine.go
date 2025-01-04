package posts

import (
	"app/model"
	"time"

	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

var whatIsQuine = model.QuinePost{
	Title:       "What is QUINE?",
	Id:          "what-is-quine",
	Teaser:      "Answering the question at last",
	Created:     time.Date(2025, 1, 4, 0, 0, 0, 0, time.UTC),
	Updated:     time.Date(2025, 1, 4, 0, 0, 0, 0, time.UTC),
	FooterSegue: "If you'd like to hear more from Quine in the future",
	Content: []Node{
		P(Text("At the heart of all software is a special space carved out by the famous “Quine” paradigm. A quine is a program that when run, produces its own source code as output.")),

		Div(
			Style("margin-top: 2em; margin-bottom: 2em"),

			P(Code(Text("((lambda (x) (list x (list 'quote x))) '(lambda (x) (list x (list 'quote x))))"))),

			P(Style("text-align: center; font-style: italic"), Text("A quine written in Common Lisp")),
		),

		P(Text("The quine then, serves as a powerful metaphor for our world. “We’re here because we’re here,” as some "), A(Style("color: inherit"), Href("https://www.youtube.com/watch?v=jSrqC_angdc"), Text("have said")), Text(".")),

		P(Text("Global health does not occur in a vacuum. It requires the attentive energy of millions, and determination that can only be attained through vivid self-reflection.")),

		P(Text("It is with these insights that QUINE Global has been formed and sets off into the wild unknown. What are the major blockers to human health? There is much yet left to discover. QUINE therefore then dedicates itself to uncovering truth and insights through the form of simple business blogs. But we’re not just about blogging. We also produce and open-source our own software, and hold a strong conviction that software is both, in some ways, what got us into this mess – and what will get us out of it.")),

		P(Text("The principal crises of this world include pollution, lack of access to solutions, and lack of clarity. Quine dedicates itself to solve all three. But we can only do it with your help.")),
	},
}
