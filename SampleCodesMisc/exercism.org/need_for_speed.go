package main

import "log"

// TODO: define the 'Car' type struct
type Car struct {
    battery int
    batteryDrain int
    speed int
    distance int
}

// NewCar creates a new remote controlled car with full battery and given specifications.
func NewCar(speed, batteryDrain int) Car {
	car := Car {
	    batteryDrain : batteryDrain,
	    speed : speed,
	    battery: 100,
	}
	return car
}

// TODO: define the 'Track' type struct

type Track struct {
    distance int
}
// NewTrack creates a new track
func NewTrack(distance int) Track {
	return Track{distance}
}

// Drive drives the car one time. If there is not enough battery to drive one more time,
// the car will not move.
func Drive(car Car) Car {
    if car.battery >= car.batteryDrain {
        car.battery -= car.batteryDrain
        car.distance += car.speed
    }
	return car
}

// CanFinish checks if a car is able to finish a certain track.
func CanFinish(car Car, track Track) bool {
	if ((track.distance/car.speed)*car.batteryDrain) <= car.battery {
	    return true
	}
	return false
}

func main() {
    log.Println(NewCar(5, 2))
    log.Println(NewTrack(800))
    log.Println(Drive(NewCar(5, 2)))

    speed := 5
    batteryDrain := 2
    car := NewCar(speed, batteryDrain)

    distance := 100
    track := NewTrack(distance)

    log.Println(CanFinish(car, track))

    speed = 5
    batteryDrain = 5
    car = NewCar(speed, batteryDrain)
    distance = 20
    track = NewTrack(distance)

    log.Println(CanFinish(car, track))
}
