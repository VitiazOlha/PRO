package main

import (
	"KPI/PRO/uno"
	tl "github.com/JoelOtter/termloop"
)

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


type Button struct {
	r *tl.Rectangle
	text *tl.Text
	f func()
}

func NewButton(x, y, w, h int, col tl.Attr, t string, fn func()) *Button {
	startPosX := x + int((w - len(t)) / 2)
	startPosY := y + int(h / 2)
	return &Button{
		r: tl.NewRectangle(x, y, w, h, col), 
		text: tl.NewText(startPosX, startPosY, t, tl.ColorBlack, tl.ColorWhite),
		f : fn,
	}
}

func (b *Button) Draw(s *tl.Screen) {
	b.r.Draw(s)
	b.text.Draw(s)
}

func (b *Button) Tick(ev tl.Event) {
	x, y := b.r.Position()
	h, w := b.r.Size()
	if ev.Type == tl.EventMouse && ev.Key == tl.MouseLeft && (ev.MouseX > x && ev.MouseX < x + h) && (ev.MouseY > y && ev.MouseY < y + w) {
		b.f()
	}
}

func main() {
	g := tl.NewGame()
/*	g.Screen().EnablePixelMode()*/
	g.Screen().AddEntity(NewImage(tl.BackgroundCanvasFromFile("uno.png")))
	//g.Screen().AddEntity(NewImage(tl.BackgroundCanvasFromFile("cards.png")))
	g.Screen().AddEntity(NewButton(10, 20, 20, 5, tl.ColorWhite, "New game", uno.NewGame))
	g.Screen().AddEntity(NewButton(50, 20, 20, 5, tl.ColorWhite, "Help", uno.Help))
	g.Screen().AddEntity(NewButton(10, 30, 20, 5, tl.ColorWhite, "Settings", uno.Settings))
	g.Screen().AddEntity(NewButton(50, 30, 20, 5, tl.ColorWhite, "Exit", uno.Exit))
	g.Start()
}
