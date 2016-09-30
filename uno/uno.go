package uno

import (
	"fmt"
	"time"
	"os"
	"github.com/nsf/termbox-go"
	)
var playersCounter = 2

func NewGame() {
	//deck := newCardDeck()
	players := make([][]Card, playersCounter) 

	for index, _ := range players {
		players[index] = make([]Card, 0, 10)
	}

	time.Sleep(1000 * time.Millisecond)
	fmt.Println("This is new game")
}

func Help() {
	fmt.Println("This is help")
}

func Settings() {
	fmt.Println("This is settings")
}

func Exit() {
	termbox.Close()
	os.Exit(0)
}
