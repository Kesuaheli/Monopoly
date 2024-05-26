package monopoly

import (
	"fmt"
	"math/rand"
	"strings"

	"github.com/Kesuaheli/monopoly/lang"
	"golang.org/x/text/language"
)

// Game represents a game of Monopoly.
type Game struct {
	// Language is the language used when printing names and messages
	Language language.Tag

	players     []*Player
	currentTurn int

	lastRoll      uint8 // 2 dice encoded in 2 blocks of 4 bit
	doubblesCount int
	state         GameState
}

// NewGame creates a new game of Monopoly and initializes it with the default state.
func NewGame(players ...Token) *Game {
	if len(players) < 2 {
		return nil
	}

	g := &Game{
		players:     make([]*Player, 0, len(players)),
		currentTurn: rand.Intn(len(players)),
	}
	for _, t := range players {
		g.players = append(g.players, InitPlayer(g, t))
	}
	return g
}

func (g Game) String() string {
	var players []string
	for _, player := range g.players {
		players = append(players, player.String())
	}
	return fmt.Sprintf("%d players\n- %s", len(g.players), strings.Join(players, "\n- "))
}

func (g Game) GoString() string {
	var players []string
	for _, player := range g.players {
		players = append(players, player.GoString())
	}
	return "{players: " + strings.Join(players, ", ") + "}"
}

// FormatCurrency is a helper function to print the given amount of money with the currency symbol
// for the selected language.
func (g Game) FormatCurrency(a int) string {
	return fmt.Sprintf(lang.MustLocalize("monopoly.currency", g.Language), a)
}

// SetLanguage sets the language used when printing names and messages
func (g *Game) SetLanguage(langTag language.Tag) {
	g.Language = langTag
}

func (g Game) GetPlayer(t Token) *Player {
	for _, player := range g.players {
		if player.token == t {
			return player
		}
	}
	return nil
}

func (g Game) GetCurrentPlayer() (*Player, GameState) {
	return g.players[g.currentTurn], g.state
}

func (g Game) GetPlayerForProperty(prop Property) (*Player, PropertyState, bool) {
	for _, player := range g.players {
		player.invLock.Lock()
		defer player.invLock.Unlock()
		for invProp := range player.inventory {
			if invProp == prop {
				return player, player.inventory[prop], true
			}
		}
	}
	return nil, -1, false
}

func (g Game) IsPropertyAvailable(prop Property) bool {
	_, _, isSold := g.GetPlayerForProperty(prop)
	return !isSold
}

func (g *Game) rollDice() (int, int) {
	d1, d2 := RollDice()
	g.lastRoll = uint8(d1<<4 | d2&(1<<4-1))
	return d1, d2
}

func (g *Game) setLastRoll(d1 int, d2 int) {
	g.lastRoll = uint8(d1<<4 | d2&(1<<4-1))
}

func (g Game) getLastRoll() (int, int) {
	return int(g.lastRoll >> 4), int(g.lastRoll & (1<<4 - 1))
}

func RollDice() (int, int) {
	return rand.Intn(6) + 1, rand.Intn(6) + 1
}

func (g *Game) nextPlayer() {
	g.currentTurn++
	if g.currentTurn >= len(g.players) {
		g.currentTurn = 0
	}
}
