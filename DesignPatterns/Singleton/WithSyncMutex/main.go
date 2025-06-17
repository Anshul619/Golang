package main

/*
- Reference - https://refactoring.guru/design-patterns/singleton/go/example
*/

import (
	"sync"
)

func main() {

	var wg sync.WaitGroup

	for i := 0; i < 30; i++ {
		wg.Add(1)
		go func() {
			getInstance()
			wg.Done()
		}()
	}

	wg.Wait()
}
