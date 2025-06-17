package main

import (
	"log"
	"sync"
)

type single struct{}

var singleInstance *single

func getInstance() *single {
	if singleInstance == nil {
        log.Println("Creating single instance now")
        singleInstance = new(single)
    } else {
        log.Println("Single instance already created")
    }

	return singleInstance
}
