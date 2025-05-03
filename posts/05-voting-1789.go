package posts

import (
	"app/model"
	"time"

	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

var voting1789 = model.QuinePost{
	Title:       "Voting like it's 1789",
	Id:          "voting-like-1789",
	Teaser:      "The majority has vanished.",
	Created:     time.Date(2025, 5, 3, 0, 0, 0, 0, time.UTC),
	Updated:     time.Date(2025, 5, 3, 0, 0, 0, 0, time.UTC),
	FooterSegue: "To hear more political coverage from Quine in the future",
	Content: []Node{
		P(Text("Imagine if the largest political group in America — the one with the most members — had no path to the presidency. Not because it lacked ideas, or candidates, or popular support, but because the system was never designed to acknowledge its existence. This is not a thought experiment. This is the condition of American democracy in 2025.")),

		P(Text("43% of Americans now identify as politically Independent. Only 28% are Republican. Another 28% are Democrat. The numbers are not ambiguous. If elections reflected reality, we would have an Independent president, or at least a contest that allowed for one — but we don't. Instead, we have a legacy system, \"first-past-the-post voting\", which continues because of its simplicity, not because it gives representation.")),

		P(Text("There is a rot that sets in when systems persist too long without iteration or modification. In software, this becomes technical debt. In governance, it becomes misalignment so severe that even consensus is invisible. Voters are told they must choose between two options they do not prefer. Any third option is \"spoiler,\" \"vanity,\" \"waste.\" The will of the majority is not lost — it is structurally excluded.")),

		P(Text("There are ways out: Ranked-choice voting. Approval voting. Basic patches to a system now brittle from neglect. These aren't theoretical; they are mathematically sound, already deployed in cities and countries that understand the stakes. They allow systems to reflect who is actually there, not just what legacy powers once were.")),

		P(Text("Democracy without feedback is just inertia. And inertia, left unchecked, becomes entropy. A misrepresented public cannot respond to crisis. Not ecological, not economic, not moral. The clock is running. The majority is here. The system is late.")),
	},
}
