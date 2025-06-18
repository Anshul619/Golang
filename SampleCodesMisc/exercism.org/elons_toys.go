package main

import (
	"fmt"
	"log"
)

type Car struct {
	battery      int
	batteryDrain int
	speed        int
	distance     int
}

// NewCar creates a new remote controlled car with full battery and given specifications.
func NewCar(speed, batteryDrain int) Car {
	car := Car{
		batteryDrain: batteryDrain,
		speed:        speed,
		battery:      100,
	}
	return car
}

// TODO: define the 'Drive()' method
func (c *Car) Drive() {
	if c.battery >= c.batteryDrain {
		c.battery -= c.batteryDrain
		c.distance += c.speed
	}
}

// TODO: define the 'DisplayDistance() string' method
func (c Car) DisplayDistance() string {
	return fmt.Sprintf("Driven %d meters", c.distance)
}

// TODO: define the 'DisplayBattery() string' method
func (c Car) DisplayBattery() string {
	return fmt.Sprintf("Battery at %d%%", c.battery)
}

// TODO: define the 'CanFinish(trackDistance int) bool' method
func (c *Car) CanFinish(trackDistance int) bool {
	if (float64(trackDistance) / float64(c.speed) * float64(c.batteryDrain)) <= float64(c.battery) {
		return true
	}
	return false
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.

func main() {

	speed := 3
	batteryDrain := 3
	car := NewCar(speed, batteryDrain)
	//car.Drive()
	// log.Println(car)
	car.battery = 60

	log.Println(car.DisplayDistance())
	log.Println(car.DisplayBattery())
	log.Println(car.CanFinish(61))

	// speed := 3
	// batteryDrain := 3

	// car = NewCar(5, 2)
	// car.battery = 40
	// log.Println(car.CanFinish(100))
}
