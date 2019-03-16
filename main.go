package main

import (
	"fmt"
	"github.com/josephthomashines/TicTacGo/board"
)

func main() {
	board := board.Init(3,3)
	bp := &board

	fmt.Println(bp)
}
