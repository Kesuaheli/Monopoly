package monopoly

import (
	"strconv"
	"testing"
)

func TestGame_setLastRoll(t *testing.T) {
	g := &Game{}
	for d1 := 1; d1 <= 6; d1++ {
		for d2 := 1; d2 <= 6; d2++ {
			g.setLastRoll(d1, d2)
			gotD1, gotD2 := g.getLastRoll()
			if gotD1 != d1 || gotD2 != d2 {
				t.Errorf("Game.getLastRoll() got = (%d, %d), want = (%d, %d), g.lastRoll: 0b%08s", gotD1, gotD2, d1, d2, strconv.FormatUint(uint64(g.lastRoll), 2))
			}
		}
	}
}

func TestRollDice(t *testing.T) {
	const testCount = 50
	for n := 0; n < testCount; n++ {
		d1, d2 := RollDice()
		if d1 < 1 {
			t.Errorf("RollDice() got = (%d, %d), d1 is smaller than 1", d1, d2)
		}
		if d1 > 6 {
			t.Errorf("RollDice() got = (%d, %d), d1 is greater than 6", d1, d2)
		}
		if d2 < 1 {
			t.Errorf("RollDice() got = (%d, %d), d2 is smaller than 1", d1, d2)
		}
		if d2 > 6 {
			t.Errorf("RollDice() got = (%d, %d), d2 is greater than 6", d1, d2)
		}
	}
}
