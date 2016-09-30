package uno

type CardDeck [] Card

const width  = 4
const height = 5
// Provides a general Drawable to be rendered.

type Card struct {
	color string
	value string
	cost int
	f func()
}

func NewCard(color, value string, cost int, _ func() ) *Card {
	return &Card{color: color, value: value, cost: cost, f: Exit,}
}
