package main

import "log"

func main() {
	log.Println("main routine start")
	go printNumbers()
	log.Println("main routine end")
}

func printNumbers() {
	for i := 0; i < 100; i++ {
		log.Println("go routine %v", i)
	}
}
