package main

import (
	"fmt"

	"github.com/Kesuaheli/monopoly"
	"golang.org/x/text/language"
)

func main() {
	const turns = 100
	g := monopoly.NewGame(monopoly.CAT, monopoly.UNICORN, monopoly.CAR)
	g.SetLanguage(language.German)
	fmt.Printf("\n\nNew game of Monopoly\n%s\n\nsimulating %d random turns...\n\n", g, turns)

	turn := 0
	var oldP *monopoly.Player
	for turn < turns {
		p, state := g.GetCurrentPlayer()
		if p != oldP {
			fmt.Printf("Turn %d/%d:\n%s\n", turn+1, turns, p)
			oldP = p
		}
		fmt.Printf("- %s:\n", state)
		switch state {
		case monopoly.GAME_TURN_START:
			d1, d2, prop := p.RollDice()
			fmt.Printf("  - Rolled [%d] [%d]: would land on %s\n", d1, d2, prop.Localize(g.Language))
		case monopoly.GAME_ROLLED_DICE:
			prop := p.Move()
			buyable, _ := p.CanBuyProperty()
			fmt.Printf("  - Landed on %s! Buyable: %t\n", prop.Localize(g.Language), buyable)
		case monopoly.GAME_MOVED_TO_NEW_FIELD:
			if buyable, prop := p.CanBuyProperty(); buyable {
				p.BuyProperty()
				fmt.Printf("  - Bought %s\n", prop.Localize(g.Language))
			} else if prop != -1 {
				fmt.Printf("  - Didn't buy %s\n", prop.Localize(g.Language))
				p.Continue()
			} else {
				p.Continue()
			}
		case monopoly.GAME_TURN:
			fmt.Print("  - Ended turn\n")
			if again := p.EndTurn(); !again {
				turn++
			}
		}

	}

	fmt.Printf("\n\nEnd of %d turns\n%s\n", turns, g)
}
