package posts

import (
	"app/model"
	"time"

	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

var lookWhereYoureHeaded = model.QuinePost{
	Title:   "Look Where You're Headed",
	Id:      "look-where-youre-headed",
	Created: time.Date(2024, 8, 25, 0, 23, 26, 0, time.UTC),
	Updated: time.Date(2024, 8, 25, 0, 23, 26, 0, time.UTC),
	Content: []Node{
		Div(
			P(Text("Years ago, I was standing in a cage in some guy's basement. The guy was someone I had hired to do personal training in San Francisco, so it wasn't quite as weird or terrifying as it sounds.")),

			P(Text("As I was standing there, lifting weights improperly, the trainer said something that stuck with me years later and transcended the context of gym workouts, joining the world of software development: \"Look straight ahead.\"")),

			P(Text("It doesn't seem like it should matter where you look when doing something unrelated to your eyeballs, such as making a salad, or riding a bike. After all, the eyes move independent of the rest of the body, right? So why is it relevant where we look when lifting a weight, if all the other parts of the body are moving properly?")),

			P(Text("An experiment you can try: next time you are on a bicycle, start biking straight. Then, while going straight, look over to your left and focus on an object in the distance. Despite your best efforts, the wheel of the bike will gradually tend toward the left. Why?")),

			P(Text("As humans, we want to know where we're going. But also, due to our psychology, we tend to "), Em(Text("go where we are knowing")), Text(". If we know about left, we will go left, almost involuntarily. And if you look straight down while lifting a weight, your back and neck will do bad things they were never supposed to do.")),

			P(Text("This insight is relevant to the workplace. If we know where we are going, we will go there, even if unwittingly. But if we don't know, we might look somewhere else and land ourselves in trouble by going somewhere else.")),

			P(Text("And for managers, if your reports cannot be given any time specifically to attaining a goal, just having them know the direction alone, and knowing your team's goals, can get your organization closer to achieving them.")),
		),
	},
}
