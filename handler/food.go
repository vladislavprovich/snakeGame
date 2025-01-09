package handler

import "math/rand"

func SpawnFood(state *GameState) {
	for {
		state.Food = Cords{X: rand.Intn(state.BoardX), Y: rand.Intn(state.BoardY)}
		valid := true
		for _, s := range state.Snake {
			if s == state.Food {
				valid = false
				break
			}
		}
		if valid {
			break
		}
	}
}
