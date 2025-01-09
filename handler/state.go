package handler

type Cords struct {
	X, Y int
}

type GameState struct {
	Snake     []Cords
	Food      Cords
	Direction Cords
	BoardX    int
	BoardY    int
	GameOver  bool
	Score     int
}

func NewGameState(boardX, boardY int) GameState {
	return GameState{
		Snake:     []Cords{{boardX / 2, boardY / 2}},
		Direction: Cords{0, 0},
		BoardX:    boardX,
		BoardY:    boardY,
		GameOver:  false,
		Score:     0,
	}
}
