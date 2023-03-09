package main

import "fmt"

func main() {
	addidasFactory, _ := GetSportsAndFactory("addidas")
	nikeFactory, _ := GetSportsAndFactory("nike")

	nikeShoe := nikeFactory.makeShoe()
	nikeShirt := nikeFactory.makeShirt()

	addidasShirt := addidasFactory.makeShirt()
	addidasShoe := addidasFactory.makeShirt()

	printShoeDetails(nikeShoe)
	printShoeDetails(addidasShoe)

	printShirtDetails(nikeShirt)
	printShirtDetails(addidasShirt)
}

func printShoeDetails(s IShoe) {
	fmt.Printf("Logo: %s", s.getLogo())
	fmt.Println()
	fmt.Printf("Size: %d", s.getSize())
	fmt.Println()
}

func printShirtDetails(s IShirt) {
	fmt.Printf("Logo: %s", s.getLogo())
	fmt.Println()
	fmt.Printf("Size: %d", s.getSize())
	fmt.Println()
}
