package handler

import (
	"fmt"
	"os"
	"os/exec"
)

func clearScreen() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func DrawBoard(state *GameState) {
	clearScreen()
	for x := 0; x < state.BoardX; x++ {
		for y := 0; y < state.BoardY; y++ {
			isSnake := false
			for _, s := range state.Snake {
				if s.X == x && s.Y == y {
					isSnake = true
					break
				}
			}
			if isSnake {
				fmt.Print("□")
			} else if state.Food.X == x && state.Food.Y == y {
				fmt.Print("o")
			} else {
				fmt.Print("~")
			}
		}
		fmt.Println()
	}
	fmt.Printf("Score: %d\n", state.Score)
}
