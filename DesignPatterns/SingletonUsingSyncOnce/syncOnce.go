package main

/*
- Reference - https://refactoring.guru/design-patterns/singleton/go/example#example-1
*/
import (
	"log"
	"sync"
)

var once sync.Once

type single struct{}

var singleInstance *single

func getInstance() *single {
	if singleInstance == nil {
		once.Do(
			func() {
				log.Println("Creating instance")
				singleInstance = new(single)
			})
	} else {
		log.Println("Instance already created")
	}

	return singleInstance
}
