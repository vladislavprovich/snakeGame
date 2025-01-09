package main

import (
	"fmt"
	"math/rand"
	"ojhguhsfdiugdfihsghidf/handler"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	state := handler.NewGameState(10, 20)
	handler.SpawnFood(&state)

	go handler.ListenInput(&state)

	ticker := time.NewTicker(300 * time.Millisecond)
	defer ticker.Stop()

	for !state.GameOver {
		<-ticker.C
		handler.Update(&state)
		handler.DrawBoard(&state)
	}

	fmt.Println("Game Over! Final Score:", state.Score)
}
