package main

import (
    "log"
    "fmt"
)
// NeedsLicense determines whether a license is needed to drive a type of vehicle. Only "car" and "truck" require a license.
func NeedsLicense(kind string) bool {
    if kind == "car" || kind == "truck" {
        return true
    }

    return false
}

// ChooseVehicle recommends a vehicle for selection. It always recommends the vehicle that comes first in lexicographical order.
func ChooseVehicle(option1, option2 string) string {
    if option1 < option2 {
        return fmt.Sprintf("%s is clearly the better choice.", option1)
    }
    return fmt.Sprintf("%s is clearly the better choice.", option2)
}

// CalculateResellPrice calculates how much a vehicle can resell for at a certain age.
func CalculateResellPrice(originalPrice, age float64) float64 {
	if age <= 3 {
	    return originalPrice*0.8
	} else if age >= 10 {
	    return originalPrice*0.5
	}
	return originalPrice*0.7
}

func main(){
    log.Println(NeedsLicense("car"))
    log.Println(NeedsLicense("bike"))
    log.Println(NeedsLicense("truck"))
    log.Println(ChooseVehicle("Wuling Hongguang", "Toyota Corolla"))
    log.Println(ChooseVehicle("Volkswagen Beetle", "Volkswagen Golf"))
    log.Println(CalculateResellPrice(1000, 1))
    log.Println(CalculateResellPrice(1000, 5))
    log.Println(CalculateResellPrice(1000, 15))
}