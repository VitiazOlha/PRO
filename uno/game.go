package uno

import (
	"encoding/json"
	"io/ioutil"
	"time"
	"math/rand"
	"math"
	)


var playersCounter = 2

func checkErr(e error) {
	if e != nil {
		panic(e)
	}
}

type Game struct {
	deck CardDeck
	pile CardDeck
	players []CardDeck
	playerId int
	top *Card
	route int
} 

var functions = [...] func(deck CardDeck) {
 Basic,
 Stop,
 Reverse,
 StopPlus,
 Color,
 ColorPlus,
}

func NewGame() {

	deck := newCardDeck()
	deck = deck.shuffleDeck()

	players := make([]CardDeck, playersCounter) 

	for index, _ := range players {
		players[index] = deck[0:7]
		deck = deck[8:]
	}
	newGame := Game{deck : deck[1:], pile : make(CardDeck, 0), players : players, top : deck[0], playerId : 0,route : 1,}
	//crutch
	for _, card := range newGame.deck {
		card.SetGame(&newGame)
	}
	for _, player := range newGame.players {
		for _, card := range player {
			card.SetGame(&newGame)
		}
	}
	newGame.top.SetGame(&newGame)
	
	newGame.playGame()
}

func (game Game) playGame() {
	game.drawGame()

}

func (game Game) drawGame() {

}

func (deck CardDeck) shuffleDeck() CardDeck {
	var newDeck CardDeck

	for len(deck) > 0 {
		i := int(math.Abs(float64(int(time.Now().UnixNano()) * rand.Intn(100)))) % len(deck)
		newDeck = append(newDeck, deck[i])
		copy(deck[i:], deck[i+1:])
		deck[len(deck)-1] = nil
		deck = deck[:len(deck)-1]
	}
	return newDeck
}


func newCardDeck() CardDeck {
	var deck CardDeck

	ldeck, err := ioutil.ReadFile("uno/deck.json")
	checkErr(err)

	var f interface{}
    err = json.Unmarshal(ldeck, &f)
    checkErr(err)

    m := f.([]interface{})
	
	for _, v := range m {
		vv := v.(map[string]interface{})
		q := int(vv["quantity"].(float64))
		for i := 0; i < q; i++ {
   			deck = append(deck, parseCard(vv))
   		}
	}
	return deck
}

func parseCard(data map[string]interface{}) *Card{
	f_num := int(data["f"].(float64)) - 1
	return NewCard(
		data["color"].(string),
		data["value"].(string),
		int(data["cost"].(float64)),
		functions[f_num],
	)
}