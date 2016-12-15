package uno

import (
	"testing"
)

func CreateTestGame() *Game {
	var deck CardDeck
	deck = append (deck, NewCard("red","1", 1, functions[0]))
	deck = append (deck, NewCard("blue","1", 1, functions[0]))
	deck = append (deck, NewCard("green","1", 1, functions[0]))
	deck = append (deck, NewCard("yellow","1", 1, functions[0]))
	deck = append (deck, NewCard("red","2", 1, functions[0]))
	deck = append (deck, NewCard("blue","⨂", 1, functions[1]))
	deck = append (deck, NewCard("green","⮂", 1, functions[2]))
	deck = append (deck, NewCard("red","+2", 1, functions[3]))
	deck = append (deck, NewCard("red","1", 1, functions[0]))
	deck = append (deck, NewCard("blue","1", 1, functions[0]))
	deck = append (deck, NewCard("green","1", 1, functions[0]))
	deck = append (deck, NewCard("yellow","1", 1, functions[0]))

	players := make([]CardDeck, playersCounter) 

	for index, _ := range players {
		players[index] = deck[0:1]
		deck = deck[2:]
	}

	newGame := Game{deck : deck[1:], pile : make(CardDeck, 0), players : players, top : deck[0], playerId : 0,route : 1, uno : false}

	for _, card := range newGame.deck {
		card.SetGame(&newGame)
	}

	for _, player := range newGame.players {
		for _, card := range player {
			card.SetGame(&newGame)
		}
	}

	newGame.top.SetGame(&newGame)

	return &newGame
}

func TestGameInteractions(t *testing.T) {
	
	newGame := CreateTestGame()
	
	oldId := newGame.playerId
	newGame.players[newGame.playerId][0].f(newGame, newGame.players[newGame.playerId][0])

	if  newGame.playerId == (len(newGame.players) + oldId + newGame.route) % len(newGame.players) {
		t.Log("Basic card works correct")
	} else {
		t.Error("Basic card works incorrect")
	}
}

func TestGameInteractions2(t *testing.T) {
	
	newGame := CreateTestGame()
	// fmt.Println("Test 2")
	oldId := newGame.playerId
	card := NewCard("blue","⨂", 1, functions[1])
	card.f(newGame, card)
	if  newGame.playerId == (len(newGame.players) + oldId + newGame.route * 2) % len(newGame.players) {
		t.Log("Skip card works correct")
	} else {
		t.Error("Skip card works incorrect")
	}
}
