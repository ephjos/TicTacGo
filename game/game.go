package game

import (
	"fmt"
	"github.com/josephthomashines/TicTacGo/board"
)

// Errors
/*
type BoardOutOfBoundsError struct {
	received, limit int
	axis            string
}

func (e *BoardOutOfBoundsError) Error() string {
	out := fmt.Sprintf("Error: Index %d must be %s to %d on the %s axis", e.received, "%s", e.limit, e.axis)

	if e.received > e.limit {
		out = fmt.Sprintf(out, "<=")
	}

	if e.received < e.limit {
		out = fmt.Sprintf(out, ">=")
	}

	return out

}
*/

// Core
type Game struct {
	board board.Board
	turn  string
	move  int
}

func (g *Game) String() string {
	out := "\n"

	out += fmt.Sprintf("%s's Turn\n%s\n", g.turn, &g.board)

	return out
}

func Init() Game {
	return Game{board.Init(3, 3), "X", 0}
}

// Read console input, let user play
func Play() {
	game := Init()
	gp := &game

	fmt.Println(gp)
}

// Play n games randomly
// Report win rate of each player
// Average number of moves
func Simulate(n int) {

}
