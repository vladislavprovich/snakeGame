package handler

func Update(state *GameState) {
	if state.Direction == (Cords{0, 0}) {
		return
	}

	head := state.Snake[0]
	newHead := Cords{head.X + state.Direction.X, head.Y + state.Direction.Y}

	if newHead.X < 0 || newHead.X >= state.BoardX || newHead.Y < 0 || newHead.Y >= state.BoardY {
		state.GameOver = true
		return
	}

	for _, s := range state.Snake {
		if s == newHead {
			state.GameOver = true
			return
		}
	}

	if newHead == state.Food {
		state.Score++
		state.Snake = append([]Cords{newHead}, state.Snake...)
		SpawnFood(state)
	} else {
		state.Snake = append([]Cords{newHead}, state.Snake[:len(state.Snake)-1]...)
	}
}
