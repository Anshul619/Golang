package main

import (
	"log"
	"sync"
)

func main() {
	log.Println("main routine start")
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		printNumbers()
		wg.Done()
	}()
	wg.Wait()

	log.Println("main routine end")
}

func printNumbers() {
	for i := 0; i < 100; i++ {
		log.Println("go routine %v", i)
	}
}
