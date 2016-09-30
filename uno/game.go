package uno

import (
	"fmt"
	)


var playersCounter = 2

func NewGame() {
	//deck := newCardDeck()
	players := make([][]Card, playersCounter) 

	for index, _ := range players {
		players[index] = make([]Card, 0, 10)
	}

	fmt.Println("This is new game")
}