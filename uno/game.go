package uno

import (
	"fmt"
	"encoding/json"
	"io/ioutil"
	)


var playersCounter = 2

func NewGame() {

	deck := newCardDeck()
	players := make([][]Card, playersCounter) 

	for index, _ := range players {
		players[index] = make([]Card, 0, 10)
	}

	fmt.Println("This is new game")
}


func newCardDeck() {
	
	ldeck, err := ioutil.ReadFile("deck.json")
	checkErr(err)
	parsers := make(map[string]tl.EntityParser)
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
		data["color"].(string)),
		data["value"].(string)),
		int(data["cost"].(int)),
		data["f"].(string)),
	)
}


func parseText(data map[string]interface{}) *Text {
	return NewText(
		int(data["x"].(float64)),
		int(data["y"].(float64)),
		data["text"].(string),
		Attr(data["fg"].(float64)),
		Attr(data["bg"].(float64)),
	)
}


