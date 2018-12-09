package main

import (
	"gof23/creation"
	"fmt"
)

func main() {
	factory := creation.PearFactory{}
	pearPicker := factory.MakePicker()
	pear := factory.MakePlant()

	fmt.Println(pear)
	fmt.Println(pearPicker)
}