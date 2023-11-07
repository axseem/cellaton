package rules

import "github.com/axseem/cellaton"

var Gof = cellaton.Rule{
	Born: []int{3},
	Surv: []int{2, 3},
}
var DayAndNight = cellaton.Rule{
	Born: []int{3, 6, 7, 8},
	Surv: []int{3, 4, 6, 7, 8},
}
var Caves = cellaton.Rule{
	Born: []int{5, 6, 7, 8},
	Surv: []int{4, 5, 6, 7, 8},
}
var Anneal = cellaton.Rule{
	Born: []int{4, 6, 7, 8},
	Surv: []int{3, 5, 6, 7, 8},
}
var Seeds = cellaton.Rule{
	Born: []int{2},
	Surv: []int{},
}
var Diamoeba = cellaton.Rule{
	Born: []int{3, 5, 6, 7, 8},
	Surv: []int{5, 6, 7, 8},
}
