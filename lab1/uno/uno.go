package uno

import (
	"fmt"
	"os"
	"github.com/nsf/termbox-go"
	tl "github.com/JoelOtter/termloop"
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

type Image struct {
	e *tl.Entity
}

func NewImage(c *tl.Canvas) *Image {
	i := Image{e: tl.NewEntity(0, 0, len(*c), len((*c)[0]))}
	i.e.ApplyCanvas(c)
	return &i
}

func (i *Image) Draw(s *tl.Screen) { i.e.Draw(s) }

func (i *Image) Tick(ev tl.Event) {	i.e.SetPosition(i.e.Position()) }

func Main() {
	g := tl.NewGame()
	g.Screen().AddEntity(NewImage(tl.BackgroundCanvasFromFile("uno.png")))
	g.Screen().AddEntity(NewButton(10, 20, 112, 7, tl.ColorWhite, "New game", NewGame))
	//g.Screen().AddEntity(uno.NewButton(50, 20, 20, 5, tl.ColorWhite, "Help", uno.Help))
	//g.Screen().AddEntity(uno.NewButton(10, 30, 20, 5, tl.ColorWhite, "Settings", uno.Settings))
	g.Screen().AddEntity(NewButton(10, 30, 112, 7, tl.ColorWhite, "Exit", Exit))
	g.Start()
}
