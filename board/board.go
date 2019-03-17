package board

import (
	"fmt"
	"os"
)

const PLACEHOLDER = "_"

// Errors
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

// Core
type Board struct {
	x, y int
	data [][]string
}

func (b *Board) String() string {
	out := ""

	for _, row := range (*b).data {
		for _, val := range row {
			out += fmt.Sprintf("%s ", val)
		}
		out += fmt.Sprintf("\n")
	}

	return out
}

func (b *Board) Set(x, y int, char string) {
	if x > b.x-1 {
		fmt.Println(&BoardOutOfBoundsError{x, b.x - 1, "x"})
		os.Exit(1)
	} else if x < 0 {
		fmt.Println(&BoardOutOfBoundsError{x, 0, "x"})
		os.Exit(1)
	} else if y > b.y-1 {
		fmt.Println(&BoardOutOfBoundsError{y, b.y - 1, "y"})
		os.Exit(1)
	} else if y < 0 {
		fmt.Println(&BoardOutOfBoundsError{y, 0, "y"})
		os.Exit(1)
	}

	(*b).data[x][y] = char
}

func (b *Board) Get() [][]string {
	return b.data
}

func Init(x, y int) Board {
	data := make([][]string, x)

	for i := range data {
		data[i] = make([]string, y)
	}

	for i, row := range data {
		for j := range row {
			data[i][j] = PLACEHOLDER
		}
	}

	return Board{x, y, data}
}
