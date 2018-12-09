package state

import "fmt"

//抽象状态角色
type State interface {
	NextState() State
	Update()
}

//具体状态角色
type GameStartState struct {
}

type GameRunState struct {
}

type GameEndState struct {
}

func (this *GameStartState) NextState() State {
	fmt.Println("Start next...")
	return new(GameRunState)
}

func (this *GameStartState) Update() {
	fmt.Println("Game start...")
}

func (this *GameRunState) NextState() State {
	fmt.Println("Run next...")
	return new(GameEndState)
}

func (this *GameRunState) Update() {
	fmt.Println("Game run...")
}

func (this *GameEndState) NextState() State {
	fmt.Println("End next...")
	return new(GameStartState)
}

func (this *GameEndState) Update() {
	fmt.Println("End")
}