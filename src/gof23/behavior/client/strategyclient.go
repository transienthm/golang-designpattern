package main

import (
	"gof23/behavior/strategy"
	"fmt"
)

func main() {
	var total float64 = 0

	context := strategy.NewCashContext("满300返100")
	total += context.Stratege.AcceptCash(1 * 10000)

	context = strategy.NewCashContext("正常收费")
	total += context.Stratege.AcceptCash(1 * 10000)

	context = strategy.NewCashContext("打八折")
	total += context.Stratege.AcceptCash(1 * 10000)
	fmt.Println(total)
}
