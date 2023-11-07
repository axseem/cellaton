package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/axseem/cellaton"
	"github.com/axseem/cellaton/neighborhood"
	"github.com/axseem/cellaton/rules"
)

func main() {
	field := cellaton.NewField(15, 15, rules.Gof, neighborhood.Moore)
	field.SetCells([][]int{
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1, 0, 0},
		{0, 0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1, 0},
		{0, 1, 1, 1, 0, 0, 1, 1, 1, 0, 0, 1, 1, 1, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1, 0, 0},
		{0, 0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1, 0},
		{0, 1, 1, 1, 0, 0, 1, 1, 1, 0, 0, 1, 1, 1, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1, 0, 0},
		{0, 0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1, 0},
		{0, 1, 1, 1, 0, 0, 1, 1, 1, 0, 0, 1, 1, 1, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	})

	for {
		field.NextGen()
		fmt.Printf(
			"generation: %v\n%v",
			strconv.Itoa(field.GetGeneration()),
			cellaton.CellsToString(field.GetCells(), "[]", " ."),
		)
		time.Sleep(100 * time.Millisecond)
	}
}
