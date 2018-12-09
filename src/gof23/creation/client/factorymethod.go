package main

import (
	"gof23/creation"
	"fmt"
)

func main() {
	bananaPicker := creation.BananaPicker{}

	banana := bananaPicker.MakeFruit("banana")
	fmt.Println(banana)
}
