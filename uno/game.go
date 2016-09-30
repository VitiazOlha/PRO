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
	players := make([][]Card, playersCounter) 

	for index, _ := range players {
		players[index] = make([]Card, 0, 10)
	}

	fmt.Println("This is new game")
}


func newCardDeck() *CardDeck {
	
	ldeck, err := ioutil.ReadFile("deck.json")
	checkErr(err)
	parsers := make(map[string] tl.EntityParser) //todo wtf?
	err = LoadDeckFromJson(string(ldeck), parsers)
	checkErr(err)
}


func LoadDeckFromJson(jsonMap string, parsers map[string]EntityParser) error {
	
	var deck CardDeck
	parsedMap := []levelMap{}
	if err := json.Unmarshal([]byte(jsonMap), &parsedMap); err != nil {
		return err
	}
	for _, lm := range parsedMap {

		var card Card 
		for i := 0; i < int(data["quantity"].(int)); i++ {
			card = parseCard(lm.Data)
			deck = append(deck, card)
		}
	}
	return nil
}

func parseCard(data map[string]interface{}) *Card{
	return NewCard(
		data["color"].(string),
		data["value"].(string),
		int(data["cost"].(int)),
		data["f"].(string),
	)
}