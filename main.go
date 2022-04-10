package main

import "gol/world"

func main() {
	var w world.World

	// w.Init(5, 8)
	// w.RunTest()

	w.Init(50, 200)
	w.Run()
}
