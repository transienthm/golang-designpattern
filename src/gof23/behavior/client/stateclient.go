package main

import (
	"gof23/behavior/state"
	"time"
)

//context角色
func stateMechine(state state.State, ch chan int)  {
	for  {
		select {
			case i := <-ch :
				if i == 1 {
					state.Update()
					state = state.NextState()
				} else if i == 0 {
					return
				}
			default:

		}
	}
}

func main() {
	st := new(state.GameStartState)
	ch := make(chan int)

	go stateMechine(st, ch)
	time.Sleep(time.Microsecond * 3)
	ch <- 1
	time.Sleep(time.Microsecond * 3)
	ch <- 1
	time.Sleep(time.Microsecond * 3)
	ch <- 1
	time.Sleep(time.Microsecond * 3)
	ch <- 0
}