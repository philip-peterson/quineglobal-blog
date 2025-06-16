package posts

import (
	"app/model"
	"time"

	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

var netua = model.QuinePost{
	Title:       "The People of Netua",
	Id:          "netua",
	Teaser:      "A declaration",
	Created:     time.Date(2019, 12, 5, 0, 0, 0, 0, time.UTC),
	Updated:     time.Date(2019, 12, 5, 0, 0, 0, 0, time.UTC),
	FooterSegue: "For more from Quine",
	Content: []Node{
		P(Text("Written 2019-12-05")),

		P(Text("We the people of Netua, do not necessarily agree on the things, and an attack on one of us is not necessarily an attack on all of us. We believe in the nonspecific and the nonexhaustive. Contradict one of us, and you have contradicted, well... only one of us, because we do not fight for each other or ourselves. We do not fight for any nations. We respect the high bidder and believe that wars are won mostly out of chance.")),

		P(Text("We stand together on our lack of unification. Our lack of group identity is our distinct identity. We are not Anonymous, since we have a name, but the name must mean nothing to all people or the name must change.")),

		P(Text("We have no allegiances. We produce no product. We believe, of course, in production, but also reject that belief. We always invalidate our own arguments and disclaim every claim.")),

		P(Text("We respect data, but choose not to collect it. We respect conclusions, but refuse to make any. We respect authority, but concede that we should not be the ones with any.")),

		P(Text("Anything you hate about us is automatically invalid, since we are both everything and nothing.")),
	},
}
