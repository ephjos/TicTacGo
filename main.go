package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/josephthomashines/TicTacGo/game"
)

func main() {
	args := os.Args[1:]

	if len(args) > 0 {
		if args[0] == "-h" || args[0] == "--help" {
			fmt.Println("\n./TicTacGo [n] -h\n  n  | toggles simulation mode\n       int number of simulations to run\n\n  -h | show this message\n")
			return
		} else {
			sims, simError := strconv.Atoi(args[0])

			if simError == nil && sims >= 1 {
				game.Simulate(sims)
				return
			}
			fmt.Println("Argument must be >= 1")
			return
		}
	}

	game.Play()
}
