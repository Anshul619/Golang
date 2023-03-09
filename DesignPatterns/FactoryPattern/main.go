package main

/*
- Reference - https://refactoring.guru/design-patterns/factory-method/go/example
*/
import "log"

func main() {
	log.Println("Test")

	ak47, _ := getGun("ak47")
	musket, _ := getGun("musket")

	print(ak47)
	print(musket)
}

func print(i IGun) {
	log.Println(i.getName())
	log.Println(i.getPower())
}
