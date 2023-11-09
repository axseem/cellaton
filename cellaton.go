package cellaton

import (
	"math/rand"
	"slices"
)

type cell struct {
	state      bool
	neighbours int
}

type Rule struct {
	Born []int
	Surv []int
}

type field struct {
	width        int
	height       int
	cells        *[2][][]cell
	currentCells int
	rule         Rule
	generation   int
	neighborhood [][2]int
}

func NewField(w, h int, r Rule, n [][2]int) *field {
	c1 := make([][]cell, h)
	c2 := make([][]cell, h)
	for i := range c1 {
		c1[i] = make([]cell, w)
		c2[i] = make([]cell, w)
	}

	return &field{
		width:        w,
		height:       h,
		cells:        &[2][][]cell{c1, c2},
		rule:         r,
		neighborhood: n,
	}
}

func (f *field) GetGeneration() int {
	return f.generation
}

func (f *field) GetCells() *[][]cell {
	return &f.cells[f.currentCells]
}

func (f *field) getNextCells() *[][]cell {
	return &f.cells[(f.currentCells+1)%2]
}

func (f *field) aliveCell(x, y int, s bool) {
	(*f.getNextCells())[y][x].state = s
	for _, p := range f.neighborhood {
		(*f.getNextCells())[backToBounds(y+p[1], f.height)][backToBounds(x+p[0], f.width)].neighbours += 1
	}
}

func (f *field) NextGen() {
	for y, row := range *f.GetCells() {
		for x, c := range row {
			if (!c.state && slices.Contains(f.rule.Born, c.neighbours)) ||
				(c.state && slices.Contains(f.rule.Surv, c.neighbours)) {
				f.aliveCell(x, y, true)
			}
			(*f.GetCells())[y][x] = cell{}
		}
	}
	f.currentCells = (f.currentCells + 1) % 2
	(*f).generation += 1
}

func (f *field) SetCells(cells [][]int) {
	for y, row := range *f.GetCells() {
		for x := range row {
			if cells[y][x] == 1 {
				f.aliveCell(x, y, true)
			}
		}
	}
}

func (f *field) FillRand(density int, seed int64) {
	s := rand.NewSource(seed)
	r := rand.New(s)

	for y, row := range *f.GetCells() {
		for x := range row {
			if r.Intn(100) < density {
				f.aliveCell(x, y, true)
			}
		}
	}
}

func backToBounds(n, b int) int {
	return (n + b) % b
}
