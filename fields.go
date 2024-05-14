package monopoly

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
	ST_CHARIES_PLACE        Property = iota
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

func (f Field) String() string {
	if p, ok := f.Property(); ok {
		return p.String()
	}
	switch f {
	case GO:
		return "GO"
	case JUST_VISITING:
		return "Just visting"
	case IN_JAIL:
		return "In Jail"
	case FREE_PARKING:
		return "Free Parking"
	case GO_TO_JAIL:
		return "GO to Jail"

	case INCOME_TAX:
		return "Income Tax"
	case LUXERY_TAX:
		return "Luxery Tax"

	case COMMUNITY_CHEST_1, COMMUNITY_CHEST_2, COMMUNITY_CHEST_3:
		return "Community Chest"
	case CHANCE_1, CHANCE_2, CHANCE_3:
		return "Chance"
	default:
		return "UNKNOWN"
	}
}

func (f Field) GoString() string {
	if p, ok := f.Property(); ok {
		return p.GoString()
	}
	switch f {
	case GO:
		return "GO"
	case JUST_VISITING:
		return "JUST_VISITING"
	case IN_JAIL:
		return "IN_JAIL"
	case FREE_PARKING:
		return "FREE_PARKING"
	case GO_TO_JAIL:
		return "GO_TO_JAIL"

	case INCOME_TAX:
		return "INCOME_TAX"
	case LUXERY_TAX:
		return "LUXERY_TAX"

	case COMMUNITY_CHEST_1, COMMUNITY_CHEST_2, COMMUNITY_CHEST_3:
		return "COMMUNITY_CHEST"
	case CHANCE_1, CHANCE_2, CHANCE_3:
		return "CHANCE"
	default:
		return "UNKNOWN"
	}
}

func (f Field) Property() (Property, bool) {
	switch p := Property(f); p {
	case MEDITERRANEAN_AVENUE, BALTIC_AVENUE,
		ORIENTAL_AVENUE, VERMONT_AVENUE, CONNECTICUT_AVENUE,
		ST_CHARIES_PLACE, STATES_AVENUE, VIRGINIA_AVENUE,
		ST_JAMES_PLACE, TENNESSEE_AVENUE, NEW_YORK_AVENUE,
		KENTUCKY_AVENUE, INDIANA_AVENUE, ILLINOIS_AVENUE,
		ATLANTIC_AVENUE, VENTNOR_AVENUE, MARVIN_GARDENS,
		PACIFIC_AVENUE, NORTH_CAROLINA_AVENUE, PENNSYLVANIA_AVENUE,
		PARK_PLACE, BOARDWALK,
		Property(READING_RAILROAD), Property(PENNSYLVANIA_RAILROAD), Property(BALTIMORE_OHIO_RAILROAD), Property(SHORT_LINE),
		Property(ELECTRIC_COMPANY), Property(WATER_WORKS):
		return p, true
	default:
		return -1, false
	}
}

func (p Property) String() string {
	if rr, ok := p.Railroad(); ok {
		return rr.String()
	} else if u, ok := p.Utility(); ok {
		return u.String()
	}

	switch p {
	case MEDITERRANEAN_AVENUE:
		return "Mediterranean Avenue"
	case BALTIC_AVENUE:
		return "Baltic Avenue"
	case ORIENTAL_AVENUE:
		return "Oriental Avenue"
	case VERMONT_AVENUE:
		return "Vermont Avenue"
	case CONNECTICUT_AVENUE:
		return "Connecticut Avenue"
	case ST_CHARIES_PLACE:
		return "St. Charles Place"
	case STATES_AVENUE:
		return "States Avenue"
	case VIRGINIA_AVENUE:
		return "Virginia Avenue"
	case ST_JAMES_PLACE:
		return "St. James Place"
	case TENNESSEE_AVENUE:
		return "Tennessee Avenue"
	case NEW_YORK_AVENUE:
		return "New York Avenue"
	case KENTUCKY_AVENUE:
		return "Kentucky Avenue"
	case INDIANA_AVENUE:
		return "Indiana Avenue"
	case ILLINOIS_AVENUE:
		return "Illinois Avenue"
	case ATLANTIC_AVENUE:
		return "Atlantic Avenue"
	case VENTNOR_AVENUE:
		return "Ventnor Avenue"
	case MARVIN_GARDENS:
		return "Marvin Gardens"
	case PACIFIC_AVENUE:
		return "Pacific Avenue"
	case NORTH_CAROLINA_AVENUE:
		return "North Carolina Avenue"
	case PENNSYLVANIA_AVENUE:
		return "Pennsylvania Avenue"
	case PARK_PLACE:
		return "Park Place"
	case BOARDWALK:
		return "Boardwalk"
	default:
		return "UNKNOWN"
	}
}

func (p Property) GoString() string {
	if rr, ok := p.Railroad(); ok {
		return rr.GoString()
	} else if u, ok := p.Utility(); ok {
		return u.GoString()
	}

	switch p {
	case MEDITERRANEAN_AVENUE:
		return "MEDITERRANEAN_AVENUE"
	case BALTIC_AVENUE:
		return "BALTIC_AVENUE"
	case ORIENTAL_AVENUE:
		return "ORIENTAL_AVENUE"
	case VERMONT_AVENUE:
		return "VERMONT_AVENUE"
	case CONNECTICUT_AVENUE:
		return "CONNECTICUT_AVENUE"
	case ST_CHARIES_PLACE:
		return "ST_CHARIES_PLACE"
	case STATES_AVENUE:
		return "STATES_AVENUE"
	case VIRGINIA_AVENUE:
		return "VIRGINIA_AVENUE"
	case ST_JAMES_PLACE:
		return "ST_JAMES_PLACE"
	case TENNESSEE_AVENUE:
		return "TENNESSEE_AVENUE"
	case NEW_YORK_AVENUE:
		return "NEW_YORK_AVENUE"
	case KENTUCKY_AVENUE:
		return "KENTUCKY_AVENUE"
	case INDIANA_AVENUE:
		return "INDIANA_AVENUE"
	case ILLINOIS_AVENUE:
		return "ILLINOIS_AVENUE"
	case ATLANTIC_AVENUE:
		return "ATLANTIC_AVENUE"
	case VENTNOR_AVENUE:
		return "VENTNOR_AVENUE"
	case MARVIN_GARDENS:
		return "MARVIN_GARDENS"
	case PACIFIC_AVENUE:
		return "PACIFIC_AVENUE"
	case NORTH_CAROLINA_AVENUE:
		return "NORTH_CAROLINA_AVENUE"
	case PENNSYLVANIA_AVENUE:
		return "PENNSYLVANIA_AVENUE"
	case PARK_PLACE:
		return "PARK_PLACE"
	case BOARDWALK:
		return "BOARDWALK"
	default:
		return "UNKNOWN"
	}
}

func (p Property) Railroad() (Railroad, bool) {
	switch rr := Railroad(p); rr {
	case READING_RAILROAD, PENNSYLVANIA_RAILROAD, BALTIMORE_OHIO_RAILROAD, SHORT_LINE:
		return rr, true
	default:
		return -1, false
	}
}

func (p Property) Utility() (Utility, bool) {
	switch u := Utility(p); u {
	case ELECTRIC_COMPANY, WATER_WORKS:
		return u, true
	default:
		return -1, false
	}
}

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

	case ST_CHARIES_PLACE, STATES_AVENUE:
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
	case ST_CHARIES_PLACE, STATES_AVENUE:
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
	case ST_CHARIES_PLACE, STATES_AVENUE, VIRGINIA_AVENUE, ST_JAMES_PLACE, TENNESSEE_AVENUE, NEW_YORK_AVENUE:
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

func (rr Railroad) String() string {
	switch rr {
	case READING_RAILROAD:
		return "Reading Railroad"
	case PENNSYLVANIA_RAILROAD:
		return "Pennsylvania Railroad"
	case BALTIMORE_OHIO_RAILROAD:
		return "Baltimore & Ohio Railroad"
	case SHORT_LINE:
		return "Short Line"
	default:
		return "UNKNOWN"
	}
}

func (rr Railroad) GoString() string {
	switch rr {
	case READING_RAILROAD:
		return "READING_RAILROAD"
	case PENNSYLVANIA_RAILROAD:
		return "PENNSYLVANIA_RAILROAD"
	case BALTIMORE_OHIO_RAILROAD:
		return "BALTIMORE_OHIO_RAILROAD"
	case SHORT_LINE:
		return "SHORT_LINE"
	default:
		return "UNKNOWN"
	}
}

func (u Utility) String() string {
	switch u {
	case ELECTRIC_COMPANY:
		return "Electric Company"
	case WATER_WORKS:
		return "Water Works"
	default:
		return "UNKNOWN"
	}
}

func (u Utility) GoString() string {
	switch u {
	case ELECTRIC_COMPANY:
		return "ELECTRIC_COMPANY"
	case WATER_WORKS:
		return "WATER_WORKS"
	default:
		return "UNKNOWN"
	}
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
	switch ps {
	case STATE_MORTGAGE:
		return "mortgaged"
	case STATE_NORMAL:
		return "without houses"
	case STATE_HOUSE_1:
		return "with 1 house"
	case STATE_HOUSE_2:
		return "with 2 houses"
	case STATE_HOUSE_3:
		return "with 3 houses"
	case STATE_HOUSE_4:
		return "with 4 houses"
	case STATE_HOTEL:
		return "with a hotel"
	default:
		return "UNKNOWN"
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
