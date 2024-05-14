package monopoly

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
	switch t {
	case BOOT:
		return "Boot"
	case CAKE:
		return "Cake"
	case CAR:
		return "Car"
	case CAT:
		return "Cat"
	case DOG:
		return "Dog"
	case DUCK:
		return "Duck"
	case HAT:
		return "Hat"
	case HORSE:
		return "Horse"
	case IRON:
		return "Iron"
	case MONEY_BAG:
		return "Money Bag"
	case PENGUIN:
		return "Penguin"
	case SHIP:
		return "Ship"
	case THIMBLE:
		return "Thimble"
	case TRAIN:
		return "Train"
	case UNICORN:
		return "Unicorn"
	case WEELBARROW:
		return "Weelbarrow"
	default:
		return "UNKNOWN"
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

func (t Token) Description() string {
	switch t {
	case BOOT:
		return "A single boot."
	case CAKE:
		return "A delicious cake."
	case CAR:
		return "A fast car (actually not necessarily faster)."
	case CAT:
		return "A cute cat."
	case DOG:
		return "A not so cute dog."
	case DUCK:
		return "A rubber duck."
	case HAT:
		return "A fancy top hat."
	case HORSE:
		return "A jumping horse with a rider."
	case IRON:
		return "A hot iron."
	case MONEY_BAG:
		return "A bag full of money."
	case PENGUIN:
		return "A clumsy penguin."
	case SHIP:
		return "A heavy battleship."
	case THIMBLE:
		return "A protecting thimble."
	case TRAIN:
		return "A cool steam locomotive."
	case UNICORN:
		return "A rare unicorn."
	case WEELBARROW:
		return "A supporting weelbarrow."
	default:
		return "UNKNOWN."
	}
}
