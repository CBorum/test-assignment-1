package monopoly

import (
	"math/rand"
	"time"
)

func init() {
	println("💩")
}

type monopolyPlayer interface {
	getPiece() *piece
}

type player struct {
	piece *piece
	board *board
}

type monopolyPiece interface {
	getLocation() int
	setLocation(int)
}

type piece struct {
	loc int
}

type board struct {
	squares int
}

type dice struct{
	seed int64
}

func (p *player) getPiece() *piece {
	return p.piece
}

func move(piece monopolyPiece, diceVal int) {
	newLoc := piece.getLocation() + diceVal // b) Calculating the new square location
	piece.setLocation(newLoc) // c) Moving the player’s piece from the old location to the new square location
}

func (p *piece) getLocation() int {
	return p.loc
}

func (p *piece) setLocation(newLoc int) {
	p.loc = newLoc
}

func takeTurn(p monopolyPlayer, d *dice) {
	diceVal := d.roll()
	diceVal += d.roll() // a) Calculating a random number total between 2 and 12 (the range of two dice)
	println("dice value:", diceVal)
	mPiece := p.getPiece()
	move(mPiece, diceVal)
	println("new postition", p.getPiece().getLocation())
	println("👌")
}

func (d *dice) roll() int {
	if d.seed == 0 {
		println("seed")
		rand.Seed(time.Now().Unix())
	}
	return 1 + rand.Intn(6)
}
