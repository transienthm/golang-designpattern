package main

import (
	"fmt"
)

func main() {
	duck := &MallardDuck{}
	turkey := &WildTurkey{}

	turkeyAdapter := NewTurkeyAdapter(turkey)

	fmt.Println("The Turkey says...")
	turkey.gobble()
	turkey.fly()

	fmt.Println("The Duck says...")
	duck.quack()
	duck.fly()

	fmt.Println("The Turkey Adapter says...")
	turkeyAdapter.quack()
	turkeyAdapter.fly()
}

type Duck interface {
	quack()
	fly()
}

type Turkey interface {
	gobble()
	fly()
}

type MallardDuck struct {
}

func (*MallardDuck) quack() {
	fmt.Println("Quark...")
}

func (*MallardDuck) fly() {
	fmt.Println("flying...")
}

type WildTurkey struct {
}

func (*WildTurkey) gobble() {
	fmt.Println("Gobble...")
}

func (*WildTurkey) fly() {
	fmt.Println("flying a short distance")
}

type TurkeyAdapter struct {
	turkey Turkey
}

func NewTurkeyAdapter(turkey Turkey) *TurkeyAdapter {
	return &TurkeyAdapter{turkey}
}

func (this *TurkeyAdapter) quack() {
	this.turkey.gobble()
}

func (this *TurkeyAdapter) fly() {
	for i := 0; i < 5; i++ {
		this.turkey.fly()
	}
}