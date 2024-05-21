package monopoly

import (
	"github.com/Kesuaheli/monopoly/lang"
	"golang.org/x/text/language"
)

// String returns the english name for f.
// String implements [fmt.Stringer] interface.
func (f Field) String() string {
	return f.Localize(language.English)
}

// Localize returns the localized name for f in the language langTag.
func (f Field) Localize(langTag language.Tag) string {
	if p, ok := f.Property(); ok {
		return p.Localize(langTag)
	}

	switch f {
	case GO:
		return lang.MustLocalize("monopoly.field.go", langTag)
	case COMMUNITY_CHEST_1:
		return lang.MustLocalize("monopoly.field.community_chest_1", langTag)
	case INCOME_TAX:
		return lang.MustLocalize("monopoly.field.income_tax", langTag)
	case CHANCE_1:
		return lang.MustLocalize("monopoly.field.chance_1", langTag)
	case JUST_VISITING:
		return lang.MustLocalize("monopoly.field.just_visiting", langTag)
	case COMMUNITY_CHEST_2:
		return lang.MustLocalize("monopoly.field.community_chest_2", langTag)
	case FREE_PARKING:
		return lang.MustLocalize("monopoly.field.free_parking", langTag)
	case CHANCE_2:
		return lang.MustLocalize("monopoly.field.chance_2", langTag)
	case GO_TO_JAIL:
		return lang.MustLocalize("monopoly.field.go_to_jail", langTag)
	case COMMUNITY_CHEST_3:
		return lang.MustLocalize("monopoly.field.community_chest_3", langTag)
	case CHANCE_3:
		return lang.MustLocalize("monopoly.field.chance_3", langTag)
	case LUXERY_TAX:
		return lang.MustLocalize("monopoly.field.luxery_tax", langTag)
	case IN_JAIL:
		return lang.MustLocalize("monopoly.field.in_jail", langTag)
	default:
		return "UNKNOWN"
	}
}

// GoString implements [fmt.GoStringer] interface.
func (f Field) GoString() string {
	if p, ok := f.Property(); ok {
		return p.GoString()
	}

	switch f {
	case GO:
		return "GO"
	case COMMUNITY_CHEST_1:
		return "COMMUNITY_CHEST_1"
	case INCOME_TAX:
		return "INCOME_TAX"
	case CHANCE_1:
		return "CHANCE_1"
	case JUST_VISITING:
		return "JUST_VISITING"
	case COMMUNITY_CHEST_2:
		return "COMMUNITY_CHEST_2"
	case FREE_PARKING:
		return "FREE_PARKING"
	case CHANCE_2:
		return "CHANCE_2"
	case GO_TO_JAIL:
		return "GO_TO_JAIL"
	case COMMUNITY_CHEST_3:
		return "COMMUNITY_CHEST_3"
	case CHANCE_3:
		return "CHANCE_3"
	case LUXERY_TAX:
		return "LUXERY_TAX"
	case IN_JAIL:
		return "IN_JAIL"
	default:
		return "UNKNOWN"
	}
}

// Property converts a [Field] into a [Property] and reports weather f is a Property.
func (f Field) Property() (Property, bool) {
	switch p := Property(f); p {
	case MEDITERRANEAN_AVENUE,
		BALTIC_AVENUE,
		ORIENTAL_AVENUE,
		VERMONT_AVENUE,
		CONNECTICUT_AVENUE,
		ST_CHARIES_PLACE,
		STATES_AVENUE,
		VIRGINIA_AVENUE,
		ST_JAMES_PLACE,
		TENNESSEE_AVENUE,
		NEW_YORK_AVENUE,
		KENTUCKY_AVENUE,
		INDIANA_AVENUE,
		ILLINOIS_AVENUE,
		ATLANTIC_AVENUE,
		VENTNOR_AVENUE,
		MARVIN_GARDENS,
		PACIFIC_AVENUE,
		NORTH_CAROLINA_AVENUE,
		PENNSYLVANIA_AVENUE,
		PARK_PLACE,
		BOARDWALK,
		Property(READING_RAILROAD),
		Property(PENNSYLVANIA_RAILROAD),
		Property(BALTIMORE_OHIO_RAILROAD),
		Property(SHORT_LINE),
		Property(ELECTRIC_COMPANY),
		Property(WATER_WORKS):
		return p, true
	default:
		return -1, false
	}
}

// String returns the english name for p.
// String implements [fmt.Stringer] interface.
func (p Property) String() string {
	return p.Localize(language.English)
}

// Localize returns the localized name for p in the language langTag.
func (p Property) Localize(langTag language.Tag) string {
	if r, ok := p.Railroad(); ok {
		return r.Localize(langTag)
	}

	if u, ok := p.Utility(); ok {
		return u.Localize(langTag)
	}

	switch p {
	case MEDITERRANEAN_AVENUE:
		return lang.MustLocalize("monopoly.field.mediterranean_avenue", langTag)
	case BALTIC_AVENUE:
		return lang.MustLocalize("monopoly.field.baltic_avenue", langTag)
	case ORIENTAL_AVENUE:
		return lang.MustLocalize("monopoly.field.oriental_avenue", langTag)
	case VERMONT_AVENUE:
		return lang.MustLocalize("monopoly.field.vermont_avenue", langTag)
	case CONNECTICUT_AVENUE:
		return lang.MustLocalize("monopoly.field.connecticut_avenue", langTag)
	case ST_CHARIES_PLACE:
		return lang.MustLocalize("monopoly.field.st_charies_place", langTag)
	case STATES_AVENUE:
		return lang.MustLocalize("monopoly.field.states_avenue", langTag)
	case VIRGINIA_AVENUE:
		return lang.MustLocalize("monopoly.field.virginia_avenue", langTag)
	case ST_JAMES_PLACE:
		return lang.MustLocalize("monopoly.field.st_james_place", langTag)
	case TENNESSEE_AVENUE:
		return lang.MustLocalize("monopoly.field.tennessee_avenue", langTag)
	case NEW_YORK_AVENUE:
		return lang.MustLocalize("monopoly.field.new_york_avenue", langTag)
	case KENTUCKY_AVENUE:
		return lang.MustLocalize("monopoly.field.kentucky_avenue", langTag)
	case INDIANA_AVENUE:
		return lang.MustLocalize("monopoly.field.indiana_avenue", langTag)
	case ILLINOIS_AVENUE:
		return lang.MustLocalize("monopoly.field.illinois_avenue", langTag)
	case ATLANTIC_AVENUE:
		return lang.MustLocalize("monopoly.field.atlantic_avenue", langTag)
	case VENTNOR_AVENUE:
		return lang.MustLocalize("monopoly.field.ventnor_avenue", langTag)
	case MARVIN_GARDENS:
		return lang.MustLocalize("monopoly.field.marvin_gardens", langTag)
	case PACIFIC_AVENUE:
		return lang.MustLocalize("monopoly.field.pacific_avenue", langTag)
	case NORTH_CAROLINA_AVENUE:
		return lang.MustLocalize("monopoly.field.north_carolina_avenue", langTag)
	case PENNSYLVANIA_AVENUE:
		return lang.MustLocalize("monopoly.field.pennsylvania_avenue", langTag)
	case PARK_PLACE:
		return lang.MustLocalize("monopoly.field.park_place", langTag)
	case BOARDWALK:
		return lang.MustLocalize("monopoly.field.boardwalk", langTag)
	default:
		return "UNKNOWN"
	}
}

// GoString implements [fmt.GoStringer] interface.
func (p Property) GoString() string {
	if r, ok := p.Railroad(); ok {
		return r.GoString()
	}

	if u, ok := p.Utility(); ok {
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

// Railroad converts a [Property] into a [Railroad] and reports weather p is a Railroad.
func (p Property) Railroad() (Railroad, bool) {
	switch r := Railroad(p); r {
	case READING_RAILROAD,
		PENNSYLVANIA_RAILROAD,
		BALTIMORE_OHIO_RAILROAD,
		SHORT_LINE:
		return r, true
	default:
		return -1, false
	}
}

// Utility converts a [Property] into a [Utility] and reports weather p is a Utility.
func (p Property) Utility() (Utility, bool) {
	switch u := Utility(p); u {
	case ELECTRIC_COMPANY,
		WATER_WORKS:
		return u, true
	default:
		return -1, false
	}
}

// String returns the english name for r.
// String implements [fmt.Stringer] interface.
func (r Railroad) String() string {
	return r.Localize(language.English)
}

// Localize returns the localized name for r in the language langTag.
func (r Railroad) Localize(langTag language.Tag) string {
	switch r {
	case READING_RAILROAD:
		return lang.MustLocalize("monopoly.field.reading_railroad", langTag)
	case PENNSYLVANIA_RAILROAD:
		return lang.MustLocalize("monopoly.field.pennsylvania_railroad", langTag)
	case BALTIMORE_OHIO_RAILROAD:
		return lang.MustLocalize("monopoly.field.baltimore_ohio_railroad", langTag)
	case SHORT_LINE:
		return lang.MustLocalize("monopoly.field.short_line", langTag)
	default:
		return "UNKNOWN"
	}
}

// GoString implements [fmt.GoStringer] interface.
func (r Railroad) GoString() string {
	switch r {
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

// String returns the english name for u.
// String implements [fmt.Stringer] interface.
func (u Utility) String() string {
	return u.Localize(language.English)
}

// Localize returns the localized name for u in the language langTag.
func (u Utility) Localize(langTag language.Tag) string {
	switch u {
	case ELECTRIC_COMPANY:
		return lang.MustLocalize("monopoly.field.electric_company", langTag)
	case WATER_WORKS:
		return lang.MustLocalize("monopoly.field.water_works", langTag)
	default:
		return "UNKNOWN"
	}
}

// GoString implements [fmt.GoStringer] interface.
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
