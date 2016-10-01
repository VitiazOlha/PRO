package uno

import (
	"fmt"
	"encoding/json"
	"io/ioutil"
	)


var playersCounter = 2

func checkErr(e error) {
	if e != nil {
		panic(e)
	}
}

func NewGame() {

	deck := newCardDeck()

	for _, card := range deck {
		card.Print()
	}
	
	players := make([][]Card, playersCounter) 

	for index, _ := range players {
		players[index] = make([]Card, 0, 10)
	}

	fmt.Println("This is new game")
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
	return NewCard(
		data["color"].(string),
		data["value"].(string),
		int(data["cost"].(float64)),
		data["f"].(string),
	)
}