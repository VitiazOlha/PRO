package uno

import (
	"fmt"
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

func NewGame() {

	deck := newCardDeck()
	deck = deck.shuffleDeck()

	//out only for test
	for _, card := range deck {
		card.Print()
	}

	fmt.Println(len(deck))

	players := make([]CardDeck, playersCounter) 

	for index, _ := range players {
		players[index] = make(CardDeck, 0, 10)
	}



	fmt.Println("This is new game")
}


func (deck CardDeck) shuffleDeck() CardDeck {
	var newDeck CardDeck

	for len(deck) > 0 {
		i := ((int(time.Now().UnixNano()) * rand.Intn(100)) % len(deck))
		fmt.Println(i)
		i = int(math.Abs(float64(i)))
//		time.Sleep(1000 * time.Millisecond)


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
	return NewCard(
		data["color"].(string),
		data["value"].(string),
		int(data["cost"].(float64)),
		data["f"].(string),
	)
}