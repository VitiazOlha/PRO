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
