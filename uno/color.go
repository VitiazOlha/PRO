package uno

import (
	tl "github.com/JoelOtter/termloop"
)

type ColorButton struct {
	r *tl.Rectangle
	color string
	game *Game
}

func NewColorButton(x, y int, col tl.Attr, color string, g *Game) *ColorButton {
	return &ColorButton{
		r: tl.NewRectangle(x, y, 4, 4, col), 
		color : color,
		game : g,
	}
}

func (b *ColorButton) Draw(s *tl.Screen) {
	b.r.Draw(s)
}

func (b *ColorButton) Tick(ev tl.Event) {
	x, y := b.r.Position()
	h, w := b.r.Size()
	if ev.Type == tl.EventMouse && ev.Key == tl.MouseLeft && (ev.MouseX >= x && ev.MouseX <= x + h) && (ev.MouseY >= y && ev.MouseY <= y + w) {
		b.game.nextColor = b.color
		b.game.playGame()
	}
}