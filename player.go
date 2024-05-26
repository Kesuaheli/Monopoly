package monopoly

import (
	"fmt"
	"strings"
	"sync"

	"golang.org/x/text/language"
)

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
	return fmt.Sprintf("%s (%s) is on %s and owns %s.", p.Token(), p.game.FormatCurrency(p.money), p.position.Localize(p.game.Language), p.inventory.Localize(p.game.Language))
}

func (p *Player) GoString() string {
	return fmt.Sprintf("{token: %#v, money: %d, position: %#v, inventory: %#v}", p.token, p.money, p.position, p.inventory)
}

// Money returns the amount of money the player currently has, formatted with the currency symbol of
// the selected language.
func (p *Player) Money() string {
	return p.game.FormatCurrency(p.money)
}

// Token returns the localized name of the token the player has.
func (p *Player) Token() string {
	return p.token.Localize(p.game.Language)
}

// Railroads returns the amount of railroads the player owns.
func (p *Player) Railroads() int {
	var rrCount int
	for prop := range p.inventory {
		if _, isRR := prop.Railroad(); isRR {
			rrCount++
		}
	}
	return rrCount
}

// Utilities returns the amount of utilities the player owns.
func (p *Player) Utilities() int {
	var utilCount int
	for prop := range p.inventory {
		if _, isUtil := prop.Utility(); isUtil {
			utilCount++
		}
	}
	return utilCount
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
		if p.game.doubblesCount == doubblesCountToJail {
			return d1, d2, IN_JAIL
		}
	} else {
		p.game.doubblesCount = 0
	}

	return d1, d2, Field((int(p.position) + d1 + d2) % numberOfFields)
}

func (p *Player) Move() Field {
	if p.game.state != GAME_ROLLED_DICE {
		return -1
	}
	if curr, _ := p.game.GetCurrentPlayer(); curr != p {
		return -1
	}

	d1, d2 := p.game.getLastRoll()
	p.game.state = GAME_MOVED_TO_NEW_FIELD
	if d1 == d2 {
		if p.game.doubblesCount == doubblesCountToJail {
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
		fmt.Printf("Player %s crossed %s\n", p.Token(), GO.Localize(p.game.Language))
		p.position %= Field(numberOfFields)
		p.money += moneyOnGo
	}

	if prop, isProp := p.position.Property(); isProp {
		propOwner, propState, ok := p.game.GetPlayerForProperty(prop)
		if !ok || propOwner == p {
			return p.position
		}
		rent := prop.GetRentCost(propState)
		if _, isRR := prop.Railroad(); isRR {
			rent *= propOwner.Railroads()
		} else if _, isUtil := prop.Utility(); isUtil {
			rent = (propOwner.Utilities()*6 - 2) * (d1 + d2)
		}
		fmt.Printf("Player %s owes player %s %s for landing on %s\n",
			p.Token(),
			propOwner.Token(),
			p.game.FormatCurrency(rent),
			prop.Localize(p.game.Language),
		)
		p.money -= rent
		propOwner.money += rent
	} else {
		switch p.position {
		case INCOME_TAX:
			p.money -= incomeTax
		case LUXERY_TAX:
			p.money -= luxeryTax
		case CHANCE_1, CHANCE_2, CHANCE_3:
			fmt.Println("Drawing Chance card...")
		case COMMUNITY_CHEST_1, COMMUNITY_CHEST_2, COMMUNITY_CHEST_3:
			fmt.Println("Drawing Community Chest card...")
		case FREE_PARKING:
			p.money += moneyOnFreeParking
		case GO_TO_JAIL:
			fmt.Printf("Player %s went to jail\n", p.Token())
			p.position = IN_JAIL
		}
	}

	return p.position
}

// Continue advances the players current turn to the next state.
func (p *Player) Continue() {
	if curr, _ := p.game.GetCurrentPlayer(); curr != p {
		return
	}

	switch p.game.state {
	case GAME_MOVED_TO_NEW_FIELD:
		p.game.state = GAME_TURN
	}
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
