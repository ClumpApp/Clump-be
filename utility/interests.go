package utility

//var interestDiagram = make(map[string]map[string]string)

//interestDiagram["sports"] =

var interestDiagram = map[string]map[string]string{
	//sports
	"basketball": {
		"team sports": "sports"},
	"football": {
		"team sports": "sports"},
	"soccer": {
		"team sports": "sports"},
	"volleyball": {
		"team sports": "sports"},
	"hiking": {
		"individual sports": "sports"},
	"body building": {
		"individual sports": "sports"},
	"tennis": {
		"individual sports": "sports"},
	"extreme sports": {
		"individual sports": "sports"},
	//movies
	"action movies": {
		"brawl&guns": "movies"},
	"comedy movies": {
		"brawl&guns": "movies"},
	"drama movies": {
		"love&cry": "movies"},
	"romance movies": {
		"love&cry": "movies"},
	"horror movies": {
		"scream&solve": "movies"},
	"detective&mystery movies": {
		"scream&solve": "movies"},
	"sci-fi movies": {
		"magic&stuff": "movies"},
	"fantasy movies": {
		"magic&stuff": "movies"},
	//tv shows
	"action series": {
		"brawl&guns": "tv series"},
	"comedy series": {
		"brawl&guns": "tv series"},
	"drama series": {
		"love&cry": "tv series"},
	"romance series": {
		"love&cry": "tv series"},
	"horror series": {
		"scream&solve": "tv series"},
	"detective&mystery series": {
		"scream&solve": "tv series"},
	"sci-fi series": {
		"magic&stuff": "tv series"},
	"fantasy series": {
		"magic&stuff": "tv series"},
	//books
	"action books": {
		"brawl&guns": "books"},
	"comedy books": {
		"brawl&guns": "books"},
	"drama books": {
		"love&cry": "books"},
	"romance books": {
		"love&cry": "books"},
	"horror books": {
		"scream&solve": "books"},
	"detective&mystery books": {
		"scream&solve": "tv books"},
	"sci-fi books": {
		"magic&stuff": "books"},
	"fantasy books": {
		"magic&stuff": "books"},
	"young adult": {
		"young adult": "books"},
	//music
	"rock": {
		"bass&drums": "music"},
	"metal": {
		"bass&drums": "music"},
	"alternative": {
		"bass&drums": "music"},
	"pop": {
		"rhythm&beat": "music"},
	"electronic": {
		"rhythm&beat": "music"},
	"indie": {
		"rhythm&beat": "music"},
	"symphonic": {
		"strings&brass": "music"},
	"classical": {
		"strings&brass": "music"},
	//science&engineering
	"chemistry": {
		"fundamental": "science&engineering"},
	"biology": {
		"fundamental": "science&engineering"},
	"physics": {
		"fundamental": "science&engineering"},
	"maths": {
		"fundamental": "science&engineering"},
	"mechanical engineering": {
		"engineering": "science&engineering"},
	"electrical engineering": {
		"engineering": "science&engineering"},
	"computer engineering": {
		"engineering": "science&engineering"},
	//arts
	"painting": {
		"2d": "arts"},
	"photography": {
		"2d": "arts"},
	"sculpting": {
		"3d": "arts"},
	"crafting": {
		"3d": "arts"},
	//comics
	"movies": {
		"broadcast": "comics"},
	"tv shows": {
		"broadcast": "comics"},
	"comic books": {
		"print": "comics"},
	//anime&manga
	"shonen": {
		"brawl&guns": "movies"},
	"slice of life": {
		"love&cry": "movies"},
	"shojo": {
		"love&cry": "movies"},
	"psychological": {
		"scream&solve": "movies"},
	"seinen": {
		"scream&solve": "movies"},
	"sci-fi anime": {
		"magic&stuff": "movies"},
	"fantasy anime": {
		"magic&stuff": "movies"},
	//philosophy&psychology
	"ancient philosophy": {
		"philosophy": "philosophy&psychology"},
	"modern philosophy": {
		"philosophy": "philosophy&psychology"},
	"behavioral psychology": {
		"psychology": "philosophy&psychology"},
	"experimental psychology": {
		"psychology": "philosophy&psychology"},
}

//this is not exactly the same but sufficiently similar
/* var topInterests = []string{
	"sports",
	"movies",
	"tv shows",
	"books",
	"music",
	"science&engineering",
	"arts",
	"comics",
	"anime&manga",
	"philosophy&psychology",
}

var midInterests = [][]string{
	{"team sports", "individual sports"},
	{"love&cry", "brawl&guns", "magic&stuff"},
	{"love&cry", "brawl&guns", "magic&stuff"},
	{"love&cry", "brawl&guns", "magic&stuff", "young adult"},
	{"guitar&bass", "rhythm&beat", "strings&brass"},
	{"fundamental sciences", "engineering"},
	{"2d", "3d"},
	{"broadcast media", "print media"},
	{"love&cry", "brawl&guns", "magic&stuff"},
	{"philosophy", "psychology"},
}

var subInterests = [][]string{
	{"basketball", "football", "volleyball", "tennis", "soccer", "handball", "swimming", "hiking", "gymnastics", "crossfit", "extreme sports"},
	{"action", "horror", "comedy", "drama", "sci-fi", "fantasy", "romance"},
	{"action", "horror", "comedy", "drama", "sci-fi", "fantasy", "romance", "detective&mystery"},
	{"science/fiction", "fantasy", "classics", "romance", "young adult", "horror", "detective&mystery"},
	{"rock", "metal", "pop", "indie", "electronic", "alternative", "classic", "symphonic"},
	{"chemistry", "biology", "physics", "mathematics", "computer science", "mechanical engineering", "electrical engineering"},
	{"painting", "crafting", "sculpturing", "photography"},
	{"movies", "comics"},
	{"slice of life", "shojo", "shonen", "seinen", "horror", "fantasy", "drama", "comedy", "psychological", "slice of life"},
	{"ancient philosophy", "modern philosophy", "behavioral psychology", "experimental psychology"},
} */
