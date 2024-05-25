package monopoly

import (
	"github.com/Kesuaheli/monopoly/lang"
	"golang.org/x/text/language"
)

//go:generate go run genfields.go -type=Field

type Field int8
type Property Field
type Railroad Property
type Utility Property

const (
	GO                      Field    = iota
	MEDITERRANEAN_AVENUE    Property = iota
	COMMUNITY_CHEST_1       Field    = iota
	BALTIC_AVENUE           Property = iota
	INCOME_TAX              Field    = iota
	READING_RAILROAD        Railroad = iota
	ORIENTAL_AVENUE         Property = iota
	CHANCE_1                Field    = iota
	VERMONT_AVENUE          Property = iota
	CONNECTICUT_AVENUE      Property = iota
	JUST_VISITING           Field    = iota
	ST_CHARLES_PLACE        Property = iota
	ELECTRIC_COMPANY        Utility  = iota
	STATES_AVENUE           Property = iota
	VIRGINIA_AVENUE         Property = iota
	PENNSYLVANIA_RAILROAD   Railroad = iota
	ST_JAMES_PLACE          Property = iota
	COMMUNITY_CHEST_2       Field    = iota
	TENNESSEE_AVENUE        Property = iota
	NEW_YORK_AVENUE         Property = iota
	FREE_PARKING            Field    = iota
	KENTUCKY_AVENUE         Property = iota
	CHANCE_2                Field    = iota
	INDIANA_AVENUE          Property = iota
	ILLINOIS_AVENUE         Property = iota
	BALTIMORE_OHIO_RAILROAD Railroad = iota
	ATLANTIC_AVENUE         Property = iota
	VENTNOR_AVENUE          Property = iota
	WATER_WORKS             Utility  = iota
	MARVIN_GARDENS          Property = iota
	GO_TO_JAIL              Field    = iota
	PACIFIC_AVENUE          Property = iota
	NORTH_CAROLINA_AVENUE   Property = iota
	COMMUNITY_CHEST_3       Field    = iota
	PENNSYLVANIA_AVENUE     Property = iota
	SHORT_LINE              Railroad = iota
	CHANCE_3                Field    = iota
	PARK_PLACE              Property = iota
	LUXERY_TAX              Field    = iota
	BOARDWALK               Property = iota

	IN_JAIL Field = iota
)

func (p Property) GetRentCost(ps PropertyState) int {
	if ps == STATE_MORTGAGE {
		return 0
	}

	if _, ok := p.Railroad(); ok {
		return 25
	} else if _, ok := p.Utility(); ok {
		return 0 // must be calculated seperately
	}

	switch p {
	case MEDITERRANEAN_AVENUE:
		return []int{2, 10, 30, 90, 160, 250}[ps]
	case BALTIC_AVENUE:
		return []int{4, 20, 60, 180, 320, 450}[ps]

	case ORIENTAL_AVENUE, VERMONT_AVENUE:
		return []int{6, 30, 90, 270, 400, 550}[ps]
	case CONNECTICUT_AVENUE:
		return []int{8, 40, 100, 300, 450, 600}[ps]

	case ST_CHARLES_PLACE, STATES_AVENUE:
		return []int{10, 50, 150, 450, 625, 750}[ps]
	case VIRGINIA_AVENUE:
		return []int{12, 60, 180, 500, 700, 900}[ps]

	case ST_JAMES_PLACE, TENNESSEE_AVENUE:
		return []int{14, 70, 200, 550, 750, 950}[ps]
	case NEW_YORK_AVENUE:
		return []int{16, 80, 220, 600, 800, 1000}[ps]

	case KENTUCKY_AVENUE, INDIANA_AVENUE:
		return []int{18, 90, 250, 700, 875, 1050}[ps]
	case ILLINOIS_AVENUE:
		return []int{20, 100, 300, 750, 925, 1100}[ps]

	case ATLANTIC_AVENUE, VENTNOR_AVENUE:
		return []int{22, 110, 330, 800, 975, 1150}[ps]
	case MARVIN_GARDENS:
		return []int{24, 120, 360, 850, 1025, 1200}[ps]

	case PACIFIC_AVENUE, NORTH_CAROLINA_AVENUE:
		return []int{26, 130, 390, 900, 1100, 1275}[ps]
	case PENNSYLVANIA_AVENUE:
		return []int{28, 150, 450, 1000, 1200, 1400}[ps]

	case PARK_PLACE:
		return []int{35, 175, 500, 1100, 1300, 1500}[ps]
	case BOARDWALK:
		return []int{50, 200, 600, 1400, 1700, 2000}[ps]

	default:
		return 0
	}
}

func (p Property) GetBaseCost() int {
	if _, ok := p.Railroad(); ok {
		return 200
	} else if _, ok := p.Utility(); ok {
		return 150
	}

	switch p {
	case MEDITERRANEAN_AVENUE, BALTIC_AVENUE:
		return 60
	case ORIENTAL_AVENUE, VERMONT_AVENUE:
		return 100
	case CONNECTICUT_AVENUE:
		return 120
	case ST_CHARLES_PLACE, STATES_AVENUE:
		return 140
	case VIRGINIA_AVENUE:
		return 160
	case ST_JAMES_PLACE, TENNESSEE_AVENUE:
		return 180
	case NEW_YORK_AVENUE:
		return 200
	case KENTUCKY_AVENUE, INDIANA_AVENUE:
		return 220
	case ILLINOIS_AVENUE:
		return 240
	case ATLANTIC_AVENUE, VENTNOR_AVENUE:
		return 160
	case MARVIN_GARDENS:
		return 280
	case PACIFIC_AVENUE, NORTH_CAROLINA_AVENUE:
		return 300
	case PENNSYLVANIA_AVENUE:
		return 320
	case PARK_PLACE:
		return 350
	case BOARDWALK:
		return 400
	default:
		return 0
	}
}

func (p Property) GetHouseCost() int {
	const baseHouseCost = 50
	switch p {
	case MEDITERRANEAN_AVENUE, BALTIC_AVENUE, ORIENTAL_AVENUE, VERMONT_AVENUE, CONNECTICUT_AVENUE:
		return baseHouseCost
	case ST_CHARLES_PLACE, STATES_AVENUE, VIRGINIA_AVENUE, ST_JAMES_PLACE, TENNESSEE_AVENUE, NEW_YORK_AVENUE:
		return baseHouseCost * 2
	case KENTUCKY_AVENUE, INDIANA_AVENUE, ILLINOIS_AVENUE, ATLANTIC_AVENUE, VENTNOR_AVENUE, MARVIN_GARDENS:
		return baseHouseCost * 3
	case PACIFIC_AVENUE, NORTH_CAROLINA_AVENUE, PENNSYLVANIA_AVENUE, PARK_PLACE, BOARDWALK:
		return baseHouseCost * 4
	default:
		return 0
	}
}

func (p Property) GetMortgageValue() int {
	return p.GetBaseCost() / 2
}

type PropertyState int8

const (
	STATE_MORTGAGE PropertyState = iota - 1
	STATE_NORMAL
	STATE_HOUSE_1
	STATE_HOUSE_2
	STATE_HOUSE_3
	STATE_HOUSE_4
	STATE_HOTEL
)

func (ps PropertyState) String() string {
	return ps.Localize(language.English)
}

func (ps PropertyState) Localize(langTag language.Tag) string {
	switch ps {
	case STATE_MORTGAGE:
		return lang.MustLocalize("monopoly.property_state.mortgaged", langTag)
	case STATE_NORMAL:
		return lang.MustLocalize("monopoly.property_state.normal", langTag)
	case STATE_HOUSE_1:
		return lang.MustLocalize("monopoly.property_state.house.1", langTag)
	case STATE_HOUSE_2:
		return lang.MustLocalize("monopoly.property_state.house.2", langTag)
	case STATE_HOUSE_3:
		return lang.MustLocalize("monopoly.property_state.house.3", langTag)
	case STATE_HOUSE_4:
		return lang.MustLocalize("monopoly.property_state.house.4", langTag)
	case STATE_HOTEL:
		return lang.MustLocalize("monopoly.property_state.hotel", langTag)
	default:
		return lang.MustLocalize("unknown", langTag)
	}
}

func (ps PropertyState) GoString() string {
	switch ps {
	case STATE_MORTGAGE:
		return "STATE_MORTGAGE"
	case STATE_NORMAL:
		return "STATE_NORMAL"
	case STATE_HOUSE_1:
		return "STATE_HOUSE_1"
	case STATE_HOUSE_2:
		return "STATE_HOUSE_2"
	case STATE_HOUSE_3:
		return "STATE_HOUSE_3"
	case STATE_HOUSE_4:
		return "STATE_HOUSE_4"
	case STATE_HOTEL:
		return "STATE_HOTEL"
	default:
		return "UNKNOWN"
	}
}
