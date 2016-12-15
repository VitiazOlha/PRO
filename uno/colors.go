package uno

type Colors struct{
 	Blues int
 	Reds int
 	Yellows int
 	Greens int
}

func SetCardsQuantities(str string, deck CardDeck) Colors {
	colors := Colors{ Blues: 0, Reds: 0, Yellows: 0, Greens: 0}
	for i:=0; i < len(deck); i++ {
		if deck[i].Value == str {
			switch deck[i].color {
				case "blue":
					colors.Blues ++
 				case "red":
 					colors.Reds ++
 				case "yellow":
 					colors.Yellows ++
 				case "green":
 					colors.Greens ++
 			}	
 		}
 	}
 	return colors
}
