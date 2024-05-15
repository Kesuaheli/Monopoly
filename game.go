package monopoly

import (
	"fmt"
	"math/rand"
	"strings"
)

// Game represents a game of Monopoly.
type Game struct {
	players     []*Player
	currentTurn int
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
		g.players = append(g.players, InitPlayer(t))
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

func (g Game) GetPlayer(t Token) *Player {
	for _, player := range g.players {
		if player.token == t {
			return player
		}
	}
	return nil
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
