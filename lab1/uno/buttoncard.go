package uno

import (
	tl "github.com/JoelOtter/termloop"
)

const width  = 4
const height = 5

type CardButton struct {
	r *tl.Entity
	card *Card
}

func NewCardButton(x, y int, c *Card) *CardButton {
	cb := CardButton{
		r: tl.NewEntity(x, y, width, height), 
		card: c,
	}
	cb.r.ApplyCanvas(tl.BackgroundCanvasFromFile(c.color + ".png"))

	return &cb
}


func (c *CardButton) Draw(s *tl.Screen) {
	x, y := c.r.Position()
	c.r.Draw(s)
	tl.NewText(x + 1, y + 2, c.card.Value, 1, 0).Draw(s)
}

func (c *CardButton) Tick(ev tl.Event) {
	x, y := c.r.Position()
	if ev.Type == tl.EventMouse && ev.Key == tl.MouseLeft && (ev.MouseX >= x && ev.MouseX <= x + height) && (ev.MouseY >= y && ev.MouseY <= y + width) {
		if c.card.color == c.card.game.nextColor || c.card.Value == c.card.game.top.Value || c.card.color == "black" {
			c.card.game.DoStep(c.card)
		}
	}
}