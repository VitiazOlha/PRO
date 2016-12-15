package uno

import (
	"testing"
	"fmt"
	// "reflect"
)

// type Colors struct{
// 	blues int
// 	reds int
// 	yellows int
// 	greens int
// }


// func SetCardsQuantities(str string, deck CardDeck) *Colors {
// 	colors := Colors{ blues: 0, reds: 0, yellows: 0, greens: 0}
// 	for i:=0; i < len(deck); i++ {
// 		if deck[i].value == str {
// 			switch deck[i].color {
// 				case "blue":
// 					colors.blues ++
// 				case "red":
// 					colors.reds ++
// 				case "yellow":
// 					colors.yellows ++
// 				case "green":
// 					colors.greens ++
// 			}	
// 		}
// 	}

// 	return &colors
// }
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



// func TestCardsQuantity(t *testing.T) {
// 	deck := newCardDeck()
// 	if len(deck) != 108 {
// 		t.Error("Incorrect quantity of cards in deck.")
// 	} else {
// 		t.Log("Deck contains 108 cards.")
// 	}
// }

// func TestDifferentCardsQuantities(t *testing.T) {
// 	deck := newCardDeck()
// 	t.Run("Zeros", func(t *testing.T){
// 		expZeros := Colors{1,1,1,1}
// 		zeros := SetCardsQuantities("0", deck)

// 		if reflect.DeepEqual(expZeros, zeros) {
// 			t.Log("Success")
// 		} else {
// 			t.Error("Failed")
// 		}

// 		})
// 	t.Run("Ones", func(t *testing.T){
// 		expOnes := Colors{2,2,2,2}
// 		ones := SetCardsQuantities("1", deck)

// 		if reflect.DeepEqual(expZeros, zeros) {
// 			t.Log("Success")
// 		} else {
// 			t.Error("Failed")
// 		}

// 		})
// 	t.Run("Twos", func(t *testing.T){
// 		expTwos := Colors{2,2,2,2}
// 		twos := SetCardsQuantities("2", deck)

// 		if reflect.DeepEqual(expZeros, zeros) {
// 			t.Log("Success")
// 		} else {
// 			t.Error("Failed")
// 		}

// 		})
// 	t.Run("Threes", func(t *testing.T){
// 		expThrees := Colors{2,2,2,2}
// 		threes := SetCardsQuantities("3", deck)

// 		if reflect.DeepEqual(expZeros, zeros) {
// 			t.Log("Success")
// 		} else {
// 			t.Error("Failed")
// 		}

// 		})
// 	t.Run("Fours", func(t *testing.T){
// 		expFours := Colors{2,2,2,2}
// 		fours := SetCardsQuantities("4", deck)

// 		if reflect.DeepEqual(expZeros, zeros) {
// 			t.Log("Success")
// 		} else {
// 			t.Error("Failed")
// 		}

// 		})
// 	t.Run("Fives", func(t *testing.T){
// 		expFives := Colors{2,2,2,2}
// 		fives := SetCardsQuantities("5", deck)

// 		if reflect.DeepEqual(expZeros, zeros) {
// 			t.Log("Success")
// 		} else {
// 			t.Error("Failed")
// 		}

// 		})
// 	t.Run("Sixs", func(t *testing.T){
// 		expSixs := Colors{2,2,2,2}
// 		sixs := SetCardsQuantities("6", deck)

// 		if reflect.DeepEqual(expZeros, zeros) {
// 			t.Log("Success")
// 		} else {
// 			t.Error("Failed")
// 		}

// 		})
// 	t.Run("Sevens", func(t *testing.T){
// 		expSevens := Colors{2,2,2,2}
// 		sevens := SetCardsQuantities("7", deck)

// 		if reflect.DeepEqual(expZeros, zeros) {
// 			t.Log("Success")
// 		} else {
// 			t.Error("Failed")
// 		}

// 		})
// 	t.Run("Eights", func(t *testing.T){
// 		expEights := Colors{2,2,2,2}
// 		eights := SetCardsQuantities("8", deck)

// 		if reflect.DeepEqual(expZeros, zeros) {
// 			t.Log("Success")
// 		} else {
// 			t.Error("Failed")
// 		}

// 		})
// 	t.Run("Nines", func(t *testing.T){
// 		expNines := Colors{2,2,2,2}
// 		nines := SetCardsQuantities("9", deck)

// 		if reflect.DeepEqual(expZeros, zeros) {
// 			t.Log("Success")
// 		} else {
// 			t.Error("Failed")
// 		}

// 		})
// 	t.Run("Reverses", func(t *testing.T){
// 		expReverses := Colors{2,2,2,2}
// 		reverses := SetCardsQuantities("⮂", deck)

// 		if reflect.DeepEqual(expZeros, zeros) {
// 			t.Log("Success")
// 		} else {
// 			t.Error("Failed")
// 		}

// 		})
// 	t.Run("+2s", func(t *testing.T){
// 		expDraw2s := Colors{2,2,2,2}
// 		draw2s := SetCardsQuantities("+2", deck)

// 		if reflect.DeepEqual(expZeros, zeros) {
// 			t.Log("Success")
// 		} else {
// 			t.Error("Failed")
// 		}

// 		})
// 	t.Run("Skips", func(t *testing.T){
// 		expSkips := Colors{2,2,2,2}
// 		skips := SetCardsQuantities("⨂", deck)

// 		if reflect.DeepEqual(expZeros, zeros) {
// 			t.Log("Success")
// 		} else {
// 			t.Error("Failed")
// 		}
		
// 		})

	// t.Run("Wilds", func(t *testing.T){
	// 	count := 0
	// 	for i:=0; i < len(deck); i++ {
	// 		if deck[i].value == "" {
	// 			count ++
	// 		} 
	// 	}
	// 	if count == 4 {
	// 		t.Log("Success")
	// 	} else {
	// 		t.Error("Failed")
	// 	}
	// 	)
	// t.Run("(Wild + 4)s", func(t *testing.T){
	// 	count := 0
	// 	for i:=0; i < len(deck); i++ {
	// 	 	if deck[i].value == "+4" {
	// 	 		count ++
	// 	 	}
	// 	}

	// 	if count == 4 {
	// 		t.Log("Success")
	// 	} else {
	// 		t.Error("Failed")
	// 	}
	// 	)



// }
