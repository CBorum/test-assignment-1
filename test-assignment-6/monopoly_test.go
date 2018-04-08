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

func TestMockRoll(t *testing.T) {
	d := &mockedDice{}
	d.On("roll").Return(123)
	d.roll()
	d.roll()
	d.AssertNumberOfCalls(t, "roll", 2)
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

func TestTakeTurn2(t *testing.T) {
	p := &player{
		piece: &piece{
			loc: 1,
		},
	}
	d := &mockedDice{}
	// mocked responses -> piece location will always be the same
	d.On("roll").Return(2).Once()
	d.On("roll").Return(4).Once()
	takeTurn(p, d)
	loc := p.getPiece().getLocation()
	assert.Equal(t, 7, loc)
}

func TestTakeTurn3(t *testing.T) {
	p := &mockedPlayer{}
	p.On("getPiece").Return(&piece{123})
	d := &mockedDice{}
	// mocked responses -> piece location will always be the same
	d.On("roll").Return(2).Once()
	d.On("roll").Return(4).Once()
	takeTurn(p, d)
	loc := p.getPiece().getLocation()
	assert.Equal(t, 129, loc)
}
