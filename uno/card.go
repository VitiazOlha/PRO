package uno

type CardDeck [] *Card

const width  = 4
const height = 5

type Card struct {
	color string
	value string
	cost int
	f func(deck CardDeck)
	game *Game
}

func Basic(deck CardDeck){
}

func Stop(deck CardDeck){
}

func Reverse(deck CardDeck){
}

func StopPlus(deck CardDeck){
}

func Color(deck CardDeck){
}

func ColorPlus(deck CardDeck){
}

func NewCard(color, value string, cost int, f func(deck CardDeck)) *Card {
	return &Card{color: color, value: value, cost: cost, f: f, game: nil,}
}

func (c *Card) SetGame(game *Game) {
	c.game = game
}
/*
func (c *Card) Draw(s *tl.Screen) {
	b.r.Draw(s)
	b.text.Draw(s)
}

func (c *Card) Tick(ev tl.Event) {
	x, y := b.r.Position()
	h, w := b.r.Size()
	if ev.Type == tl.EventMouse && ev.Key == tl.MouseLeft && (ev.MouseX > x && ev.MouseX < x + h) && (ev.MouseY > y && ev.MouseY < y + w) {
		
		c.f()
	}
}*/