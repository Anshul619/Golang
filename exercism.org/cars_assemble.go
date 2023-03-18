package main

import "log"

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	return float64(productionRate) * (successRate / 100)
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	return int(CalculateWorkingCarsPerHour(productionRate, successRate) / 60)
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
	return (uint)(int(carsCount/10)*95000 + ((carsCount % 10) * 10000))
}

func main() {
	log.Println(CalculateWorkingCarsPerHour(1547, 90))
	log.Println(CalculateWorkingCarsPerMinute(1105, 90))
	log.Println(CalculateCost(37))
	log.Println(CalculateCost(21))
}
