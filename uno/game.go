package uno

import (
	"encoding/json"
	"io/ioutil"
	"time"
	"math/rand"
	"math"
	"strconv"
	tl "github.com/JoelOtter/termloop"
	"github.com/nsf/termbox-go"
	)

var playersCounter = 4

func checkErr(e error) {
	if e != nil {
		panic(e)
	}
}

type Game struct {
	deck CardDeck
	pile CardDeck
	players []CardDeck
	playerId int
	top *Card
	nextColor string
	route int
	g *tl.Game
	uno bool
	isGet bool
} 

var functions = [...] func(game *Game, card *Card) {
 Basic,
 Stop,
 Reverse,
 StopPlus,
 Color,
 ColorPlus,
}

func NewGame() {
	deck := newCardDeck()
	deck = deck.shuffleDeck()

	players := make([]CardDeck, playersCounter) 

	for index, _ := range players {
		players[index] = deck[0:7]
		deck = deck[8:]
	}
	
	newGame := Game{deck : deck[1:], pile : make(CardDeck, 0), players : players, top : deck[0], playerId : 0,route : 1, uno : false}
	newGame.nextColor = newGame.top.color
	for _, card := range newGame.deck {
		card.SetGame(&newGame)
	}

	for _, player := range newGame.players {
		for _, card := range player {
			card.SetGame(&newGame)
		}
	}
	newGame.top.SetGame(&newGame)
	newGame.playGame()
}

func (game *Game) playGame() {
	g := tl.NewGame()
	game.drawGame(g)
	g.Start()
}

func (game *Game) drawGame(g *tl.Game) {
	l := tl.NewBaseLevel(tl.Cell{Bg: 1,})
	g.Screen().SetLevel(l)

	if game.top.color != game.nextColor {
		g.Screen().AddEntity(tl.NewText(0, 0, "Selected color: " + game.nextColor, 0, 1))
	}

	g.Screen().AddEntity(tl.NewEntityFromCanvas(5, 2, *tl.BackgroundCanvasFromFile(game.top.color + ".png")))
	g.Screen().AddEntity(tl.NewText(6, 4, game.top.value, 1, 0))

	g.Screen().AddEntity(tl.NewEntityFromCanvas(10, 2, *tl.BackgroundCanvasFromFile("black.png")))
	g.Screen().AddEntity(tl.NewText(14, 4, " X " + strconv.Itoa(len(game.deck)), 0, 1))
	g.Screen().AddEntity(tl.NewText(11, 4, "տօ", 0, 1))

	g.Screen().AddEntity(NewButton(20, 2, 5, 1, tl.ColorWhite, "Exit", Exit))
	
	if game.isGet == true {
		g.Screen().AddEntity(NewButton(20, 4, 5, 1, tl.ColorWhite, "Next", game.NextStep))
	} else {
		g.Screen().AddEntity(NewButton(20, 4, 5, 1, tl.ColorWhite, "Get", game.GetCard))
	}
	g.Screen().AddEntity(NewButton(20, 6, 5, 1, tl.ColorWhite, "UNO", game.TalkUno))

	for i, player := range game.players {
		g.Screen().AddEntity(tl.NewEntityFromCanvas(5, 10 + i * 7, *tl.BackgroundCanvasFromFile("black.png")))
		g.Screen().AddEntity(tl.NewText(5, 9 + i * 7, "Player" + strconv.Itoa(i), 0, 1))	
		g.Screen().AddEntity(tl.NewText(6, 12 + i * 7, "տօ", 0, 1))
		if i == game.playerId {
			for j, card := range player {
				g.Screen().AddEntity(NewCardButton(10 + 5 * j, 10 + i * 7, card))
			}
		} else {
			g.Screen().AddEntity(tl.NewText(10, 12 + i * 7, " X " + strconv.Itoa(len(player)), 0, 1))	
		}
	}
}

func (deck CardDeck) shuffleDeck() CardDeck {
	var newDeck CardDeck

	for len(deck) > 0 {
		i := int(math.Abs(float64(int(time.Now().UnixNano()) * rand.Intn(100)))) % len(deck)
		newDeck = append(newDeck, deck[i])
		copy(deck[i:], deck[i+1:])
		deck[len(deck)-1] = nil
		deck = deck[:len(deck)-1]
	}
	return newDeck
}

func newCardDeck() CardDeck {
	var deck CardDeck

	ldeck, err := ioutil.ReadFile("uno/deck.json")
	checkErr(err)

	var f interface{}
    err = json.Unmarshal(ldeck, &f)
    checkErr(err)

    m := f.([]interface{})
	
	for _, v := range m {
		vv := v.(map[string]interface{})
		q := int(vv["quantity"].(float64))
		for i := 0; i < q; i++ {
   			deck = append(deck, parseCard(vv))
   		}
	}
	return deck
}

func parseCard(data map[string]interface{}) *Card{
	f_num := int(data["f"].(float64)) - 1
	return NewCard(
		data["color"].(string),
		data["value"].(string),
		int(data["cost"].(float64)),
		functions[f_num],
	)
}

func (game *Game) Basic(card *Card){
	game.playerId = (len(game.players) + game.playerId + game.route) % len(game.players)
}

func (game *Game) Stop(card *Card){
	game.playerId = (len(game.players) + game.playerId + game.route * 2) % len(game.players)
}

func (game *Game) Reverse(card *Card){
	game.route = -1 * game.route
	game.playerId = (len(game.players) + game.playerId + game.route) % len(game.players)
}

func (game *Game) StopPlus(card *Card){
	nextPlayerID := (len(game.players) + game.playerId + game.route) % len(game.players)
	game.GetCardFromDeck(nextPlayerID, 2)
	game.playerId = (len(game.players) + game.playerId + game.route * 2) % len(game.players)
}

func (game *Game) Color(card *Card){
	game.playerId = (len(game.players) + game.playerId + game.route) % len(game.players)
	g := tl.NewGame()
	l := tl.NewBaseLevel(tl.Cell{Bg: 1,})
	g.Screen().SetLevel(l)
	g.Screen().AddEntity(tl.NewText(0, 0, "What color you choose?", 0, 1))	
	g.Screen().AddEntity(NewColorButton(2, 2, tl.ColorRed, "red", game))	
	g.Screen().AddEntity(NewColorButton(7, 2, tl.ColorBlue, "blue", game))
	g.Screen().AddEntity(NewColorButton(12, 2, tl.ColorGreen, "green", game))
	g.Screen().AddEntity(NewColorButton(17, 2, tl.ColorYellow, "yellow", game))
	g.Start()
}

func (game *Game) ColorPlus(card *Card){
	nextPlayerID := (len(game.players) + game.playerId + game.route) % len(game.players)
	game.GetCardFromDeck(nextPlayerID, 4)
	game.playerId = (len(game.players) + game.playerId + game.route) % len(game.players)
	game.Color(card)
}

func (game *Game) NextStep(){
	game.playerId = (len(game.players) + game.playerId + game.route) % len(game.players)
	game.isGet = false
	game.playGame()
}

func (game *Game) GetCard(){
	game.GetCardFromDeck(game.playerId, 1)
	game.isGet = true
	game.playGame()
}

func (game *Game) TalkUno(){
	if len(game.players[game.playerId]) != 2 && game.uno != true {
		game.GetCardFromDeck(game.playerId, 2)
	}
	game.uno = true
	game.playGame()
}

func (game *Game) GetCardFromDeck(playerId, cardCount int) {
	if len(game.deck) < cardCount {
		prevLen := len(game.deck)
		game.players[playerId] = append(game.players[playerId], game.deck...)
		game.deck = game.pile.shuffleDeck()
		game.pile = make(CardDeck, 0)
		game.GetCardFromDeck(playerId, cardCount - prevLen)
	} else {
		game.players[playerId] = append(game.players[playerId], game.deck[0:cardCount]...)
		game.deck = game.deck[cardCount:]
	}
}

func(game *Game) DoStep(card *Card){
	game.pile = append(game.pile, game.top)
	game.top = card
	game.nextColor = card.color
	for i, ccard := range game.players[game.playerId] {
		if ccard == card {
			game.players[game.playerId] = append(game.players[game.playerId][:i], game.players[game.playerId][i+1:]...)
			break
		}
	}

	if len(game.players[game.playerId]) == 0 {
		game.Winner(game.playerId)
	}

	if len(game.players[game.playerId]) == 1 && game.uno == false {
		game.GetCardFromDeck(game.playerId, 2)
	}

	card.f(game, card)
	game.uno = false
	game.isGet = false
	game.playGame()
}

func (game *Game) Winner(playerId int) {
	termbox.Close()
	g := tl.NewGame()
	l := tl.NewBaseLevel(tl.Cell{Bg: 1,})
	g.Screen().SetLevel(l)
	g.Screen().AddEntity(NewButton(0, 0, 132, 43, 1, "Player"+ strconv.Itoa(playerId) +" Winner", Main))
	g.Start()
}
