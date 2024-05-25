package monopoly

import (
	"fmt"
	"strings"
	"sync"

	"github.com/Kesuaheli/monopoly/lang"
	"golang.org/x/text/language"
)

const startMoney = 5000

type Inventory map[Property]PropertyState

func (inv Inventory) String() string {
	return inv.Localize(language.English)
}

func (inv Inventory) Localize(langTag language.Tag) string {
	if len(inv) == 0 {
		return "no properties"
	}
	var props []string
	for prop, state := range inv {
		if state == STATE_NORMAL {
			props = append(props, prop.Localize(langTag))
		} else {
			props = append(props, prop.Localize(langTag)+" "+state.Localize(langTag))
		}
	}
	return strings.Join(props, ", ")
}

func (inv Inventory) GoString() string {
	if len(inv) == 0 {
		return "[]"
	}
	var props []string
	for prop, state := range inv {
		if state == STATE_NORMAL {
			props = append(props, prop.GoString())
		} else {
			props = append(props, prop.GoString()+":"+state.GoString())
		}
	}
	return "[" + strings.Join(props, ", ") + "]"
}

type Player struct {
	game         *Game
	token        Token
	position     Field
	money        int
	inventory    Inventory
	invLock      sync.Mutex
	roundsInJail int
}

func InitPlayer(g *Game, t Token) *Player {
	return &Player{
		game:      g,
		token:     t,
		position:  GO,
		money:     startMoney,
		inventory: map[Property]PropertyState{},
	}
}

func (p *Player) String() string {
	return fmt.Sprintf("%s (%s) is on %s and owns %s.", p.token.Localize(p.game.Language), p.GetBalanceString(), p.position.Localize(p.game.Language), p.inventory.Localize(p.game.Language))
}

func (p *Player) GoString() string {
	return fmt.Sprintf("{token: %#v, money: %d, position: %#v, inventory: %#v}", p.token, p.money, p.position, p.inventory)
}

func (p *Player) GetBalance() int {
	return p.money
}

func (p *Player) GetBalanceString() string {
	return fmt.Sprintf(lang.MustLocalize("monopoly.currency", p.game.Language), p.money)
}

func (p *Player) CanBuyProperty() (bool, Property) {
	prop, ok := p.position.Property()
	if !ok {
		return false, -1
	}
	if p.money < prop.GetBaseCost() {
		return false, -1
	}

	return p.game.IsPropertyAvailable(prop), prop
}

func (p *Player) BuyProperty() bool {
	canBuy, prop := p.CanBuyProperty()
	if !canBuy {
		return false
	}
	p.money -= prop.GetBaseCost()

	p.invLock.Lock()
	defer p.invLock.Unlock()
	p.inventory[prop] = STATE_NORMAL
	return true
}

func (p *Player) TransferProperty(toPlayer *Player, prop Property, money int) bool {
	p.invLock.Lock()
	defer p.invLock.Unlock()
	if _, hasProp := p.inventory[prop]; !hasProp || toPlayer.money < money {
		return false
	}

	p.money += money
	toPlayer.money -= money

	toPlayer.invLock.Lock()
	defer toPlayer.invLock.Unlock()
	toPlayer.inventory[prop] = p.inventory[prop]
	delete(p.inventory, prop)
	return true
}

func (p *Player) MortgageProperty(prop Property) bool {
	p.invLock.Lock()
	defer p.invLock.Unlock()
	if state, hasProp := p.inventory[prop]; !hasProp || state != STATE_NORMAL {
		return false
	}
	p.money += prop.GetMortgageValue()
	p.inventory[prop] = STATE_MORTGAGE
	return true
}

func (p *Player) CancelMortgageProperty(prop Property) bool {
	cost := int(float32(prop.GetMortgageValue()) * 1.1)
	p.invLock.Lock()
	defer p.invLock.Unlock()
	if state, hasProp := p.inventory[prop]; !hasProp || state != STATE_MORTGAGE || p.money < cost {
		return false
	}

	p.money -= cost
	p.inventory[prop] = STATE_NORMAL
	return true
}

func (p *Player) CanBuildHouse(prop Property) bool {
	p.invLock.Lock()
	defer p.invLock.Unlock()
	state, hasProp := p.inventory[prop]
	return hasProp && (state == STATE_NORMAL || state == STATE_HOUSE_1 || state == STATE_HOUSE_2 || state == STATE_HOUSE_3 || state == STATE_HOUSE_4)
}

func (p *Player) CanBuyHouse(prop Property) bool {
	return p.CanBuildHouse(prop) && p.money >= prop.GetHouseCost()
}

func (p *Player) BuyHouse(prop Property) (PropertyState, bool) {
	if !p.CanBuyHouse(prop) {
		return -1, false
	}

	p.money -= prop.GetHouseCost()
	p.invLock.Lock()
	defer p.invLock.Unlock()
	p.inventory[prop] += 1

	return p.inventory[prop], true
}

func (p *Player) CanSellHouse(prop Property) bool {
	p.invLock.Lock()
	defer p.invLock.Unlock()
	state, hasProp := p.inventory[prop]
	return hasProp && (state == STATE_HOUSE_1 || state == STATE_HOUSE_2 || state == STATE_HOUSE_3 || state == STATE_HOUSE_4 || state == STATE_HOTEL)
}

func (p *Player) SellHouse(prop Property) (PropertyState, bool) {
	if !p.CanSellHouse(prop) {
		return -1, false
	}

	p.money += prop.GetHouseCost() / 2
	p.invLock.Lock()
	defer p.invLock.Unlock()
	p.inventory[prop] -= 1

	return p.inventory[prop], true
}

func (p *Player) RollDice() (int, int, Field) {
	if p.game.state != GAME_TURN_START {
		return -1, -1, -1
	}
	if curr, _ := p.game.GetCurrentPlayer(); curr != p {
		return -1, -1, -1
	}

	d1, d2 := p.game.rollDice()
	p.game.state = GAME_ROLLED_DICE

	if d1 == d2 {
		p.game.doubblesCount++
		if p.game.doubblesCount == 4 {
			return d1, d2, IN_JAIL
		}
	} else {
		p.game.doubblesCount = 0
	}

	return d1, d2, (p.position + Field(d1+d2)) % IN_JAIL
}

func (p *Player) Move() Field {
	if p.game.state != GAME_ROLLED_DICE {
		return -1
	}
	if curr, _ := p.game.GetCurrentPlayer(); curr != p {
		return -1
	}

	d1, d2 := p.game.getLastRoll()
	p.game.state = GAME_TURN
	if d1 == d2 {
		if p.game.doubblesCount == 4 {
			p.position = IN_JAIL
			p.game.doubblesCount = 0
			return IN_JAIL
		}
		if p.position == IN_JAIL {
			p.position = JUST_VISITING
		}
	}

	p.position = p.position + Field(d1+d2)
	if p.position >= IN_JAIL {
		fmt.Printf("Player %s crossed %s\n", p.token.Localize(p.game.Language), GO.Localize(p.game.Language))
		p.position = p.position % IN_JAIL
	}
	return p.position
}

func (p *Player) EndTurn() (again bool) {
	if p.game.state != GAME_TURN {
		return false
	}
	if curr, _ := p.game.GetCurrentPlayer(); curr != p {
		return false
	}

	if p.game.doubblesCount == 0 {
		p.game.nextPlayer()
		again = false
	} else {
		again = true
	}
	p.game.state = GAME_TURN_START
	return again
}
