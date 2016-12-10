package uno

type CardDeck [] *Card

type Card struct {
	color string
	value string
	cost int
	f func(game *Game, card *Card)
	game *Game
}

func Basic(game *Game, card *Card){
	game.Basic(card)
}

func Stop(game *Game, card *Card){
	game.Stop(card)
}

func Reverse(game *Game, card *Card){
	game.Reverse(card)
}

func StopPlus(game *Game, card *Card){
	game.StopPlus(card)
}

func Color(game *Game, card *Card){
	game.Color(card)
}

func ColorPlus(game *Game, card *Card){
	game.ColorPlus(card)
}

func NewCard(color, value string, cost int, f func(game *Game, card *Card)) *Card {
	return &Card{color: color, value: value, cost: cost, f: f, game: nil,}
}

func (c *Card) SetGame(game *Game) {
	c.game = game
}
