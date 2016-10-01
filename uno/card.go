package uno

import (
	"fmt"
)

type CardDeck [] *Card

const width  = 4
const height = 5

type Card struct {
	color string
	value string
	cost int
	f func()
}

func NewCard(color, value string, cost int, _ string) *Card {
	return &Card{color: color, value: value, cost: cost, f: Exit,}
}

func (c *Card) Print () {
	fmt.Println(c.color, c.value, c.cost)
}