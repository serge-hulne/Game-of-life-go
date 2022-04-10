package world

import (
	"fmt"
	"time"
)

// Clear terminal
func clear() {
	fmt.Print("\033[H\033[2J")
}

// Move cursor
func move(l int, c int) {
	fmt.Printf("\033[%d;%dH", l, c)
}

// Show cursor
func show() {
	fmt.Print("\033[?25h")
}

// Hide cursor
func hide() {
	fmt.Print("\033[?25l")
}

// Test for the terminal IO functions
func moveAndPrint() {
	clear()
	hide()
	defer show()
	for i := 0; i < 3; i++ {
		move(1, 10)
		print("Testing term IO")
		move(10, 1)
		print(i)
		time.Sleep(time.Second)
	}
}
