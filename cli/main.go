package main

import (
	"fmt"
	"math"
	"slices"
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
	selectedLang = selectableInput(lang.MustLocalize("monopoly.word.language.singular.article.indefinite", selectedLang), lang.AllLangs(), func(l language.Tag) { selectedLang = l })
	fmt.Printf("\n"+lang.MustLocalize("cli.input.choose", selectedLang)+"\n", lang.MustLocalize("monopoly.word.player.plural", selectedLang))
	players := make([]monopoly.Token, numberInput(2, len(monopoly.AllTokens())))
	for i := range players {
		id := fmt.Sprintf("%s %d", ToUpperFirst(lang.MustLocalize("monopoly.word.player.singular", selectedLang)), i+1)
		fmt.Printf("\n/%s\\\n| %s |\n\\%s/\n", strings.Repeat("=", len(id)+2), id, strings.Repeat("=", len(id)+2))

		players[i] = selectableInput(lang.MustLocalize("monopoly.word.token.singular.article.indefinite", selectedLang), SliceDeleteElements(monopoly.AllTokens(), players[:i]), nil)
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

	selected := all[numberInput(1, len(all))-1]
	if afterSelection != nil {
		afterSelection(selected)
	}
	fmt.Printf(lang.MustLocalize("cli.input.selected", selectedLang)+"\n", fmt.Sprintf("%s", lang.LocalizeInterface(selected, selectedLang)))
	return selected
}

func numberInput(min, max int) (selected int) {
	for {
		fmt.Printf(lang.MustLocalize("cli.input.select", selectedLang), min, max)
		var bScan []byte
		fmt.Scan(&bScan)
		n, err := strconv.ParseInt(string(bScan), 10, 0)

		if err == nil && n >= int64(min) && n <= int64(max) {
			return int(n)
		}
	}
}

// SliceDeleteElements removes every element from base that is in delete.
//
// When SliceDeleteElements removes m elements, it might not modify the elements
// base[len(base)-m:len(base)]. If those elements contain pointers you might consider zeroing those
// elements so that objects they reference can be garbage collected.
func SliceDeleteElements[S ~[]E, E comparable](base, delete S) S {
	return slices.DeleteFunc(base, func(v E) bool {
		return slices.Contains(delete, v)
	})
}

// ToUpperFirst returns s with only the first character mapped to their upper case.
func ToUpperFirst(s string) string {
	return strings.ToUpper(s[0:1]) + s[1:]
}
