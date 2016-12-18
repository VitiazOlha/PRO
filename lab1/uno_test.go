package main
import (
	"testing"
	"KPI/PRO/lab1/uno"
	"reflect"
)

func TestCardsQuantity(t *testing.T) {
	deck := uno.NewCardDeck()
	deck = deck.ShuffleDeck()
	if len(deck) != 108 {
		t.Error("Incorrect quantity of cards in deck.")
	} else {
		t.Log("Deck contains 108 cards.")
	}
}

func func_name(t *testing.T, exp uno.Colors, deck uno.CardDeck, card_type string) {
	test := uno.SetCardsQuantities(card_type, deck)
		if reflect.DeepEqual(exp, test) {
			t.Log("Success")
		} else {
			t.Error("Failed")
		}
}

func TestDifferentCardsQuantities(t *testing.T) {
	deck := uno.NewCardDeck()
	deck = deck.ShuffleDeck()

	t.Run("Zeros", func(t *testing.T){
	func_name(t, uno.Colors{1,1,1,1}, deck, "0")
	})

	t.Run("Ones", func(t *testing.T){
	func_name(t, uno.Colors{2,2,2,2}, deck, "1")
	})

	t.Run("Twos", func(t *testing.T){
	func_name(t, uno.Colors{2,2,2,2}, deck, "2")
	})

	t.Run("Threes", func(t *testing.T){
	func_name(t, uno.Colors{2,2,2,2}, deck, "3")
	})

	t.Run("Fours", func(t *testing.T){
	func_name(t, uno.Colors{2,2,2,2}, deck, "4")
	})

	t.Run("Fives", func(t *testing.T){
	func_name(t, uno.Colors{2,2,2,2}, deck, "5")
	})

	t.Run("Sixs", func(t *testing.T){
	func_name(t, uno.Colors{2,2,2,2}, deck, "6")
	})

	t.Run("Sevens", func(t *testing.T){
	func_name(t, uno.Colors{2,2,2,2}, deck, "7")
	})

	t.Run("Eights", func(t *testing.T){
	func_name(t, uno.Colors{2,2,2,2}, deck, "8")
	})

	t.Run("Nines", func(t *testing.T){
	func_name(t, uno.Colors{2,2,2,2}, deck, "9")
	})

	t.Run("Reverses", func(t *testing.T){
	func_name(t, uno.Colors{2,2,2,2}, deck, "⮂")
	})

	t.Run("+2s", func(t *testing.T){
	func_name(t, uno.Colors{2,2,2,2}, deck, "+2")
	})

	t.Run("Skips", func(t *testing.T){
	func_name(t, uno.Colors{2,2,2,2}, deck, "⨂")
	})
	
	t.Run("Wilds", func(t *testing.T) {
	 	count := 0
	 	for i:=0; i < len(deck); i++ {
	 		if deck[i].Value == "" {
	 			count ++
	 		} 
	 	}
	 	if count == 4 {
	 		t.Log("Success")
	 	} else {
	 		t.Error("Failed")
 		}
 	})
 	t.Run("(Wild + 4)s", func(t *testing.T){
	 	count := 0
	 	for i:=0; i < len(deck); i++ {
	 	 	if deck[i].Value == "+4" {
	 	 		count ++
	 	 	}
	 	}
	 	if count == 4 {
	 		t.Log("Success")
	 	} else {
	 		t.Error("Failed")
	 	}
 	})
}
