package cellaton

func CellsToString(cells *[][]cell, alive, dead string) string {
	var s string
	for _, row := range *cells {
		for _, c := range row {
			if c.state {
				s += alive
			} else {
				s += dead
			}
		}
		s += "\n"
	}
	return s
}
