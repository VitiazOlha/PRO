package uno

type CardDeck [] Card

const width  = 4
const height = 5

type Card struct {
	color string
	value string
	cost int
	f func()
}
