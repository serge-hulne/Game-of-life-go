package world

import (
	"fmt"
	"math/rand"
	"time"
)

type World struct {
	Cells   [][]Cell
	Lines   int
	Columns int
}

// Initialize World w
func (w *World) Init(lines int, columns int) {
	w.Lines = lines
	w.Columns = columns
	var rnd bool
	rand.Seed(time.Now().UnixNano())
	for l := 0; l < w.Lines; l++ {
		var cells_line []Cell
		for c := 0; c < w.Columns; c++ {
			if rand.Intn(2) == 0 {
				rnd = false
			} else {
				rnd = true
			}
			cells_line = append(cells_line, Cell{rnd, l, c, 0})
		}
		w.Cells = append(w.Cells, cells_line)
	}
}

// Print World w (data)
func (w *World) PrintValues() {
	for _, l := range w.Cells {
		for _, c := range l {
			fmt.Printf("%v,", c)
		}
		println()
	}
}

// Print World w (data)
func (w *World) PrintValues1() {
	for _, l := range w.Cells {
		for _, c := range l {
			// fmt.Printf("(%s %d)", c.ReprTest(), c.N)
			fmt.Printf("(%d)", c.N)
		}
		println()
	}
}

// Print World w (representation)
func (w *World) Print() {
	for _, l := range w.Cells {
		for _, c := range l {
			fmt.Printf("%s", c.Repr())
		}
		println()
	}
}

// Print World w (representation)
func (w *World) Print1() {
	for _, l := range w.Cells {
		for _, c := range l {
			fmt.Printf("%s", c.ReprTest())
		}
		println()
	}
}

func (w *World) Print2() {
	for _, l := range w.Cells {
		for _, c := range l {
			fmt.Printf("(%s)", c.ReprTest())
		}
		println()
	}
}

func (w *World) ModifyCell(c Cell) {
	w.Cells[c.X][c.Y] = c
}

// Run simulation

func (w World) RunTest() {

	clear()

	// Loop over generations (for test)
	for i := 0; i < 3; i++ {
		time.Sleep(1 * time.Millisecond)

		w.Print2()
		println()

		var w1 World
		w1.Init(w.Lines, w.Columns)
		for _, l := range w.Cells {
			for _, c := range l {
				w1.Cells[c.X][c.Y] = CheckNeighbours(w, c)
			}
		}
		w = w1

		w.PrintValues1()
		println("- - -")
	}
}

func (w World) Run() {
	clear()
	// Loop over generations (endlessly)
	for {
		clear()
		w.Print()
		// time.Sleep(1 * time.Millisecond)

		var w1 World
		w1.Init(w.Lines, w.Columns)
		for _, l := range w.Cells {
			for _, c := range l {
				w1.Cells[c.X][c.Y] = CheckNeighbours(w, c)
			}
		}
		w = w1

	}
}
