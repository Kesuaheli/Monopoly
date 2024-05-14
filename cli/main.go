package main

import (
	"fmt"

	"github.com/Kesuaheli/monopoly"
)

func main() {
	g := monopoly.NewGame(monopoly.CAT, monopoly.UNICORN)

	p1 := g.GetPlayer(monopoly.CAT)
	p2 := g.GetPlayer(monopoly.UNICORN)

	p1.BuyProperty(g, monopoly.Property(monopoly.READING_RAILROAD))
	p1.BuyHouse(monopoly.Property(monopoly.READING_RAILROAD))
	p1.BuyHouse(monopoly.Property(monopoly.READING_RAILROAD))
	p1.BuyHouse(monopoly.Property(monopoly.READING_RAILROAD))

	p2.BuyProperty(g, monopoly.BOARDWALK)
	p2.BuyProperty(g, monopoly.PENNSYLVANIA_AVENUE)
	p2.BuyHouse(monopoly.PENNSYLVANIA_AVENUE)
	p2.BuyHouse(monopoly.PENNSYLVANIA_AVENUE)
	p2.BuyHouse(monopoly.PENNSYLVANIA_AVENUE)
	p2.BuyHouse(monopoly.PENNSYLVANIA_AVENUE)
	p2.BuyHouse(monopoly.PENNSYLVANIA_AVENUE)

	fmt.Printf("Monopoly\n--------\n\nprinting using %%s:\n%s\n\nprintign using %%#v:\n%#v\n", g, g)
}
