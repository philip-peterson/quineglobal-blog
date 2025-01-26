package posts

import (
	"app/model"
	"time"

	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

var screwsAndSoftware = model.QuinePost{
	Title:       "Screws and Software",
	Id:          "screws-and-software",
	Teaser:      "What we fail to connect from the physical world to the software world",
	Created:     time.Date(2024, 7, 12, 0, 0, 0, 0, time.UTC),
	Updated:     time.Date(2024, 7, 12, 0, 0, 0, 0, time.UTC),
	FooterSegue: "For more insights about Software Development Lifecycle (SDLC) or thoughts about tech and business",
	Content: []Node{
		P(Text("What can screws teach us about coding?")),

		P(Text("When screws are being made, they undergo rigorous bending and rolling. This is often known as working the metal. Once fully wrought, these metal pellets are fired in an oven, which might seem strange at first. After all, why take something after it’s been shaped and soften it up by heating?")),

		P(Figure(
			A(
				Href("https://www.youtube.com/watch?v=3kxcw08p_oY"),
				Img(Src("https://i3.ytimg.com/vi/3kxcw08p_oY/hqdefault.jpg"), Width("480"), Height("360"), Alt("Youtube video about how bolts are made")),
			),
			FigCaption(Text("How It's Made: Nuts and Bolts")),
		)),

		P(Text("Well, the process of forming the screws"), Sup(Text("1")), Text(", while necessary, introduces microscopic regions of high torque, high compression, and high rarefaction. These stresses and imperfections make the screws problematic to work with, since they can cause the fasteners to break prematurely, leading to issues or even injury. You wouldn’t want to build a building with these overworked screws.")),

		P(Text("Luckily, there is a solution: heat. Since metal is in many ways a kind of crystal, the increased motion of the atoms with heat allows the metal to recrystallize. Regions of high stress and high internal torque are allowed to reduce and even disappear. (This is one of the rare cases in life where two wrongs can make a right.) Once the screws have been heated, they are returned to normal temperature, and are significantly stronger than before. Now, they are ready to be used in projects.")),

		P(Figure(
			A(Href("https://en.wikipedia.org/wiki/Recovery_(metallurgy)#Process"), Img(Src("https://upload.wikimedia.org/wikipedia/commons/3/38/Polygonization_animation.gif"), Width("522"), Height("300"), Alt("Image illustrating that several opposite defects may join to cancel each other out in a metal grain structure."))),

			FigCaption(Text("When two opposite dislocations are encouraged to meet up, they cancel out.")),
		)),

		P(Div(
			Text("This is not so different from what happens to software as we work on it. Exerting our will onto the code, while of utmost necessity, gradually makes the code harder to work with. Sticking points, such as:"),
			Ul(
				Li(Text("mismatched interfaces")),
				Li(Text("fudged types")),
				Li(Text("unvalidated assumptions")),
				Li(Text("stringified values, and")),
				Li(Text("too much responsibility per module")),
			),
		)),

		P(Text("are easily introduced, increasing obstacles and decreasing project velocity. You could even think of this like laundry or dishes accumulating in the home; a bit is fine, but too much and everything becomes unnavigable.")),

		P(Text("If we can add even a low-effort refactor phase in our development cycle, we can consistently cause these sticking points to be relaxed, making our next iteration faster and more effective. The next time you are planning a project, I would recommend to try budgeting at least 15% of time to performing refactors and micro-rewrites, and see how the project velocity responds.")),

		P(Text("Footnotes:")),
		P(Text("----")),

		P(
			Text("1. For our purposes, it’s simplest to call them screws, but actually we are speaking of bolts, and you can watch the manufacturing process "),
			A(Href("https://www.youtube.com/watch?v=3kxcw08p_oY"), Text("in this video")),
			Text("."),
		),
	},
}
