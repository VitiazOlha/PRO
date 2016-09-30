package uno

import (
	"fmt"
	"os"
	"github.com/nsf/termbox-go"
	)


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
