package main

import (
	"fmt"
	"gof23/creation"
)

func main() {
	var builder creation.Builder = &creation.CharacterBuilder{}
	var director *creation.Director = &creation.Director {builder}
	var character *creation.Character = director.Create("loader", "AK47")
	fmt.Println(character.GetName() + "," + character.GetArms())
}