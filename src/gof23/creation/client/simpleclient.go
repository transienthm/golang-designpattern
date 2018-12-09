package main

import (
	"gof23/creation"
	"fmt"
)

func main() {
	factory := creation.Factory{}

	apple := factory.MakeApple()
	orange := factory.MakeOrange()

	fmt.Println(apple)
	fmt.Println(orange)
}
