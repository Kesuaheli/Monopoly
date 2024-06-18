package main

import (
	"fmt"
	"strings"

	"github.com/Kesuaheli/monopoly"
	"github.com/Kesuaheli/monopoly/cli/util"
	"github.com/Kesuaheli/monopoly/lang"
	"golang.org/x/text/language"
)

var (
	selectedLang language.Tag
)

func main() {
	selectedLang, _ = util.SelectableInput(lang.MustLocalize("monopoly.word.language.singular.article.indefinite", selectedLang), lang.AllLangs(), true, func(l language.Tag, i int) bool { util.SelectedLanguage = l; return false })
	fmt.Printf("\n"+lang.MustLocalize("cli.input.choose", selectedLang)+"\n", lang.MustLocalize("monopoly.word.player.plural", selectedLang))
	players := make([]monopoly.Token, util.NumberInput(2, len(monopoly.AllTokens())))
	for i := range players {
		id := fmt.Sprintf("%s %d", util.ToUpperFirst(lang.MustLocalize("monopoly.word.player.singular", selectedLang)), i+1)
		fmt.Printf("\n/%s\\\n| %s |\n\\%s/\n", strings.Repeat("=", len(id)+2), id, strings.Repeat("=", len(id)+2))

		players[i], _ = util.SelectableInput(
			lang.MustLocalize("monopoly.word.token.singular.article.indefinite", selectedLang),
			util.SliceDeleteElements(monopoly.AllTokens(), players[:i]),
			true,
			nil,
		)
	}

	const turns = 0
	g := monopoly.NewGame(players...)
	g.SetLanguage(selectedLang)
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
