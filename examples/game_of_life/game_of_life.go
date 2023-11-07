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
	field := cellaton.NewField(48, 24, rules.Gof, neighborhood.Moore)
	field.FillRand(50)

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
