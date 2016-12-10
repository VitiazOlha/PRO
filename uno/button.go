package uno

import (
	tl "github.com/JoelOtter/termloop"
)

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
	if ev.Type == tl.EventMouse && ev.Key == tl.MouseLeft && (ev.MouseX >= x && ev.MouseX <= x + h) && (ev.MouseY >= y && ev.MouseY <= y + w) {
		b.f()
	}
}