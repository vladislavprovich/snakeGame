package handler

import (
	"syscall"
	"time"
)

var (
	user32               = syscall.NewLazyDLL("user32.dll")
	procGetAsyncKeyState = user32.NewProc("GetAsyncKeyState")
)

func checkKeyState(key int) bool {
	ret, _, _ := procGetAsyncKeyState.Call(uintptr(key))
	return ret != 0
}

func ListenInput(state *GameState) {
	for !state.GameOver {
		if checkKeyState(0x41) { // A
			if state.Direction.Y == 0 {
				state.Direction = Cords{0, -1}
			}
		}
		if checkKeyState(0x44) { // D
			if state.Direction.Y == 0 {
				state.Direction = Cords{0, 1}
			}
		}
		if checkKeyState(0x57) { // W
			if state.Direction.X == 0 {
				state.Direction = Cords{-1, 0}
			}
		}
		if checkKeyState(0x53) { // S
			if state.Direction.X == 0 {
				state.Direction = Cords{1, 0}
			}
		}
		if checkKeyState(0x51) { // Q
			state.GameOver = true
		}
		time.Sleep(50 * time.Millisecond)
	}
}
