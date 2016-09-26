package termloop

// Provides a general Drawable to be rendered.
type Card struct {
	canvas Canvas
	x      int
	y      int
	width  int
	height int
}

// NewCard creates a new Card, with position (x, y) and size
// (width, height).
// Returns a pointer to the new Card.
func NewCard(x, y, width, height int) *Card {
	canvas := NewCanvas(width, height)
	e := Card{x: x, y: y, width: width, height: height,
		canvas: canvas}
	return &e
}

// NewCardFromCanvas returns a pointer to a new Card, with
// position (x, y) and Canvas c. Width and height are calculated
// using the Canvas.
func NewCardFromCanvas(x, y int, c Canvas) *Card {
	e := Card{
		x:      x,
		y:      y,
		canvas: c,
		width:  len(c),
		height: len(c[0]),
	}
	return &e
}

// Draw draws the Card to its current position on the screen.
// This is usually called every frame.
func (e *Card) Draw(s *Screen) {
	for i := 0; i < e.width; i++ {
		for j := 0; j < e.height; j++ {
			s.RenderCell(e.x+i, e.y+j, &e.canvas[i][j])
		}
	}
}

func (e *Card) Tick(ev Event) {}

// Position returns the (x, y) coordinates of the Card.
func (e *Card) Position() (int, int) {
	return e.x, e.y
}

// Size returns the width and height of the Card, in characters.
func (e *Card) Size() (int, int) {
	return e.width, e.height
}

// SetPosition sets the x and y coordinates of the Card.
func (e *Card) SetPosition(x, y int) {
	e.x = x
	e.y = y
}

// SetCell updates the attribute of the Cell at x, y to match those of c.
// The coordinates are relative to the Card itself, not the Screen.
func (e *Card) SetCell(x, y int, c *Cell) {
	renderCell(&e.canvas[x][y], c)
}

// Fill fills the canvas of the Card with
// a Cell c.
func (e *Card) Fill(c *Cell) {
	for i := range e.canvas {
		for j := range e.canvas[i] {
			renderCell(&e.canvas[i][j], c)
		}
	}
}

// ApplyCanvas takes a pointer to a Canvas, c, and applies this canvas
// over the top of the Card's canvas. Any new values in c will overwrite
// those in the Card.
func (e *Card) ApplyCanvas(c *Canvas) {
	for i := 0; i < min(len(e.canvas), len(*c)); i++ {
		for j := 0; j < min(len(e.canvas[0]), len((*c)[0])); j++ {
			renderCell(&e.canvas[i][j], &(*c)[i][j])
		}
	}
}
