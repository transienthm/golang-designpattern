package main

import (
	"sync"
)

var (
	instance *Singleton
	lock     *sync.Mutex = &sync.Mutex{}
	once sync.Once
)

type Singleton struct {
}

func GetInstance() *Singleton {

	once.Do(func() {
		instance = &Singleton{}
	})
	return instance
}
