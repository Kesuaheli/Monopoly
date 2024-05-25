package monopoly

import (
	"github.com/Kesuaheli/monopoly/lang"
	"golang.org/x/text/language"
)

type Token uint8

const (
	BOOT       Token = iota // a single boot
	CAKE                    // a delicious cake
	CAR                     // a fast car (actually not necessarily faster)
	CAT                     // a cute cat
	DOG                     // a not so cute dog
	DUCK                    // a rubber duck
	HAT                     // a fancy top hat
	HORSE                   // a jumping horse with a rider
	IRON                    // a hot iron
	MONEY_BAG               // a bag full of money
	PENGUIN                 // a clumsy penguin
	SHIP                    // a heavy battleship
	THIMBLE                 // a protecting thimble
	TRAIN                   // a cool steam locomotive
	UNICORN                 // a rare unicorn
	WEELBARROW              // a supporting weelbarrow
)

func (t Token) String() string {
	return t.Localize(language.English)
}

func (t Token) Localize(langTag language.Tag) string {
	switch t {
	case BOOT:
		return lang.MustLocalize("monopoly.token.boot", langTag)
	case CAKE:
		return lang.MustLocalize("monopoly.token.cake", langTag)
	case CAR:
		return lang.MustLocalize("monopoly.token.car", langTag)
	case CAT:
		return lang.MustLocalize("monopoly.token.cat", langTag)
	case DOG:
		return lang.MustLocalize("monopoly.token.dog", langTag)
	case DUCK:
		return lang.MustLocalize("monopoly.token.duck", langTag)
	case HAT:
		return lang.MustLocalize("monopoly.token.hat", langTag)
	case HORSE:
		return lang.MustLocalize("monopoly.token.horse", langTag)
	case IRON:
		return lang.MustLocalize("monopoly.token.iron", langTag)
	case MONEY_BAG:
		return lang.MustLocalize("monopoly.token.money_bag", langTag)
	case PENGUIN:
		return lang.MustLocalize("monopoly.token.penguin", langTag)
	case SHIP:
		return lang.MustLocalize("monopoly.token.ship", langTag)
	case THIMBLE:
		return lang.MustLocalize("monopoly.token.thimble", langTag)
	case TRAIN:
		return lang.MustLocalize("monopoly.token.train", langTag)
	case UNICORN:
		return lang.MustLocalize("monopoly.token.unicorn", langTag)
	case WEELBARROW:
		return lang.MustLocalize("monopoly.token.weelbarrow", langTag)
	default:
		return lang.MustLocalize("unknown", langTag)
	}
}

func (t Token) GoString() string {
	switch t {
	case BOOT:
		return "BOOT"
	case CAKE:
		return "CAKE"
	case CAR:
		return "CAR"
	case CAT:
		return "CAT"
	case DOG:
		return "DOG"
	case DUCK:
		return "DUCK"
	case HAT:
		return "HAT"
	case HORSE:
		return "HORSE"
	case IRON:
		return "IRON"
	case MONEY_BAG:
		return "MONEY_BAG"
	case PENGUIN:
		return "PENGUIN"
	case SHIP:
		return "SHIP"
	case THIMBLE:
		return "THIMBLE"
	case TRAIN:
		return "TRAIN"
	case UNICORN:
		return "UNICORN"
	case WEELBARROW:
		return "WEELBARROW"
	default:
		return "UNKNOWN"
	}
}

func (t Token) Description(langTag language.Tag) string {
	switch t {
	case BOOT:
		return lang.MustLocalize("monopoly.token.boot.description", langTag)
	case CAKE:
		return lang.MustLocalize("monopoly.token.cake.description", langTag)
	case CAR:
		return lang.MustLocalize("monopoly.token.car.description", langTag)
	case CAT:
		return lang.MustLocalize("monopoly.token.cat.description", langTag)
	case DOG:
		return lang.MustLocalize("monopoly.token.dog.description", langTag)
	case DUCK:
		return lang.MustLocalize("monopoly.token.duck.description", langTag)
	case HAT:
		return lang.MustLocalize("monopoly.token.hat.description", langTag)
	case HORSE:
		return lang.MustLocalize("monopoly.token.horse.description", langTag)
	case IRON:
		return lang.MustLocalize("monopoly.token.iron.description", langTag)
	case MONEY_BAG:
		return lang.MustLocalize("monopoly.token.money_bag.description", langTag)
	case PENGUIN:
		return lang.MustLocalize("monopoly.token.penguin.description", langTag)
	case SHIP:
		return lang.MustLocalize("monopoly.token.ship.description", langTag)
	case THIMBLE:
		return lang.MustLocalize("monopoly.token.thimble.description", langTag)
	case TRAIN:
		return lang.MustLocalize("monopoly.token.train.description", langTag)
	case UNICORN:
		return lang.MustLocalize("monopoly.token.unicorn.description", langTag)
	case WEELBARROW:
		return lang.MustLocalize("monopoly.token.weelbarrow.description", langTag)
	default:
		return lang.MustLocalize("unknown", langTag)
	}
}
