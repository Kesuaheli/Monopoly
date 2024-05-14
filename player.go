package monopoly

import (
	"fmt"
	"strings"
	"sync"
)

const startMoney = 5000

type Inventory map[Property]PropertyState

func (inv Inventory) String() string {
	if len(inv) == 0 {
		return "no properties"
	}
	var props []string
	for prop, state := range inv {
		if state == STATE_NORMAL {
			props = append(props, prop.String())
		} else {
			props = append(props, prop.String()+" "+state.String())
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
	token     Token
	position  Field
	money     int
	inventory Inventory
	invLock   sync.Mutex
}

func InitPlayer(t Token) *Player {
	return &Player{
		token:     t,
		position:  GO,
		money:     startMoney,
		inventory: map[Property]PropertyState{},
	}
}

func (p *Player) String() string {
	return fmt.Sprintf("%s ($%d) is on %s and owns %s.", p.token, p.money, p.position, p.inventory)
}

func (p *Player) GoString() string {
	return fmt.Sprintf("{token: %#v, money: $%d, position: %#v, inventory: %#v}", p.token, p.money, p.position, p.inventory)
}

func (p *Player) GetBalance() int {
	return p.money
}

func (p *Player) BuyProperty(g *Game, prop Property) bool {
	if !g.IsPropertyAvailable(prop) {
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
