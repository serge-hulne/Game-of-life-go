package world

import (
	"os"
)

type Cell struct {
	Alive bool
	X     int
	Y     int
	N     int
}

func (c *Cell) ReprTest() string {
	if c.Alive == true {
		return "x"
	} else {
		return "0"
	}
}

func (c *Cell) Repr() string {
	if c.Alive == true {
		return "x"
	} else {
		return " "
	}
}

// Checking neighbouring cels
func CheckNeighbours(w World, c Cell) Cell {

	// Builds a slice of neigbouring cells.
	x_max := w.Lines
	y_max := w.Columns
	border := []int{-1, 0, 1}

	// Checking for limit conditions
	var neighbours []Cell
	for _, i := range border {
		for _, j := range border {
			if (c.X+i >= 0) && (c.X+i < x_max) &&
				(c.Y+j >= 0) && (c.Y+j < y_max) &&
				(i != 0 || j != 0) {
				neighbours = append(neighbours, w.Cells[c.X+i][c.Y+j])
			}
		}
	}

	// Number of neighbours alive:
	neighbours_alive := 0
	for _, n := range neighbours {
		if n.Alive {
			neighbours_alive++
		}
	}

	c.N = neighbours_alive

	// Total nb of cells alive
	cells_alive := 0
	for _, l := range w.Cells {
		for _, c := range l {
			if c.Alive {
				cells_alive++
			}
		}
	}

	// Exit, if no cells left
	if cells_alive == 0 {
		println("All cells dead!")
		os.Exit(0)
	}

	// Conway's conditions
	// "Any live cell with fewer than two live neighbours dies, as if by underpopulation."
	if neighbours_alive < 2 && c.Alive {
		c.Alive = false
	}

	// "Any live cell with two or three live neighbours lives on to the next generation."
	if (neighbours_alive == 2) && (neighbours_alive == 3) && c.Alive {
		c.Alive = true
	}

	// "Any dead cell with exactly three live neighbours becomes a live cell, as if by reproduction."
	if (neighbours_alive == 3) && !c.Alive {
		c.Alive = true
	}

	// "Any live cell with more than three live neighbours dies, as if by overpopulation."
	if (neighbours_alive > 3) && c.Alive {
		c.Alive = false
	}

	return c

}
