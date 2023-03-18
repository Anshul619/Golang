package main

import (
	"log"
	"sync"
)

var lock = new(sync.Mutex)

type single struct{}

var singleInstance *single

func getInstance() *single {
	if singleInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if singleInstance == nil {
			log.Println("Creating single instance now")
			singleInstance = new(single)
		} else {
			log.Println("Single instance already created")
		}
	} else {
		log.Println("Single instance already created")
	}

	return singleInstance
}
