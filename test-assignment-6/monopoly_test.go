package monopoly

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRoll(t *testing.T) {
	d := &dice{}
	for i := 0; i < 100; i++ {
		val := d.roll()
		// println(val)
		assert.True(t, val > 0)
		assert.True(t, val < 7)
	}
}

func TestTakeTurn(t *testing.T) {
	p := &player{
		piece: &piece{
			loc: 1,
		},
	}
	d := &dice{
		seed: 1, // Using "fakes", same seed will always produce the same values
	}
	takeTurn(p, d)
}
