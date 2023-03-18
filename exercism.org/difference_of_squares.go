package main

import "log"

func SquareOfSum(n int) int {
	sum := 0

	for i := 1; i <= n; i++ {
		sum += i
	}

	return sum * sum
}

func SumOfSquares(n int) int {
	sum := 0

	for i := 1; i <= n; i++ {
		sum += i * i
	}

	return sum
}

func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}

func main() {
	log.Println(SquareOfSum(10))
	log.Println(SumOfSquares(10))
	log.Println(Difference(10))
}
