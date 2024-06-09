package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/Kesuaheli/monopoly"
	"github.com/Kesuaheli/monopoly/lang"
	"golang.org/x/text/language"
)

var (
	selectedLang language.Tag
)

func main() {
	selectedLang = selectableInput(lang.MustLocalize("monopoly.word.language.singular", selectedLang), lang.AllLangs(), func(l language.Tag) { selectedLang = l })
	token := selectableInput(lang.MustLocalize("monopoly.word.token.singular", selectedLang), monopoly.AllTokens(), nil)
	fmt.Println(token.Description(selectedLang))

	const turns = 100
	g := monopoly.NewGame(monopoly.CAT, monopoly.UNICORN, monopoly.CAR)
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

func selectableInput[T any](head string, all []T, afterSelection func(selected T)) T {
	if len(all) == 0 {
		var t T
		return t
	} else if len(all) == 1 {
		return all[0]
	}

	selection := strings.Builder{}
	selection.WriteString(fmt.Sprintf(lang.MustLocalize("cli.input.choose", selectedLang), head))
	selection.WriteByte('\n')
	digitsInAll := int(math.Floor(math.Log10(float64(len(all))))) + 1
	for n, item := range all {
		selection.WriteString(fmt.Sprintf("- [%*d] %s\n", digitsInAll, n+1, lang.LocalizeInterface(&item, selectedLang)))
	}
	fmt.Printf(selection.String())

	var selected T
	for {
		fmt.Printf(lang.MustLocalize("cli.input.select", selectedLang), len(all))
		var bScan []byte
		fmt.Scan(&bScan)
		n, err := strconv.ParseInt(string(bScan), 10, 0)

		if err == nil && n > 0 && n <= int64(len(all)) {
			selected = all[n-1]
			break
		}
	}
	if afterSelection != nil {
		afterSelection(selected)
	}
	fmt.Printf(lang.MustLocalize("cli.input.selected", selectedLang)+"\n", fmt.Sprintf("%s", lang.LocalizeInterface(selected, selectedLang)))
	return selected
}
