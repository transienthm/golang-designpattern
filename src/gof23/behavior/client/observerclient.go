package main

import "gof23/behavior/observer"

func main() {
	s := new(observer.Subject)
	o := new(observer.Observer)

	s.AddObservers(o)

	s.NotifyObservers()
}
